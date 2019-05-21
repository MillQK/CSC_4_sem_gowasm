package materials

import (
	"CSC_4_sem_gowasm/raytracer/entities"
)

type HitRecord struct {
	T             float64
	Point, Normal entities.Vec3
	Material      Scatter
}

func MakeHitRecord(t float64, point, normal entities.Vec3, material Scatter) HitRecord {
	return HitRecord{t, point, normal, material}
}

func NewHitRecord(t float64, point, normal entities.Vec3, material Scatter) *HitRecord {
	return &HitRecord{t, point, normal, material}
}
