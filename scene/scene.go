package scene

import (
	"CSC_4_sem_gowasm/raytracer/entities"
	"CSC_4_sem_gowasm/raytracer/hitable"
)

type Scene struct {
	Camera        entities.Camera
	RaysPerPixel  uint32
	HitableList   hitable.HitableList
	Width, Height uint32
}
