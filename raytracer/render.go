package raytracer

import (
	"CSC_4_sem_gowasm/raytracer/entities"
	"CSC_4_sem_gowasm/raytracer/hitable"
	"CSC_4_sem_gowasm/scene"
	"math"
	"math/rand"
)

func color(r *entities.Ray, world hitable.Hitable) entities.Vec3 {
	if hit := world.Hit(*r, 0.001, math.MaxFloat64); hit != nil {
		target := entities.NewZeroVec3().AddAssign(hit.Point).AddAssign(hit.Normal).AddAssign(PointOnUnitSphereSurface()).SubAssign(hit.Point)
		return color(entities.NewRay(hit.Point, *target), world).MulScalar(0.5)
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
		averageColor.AddAssign(color(&ray, &scene.HitableList))
	}

	averageColor.ModifyFields(func(x float64) float64 {
		return 255.99 * math.Sqrt(x/float64(scene.RaysPerPixel))
	})

	return *averageColor
}
