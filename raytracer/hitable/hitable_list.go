package hitable

import "CSC_4_sem_gowasm/raytracer"

type HitableList struct {
	List []Hitable
}

func NewHitableListSize(size uint32) *HitableList {
	return &HitableList{make([]Hitable, size)}
}

func NewHitableList(list []Hitable) *HitableList {
	return &HitableList{list}
}

func (list *HitableList) Hit(ray raytracer.Ray, tMin, tMax float64) *raytracer.HitRecord {
	var hitRecord *raytracer.HitRecord = nil
	closest := tMax

	for _, figure := range list.List {
		if hit := figure.Hit(ray, tMin, closest); hit != nil {
			closest = hit.T
			hitRecord = hit
		}
	}

	return hitRecord
}
