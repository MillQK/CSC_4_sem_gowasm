package hittable

import (
	"github.com/MillQK/gowasm_raytracer/raytracer/entities"
	"github.com/MillQK/gowasm_raytracer/raytracer/hittable/materials"
	"math"
)

type Sphere struct {
	Center   entities.Vec3
	Radius   float64
	Material materials.Material
}

func MakeSphere(center entities.Vec3, radius float64, material materials.Material) Sphere {
	return Sphere{center, radius, material}
}

func NewSphere(center entities.Vec3, radius float64, material materials.Material) *Sphere {
	return &Sphere{center, radius, material}
}

func (sphere *Sphere) Hit(ray *entities.Ray, tMin, tMax float64) *materials.HitRecord {
	oc := ray.Origin.Sub(sphere.Center)
	a := ray.Direction.Dot(ray.Direction)
	b := 2.0 * ray.Direction.Dot(oc)
	c := oc.Dot(oc) - sphere.Radius*sphere.Radius

	discriminant := b*b - 4.0*a*c

	if discriminant > 0.0 {
		t := (-b - math.Sqrt(discriminant)) / (2.0 * a)

		if tMin < t && t < tMax {
			point := ray.PointAtParameter(t)
			return materials.NewHitRecord(t, point, point.Sub(sphere.Center).DivScalar(sphere.Radius), sphere.Material)
		}

		t = (-b + math.Sqrt(discriminant)) / (2.0 * a)
		if tMin < t && t < tMax {
			point := ray.PointAtParameter(t)
			return materials.NewHitRecord(t, point, point.Sub(sphere.Center).DivScalar(sphere.Radius), sphere.Material)
		}
	}

	return nil
}
