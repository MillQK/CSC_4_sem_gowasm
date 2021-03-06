package entities

import "math"

type Vec3 struct {
	X, Y, Z float64
}

func MakeVec3(x, y, z float64) Vec3 {
	return Vec3{x, y, z}
}

func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{x, y, z}
}

func NewZeroVec3() *Vec3 {
	return &Vec3{0, 0, 0}
}

func (vec Vec3) Squared_length() float64 {
	x, y, z := vec.X, vec.Y, vec.Z
	return x*x + y*y + z*z
}

func (vec Vec3) Length() float64 {
	return math.Sqrt(vec.Squared_length())
}

func (vec Vec3) Dot(other Vec3) float64 {
	x1, y1, z1 := vec.X, vec.Y, vec.Z
	x2, y2, z2 := other.X, other.Y, other.Z
	return x1*x2 + y1*y2 + z1*z2
}

func (vec Vec3) Cross(other Vec3) Vec3 {
	x1, y1, z1 := vec.X, vec.Y, vec.Z
	x2, y2, z2 := other.X, other.Y, other.Z

	x := y1*z2 - z1*y2
	y := z1*x2 - x1*z2
	z := x1*y2 - y1*x2

	return Vec3{x, y, z}
}

func (vec Vec3) UnitVector() Vec3 {
	return vec.DivScalar(vec.Length())
}

func (vec *Vec3) ModifyFields(f func(float64) float64) {
	vec.X = f(vec.X)
	vec.Y = f(vec.Y)
	vec.Z = f(vec.Z)
}

func (vec Vec3) Reflect(other Vec3) Vec3 {
	return vec.Sub(other.MulScalar(2 * vec.Dot(other)))
}

func (vec Vec3) Refract(other Vec3, niOverNt float64) *Vec3 {
	unitVec := vec.UnitVector()
	dt := unitVec.Dot(other)
	discriminant := 1.0 - niOverNt*niOverNt*(1-dt*dt)
	if discriminant > 0.0 {
		nDiscr := other.MulScalar(math.Sqrt(discriminant))
		nDt := other.MulScalar(dt)

		return NewZeroVec3().AddAssign(unitVec).SubAssign(nDt).MulScalarAssign(niOverNt).SubAssign(nDiscr)
	}

	return nil
}

func (vec Vec3) Neg() Vec3 {
	return Vec3{-vec.X, -vec.Y, -vec.Z}
}

func (vec *Vec3) NegSelf() {
	vec.X = -vec.X
	vec.Y = -vec.Y
	vec.Z = -vec.Z
}

func (vec Vec3) Add(other Vec3) Vec3 {
	x1, y1, z1 := vec.X, vec.Y, vec.Z
	x2, y2, z2 := other.X, other.Y, other.Z
	return Vec3{x1 + x2, y1 + y2, z1 + z2}
}

func (vec *Vec3) AddAssign(other Vec3) *Vec3 {
	vec.X += other.X
	vec.Y += other.Y
	vec.Z += other.Z
	return vec
}

func (vec Vec3) AddScalar(scalar float64) Vec3 {
	return Vec3{vec.X + scalar, vec.Y + scalar, vec.Z + scalar}
}

func (vec *Vec3) AddScalarAssign(scalar float64) *Vec3 {
	vec.X += scalar
	vec.Y += scalar
	vec.Z += scalar
	return vec
}

func (vec Vec3) Sub(other Vec3) Vec3 {
	x1, y1, z1 := vec.X, vec.Y, vec.Z
	x2, y2, z2 := other.X, other.Y, other.Z
	return Vec3{x1 - x2, y1 - y2, z1 - z2}
}

func (vec *Vec3) SubAssign(other Vec3) *Vec3 {
	vec.X -= other.X
	vec.Y -= other.Y
	vec.Z -= other.Z
	return vec
}

func (vec *Vec3) SubScalarAssign(scalar float64) *Vec3 {
	vec.X -= scalar
	vec.Y -= scalar
	vec.Z -= scalar
	return vec
}

func (vec Vec3) Mul(other Vec3) Vec3 {
	x1, y1, z1 := vec.X, vec.Y, vec.Z
	x2, y2, z2 := other.X, other.Y, other.Z
	return Vec3{x1 * x2, y1 * y2, z1 * z2}
}

func (vec *Vec3) MulAssign(other Vec3) *Vec3 {
	vec.X *= other.X
	vec.Y *= other.Y
	vec.Z *= other.Z
	return vec
}

func (vec Vec3) MulScalar(scalar float64) Vec3 {
	return Vec3{vec.X * scalar, vec.Y * scalar, vec.Z * scalar}
}

func (vec *Vec3) MulScalarAssign(scalar float64) *Vec3 {
	vec.X *= scalar
	vec.Y *= scalar
	vec.Z *= scalar
	return vec
}

func (vec Vec3) Div(other Vec3) Vec3 {
	x1, y1, z1 := vec.X, vec.Y, vec.Z
	x2, y2, z2 := other.X, other.Y, other.Z
	return Vec3{x1 / x2, y1 / y2, z1 / z2}
}

func (vec *Vec3) DivAssign(other Vec3) *Vec3 {
	vec.X /= other.X
	vec.Y /= other.Y
	vec.Z /= other.Z
	return vec
}

func (vec Vec3) DivScalar(scalar float64) Vec3 {
	return Vec3{vec.X / scalar, vec.Y / scalar, vec.Z / scalar}
}

func (vec *Vec3) DivScalarAssign(scalar float64) *Vec3 {
	vec.X /= scalar
	vec.Y /= scalar
	vec.Z /= scalar
	return vec
}
