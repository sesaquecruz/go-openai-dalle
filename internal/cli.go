package internal

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sesaquecruz/go-openai-dalle/external"
)

func StartCli(ctx context.Context, dalle external.DallE) {
	fmt.Print("\n[Press 'ctrl + c' to exit]\n")

	fmt.Print("\n    ____  ___    __    __       ______   ___\n")
	fmt.Print("   / __ \\/   |  / /   / /      / ____/  |__ \\\n")
	fmt.Print("  / / / / /| | / /   / /      / __/     __/ /\n")
	fmt.Print(" / /_/ / ___ |/ /___/ /___   / /___    / __/\n")
	fmt.Print("/_____/_/  |_/_____/_____/  /_____/   /____/\n\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\nimage description > ")

		description, err := reader.ReadString('\n')
		if err != nil {
			log.Panicln(err)
		}

		description = description[:len(description)-1]

		if len(description) == 0 {
			continue
		}

		fmt.Print("generating...\n")

		image, err := dalle.GenerateImage(ctx, description)
		if err != nil {
			log.Panicln(err)
		}

		path, err := dalle.SaveImage(image)
		if err != nil {
			log.Panicln(err)
		}

		fmt.Printf("done, localization: %s\n", *path)
	}
}
