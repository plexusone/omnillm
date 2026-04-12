module github.com/plexusone/omnillm

go 1.26.0

require (
	github.com/plexusone/omnillm-anthropic v0.0.0
	github.com/plexusone/omnillm-core v0.1.0
	github.com/plexusone/omnillm-gemini v0.0.0
	github.com/plexusone/omnillm-openai v0.0.0
)

require (
	cloud.google.com/go v0.123.0 // indirect
	cloud.google.com/go/auth v0.18.2 // indirect
	cloud.google.com/go/compute/metadata v0.9.0 // indirect
	github.com/anthropics/anthropic-sdk-go v1.30.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/s2a-go v0.1.9 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.12 // indirect
	github.com/googleapis/gax-go/v2 v2.17.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/grokify/mogo v0.74.1 // indirect
	github.com/grokify/sogo v0.14.0 // indirect
	github.com/openai/openai-go v1.12.0 // indirect
	github.com/tidwall/gjson v1.18.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.1 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.65.0 // indirect
	go.opentelemetry.io/otel v1.40.0 // indirect
	go.opentelemetry.io/otel/metric v1.40.0 // indirect
	go.opentelemetry.io/otel/trace v1.40.0 // indirect
	golang.org/x/crypto v0.49.0 // indirect
	golang.org/x/net v0.52.0 // indirect
	golang.org/x/sync v0.20.0 // indirect
	golang.org/x/sys v0.42.0 // indirect
	golang.org/x/text v0.35.0 // indirect
	google.golang.org/genai v1.52.1 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20260226221140-a57be14db171 // indirect
	google.golang.org/grpc v1.79.1 // indirect
	google.golang.org/protobuf v1.36.11 // indirect
)

replace (
	github.com/plexusone/omnillm-anthropic => ../omnillm-anthropic
	github.com/plexusone/omnillm-core => ../omnillm-core
	github.com/plexusone/omnillm-gemini => ../omnillm-gemini
	github.com/plexusone/omnillm-openai => ../omnillm-openai
)
