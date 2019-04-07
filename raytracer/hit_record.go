package raytracer

type HitRecord struct {
	T             float64
	Point, Normal Vec3
}

func MakeHitRecord(t float64, point, normal Vec3) HitRecord {
	return HitRecord{t, point, normal}
}
