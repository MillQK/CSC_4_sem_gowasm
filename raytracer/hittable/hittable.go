package hittable

import (
	"CSC_4_sem_gowasm/raytracer/entities"
	"CSC_4_sem_gowasm/raytracer/hittable/materials"
)

type Hittable interface {
	Hit(ray *entities.Ray, tMin, tMax float64) *materials.HitRecord
}
