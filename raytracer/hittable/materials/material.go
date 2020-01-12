package materials

import (
	"github.com/MillQK/gowasm_raytracer/raytracer/entities"
)

type ScatteredRay struct {
	Ray         entities.Ray
	Attenuation entities.Vec3
}

func MakeScatteredRay(ray entities.Ray, attenuation entities.Vec3) ScatteredRay {
	return ScatteredRay{ray, attenuation}
}

func NewScatteredRay(ray entities.Ray, attenuation entities.Vec3) *ScatteredRay {
	return &ScatteredRay{ray, attenuation}
}

type Material interface {
	Scatter(ray *entities.Ray, hit *HitRecord) *ScatteredRay
}
