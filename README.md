# Custom Bento Distribution

This repository provides a unified [Bento](https://warpstreamlabs.github.io/bento/) build containing all custom plugins I'm maintaining. It brings together geographic utilities, amateur radio integrations, binary data processing, and LLM/AI capabilities into a single powerful stream-processing binary...

## Included Plugins

This distribution compiles the following Go modules into Bento:

### 1. [Geo-Bento](https://github.com/akhenakh/geo-bento)
A suite of mapping processors to transform and enrich geographic coordinates from a stream.
*   `country`: Resolves the country for a given latitude/longitude.
*   `tz`: Resolves the timezone (e.g., `Europe/Paris`) for a given latitude/longitude.
*   `h3`: Transforms latitude/longitude into an Uber H3 cell index.
*   `a5`: Transforms latitude/longitude into an A5 cell index.
*   `s2`: Transforms latitude/longitude into a Google S2 cell index.
*   `randpos` *(Input)*: Generates random coordinates within a specified bounding box (useful for debugging/testing).

### 1. [Bento-APRS](https://github.com/akhenakh/bento-aprs)
An input plugin (`aprs_is`) for connecting to APRS-IS servers and streaming parsed APRS (Automatic Packet Reporting System) packets. 
*   Connects globally to relay APRS data (Weather, Telemetry, Position tracking).
*   Outputs structured JSON messages with automatic metadata extraction (`@aprs_src`, `@aprs_dst`, `@aprs_type`).
*   Supports server-side APRS-IS filters and automatic reconnections.

### 1. [Bento-CBOR](https://github.com/akhenakh/bento-cbor)
A processor plugin (`cbor`) for converting between JSON and CBOR (Concise Binary Object Representation) formats.
*   `operator: to_json`: Converts incoming CBOR binary payloads into JSON.
*   `operator: from_json`: Converts JSON payloads into CBOR binary format.
*   Fully RFC 7049 and RFC 8949 compliant.

### 1. [Bento-LLM](https://github.com/akhenakh/bento-llm)
A processor plugin (`llm`) that allows you to query Large Language Models directly within your data pipelines.
*   Supports multiple AI providers: `openai`, `anthropic`, `openrouter`, and generic `openai-compat` endpoints (like Ollama, LM Studio).
*   Supports **Model Context Protocol (MCP)** via `stdio`, `http`, and `sse`, allowing your LLM agent to access external tools and data sources autonomously during the stream processing.
*   Fully supports Bento's Bloblang interpolation in prompts.

---

## Building the Custom Binary

To build your custom Bento binary containing all the plugins above, ensure you have Go installed, then clone this repository and run:

```bash
go mod tidy
go build -o bento-plugins main.go
```

## Example Usage

Run a pipeline utilizing multiple custom plugins:

```bash
./bento-plugins -c pipeline.yaml
```

**Example Pipeline (`pipeline.yaml`):**
```yaml
input:
  # Generate random coordinates using Geo-Bento
  randpos:
    min_lat: 46.0
    max_lat: 48.0
    min_lng: 2.0
    max_lng: 2.3

pipeline:
  processors:
    # Enrich the payload with H3 cell and Timezone using Geo-Bento
    - mapping: |
        root = this
        root.tz = tz(this.lat, this.lng)
    
    # Use an LLM to generate a quick summary string of the location
    - llm:
        provider: "openai"
        model: "gpt-4o-mini"
        api_key: "${OPENAI_API_KEY}"
        prompt: "Write a one sentence welcome message for someone located in the timezone ${! json('tz') }."
        
    #  Convert the final JSON payload to CBOR for efficient storage
    - cbor:
        operator: from_json

output:
  file:
    path: output.cbor
```
## Build the full Bento with extra
```sh
CGO_ENABLED=1 go build -tags "x_bento_extra" -ldflags="-s -w" -trimpath .
```

Zig CC build:
```sh
GOOS=linux GOARCH=amd64  CGO_ENABLED=1 CGO_CFLAGS="-I/usr/include" CC="zig cc -target x86_64-linux-musl" go build -tags "x_bento_extra" -o bento-amd64 --ldflags '-linkmode=external -extldflags=-static'  .
GOOS=linux GOARCH=riscv64  CGO_ENABLED=1  CC="zig cc -target riscv64-linux-musl" go build -tags "x_bento_extra" -o bento-risc-v --ldflags '-linkmode=external -extldflags=-static' .
GOOS=linux GOARCH=arm64  CGO_ENABLED=1  CC="zig cc -target aarch64-linux-musl" go build -tags "x_bento_extra" -o bento-aarch64 --ldflags '-linkmode=external -extldflags=-static' .
```
## License

All individual plugins and this combined distribution are licensed under the [MIT License](LICENSE). Check individual repositories for specific details.
