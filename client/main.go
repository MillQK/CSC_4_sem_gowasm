package client

import "syscall/js"

func main() {
	println("Hello world!")
	js.Global().Get("document").Call("getElementById", "runButton").Set("textContent", "Hello from go + wasm!")
}
