package main

import (
	"encoding/json"
	"fmt"
	"github.com/MillQK/gowasm_raytracer/raytracer/entities"
	"github.com/MillQK/gowasm_raytracer/web/shared"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type PixelsRangeStatus struct {
	WidthPixels, HeightPixels shared.PixelsRange
	Mutex                     sync.Mutex
	inProcess, finished       bool
}

var scene = shared.DefaultScene()
var image = entities.NewImage(scene.Width, scene.Height)
var ranges = setupRanges()

func setupRanges() []PixelsRangeStatus {
	rangeSize := uint32(250)
	ranges := make([]PixelsRangeStatus, 0, 16)

	for i := uint32(0); i < scene.Width; i += rangeSize {
		widthRange := shared.PixelsRange{From: i, To: shared.Min(i+rangeSize, scene.Width)}
		for j := uint32(0); j < scene.Height; j += rangeSize {
			heightRange := shared.PixelsRange{From: j, To: shared.Min(j+rangeSize, scene.Height)}

			rangeStatus := PixelsRangeStatus{
				WidthPixels:  widthRange,
				HeightPixels: heightRange,
			}

			ranges = append(ranges, rangeStatus)
		}
	}

	return ranges
}

func getTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Methods", "GET")
	header.Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusOK)
		return
	}

	println(fmt.Sprintf("New getTask connection from %s", r.RemoteAddr))

	var jobRange *PixelsRangeStatus = nil

	for i := 0; i < len(ranges); i++ {
		pixelsRange := &ranges[i]
		if !pixelsRange.inProcess {
			pixelsRange.Mutex.Lock()
			if !pixelsRange.inProcess {
				pixelsRange.inProcess = true

				jobRange = pixelsRange
				pixelsRange.Mutex.Unlock()
				break
			}
			pixelsRange.Mutex.Unlock()
		}
	}

	if jobRange == nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	job := shared.RayTracingJob{
		SceneId:      "sceneId",
		Scene:        *scene,
		WidthPixels:  jobRange.WidthPixels,
		HeightPixels: jobRange.HeightPixels,
	}

	b, err := json.Marshal(job)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.Write(b)
}

func taskResult(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Methods", "POST")
	header.Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusOK)
		return
	}

	println(fmt.Sprintf("New taskResult connection from %s", r.RemoteAddr))

	resultBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	var rayTracingJobResult shared.RayTracingJobResult
	err = json.Unmarshal(resultBytes, &rayTracingJobResult)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	go func() {
		for i := 0; i < len(ranges); i++ {
			pixelsRange := &ranges[i]

			if pixelsRange.WidthPixels == rayTracingJobResult.WidthPixels && pixelsRange.HeightPixels == rayTracingJobResult.HeightPixels {
				pixelsRange.Mutex.Lock()
				pixelsRange.finished = true

				for j := pixelsRange.HeightPixels.From; j < pixelsRange.HeightPixels.To; j++ {
					for i := pixelsRange.WidthPixels.From; i < pixelsRange.WidthPixels.To; i++ {
						image.GetPixel(i, j).FromColor(*rayTracingJobResult.Image.GetPixel(i-pixelsRange.WidthPixels.From, j-pixelsRange.HeightPixels.From))
					}
				}

				pixelsRange.Mutex.Unlock()
				break
			}
		}
	}()

	w.WriteHeader(http.StatusOK)
}

func sceneImage(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Methods", "GET")
	header.Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusOK)
		return
	}

	println(fmt.Sprintf("New sceneImage connection from %s", r.RemoteAddr))

	for i := 0; i < len(ranges); i++ {
		pixelsRange := &ranges[i]

		pixelsRange.Mutex.Lock()
	}

	imageJson, err := json.Marshal(image)

	for i := 0; i < len(ranges); i++ {
		pixelsRange := &ranges[i]

		pixelsRange.Mutex.Unlock()
	}

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	w.Write(imageJson)
}

func main() {
	setupRanges()

	http.HandleFunc(shared.GET_TASK, getTask)
	http.HandleFunc(shared.TASK_RESULT, taskResult)
	http.HandleFunc(shared.SCENE_IMAGE, sceneImage)
	log.Fatal(http.ListenAndServe(":"+shared.DEFAULT_PORT, nil))
}
