// Package provider re-exports the provider types from omnillm-core for API compatibility.
//
// This allows code using "github.com/plexusone/omnillm/provider" to work unchanged
// when switching between omnillm (thick) and omnillm-core (thin).
package provider

import (
	core "github.com/plexusone/omnillm-core/provider"
)

// Re-export interfaces.
type (
	// Provider is the interface for LLM providers.
	Provider = core.Provider

	// ChatCompletionStream is the interface for streaming responses.
	ChatCompletionStream = core.ChatCompletionStream
)

// Re-export types.
type (
	// Role represents the role of a message sender.
	Role = core.Role

	// Message represents a chat message.
	Message = core.Message

	// ToolCall represents a tool function call.
	ToolCall = core.ToolCall

	// ToolFunction represents the function being called.
	ToolFunction = core.ToolFunction

	// ChatCompletionRequest represents a request for chat completion.
	ChatCompletionRequest = core.ChatCompletionRequest

	// ResponseFormat specifies the format of the response.
	ResponseFormat = core.ResponseFormat

	// Tool represents a tool that can be called.
	Tool = core.Tool

	// ToolSpec defines a tool specification.
	ToolSpec = core.ToolSpec

	// ChatCompletionResponse represents a response from chat completion.
	ChatCompletionResponse = core.ChatCompletionResponse

	// ChatCompletionChoice represents a single choice in the response.
	ChatCompletionChoice = core.ChatCompletionChoice

	// Usage represents token usage information.
	Usage = core.Usage

	// ChatCompletionChunk represents a chunk in streaming response.
	ChatCompletionChunk = core.ChatCompletionChunk
)

// Re-export role constants.
const (
	RoleSystem    = core.RoleSystem
	RoleUser      = core.RoleUser
	RoleAssistant = core.RoleAssistant
	RoleTool      = core.RoleTool
)
