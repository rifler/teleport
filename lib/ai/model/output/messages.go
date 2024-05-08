/*
 * Teleport
 * Copyright (C) 2023  Gravitational, Inc.
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

package output

import "strings"

// Message represents a new message within a live conversation.
type Message struct {
	Content string
}

// StreamingMessage represents a new message that is being streamed from the LLM.
type StreamingMessage struct {
	Parts <-chan string
}

// WaitAndConsume waits until the message stream is over and returns the full message.
// This can only be called once on a message as it empties its Parts channel.
func (msg *StreamingMessage) WaitAndConsume() string {
	sb := strings.Builder{}
	for part := range msg.Parts {
		sb.WriteString(part)
	}
	return sb.String()
}

// Label represents a label returned by OpenAI's completion API.
type Label struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// CompletionCommand represents a command suggestion returned by OpenAI's completion API.
type CompletionCommand struct {
	Command string   `json:"command,omitempty"`
	Nodes   []string `json:"nodes,omitempty"`
	Labels  []Label  `json:"labels,omitempty"`
}

// GeneratedCommand represents a Bash command generated by LLM.
type GeneratedCommand struct {
	Command string `json:"command"`
}

// AccessRequest represents an access request suggestion returned by OpenAI's completion API.
type AccessRequest struct {
	Roles              []string   `json:"roles"`
	Resources          []Resource `json:"resources"`
	Reason             string     `json:"reason"`
	SuggestedReviewers []string   `json:"suggested_reviewers"`
}

// Resource represents a resource suggestion returned by OpenAI's completion API.
type Resource struct {
	// The resource type.
	Type string `json:"type"`

	// The resource name.
	Name string `json:"id"`

	// Set if a display-friendly alternative name is available.
	FriendlyName string `json:"friendlyName,omitempty"`
}