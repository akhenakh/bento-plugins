package main

import (
	"context"

	"github.com/warpstreamlabs/bento/public/service"

	// Import all standard Bento components
	_ "github.com/warpstreamlabs/bento/public/components/all"

	// Import APRS plugin
	_ "github.com/akhenakh/bento-aprs/aprs"

	// Import CBOR plugin
	_ "github.com/akhenakh/bento-cbor"

	// Import LLM plugin
	_ "github.com/akhenakh/bento-llm/llm"

	// Import Geo plugins
	_ "github.com/akhenakh/geo-bento/a5"
	_ "github.com/akhenakh/geo-bento/country"
	_ "github.com/akhenakh/geo-bento/h3"
	_ "github.com/akhenakh/geo-bento/randpos"
	_ "github.com/akhenakh/geo-bento/s2"
	_ "github.com/akhenakh/geo-bento/tz"

	// Import IRC plugin
	_ "github.com/akhenakh/bento-irc/irc"

	// AIS
	_ "github.com/akhenakh/bento-aisstream/ais"

	// Bluesky
	_ "github.com/akhenakh/bento-bluesky/bluesky"
)

func main() {
	service.RunCLI(context.Background())
}
