# OmniLLM

[![Go Reference][docs-godoc-svg]][docs-godoc-url]
[![Go Report Card][goreport-svg]][goreport-url]

 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/plexusone/omnillm
 [docs-godoc-url]: https://pkg.go.dev/github.com/plexusone/omnillm
 [goreport-svg]: https://goreportcard.com/badge/github.com/plexusone/omnillm
 [goreport-url]: https://goreportcard.com/report/github.com/plexusone/omnillm

Batteries-included LLM client that bundles [omnillm-core](https://github.com/plexusone/omnillm-core) with all thick providers.

## Installation

```bash
go get github.com/plexusone/omnillm
```

## Quick Start

```go
import "github.com/plexusone/omnillm"

client, _ := omnillm.NewClient(omnillm.ClientConfig{
    Provider: omnillm.ProviderNameOpenAI,
    APIKey:   os.Getenv("OPENAI_API_KEY"),
})
```

## Provider Support

### Thin Providers (omnillm-core)

Lightweight implementations using stdlib `net/http`:

| Provider | Streaming | Tools | JSON Mode |
|----------|-----------|-------|-----------|
| OpenAI | Yes | Yes | Yes |
| Anthropic | Yes | Yes | No |
| Gemini | Yes | No | No |
| X.AI (Grok) | Yes | Yes | Yes |
| GLM (Zhipu) | Yes | Yes | No |
| Kimi (Moonshot) | Yes | No | No |
| Qwen (Alibaba) | Yes | Yes | No |
| Ollama | Yes | Yes | Yes |

### Thick Providers (Official SDKs)

Full-featured implementations using official vendor SDKs:

| Provider | Module | Streaming | Tools | JSON Mode |
|----------|--------|-----------|-------|-----------|
| OpenAI | [omnillm-openai](https://github.com/plexusone/omnillm-openai) | Yes | Yes | Yes |
| Anthropic | [omnillm-anthropic](https://github.com/plexusone/omnillm-anthropic) | Yes | Yes | No |
| Gemini | [omnillm-gemini](https://github.com/plexusone/omnillm-gemini) | Yes | No | No |
| Bedrock | [omnillm-bedrock](https://github.com/plexusone/omnillm-bedrock) | Yes | Yes | No |

Thick providers automatically override thin providers when imported.

## Thin vs Thick

| Aspect | Thin (omnillm-core) | Thick (omnillm-*) |
|--------|---------------------|-------------------|
| Dependencies | Minimal (stdlib) | Official SDK |
| API Coverage | Core features | Full coverage |
| Retries | Manual | SDK-managed |
| Auth | API key | SDK-managed |

## Selective Import

Import only what you need:

```go
import (
    omnillm "github.com/plexusone/omnillm-core"
    _ "github.com/plexusone/omnillm-openai" // Only OpenAI thick provider
)
```

## Documentation

- [omnillm-core](https://github.com/plexusone/omnillm-core) - Core interfaces and thin providers
- [API Reference](https://pkg.go.dev/github.com/plexusone/omnillm) - GoDoc

## License

MIT
