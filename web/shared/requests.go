package shared

import (
	"github.com/MillQK/gowasm_raytracer/raytracer/entities"
	"github.com/MillQK/gowasm_raytracer/scene"
)

type PixelsRange struct {
	From, To uint32
}

type RayTracingJob struct {
	SceneId      string
	Scene        scene.Scene
	WidthPixels  PixelsRange
	HeightPixels PixelsRange
}

type RayTracingJobResult struct {
	SceneId      string
	WidthPixels  PixelsRange
	HeightPixels PixelsRange
	Image        entities.Image
}
