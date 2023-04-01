package e2e

import (
	"os"

	lemonsqueezy "github.com/NdoleStudio/lemonsqueezy-go"
	_ "github.com/joho/godotenv/autoload" // import LEMONSQUEEZY_API_KEY from .env file
)

var client = lemonsqueezy.New(lemonsqueezy.WithAPIKey(os.Getenv("LEMONSQUEEZY_API_KEY")))
