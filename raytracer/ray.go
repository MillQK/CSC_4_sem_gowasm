package raytracer

type Ray struct {
	Origin, Direction Vec3
}

func MakeRay(origin, direction Vec3) Ray {
	return Ray{origin, direction}
}

func (ray Ray) PointAtParameter(scalar float64) Vec3 {
	return ray.Direction.MulScalar(scalar).Add(ray.Origin)
}
