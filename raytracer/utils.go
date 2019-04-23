package raytracer

import (
	"math"
	"math/rand"
)

// Copy from Rust UnitSphereSurface Distribution
// https://projecteuclid.org/download/pdf_1/euclid.aoms/1177692644
func PointOnUnitSphereSufrace() Vec3 {
	for {
		// [-1.0, 1.0)
		x1 := 2.0*rand.Float64() - 1.0
		x2 := 2.0*rand.Float64() - 1.0
		sum := x1*x1 + x2*x2

		if sum >= 1.0 {
			continue
		}

		factor := 2.0 * math.Sqrt(1.0-sum)
		return Vec3{x1 * factor, x2 * factor, 1.0 - 2.0*sum}
	}
}
