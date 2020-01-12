package scene

import (
	"github.com/MillQK/gowasm_raytracer/raytracer/entities"
	"github.com/MillQK/gowasm_raytracer/raytracer/hittable"
)

type Scene struct {
	Camera        *entities.Camera
	RaysPerPixel  uint32
	HittableList  *hittable.HittableList
	Width, Height uint32
}
