package shared

import (
	"fmt"
	"github.com/MillQK/gowasm_raytracer/raytracer/entities"
	"github.com/MillQK/gowasm_raytracer/raytracer/hittable"
	"github.com/MillQK/gowasm_raytracer/raytracer/hittable/materials"
	"github.com/MillQK/gowasm_raytracer/scene"
	"math/rand"
	"runtime"
)

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tHeapAlloc = %v MiB", bToMb(m.HeapAlloc))
	fmt.Printf("\tHeapObjects = %v", m.HeapObjects)
	fmt.Printf("\tNextGC = %v MiB", bToMb(m.NextGC))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}

func randomObjects() *hittable.HittableList {
	n := 500
	list := make([]hittable.Hittable, n+1)
	list[0] = hittable.NewSphere(entities.MakeVec3(0, -1000, 0), 1000, materials.NewLambertian(entities.MakeVec3(0.5, 0.5, 0.5)))

	i := 1
	for a := -11; a < 11; a++ {
		for b := -11; b < 11; b++ {
			chooseMaterial := rand.Float64()
			center := entities.MakeVec3(float64(a)+0.9*rand.Float64(), 0.2, float64(b)+0.9*rand.Float64())
			offset := entities.MakeVec3(4, 0.2, 0)

			if center.Sub(offset).Length() > 0.9 {
				if chooseMaterial < 0.8 { // diffuse
					list[i] = hittable.NewSphere(center, 0.2,
						materials.NewLambertian(entities.MakeVec3(rand.Float64()*rand.Float64(), rand.Float64()*rand.Float64(), rand.Float64()*rand.Float64())))
					i++
				} else if chooseMaterial < 0.95 { // metal
					list[i] = hittable.NewSphere(center, 0.2,
						materials.NewMetal(entities.MakeVec3(0.5*(1+rand.Float64()), 0.5*(1+rand.Float64()), 0.5*(1+rand.Float64())), 0.5*rand.Float64()))
					i++
				} else { // glass
					list[i] = hittable.NewSphere(center, 0.2, materials.NewDielectric(1.5))
					i++
				}
			}
		}
	}

	list[i] = hittable.NewSphere(entities.MakeVec3(0, 1, 0), 1.0, materials.NewDielectric(1.5))
	i++
	list[i] = hittable.NewSphere(entities.MakeVec3(-4, 1, 0), 1.0, materials.NewLambertian(entities.MakeVec3(0.4, 0.2, 0.1)))
	i++
	list[i] = hittable.NewSphere(entities.MakeVec3(4, 1, 0), 1.0, materials.NewMetal(entities.MakeVec3(0.7, 0.6, 0.5), 0.0))
	i++

	return hittable.NewHittableList(list[:i])
}

func smallObjects() *hittable.HittableList {
	return hittable.NewHittableList([]hittable.Hittable{
		hittable.NewSphere(entities.MakeVec3(0.0, 0.0, -1.0), 0.5,
			materials.NewLambertian(entities.MakeVec3(0.1, 0.2, 0.5))),

		hittable.NewSphere(entities.MakeVec3(0.0, -100.5, -1.0), 100,
			materials.NewLambertian(entities.MakeVec3(0.8, 0.8, 0.0))),

		hittable.NewSphere(entities.MakeVec3(1.0, 0.0, -1.0), 0.5,
			materials.NewMetal(entities.MakeVec3(0.8, 0.6, 0.2), 0.1)),

		hittable.NewSphere(entities.MakeVec3(-1.0, 0.0, -1.0), 0.5,
			materials.NewDielectric(1.5)),

		hittable.NewSphere(entities.MakeVec3(-1.0, 0.0, -1.0), -0.45,
			materials.NewDielectric(1.5)),
	})
}

func DefaultScene() *scene.Scene {
	width := uint32(1000)
	height := uint32(500)

	lookfrom := entities.MakeVec3(13, 2, 3)
	lookat := entities.MakeVec3(0, 0, 0)

	return &scene.Scene{
		Camera: entities.NewCamera(
			lookfrom,
			lookat,
			entities.MakeVec3(0, 1, 0),
			20,
			float64(width)/float64(height),
			0.1, 10.0),
		RaysPerPixel: 50,
		HittableList: randomObjects(),
		Width:        width,
		Height:       height,
	}
}

func Min(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}
