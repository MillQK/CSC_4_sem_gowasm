package scene

import (
	"CSC_4_sem_gowasm/raytracer/entities"
	"CSC_4_sem_gowasm/raytracer/hittable"
)

type Scene struct {
	Camera        entities.Camera
	RaysPerPixel  uint32
	HittableList  hittable.HittableList
	Width, Height uint32
}
