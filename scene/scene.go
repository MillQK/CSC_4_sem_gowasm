package scene

import (
	rt "CSC_4_sem_gowasm/raytracer"
	"CSC_4_sem_gowasm/raytracer/hitable"
)

type Scene struct {
	Camera       rt.Camera
	RaysPerPixel int32
	HitableList  hitable.HitableList
}
