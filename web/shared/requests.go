package shared

import (
	"CSC_4_sem_gowasm/scene"
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
