package main

import (
	rt "CSC_4_sem_gowasm/raytracer"
	"CSC_4_sem_gowasm/raytracer/figures"
	"io"
	"math"
	"os"
)

func color(r *rt.Ray, sphere *figures.Sphere) rt.Vec3 {
	if hit := sphere.Hit(*r, 0.0, math.MaxFloat64); hit != nil {
		return *hit.Normal.AddScalarAssign(1.0).MulScalarAssign(0.5)
	} else {
		unitDirection := r.Direction.UnitVector()
		t := 0.5 * (unitDirection.Y + 1.0)
		return rt.MakeVec3(1.0, 1.0, 1.0).MulScalar(1.0 - t).Add(rt.MakeVec3(0.5, 0.7, 1.0).MulScalar(t))
	}
}

func printGradientAndCircle(output io.Writer) error {
	image := rt.MakeImage(200, 100)
	lowerLeftCorner := rt.MakeVec3(-2.0, -1.0, -1.0)
	horizontal := rt.MakeVec3(4.0, 0.0, 0.0)
	vertical := rt.MakeVec3(0.0, 2.0, 0.0)
	origin := rt.MakeVec3(0.0, 0.0, 0.0)
	sphere := figures.MakeSphere(rt.MakeVec3(0.0, 0.0, -1.0), 0.5)

	for j := uint32(0); j < image.Height; j++ {
		for i := uint32(0); i < image.Width; i++ {

			u := float64(i) / float64(image.Width)
			v := float64(j) / float64(image.Height)

			ray := rt.MakeRay(origin, lowerLeftCorner.Add(horizontal.MulScalar(u)).Add(vertical.MulScalar(v)))
			color := color(&ray, &sphere)

			imageColor := image.GetColor(i, image.Height-j-1)

			imageColor.R = uint8(255.99 * color.X)
			imageColor.G = uint8(255.99 * color.Y)
			imageColor.B = uint8(255.99 * color.Z)
		}
	}

	return image.WriteAsPpm(output)
}

func main() {
	args := os.Args

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
