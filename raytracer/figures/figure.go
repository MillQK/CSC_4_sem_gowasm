package figures

import rt "CSC_4_sem_gowasm/raytracer"

type Figure interface {
	Hit(ray rt.Ray, tMin, tMax float64) *rt.HitRecord
}
