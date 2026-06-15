# OmniLLM

[![Go CI][go-ci-svg]][go-ci-url]
[![Go Lint][go-lint-svg]][go-lint-url]
[![Go SAST][go-sast-svg]][go-sast-url]
[![Go Report Card][goreport-svg]][goreport-url]
[![Docs][docs-godoc-svg]][docs-godoc-url]
[![Docs][docs-mkdoc-svg]][docs-mkdoc-url]
[![Visualization][viz-svg]][viz-url]
[![License][license-svg]][license-url]

 [go-ci-svg]: https://github.com/plexusone/omnillm/actions/workflows/go-ci.yaml/badge.svg?branch=main
 [go-ci-url]: https://github.com/plexusone/omnillm/actions/workflows/go-ci.yaml
 [go-lint-svg]: https://github.com/plexusone/omnillm/actions/workflows/go-lint.yaml/badge.svg?branch=main
 [go-lint-url]: https://github.com/plexusone/omnillm/actions/workflows/go-lint.yaml
 [go-sast-svg]: https://github.com/plexusone/omnillm/actions/workflows/go-sast-codeql.yaml/badge.svg?branch=main
 [go-sast-url]: https://github.com/plexusone/omnillm/actions/workflows/go-sast-codeql.yaml
 [goreport-svg]: https://goreportcard.com/badge/github.com/plexusone/omnillm
 [goreport-url]: https://goreportcard.com/report/github.com/plexusone/omnillm
 [docs-godoc-svg]: https://pkg.go.dev/badge/github.com/plexusone/omnillm
 [docs-godoc-url]: https://pkg.go.dev/github.com/plexusone/omnillm
 [docs-mkdoc-svg]: https://img.shields.io/badge/Go-dev%20guide-blue.svg
 [docs-mkdoc-url]: https://plexusone.dev/omnillm
 [viz-svg]: https://img.shields.io/badge/Go-visualizaton-blue.svg
 [viz-url]: https://mango-dune-07a8b7110.1.azurestaticapps.net/?repo=plexusone%2Fomnillm
 [loc-svg]: https://tokei.rs/b1/github/plexusone/omnillm
 [repo-url]: https://github.com/plexusone/omnillm
 [license-svg]: https://img.shields.io/badge/license-MIT-blue.svg
 [license-url]: https://github.com/plexusone/omnillm/blob/main/LICENSE

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
| OpenAI | [omni-openai](https://github.com/plexusone/omni-openai) | Yes | Yes | Yes |
| Anthropic | [omni-anthropic](https://github.com/plexusone/omni-anthropic) | Yes | Yes | No |
| Gemini | [omni-google](https://github.com/plexusone/omni-google) | Yes | No | No |
| Bedrock | [omni-aws](https://github.com/plexusone/omni-aws) | Yes | Yes | No |
| OpenRouter | [omni-openrouter](https://github.com/plexusone/omni-openrouter) | Yes | Yes | Yes |

Thick providers automatically override thin providers when imported.

## Thin vs Thick

| Aspect | Thin (omnillm-core) | Thick (omni-*) |
|--------|---------------------|----------------|
| Dependencies | Minimal (stdlib) | Official SDK |
| API Coverage | Core features | Full coverage |
| Retries | Manual | SDK-managed |
| Auth | API key | SDK-managed |

## Selective Import

Import only what you need:

```go
import (
    omnillm "github.com/plexusone/omnillm-core"
    _ "github.com/plexusone/omni-openai/omnillm" // Only OpenAI thick provider
)
```

## Documentation

- [omnillm-core](https://github.com/plexusone/omnillm-core) - Core interfaces and thin providers
- [API Reference](https://pkg.go.dev/github.com/plexusone/omnillm) - GoDoc

## License

MIT
