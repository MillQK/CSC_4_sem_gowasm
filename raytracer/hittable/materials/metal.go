package materials

import (
	"CSC_4_sem_gowasm/raytracer/entities"
)

type Metal struct {
	Albedo entities.Vec3
	Fuzz   float64
}

func MakeMetal(albedo entities.Vec3, fuzz float64) Metal {
	return *NewMetal(albedo, fuzz)
}

func NewMetal(albedo entities.Vec3, fuzz float64) *Metal {
	if fuzz > 1.0 {
		fuzz = 1.0
	}

	return &Metal{albedo, fuzz}
}

func (mat *Metal) Scatter(ray *entities.Ray, hit *HitRecord) *ScatteredRay {

	reflected := ray.Direction.UnitVector().Reflect(hit.Normal)
	unitSpherePoint := PointOnUnitSphereSurface()

	scattered := entities.MakeRay(hit.Point, unitSpherePoint.MulScalar(mat.Fuzz).Add(reflected))

	if scattered.Direction.Dot(hit.Normal) > 0.0 {
		return NewScatteredRay(scattered, mat.Albedo)
	}

	return nil
}
