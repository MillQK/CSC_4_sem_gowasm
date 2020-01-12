# gowasm_raytracer
Go + WASM distributed raytracer

### Layout

    .
    ├── raytracer               # All raytracer entities and algorithms
    ├── scene                   # Scene description
    ├── web                     # Client and server
    │   ├── client              # GO WASM client and page server
    │   ├── server              # Raytracing task server
    │   └── shared              # Entities shared between server and client
    └── main.go                 # Local raytracing

