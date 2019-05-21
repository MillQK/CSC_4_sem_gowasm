package materials

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	MATERIAL_TYPE = "material_type"
	LAMBERTIAN    = "lambertian"
	METAL         = "metal"
)

func (lambertian *Lambertian) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m[MATERIAL_TYPE] = LAMBERTIAN
	m[LAMBERTIAN] = *lambertian
	return json.Marshal(m)
}

func (metal *Metal) MarshalJSON() ([]byte, error) {
	m := make(map[string]interface{})
	m[MATERIAL_TYPE] = METAL
	m[METAL] = *metal
	return json.Marshal(m)
}

func UnmarshalMaterial(b []byte) (Material, error) {
	var m map[string]*json.RawMessage
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}

	var materialType string
	err = json.Unmarshal(*m[MATERIAL_TYPE], &materialType)
	if err != nil {
		return nil, err
	}

	switch materialType {
	case LAMBERTIAN:
		var lambertian Lambertian
		err := json.Unmarshal(*m[LAMBERTIAN], &lambertian)
		if err != nil {
			return nil, err
		}

		return &lambertian, nil

	case METAL:
		var metal Metal
		err := json.Unmarshal(*m[METAL], &metal)
		if err != nil {
			return nil, err
		}

		return &metal, nil

	default:
		return nil, errors.New(fmt.Sprintf("Unsupported material type found: %s", materialType))
	}
}
