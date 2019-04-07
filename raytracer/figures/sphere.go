package figures

import (
	rt "CSC_4_sem_gowasm/raytracer"
	"math"
)

type Sphere struct {
	Center rt.Vec3
	Radius float64
}

func MakeSphere(center rt.Vec3, radius float64) Sphere {
	return Sphere{center, radius}
}

func (sphere *Sphere) Hit(ray rt.Ray, tMin, tMax float64) *rt.HitRecord {
	oc := ray.Origin.Sub(sphere.Center)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * ray.Direction.Dot(oc)
	c := oc.Dot(oc) - sphere.Radius*sphere.Radius

	discriminant := b*b - 4.0*a*c

	if discriminant > 0.0 {
		t := (-b - math.Sqrt(discriminant)) / (2.0 * a)

		if tMin <= t && t <= tMax {
			point := ray.PointAtParameter(t)
			hitRecord := rt.MakeHitRecord(t, point, point.Sub(sphere.Center).DivScalar(sphere.Radius))
			return &hitRecord
		}

		t = (-b + math.Sqrt(discriminant)) / (2.0 * a)
		if tMin <= t && t <= tMax {
			point := ray.PointAtParameter(t)
			hitRecord := rt.MakeHitRecord(t, point, point.Sub(sphere.Center).DivScalar(sphere.Radius))
			return &hitRecord
		}
	}

	return nil
}
