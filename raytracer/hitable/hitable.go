package hitable

import (
	"CSC_4_sem_gowasm/raytracer/entities"
)

type Hitable interface {
	Hit(ray entities.Ray, tMin, tMax float64) *entities.HitRecord
}
