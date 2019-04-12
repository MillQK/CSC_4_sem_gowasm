package raytracer

import (
	"fmt"
	"io"
)

type Image struct {
	Width, Height uint32
	Pixels        []Color
}

func MakeImageWithBackground(width, height uint32, color Color) Image {
	pixels := make([]Color, width*height)
	for i := uint32(0); i < width*height; i++ {
		pixels[i] = color
	}

	return Image{width, height, pixels}
}

func MakeImage(width, height uint32) Image {
	return MakeImageWithBackground(width, height, MakeColor(0, 0, 0))
}

func (image *Image) GetPixel(w, h uint32) *Color {
	return &image.Pixels[h*image.Width+w]
}

func (image *Image) WriteAsPpm(output io.Writer) error {
	if _, err := fmt.Fprintf(output, "P3\n%d %d\n255\n", image.Width, image.Height); err != nil {
		return err
	}

	for _, pixel := range image.Pixels {
		if _, err := fmt.Fprintf(output, "%d %d %d\n", pixel.R, pixel.G, pixel.B); err != nil {
			return err
		}
	}

	return nil
}
