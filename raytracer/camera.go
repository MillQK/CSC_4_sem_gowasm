package raytracer

type Camera struct {
	origin, lowerLeftCorner, horizontal, vertical Vec3
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
	return MakeRay(camera.origin, camera.lowerLeftCorner.Sub(camera.origin).Add(camera.horizontal.MulScalar(u)).Add(camera.vertical.MulScalar(v)))
}
