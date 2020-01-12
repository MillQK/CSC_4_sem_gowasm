package hittable

import (
	"github.com/MillQK/gowasm_raytracer/raytracer/entities"
	"github.com/MillQK/gowasm_raytracer/raytracer/hittable/materials"
)

type HittableList struct {
	List []Hittable
}

func NewHittableListSize(size uint32) *HittableList {
	return &HittableList{make([]Hittable, size)}
}

func NewHittableList(list []Hittable) *HittableList {
	return &HittableList{list}
}

func (list *HittableList) Hit(ray *entities.Ray, tMin, tMax float64) *materials.HitRecord {
	var hitRecord *materials.HitRecord = nil
	closest := tMax

	for _, figure := range list.List {
		if hit := figure.Hit(ray, tMin, closest); hit != nil {
			closest = hit.T
			hitRecord = hit
		}
	}

	return hitRecord
}
