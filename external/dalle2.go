package external

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

type DallE2 struct {
	Client *openai.Client
	Size   string
	N      int
	Path   string
}

func NewDallE2(client *openai.Client, path string) *DallE2 {
	return &DallE2{
		Client: client,
		Size:   openai.CreateImageSize512x512,
		N:      1,
		Path:   path,
	}
}

func (d *DallE2) GenerateImage(ctx context.Context, description string) (image.Image, error) {
	request := openai.ImageRequest{
		Prompt:         description,
		Size:           d.Size,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              d.N,
	}

	responseBase64, err := d.Client.CreateImage(ctx, request)
	if err != nil {
		return nil, err
	}

	responseBytes, err := base64.StdEncoding.DecodeString(responseBase64.Data[0].B64JSON)
	if err != nil {
		return nil, err
	}

	responseImage, err := png.Decode(bytes.NewReader(responseBytes))
	if err != nil {
		return nil, err
	}

	return responseImage, nil
}

func (d DallE2) SaveImage(image image.Image) (*string, error) {
	path := fmt.Sprintf("%sdalle2_%s.png", d.Path, time.Now().Format("2006-01-02_15-04-05-001"))

	err := os.MkdirAll(filepath.Dir(path), 0755)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	err = png.Encode(file, image)
	if err != nil {
		return nil, err
	}

	return &path, nil
}
