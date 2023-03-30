package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	openai "github.com/sashabaranov/go-openai"
	"github.com/sesaquecruz/go-openai-dalle/external"
	"github.com/sesaquecruz/go-openai-dalle/internal"
)

func main() {
	godotenv.Load()

	apiKey, ok := os.LookupEnv("OPENAI_API_KEY")
	if !ok {
		log.Fatalln("the API_KEY was not found")
	}

	client := openai.NewClient(apiKey)
	path := "out/"
	dalle2 := external.NewDallE2(client, path)

	ctx := context.Background()

	internal.StartCli(ctx, dalle2)
}
