/*
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package workloadattest

import (
	"context"
	"path"
	"regexp"
	"strconv"
	"strings"

	"github.com/gravitational/trace"
	v1 "k8s.io/api/core/v1"
	"k8s.io/utils/mount"
)

// KubernetesAttestation holds the Kubernetes pod information retrieved from
// the workload attestation process.
type KubernetesAttestation struct {
	// Attested is true if the PID was successfully attested to a Kubernetes
	// pod. This indicates the validity of the rest of the fields.
	Attested bool
	// Namespace is the namespace of the pod.
	Namespace string
	// ServiceAccount is the service account of the pod.
	ServiceAccount string
	// Container is the individual container that the PID resolved to within
	// the pod.
	Container string
	// PodName is the name of the pod.
	PodName string
	// PodUID is the UID of the pod.
	PodUID string
	// Labels is a map of labels on the pod.
	Labels map[string]string
}

type KubernetesAttestor struct {
	// TODO: Configurable bits...
}

// Attest resolves the Kubernetes pod information from the
// PID of the workload.
//
// From what I can tell, there's two common ways of doing this:
// - /proc/<pid>/mountinfo
// - /proc/<pid>/cgroup
//
// This implementation leverages /proc/<pid>/mountinfo
//
// Requires `hostPID: true` so we can see the /proc of other pods.
//
// We can then query the kubelet api to find the pod that this corresponds to.
func (a *KubernetesAttestor) Attest(ctx context.Context, pid int) (KubernetesAttestation, error) {
	podID, containerID, err := a.getContainerAndPodID(pid)
	if err != nil {
		return KubernetesAttestation{}, trace.Wrap(err, "determining pod and container ID")
	}

	pod, err := a.getPodForID(ctx, podID)
	if err != nil {
		return KubernetesAttestation{}, trace.Wrap(err, "finding pod by ID")
	}

	var container *v1.ContainerStatus
	for _, c := range pod.Status.ContainerStatuses {
		if c.ContainerID == containerID {
			container = &c
			break
		}
	}
	if container == nil {
		for _, c := range pod.Status.InitContainerStatuses {
			if c.ContainerID == containerID {
				container = &c
				break
			}
		}
	}
	if container == nil {
		return KubernetesAttestation{}, trace.BadParameter("container %q not found in pod %q", containerID, pod.Name)
	}

	return KubernetesAttestation{
		Attested:       true,
		Namespace:      pod.Namespace,
		ServiceAccount: pod.Spec.ServiceAccountName,
		Container:      container.Name,
		PodName:        pod.Name,
		PodUID:         string(pod.UID),
		Labels:         pod.Labels,
	}, nil
}

// getContainerAndPodID retrieves the container ID and pod ID for the provided
// PID.
func (a *KubernetesAttestor) getContainerAndPodID(pid int) (podID string, containerID string, err error) {
	info, err := mount.ParseMountInfo(
		path.Join("/proc", strconv.Itoa(pid), "mountinfo"),
	)
	if err != nil {
		return "", "", trace.Wrap(
			err, "parsing mountinfo",
		)
	}

	// Find the cgroup or cgroupv2 mount
	// TODO(noah): Is it possible for there to be multiple cgroup mounts?
	// If so, how should we handle.
	// I imagine with cgroup v1, we get one mount per controller, but all should
	// be fairly equivelant.
	var cgroupMount mount.MountInfo
	for _, m := range info {
		if m.FsType == "cgroup" || m.FsType == "cgroup2" {
			cgroupMount = m
			break
		}
	}

	podID, containerID, err = mountpointSourceToContainerAndPodID(
		cgroupMount.Source,
	)
	if err != nil {
		return "", "", trace.Wrap(
			err, "parsing cgroup mount (source: %q)", cgroupMount.Source,
		)
	}
	return podID, containerID, nil
}

var (
	// TODO: This is a super naive implementation that may only work in my
	// cluster. This needs revisiting before merging.

	// A container ID is usually a 64 character hex string, so this regex just
	// selects for that.
	containerIDRegex = regexp.MustCompile(`(?P<containerID>[[:xdigit:]]{64})`)
	// A pod ID is usually a UUID prefaced with "pod".
	// There are two main cgroup drivers:
	// - systemd , the dashes are replaced with underscores
	// - cgroupfs, the dashes are kept.
	podIDRegex = regexp.MustCompile(`pod(?P<podID>[[:xdigit:]]{8}[_-][[:xdigit:]]{4}[_-][[:xdigit:]]{4}[_-][[:xdigit:]]{4}[_-][[:xdigit:]]{12})`)
)

// mountpointSourceToContainerAndPodID takes the source of the cgroup mountpoint
// and extracts the container ID and pod ID from it.
// TODO: This is a super naive implementation that may only work in my
// cluster. This needs revisiting before merging.
func mountpointSourceToContainerAndPodID(source string) (podID string, containerID string, err error) {
	// From the mount, we need to extract the container ID and pod ID.
	// Unfortunately this process can be a little fragile, as the format of
	// the mountpoint varies across Kubernetes implementations.
	// The format of the mountpoint varies, but can look something like:
	// /kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-pod30e5e887_5bea_42fb_a256_ec9d6a76efc7.slice/cri-containerd-22985f2d7e6472530eabf5ed449b0c84899f38f60e778cbee5c1642f1b24cda6.scope

	matches := containerIDRegex.FindStringSubmatch(source)
	if len(matches) != 2 {
		return "", "", trace.BadParameter(
			"expected 2 matches searching for container ID but found %d",
			len(matches),
		)
	}
	containerID = matches[1]
	if containerID == "" {
		return "", "", trace.BadParameter(
			"source does not contain container ID",
		)
	}

	matches = podIDRegex.FindStringSubmatch(source)
	if len(matches) != 2 {
		return "", "", trace.BadParameter(
			"expected 2 matches searching for pod ID but found %d",
			len(matches),
		)
	}
	podID = matches[1]
	if podID == "" {
		return "", "", trace.BadParameter(
			"source does not contain pod ID",
		)
	}

	// When using the `systemd` cgroup driver, the dashes are replaced with
	// underscores. So let's correct that.
	podID = strings.ReplaceAll(podID, "_", "-")

	return podID, containerID, nil
}

// getPodForID retrieves the pod information for the provided pod ID.
// https://github.com/kubernetes/kubernetes/blob/master/pkg/kubelet/server/server.go#L371
func (a *KubernetesAttestor) getPodForID(ctx context.Context, podID string) (*v1.Pod, error) {
	return nil, trace.NotImplemented("method not implemented")
}
