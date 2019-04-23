package hitable

import "encoding/json"

const SPHERE_TYPE = "sphere"

func (sphere *Sphere) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m[TYPE] = SPHERE_TYPE
	m[SPHERE_TYPE] = *sphere
	return json.Marshal(m)
}
