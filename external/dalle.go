package external

import (
	"context"
	"image"
)

type DallE interface {
	GenerateImage(ctx context.Context, description string) (image.Image, error)
	SaveImage(image image.Image) (*string, error)
}
