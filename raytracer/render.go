package raytracer

import (
	"github.com/MillQK/gowasm_raytracer/raytracer/entities"
	"github.com/MillQK/gowasm_raytracer/raytracer/hittable"
	"github.com/MillQK/gowasm_raytracer/scene"
	"math"
	"math/rand"
)

func color(r *entities.Ray, world hittable.Hittable, depth uint32) entities.Vec3 {
	if hit := world.Hit(r, 0.001, math.MaxFloat64); hit != nil {

		if scatter := hit.Material.Scatter(r, hit); depth < 50 && scatter != nil {
			return scatter.Attenuation.Mul(color(&scatter.Ray, world, depth+1))
		}

		return *entities.NewZeroVec3()
	} else {
		unitDirection := r.Direction.UnitVector()
		t := 0.5 * (unitDirection.Y + 1.0)
		return *entities.NewVec3(1.0, 1.0, 1.0).MulScalarAssign(1.0 - t).AddAssign(*entities.NewVec3(0.5, 0.7, 1.0).MulScalarAssign(t))
	}
}

func RenderPixel(scene *scene.Scene, pixelWidth, pixelHeight uint32) entities.Vec3 {
	averageColor := entities.NewVec3(0, 0, 0)

	for s := uint32(0); s < scene.RaysPerPixel; s++ {
		u := (float64(pixelWidth) + rand.Float64()) / float64(scene.Width)
		v := (float64(pixelHeight) + rand.Float64()) / float64(scene.Height)

		ray := scene.Camera.GetRay(u, v)
		averageColor.AddAssign(color(&ray, scene.HittableList, 0))
	}

	averageColor.ModifyFields(func(x float64) float64 {
		return 255.99 * math.Sqrt(x/float64(scene.RaysPerPixel))
	})

	return *averageColor
}
