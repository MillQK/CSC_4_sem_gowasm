package main

import (
	rt "CSC_4_sem_gowasm/raytracer"
	"CSC_4_sem_gowasm/web/shared"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"syscall/js"
	"time"
)

var document = js.Global().Get("document")

func elementById(value js.Value, id string) js.Value {
	return value.Call("getElementById", id)
}

func addElement(value js.Value, tag string) js.Value {
	return value.Call("createElement", tag)
}

func appendChild(parent js.Value, child js.Value) {
	parent.Call("appendChild", child)
}

func getUrl() (string, error) {
	userUrl := strings.TrimSpace(elementById(document, "server_url").Get("value").String())

	if userUrl == "" {
		return fmt.Sprintf("%s:%s", shared.DEFAULT_ADDRESS, shared.DEFAULT_PORT), nil
	}

	if _, err := url.ParseRequestURI(userUrl); err != nil {
		return "", nil
	}

	return strings.TrimRight(userUrl, "/"), nil
}

func log(msg string) {
	println(msg)
	elementById(document, "info").Set("textContent", msg)
}

func draw(job *shared.RayTracingJob) {
	canvas := elementById(document, "imageCanvas")
	canvas.Set("width", job.WidthPixels.To-job.WidthPixels.From)
	canvas.Set("height", job.HeightPixels.To-job.HeightPixels.From)

	context := canvas.Call("getContext", "2d")
	imgData := context.Call("createImageData", job.WidthPixels.To-job.WidthPixels.From, job.HeightPixels.To-job.HeightPixels.From)
	imgDataData := imgData.Get("data")

	startTime := time.Now()

	var once sync.Once

	for j := job.HeightPixels.From; j < job.HeightPixels.To; j++ {
		for i := job.WidthPixels.From; i < job.WidthPixels.To; i++ {

			pixelColor := rt.RenderPixel(&job.Scene, i, j)
			imgIndex := 4 * int((job.HeightPixels.To-j-1)*(job.WidthPixels.To-job.WidthPixels.From)+(i-job.WidthPixels.From))

			imgDataData.SetIndex(imgIndex, uint8(pixelColor.X))
			imgDataData.SetIndex(imgIndex+1, uint8(pixelColor.Y))
			imgDataData.SetIndex(imgIndex+2, uint8(pixelColor.Z))
			imgDataData.SetIndex(imgIndex+3, 255)
		}

		if float64(100*j)/float64(job.HeightPixels.To-job.HeightPixels.From) >= 26.0 {
			once.Do(shared.PrintMemUsage)
		}
	}
	shared.PrintMemUsage()
	elapsed := time.Now().Sub(startTime)
	println("Elapsed time ", elapsed.String())

	context.Call("putImageData", imgData, 0, 0)
}

func main() {
	serverUrl, err := getUrl()
	if err != nil {
		println(err.Error())
		return
	}
	println("Url: ", serverUrl)

	httpClient := http.DefaultClient
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", serverUrl, shared.GET_TASK), nil)
	if err != nil {
		println(err.Error())
		return
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		println(err.Error())
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err.Error())
		return
	}

	var rayTracingJob shared.RayTracingJob
	err = json.Unmarshal(data, &rayTracingJob)
	if err != nil {
		println(err.Error())
		return
	}

	println(fmt.Sprintf("%+v\n", rayTracingJob))
	draw(&rayTracingJob)
}
