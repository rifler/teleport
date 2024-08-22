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

package autoupdate

import (
	"github.com/gravitational/trace"

	"github.com/gravitational/teleport/api/gen/proto/go/teleport/autoupdate/v1"
	headerv1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/header/v1"
	"github.com/gravitational/teleport/api/types"
)

// NewAutoupdateConfig creates a new autoupdate configuration resource.
func NewAutoupdateConfig(spec *autoupdate.AutoupdateConfigSpec) (*autoupdate.AutoupdateConfig, error) {
	config := &autoupdate.AutoupdateConfig{
		Kind:    types.KindAutoupdateConfig,
		Version: types.V1,
		Metadata: &headerv1.Metadata{
			Name: types.MetaNameAutoupdateConfig,
		},
		Spec: spec,
	}
	if err := ValidateAutoupdateConfig(config); err != nil {
		return nil, trace.Wrap(err)
	}

	return config, nil
}

// ValidateAutoupdateConfig checks that required parameters are set
// for the specified AutoupdateConfig.
func ValidateAutoupdateConfig(c *autoupdate.AutoupdateConfig) error {
	if c == nil {
		return trace.BadParameter("AutoupdateConfig is nil")
	}
	if c.Metadata == nil {
		return trace.BadParameter("Metadata is nil")
	}
	if c.Spec == nil {
		return trace.BadParameter("Spec is nil")
	}

	return nil
}