package main

import (
	rt "CSC_4_sem_gowasm/raytracer"
	"CSC_4_sem_gowasm/raytracer/entities"
	"CSC_4_sem_gowasm/web/shared"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"syscall/js"
	"time"
)

var document = js.Global().Get("document")

// TODO: make struct for this global variables
var exitChannel = make(chan struct{})
var callbacks = make([]js.Func, 0, 8)

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
	elementById(document, "status").Set("innerText", msg)
}

func showImage(image *entities.Image) {
	canvas := elementById(document, "imageCanvas")
	canvas.Set("width", image.Width)
	canvas.Set("height", image.Height)

	context := canvas.Call("getContext", "2d")
	imgData := context.Call("createImageData", image.Width, image.Height)
	imgDataData := imgData.Get("data")

	startTime := time.Now()
	for j := uint32(0); j < image.Height; j++ {
		for i := uint32(0); i < image.Width; i++ {

			pixelColor := image.GetPixel(i, j)
			imgIndex := 4 * int(j*image.Width+i)

			imgDataData.SetIndex(imgIndex, pixelColor.R)
			imgDataData.SetIndex(imgIndex+1, pixelColor.G)
			imgDataData.SetIndex(imgIndex+2, pixelColor.B)
			imgDataData.SetIndex(imgIndex+3, 255)
		}
	}
	elapsed := time.Now().Sub(startTime)
	println("Elapsed showImage time:", elapsed.String())

	context.Call("putImageData", imgData, 0, 0)
}

func draw(job *shared.RayTracingJob) *entities.Image {
	startTime := time.Now()
	image := entities.NewImage(job.WidthPixels.To-job.WidthPixels.From, job.HeightPixels.To-job.HeightPixels.From)

	for j := job.HeightPixels.From; j < job.HeightPixels.To; j++ {
		for i := job.WidthPixels.From; i < job.WidthPixels.To; i++ {

			pixelColor := rt.RenderPixel(&job.Scene, i, job.Scene.Height-j-1)
			image.GetPixel(i-job.WidthPixels.From, j-job.HeightPixels.From).FromVec(pixelColor)
		}
	}
	shared.PrintMemUsage()
	elapsed := time.Now().Sub(startTime)
	println("Elapsed draw time:", elapsed.String())

	showImage(image)
	return image
}

func sendImageToServer(job *shared.RayTracingJob, image *entities.Image) error {
	serverUrl, err := getUrl()
	if err != nil {
		println(err.Error())
		return err
	}

	jobResult := shared.RayTracingJobResult{
		SceneId:      job.SceneId,
		WidthPixels:  job.WidthPixels,
		HeightPixels: job.HeightPixels,
		Image:        *image,
	}

	jobResultJson, _ := json.Marshal(jobResult)
	jobResultJsonReader := bytes.NewReader(jobResultJson)

	httpClient := http.DefaultClient
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", serverUrl, shared.TASK_RESULT), jobResultJsonReader)
	if err != nil {
		println(err.Error())
		return err
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := httpClient.Do(request)
	if err != nil {
		println(err.Error())
		return err
	}
	defer resp.Body.Close()
	return nil
}

func setupCallbacks() {
	currentImageButtonCb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		log("Loading image...")
		go func() {
			serverUrl, err := getUrl()
			if err != nil {
				println(err.Error())
				return
			}
			println("Url: ", serverUrl)

			httpClient := http.DefaultClient
			request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", serverUrl, shared.SCENE_IMAGE), nil)
			if err != nil {
				println(err.Error())
				return
			}

			resp, err := httpClient.Do(request)
			if err != nil {
				println(err.Error())
				return
			}

			defer resp.Body.Close()

			data, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				println(err.Error())
				return
			}

			var image entities.Image
			err = json.Unmarshal(data, &image)
			if err != nil {
				println(err.Error())
				return
			}

			showImage(&image)

			log("Loading finished")
		}()
		return nil
	})

	elementById(document, "currentImage").
		Call("addEventListener", "click", currentImageButtonCb)

	runOnceButtonCb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		log("Running...")
		go func() {
			getTask()
			log("Finished")
		}()

		return nil
	})

	elementById(document, "runOnceButton").
		Call("addEventListener", "click", runOnceButtonCb)

	runButtonCb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		log("Running...")

		go func() {
			for {
				err := getTask()
				if err != nil {
					break
				}
			}

			log("Finished")
		}()

		return nil
	})

	elementById(document, "runButton").
		Call("addEventListener", "click", runButtonCb)
}

func exit() {
	for _, callback := range callbacks {
		callback.Release()
	}
	exitChannel <- struct{}{}
}

func getTask() error {
	serverUrl, err := getUrl()
	if err != nil {
		println(err.Error())
		return err
	}
	println("Url: ", serverUrl)

	httpClient := http.DefaultClient
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", serverUrl, shared.GET_TASK), nil)
	if err != nil {
		println(err.Error())
		return err
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		println(err.Error())
		return err
	}

	if resp.StatusCode == http.StatusNoContent {
		println("Jobs is over")
		return fmt.Errorf("Jobs is over")
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println(err.Error())
		return err
	}

	println(string(data))

	var rayTracingJob shared.RayTracingJob
	err = json.Unmarshal(data, &rayTracingJob)
	if err != nil {
		println(err.Error())
		return err
	}

	image := draw(&rayTracingJob)
	err = sendImageToServer(&rayTracingJob, image)
	if err != nil {
		println(err.Error())
		return err
	}

	return nil
}

func main() {
	setupCallbacks()
	<-exitChannel
}
