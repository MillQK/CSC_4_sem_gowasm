package materials

import (
	"github.com/MillQK/gowasm_raytracer/raytracer/entities"
	"math"
	"math/rand"
)

type Dielectric struct {
	RefractionIndex float64
}

func MakeDielectric(refractionIndex float64) Dielectric {
	return Dielectric{refractionIndex}
}

func NewDielectric(refractionIndex float64) *Dielectric {
	return &Dielectric{refractionIndex}
}

func schlick(cosine, refractionIndex float64) float64 {
	r0 := (1 - refractionIndex) / (1 + refractionIndex)
	r0 = r0 * r0
	return r0 + (1-r0)*math.Pow(1-cosine, 5)
}

func (mat *Dielectric) Scatter(ray *entities.Ray, hit *HitRecord) *ScatteredRay {
	attenuation := entities.MakeVec3(1.0, 1.0, 1.0)
	reflected := ray.Direction.Reflect(hit.Normal)
	var outwardNormal entities.Vec3
	var niOverNt, cosine float64

	if ray.Direction.Dot(hit.Normal) > 0.0 {
		outwardNormal = hit.Normal.MulScalar(-1.0)
		niOverNt = mat.RefractionIndex
		cosine = mat.RefractionIndex * ray.Direction.Dot(hit.Normal) / ray.Direction.Length()
	} else {
		outwardNormal = hit.Normal
		niOverNt = 1.0 / mat.RefractionIndex
		cosine = -ray.Direction.Dot(hit.Normal) / ray.Direction.Length()
	}

	refracted := ray.Direction.Refract(outwardNormal, niOverNt)
	var reflectProbability float64
	if refracted != nil {
		reflectProbability = schlick(cosine, mat.RefractionIndex)
	} else {
		reflectProbability = 1.0
	}

	if rand.Float64() < reflectProbability {
		return NewScatteredRay(entities.MakeRay(hit.Point, reflected), attenuation)
	} else {
		return NewScatteredRay(entities.MakeRay(hit.Point, *refracted), attenuation)
	}
}
