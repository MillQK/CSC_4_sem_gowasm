package materials

import (
	"CSC_4_sem_gowasm/raytracer/entities"
)

type Lambertian struct {
	Albedo entities.Vec3
}

func MakeLambertian(albedo entities.Vec3) Lambertian {
	return Lambertian{albedo}
}

func NewLambertian(albedo entities.Vec3) *Lambertian {
	return &Lambertian{albedo}
}

func (mat *Lambertian) Scatter(ray *entities.Ray, hit *HitRecord) *ScatteredRay {

	unitSpherePoint := PointOnUnitSphereSurface()
	target := unitSpherePoint.Add(hit.Normal)

	scattered := entities.MakeRay(hit.Point, target)

	return NewScatteredRay(scattered, mat.Albedo)
}
