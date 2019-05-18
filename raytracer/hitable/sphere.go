package hitable

import (
	"CSC_4_sem_gowasm/raytracer/entities"
	"math"
)

type Sphere struct {
	Center entities.Vec3
	Radius float64
}

func MakeSphere(center entities.Vec3, radius float64) Sphere {
	return Sphere{center, radius}
}

func NewSphere(center entities.Vec3, radius float64) *Sphere {
	return &Sphere{center, radius}
}

func (sphere *Sphere) Hit(ray entities.Ray, tMin, tMax float64) *entities.HitRecord {
	oc := ray.Origin.Sub(sphere.Center)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * ray.Direction.Dot(oc)
	c := oc.Dot(oc) - sphere.Radius*sphere.Radius

	discriminant := b*b - 4.0*a*c

	if discriminant > 0.0 {
		t := (-b - math.Sqrt(discriminant)) / (2.0 * a)

		if tMin < t && t < tMax {
			point := ray.PointAtParameter(t)
			return entities.NewHitRecord(t, point, point.Sub(sphere.Center).DivScalar(sphere.Radius))
		}

		t = (-b + math.Sqrt(discriminant)) / (2.0 * a)
		if tMin < t && t < tMax {
			point := ray.PointAtParameter(t)
			return entities.NewHitRecord(t, point, point.Sub(sphere.Center).DivScalar(sphere.Radius))
		}
	}

	return nil
}
