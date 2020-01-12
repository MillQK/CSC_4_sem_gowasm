package main

import (
	rt "github.com/MillQK/gowasm_raytracer/raytracer"
	"github.com/MillQK/gowasm_raytracer/raytracer/entities"
	"github.com/MillQK/gowasm_raytracer/web/shared"
	"io"
	"math/rand"
	"os"
	"time"
)

func printGradientAndCircle(output io.Writer) error {
	renderScene := shared.DefaultScene()
	image := entities.NewImage(renderScene.Width, renderScene.Height)

	start := time.Now()

	for j := uint32(0); j < image.Height; j++ {
		for i := uint32(0); i < image.Width; i++ {
			image.GetPixel(i, j).FromVec(rt.RenderPixel(renderScene, i, image.Height-j-1))
		}
	}

	shared.PrintMemUsage()

	elapsed := time.Now().Sub(start)
	println("Elapsed time: ", elapsed.String())

	return image.WriteAsPpm(output)
}

func main() {
	args := os.Args
	rand.Seed(time.Now().UnixNano())

	switch len(args) {
	case 1:
		if err := printGradientAndCircle(os.Stdout); err != nil {
			println(err)
		}
	case 2:
		file, err := os.Create(args[1])

		if err != nil {
			println(err)
			return
		}
		defer file.Close()

		err = printGradientAndCircle(file)
		if err != nil {
			println(err)
			return
		}
	default:
		println("Must be 0 or 1 args")
	}
}
