package e2e

import (
	"os"

	"github.com/NdoleStudio/lemonsqueezy-go"
	_ "github.com/joho/godotenv/autoload" // import LEMON_SQUEEZY_API_KEY from .env file
)

var client = lemonsqueezy.New(lemonsqueezy.WithAPIKey(os.Getenv("LEMON_SQUEEZY_API_KEY")))
