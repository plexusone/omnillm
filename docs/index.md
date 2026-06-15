# OmniLLM

OmniLLM is a batteries-included package that bundles [omnillm-core](https://github.com/plexusone/omnillm-core) with all thick providers for a complete LLM integration experience.

## Installation

```bash
go get github.com/plexusone/omnillm
```

## Quick Start

```go
import (
    "github.com/plexusone/omnillm"
)

client, err := omnillm.NewClient(omnillm.ClientConfig{
    Provider: omnillm.ProviderNameOpenAI,
    APIKey:   os.Getenv("OPENAI_API_KEY"),
})
```

Thick providers are automatically registered via `init()` and override thin providers.

## Provider Overview

### Thin Providers (omnillm-core)

Lightweight implementations using stdlib `net/http`:

| Provider | Thin Package | Streaming | Tools | JSON Mode |
|----------|--------------|-----------|-------|-----------|
| OpenAI | Built-in | Yes | Yes | Yes |
| Anthropic | Built-in | Yes | Yes | No |
| Gemini | Built-in | Yes | No | No |
| X.AI (Grok) | Built-in | Yes | Yes | Yes |
| GLM (Zhipu) | Built-in | Yes | Yes | No |
| Kimi (Moonshot) | Built-in | Yes | No | No |
| Qwen (Alibaba) | Built-in | Yes | Yes | No |
| Ollama | Built-in | Yes | Yes | Yes |

### Thick Providers (Official SDKs)

Full-featured implementations using official vendor SDKs:

| Provider | Module | SDK | Streaming | Tools | JSON Mode |
|----------|--------|-----|-----------|-------|-----------|
| OpenAI | `omni-openai` | [openai-go](https://github.com/openai/openai-go) | Yes | Yes | Yes |
| Anthropic | `omni-anthropic` | [anthropic-sdk-go](https://github.com/anthropics/anthropic-sdk-go) | Yes | Yes | No |
| Gemini | `omni-google` | [google.golang.org/genai](https://pkg.go.dev/google.golang.org/genai) | Yes | No | No |
| Bedrock | `omni-aws` | [aws-sdk-go-v2](https://github.com/aws/aws-sdk-go-v2) | Yes | Yes | No |

## Thin vs Thick

| Aspect | Thin (omnillm-core) | Thick (omni-*) |
|--------|---------------------|----------------|
| Dependencies | Minimal (stdlib only) | Official SDK + transitive deps |
| API Coverage | Core features | Full API coverage |
| Retries | Manual | SDK-managed |
| Auth | API key only | SDK-managed (incl. OAuth, IAM) |
| Priority | 0 (default) | 10 (overrides thin) |

When both are available, thick providers take precedence automatically.

## Documentation

- [omnillm-core Documentation](https://plexusone.github.io/omnillm-core) - Core interfaces, types, and thin providers
- [API Reference](https://pkg.go.dev/github.com/plexusone/omnillm) - GoDoc

## Individual Thick Providers

Import only what you need:

```go
import (
    omnillm "github.com/plexusone/omnillm-core"
    _ "github.com/plexusone/omni-openai/omnillm"    // OpenAI thick provider
    _ "github.com/plexusone/omni-anthropic/omnillm" // Anthropic thick provider
    _ "github.com/plexusone/omni-google/omnillm"    // Gemini thick provider
    _ "github.com/plexusone/omni-aws/omnillm"       // AWS Bedrock thick provider
)
```
