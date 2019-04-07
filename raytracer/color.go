package raytracer

type Color struct {
	R, G, B uint8
}

func MakeColor(r, g, b uint8) Color {
	return Color{r, g, b}
}
