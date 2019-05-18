package entities

type Camera struct {
	Origin, LowerLeftCorner, Horizontal, Vertical Vec3
}

func MakeCamera() Camera {
	return Camera{
		MakeVec3(0.0, 0.0, 0.0),
		MakeVec3(-2.0, -1.0, -1.0),
		MakeVec3(4.0, 0.0, 0.0),
		MakeVec3(0.0, 2.0, 0.0),
	}
}

func (camera *Camera) GetRay(u, v float64) Ray {
	return MakeRay(camera.Origin, camera.LowerLeftCorner.Sub(camera.Origin).Add(camera.Horizontal.MulScalar(u)).Add(camera.Vertical.MulScalar(v)))
}
