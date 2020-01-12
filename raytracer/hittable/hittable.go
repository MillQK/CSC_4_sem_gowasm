package hittable

import (
	"github.com/MillQK/gowasm_raytracer/raytracer/entities"
	"github.com/MillQK/gowasm_raytracer/raytracer/hittable/materials"
)

type Hittable interface {
	Hit(ray *entities.Ray, tMin, tMax float64) *materials.HitRecord
}
