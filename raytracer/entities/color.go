package entities

type Color struct {
	R, G, B uint8
}

func MakeColor(r, g, b uint8) Color {
	return Color{r, g, b}
}

func (color *Color) FromVec(vec3 Vec3) {
	color.R = uint8(vec3.X)
	color.G = uint8(vec3.Y)
	color.B = uint8(vec3.Z)
}
