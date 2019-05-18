package hitable

import (
	"CSC_4_sem_gowasm/raytracer/entities"
)

type HitableList struct {
	List []Hitable
}

func NewHitableListSize(size uint32) *HitableList {
	return &HitableList{make([]Hitable, size)}
}

func NewHitableList(list []Hitable) *HitableList {
	return &HitableList{list}
}

func (list *HitableList) Hit(ray entities.Ray, tMin, tMax float64) *entities.HitRecord {
	var hitRecord *entities.HitRecord = nil
	closest := tMax

	for _, figure := range list.List {
		if hit := figure.Hit(ray, tMin, closest); hit != nil {
			closest = hit.T
			hitRecord = hit
		}
	}

	return hitRecord
}
