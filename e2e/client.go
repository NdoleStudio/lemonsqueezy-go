package e2e

import (
	"os"

	lemonsqueezy "github.com/NdoleStudio/lemonsqueezy-go"
)

var client = lemonsqueezy.New(lemonsqueezy.WithAPIKey(os.Getenv("LEMONSQUEEZY_API_KEY")))
