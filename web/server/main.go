package main

import (
	"CSC_4_sem_gowasm/web/shared"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var job = *defaultJob()

func defaultJob() *shared.RayTracingJob {
	defaultScene := *shared.DefaultScene()

	return &shared.RayTracingJob{
		SceneId:      "sceneId",
		Scene:        defaultScene,
		WidthPixels:  shared.PixelsRange{From: 0, To: defaultScene.Width},
		HeightPixels: shared.PixelsRange{From: 0, To: defaultScene.Height},
	}
}

func getTask(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	println(fmt.Sprintf("New connection from %s", r.RemoteAddr))

	b, err := json.Marshal(job)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	header := w.Header()
	header.Set("Content-Type", "application/json")
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Methods", "GET")
	header.Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type")

	w.Write(b)
}

func main() {
	http.HandleFunc(shared.GET_TASK, getTask)
	log.Fatal(http.ListenAndServe(":"+shared.DEFAULT_PORT, nil))
}
