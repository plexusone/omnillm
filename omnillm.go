// Package omnillm provides a batteries-included LLM client with all official SDK providers.
//
// This package imports omnillm-core plus all thick (SDK-based) providers, giving you
// full official SDK support for OpenAI, Anthropic, and Gemini out of the box.
//
// The thick providers automatically override the thin (native HTTP) implementations
// in omnillm-core via the priority-based provider registry.
//
// Usage:
//
//	import "github.com/plexusone/omnillm"
//
//	client := omnillm.NewClient(omnillm.ClientConfig{
//	    Providers: []omnillm.ProviderConfig{
//	        {Provider: omnillm.ProviderNameOpenAI, APIKey: os.Getenv("OPENAI_API_KEY")},
//	        {Provider: omnillm.ProviderNameAnthropic, APIKey: os.Getenv("ANTHROPIC_API_KEY")},
//	        {Provider: omnillm.ProviderNameGemini, APIKey: os.Getenv("GEMINI_API_KEY")},
//	    },
//	})
//
// For a lightweight alternative with minimal dependencies (thin providers only),
// use github.com/plexusone/omnillm-core directly.
package omnillm

import (
	// Import core for re-exports
	core "github.com/plexusone/omnillm-core"

	// Blank imports to trigger init() registration of thick providers.
	// These override the thin providers in omnillm-core.
	_ "github.com/plexusone/omni-anthropic/omnillm"
	_ "github.com/plexusone/omni-aws/omnillm"
	_ "github.com/plexusone/omni-google/omnillm"
	_ "github.com/plexusone/omni-openai/omnillm"
	_ "github.com/plexusone/omni-openrouter/omnillm"
)

// Re-export core types for convenience.
// Users can use omnillm.Provider instead of core.Provider, etc.
type (
	// Provider is the interface for LLM providers.
	Provider = core.Provider

	// ChatCompletionStream is the interface for streaming responses.
	ChatCompletionStream = core.ChatCompletionStream

	// ClientConfig holds configuration for creating a multi-provider client.
	ClientConfig = core.ClientConfig

	// ProviderConfig holds configuration for a single provider.
	ProviderConfig = core.ProviderConfig

	// ProviderName identifies a provider.
	ProviderName = core.ProviderName

	// Message represents a chat message.
	Message = core.Message

	// Role represents the role of a message sender.
	Role = core.Role

	// ChatCompletionRequest is the request for chat completion.
	ChatCompletionRequest = core.ChatCompletionRequest

	// ChatCompletionResponse is the response from chat completion.
	ChatCompletionResponse = core.ChatCompletionResponse

	// ChatCompletionChoice is a single choice in a response.
	ChatCompletionChoice = core.ChatCompletionChoice

	// ChatCompletionChunk is a streaming response chunk.
	ChatCompletionChunk = core.ChatCompletionChunk

	// Tool represents a tool/function definition.
	Tool = core.Tool

	// ToolSpec defines a tool's function specification.
	ToolSpec = core.ToolSpec

	// ToolCall represents a tool call from the model.
	ToolCall = core.ToolCall

	// ToolFunction represents the function details of a tool call.
	ToolFunction = core.ToolFunction

	// Usage tracks token usage.
	Usage = core.Usage

	// ResponseFormat specifies the response format.
	ResponseFormat = core.ResponseFormat

	// Capabilities describes provider features.
	Capabilities = core.Capabilities

	// APIError represents an API error.
	APIError = core.APIError

	// ChatClient is the multi-provider LLM client.
	ChatClient = core.ChatClient

	// ProviderFactory creates providers from config.
	ProviderFactory = core.ProviderFactory

	// ObservabilityHook allows external packages to observe LLM calls.
	ObservabilityHook = core.ObservabilityHook

	// LLMCallInfo provides metadata about the LLM call for observability.
	LLMCallInfo = core.LLMCallInfo
)

// Re-export provider name constants.
const (
	ProviderNameOpenAI     = core.ProviderNameOpenAI
	ProviderNameAnthropic  = core.ProviderNameAnthropic
	ProviderNameGemini     = core.ProviderNameGemini
	ProviderNameXAI        = core.ProviderNameXAI
	ProviderNameGLM        = core.ProviderNameGLM
	ProviderNameKimi       = core.ProviderNameKimi
	ProviderNameQwen       = core.ProviderNameQwen
	ProviderNameOllama     = core.ProviderNameOllama
	ProviderNameBedrock    = core.ProviderNameBedrock
	ProviderNameOpenRouter = "openrouter"
)

// Re-export role constants.
const (
	RoleSystem    = core.RoleSystem
	RoleUser      = core.RoleUser
	RoleAssistant = core.RoleAssistant
	RoleTool      = core.RoleTool
)

// Re-export priority constants.
const (
	PriorityThin  = core.PriorityThin
	PriorityThick = core.PriorityThick
)

// Re-export common errors.
var (
	ErrUnsupportedProvider  = core.ErrUnsupportedProvider
	ErrInvalidConfiguration = core.ErrInvalidConfiguration
	ErrNoProviders          = core.ErrNoProviders
	ErrEmptyAPIKey          = core.ErrEmptyAPIKey
	ErrInvalidAPIKey        = core.ErrInvalidAPIKey
	ErrEmptyModel           = core.ErrEmptyModel
	ErrEmptyMessages        = core.ErrEmptyMessages
	ErrStreamClosed         = core.ErrStreamClosed
	ErrInvalidResponse      = core.ErrInvalidResponse
	ErrRateLimitExceeded    = core.ErrRateLimitExceeded
	ErrQuotaExceeded        = core.ErrQuotaExceeded
	ErrInvalidRequest       = core.ErrInvalidRequest
	ErrModelNotFound        = core.ErrModelNotFound
	ErrServerError          = core.ErrServerError
	ErrNetworkError         = core.ErrNetworkError
)

// Re-export constructor functions.
var (
	NewClient               = core.NewClient
	NewAPIError             = core.NewAPIError
	NewAPIErrorFull         = core.NewAPIErrorFull
	RegisterProvider        = core.RegisterProvider
	GetProviderFactory      = core.GetProviderFactory
	ListRegisteredProviders = core.ListRegisteredProviders
	GetProviderPriority     = core.GetProviderPriority
	ClassifyError           = core.ClassifyError
	IsRetryableError        = core.IsRetryableError
	IsNonRetryableError     = core.IsNonRetryableError
)
