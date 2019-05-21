package shared

import (
	"CSC_4_sem_gowasm/raytracer/entities"
	"CSC_4_sem_gowasm/raytracer/hittable"
	"CSC_4_sem_gowasm/raytracer/hittable/materials"
	"CSC_4_sem_gowasm/scene"
	"fmt"
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

func DefaultScene() *scene.Scene {
	return &scene.Scene{
		Camera:       entities.MakeCamera(),
		RaysPerPixel: 100,
		HittableList: *hittable.NewHittableList([]hittable.Hittable{
			hittable.NewSphere(entities.MakeVec3(0.0, 0.0, -1.0), 0.5,
				materials.NewLambertian(entities.MakeVec3(0.8, 0.3, 0.3))),
			hittable.NewSphere(entities.MakeVec3(0.0, -100.5, -1.0), 100,
				materials.NewLambertian(entities.MakeVec3(0.8, 0.8, 0.0))),
			hittable.NewSphere(entities.MakeVec3(1.0, 0.0, -1.0), 0.5,
				materials.NewMetal(entities.MakeVec3(0.8, 0.6, 0.2), 1.0)),
			hittable.NewSphere(entities.MakeVec3(-1.0, 0.0, -1.0), 0.5,
				materials.NewMetal(entities.MakeVec3(0.8, 0.8, 0.8), 0.3)),
		}),
		Width:  1000,
		Height: 500,
	}
}

func Min(a, b uint32) uint32 {
	if a < b {
		return a
	}
	return b
}
