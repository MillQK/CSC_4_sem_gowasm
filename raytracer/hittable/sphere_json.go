package hittable

import (
	"encoding/json"
	"github.com/MillQK/gowasm_raytracer/raytracer/hittable/materials"
)

const SPHERE = "sphere"

func (sphere *Sphere) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m[HITTABLE_TYPE] = SPHERE
	m[SPHERE] = *sphere
	return json.Marshal(m)
}

func (sphere *Sphere) UnmarshalJSON(b []byte) error {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*m["Center"], &sphere.Center)
	if err != nil {
		return err
	}

	err = json.Unmarshal(*m["Radius"], &sphere.Radius)
	if err != nil {
		return err
	}

	material, err := materials.UnmarshalMaterial(*m["Material"])
	if err != nil {
		return err
	}

	sphere.Material = material

	return nil
}
