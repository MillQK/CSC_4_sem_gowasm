package entities

import (
	"math"
	"math/rand"
)

type Camera struct {
	Origin, LowerLeftCorner, Horizontal, Vertical Vec3
	W, U, V                                       Vec3
	LensRadius                                    float64
}

func pointOnUnitDiskSurface() Vec3 {
	for {
		// [-1.0, 1.0)
		x1 := 2.0*rand.Float64() - 1.0
		x2 := 2.0*rand.Float64() - 1.0
		sum := x1*x1 + x2*x2

		if sum >= 1.0 {
			continue
		}

		return Vec3{x1, x2, 0}
	}
}

// vfov is top to bottom in degrees
func NewCamera(lookfrom, lookat, vup Vec3, vfov, aspect, aperture, focusDistance float64) *Camera {
	lensRadius := aperture / 2
	theta := vfov * math.Pi / 180
	halfHeight := math.Tan(theta / 2)
	halfWidth := aspect * halfHeight

	w := lookfrom.Sub(lookat).UnitVector()
	u := vup.Cross(w).UnitVector()
	v := w.Cross(u)

	origin := lookfrom
	lowerLeftCorner := origin.Sub(u.MulScalar(halfWidth * focusDistance)).Sub(v.MulScalar(halfHeight * focusDistance)).Sub(w.MulScalar(focusDistance))
	horizontal := u.MulScalar(2 * halfWidth * focusDistance)
	vertical := v.MulScalar(2 * halfHeight * focusDistance)

	return &Camera{
		origin, lowerLeftCorner, horizontal, vertical,
		w, u, v, lensRadius,
	}
}

func (camera *Camera) GetRay(s, t float64) Ray {
	rd := pointOnUnitDiskSurface().MulScalar(camera.LensRadius)
	offset := camera.U.MulScalar(rd.X).Add(camera.V.MulScalar(rd.Y))

	rayOrigin := camera.Origin.Add(offset)
	rayDirection := camera.LowerLeftCorner.Add(camera.Horizontal.MulScalar(s)).Add(camera.Vertical.MulScalar(t)).Sub(camera.Origin).Sub(offset)

	return MakeRay(rayOrigin, rayDirection)
}
