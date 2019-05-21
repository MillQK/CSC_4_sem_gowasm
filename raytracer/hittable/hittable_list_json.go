package hittable

import (
	"encoding/json"
	"errors"
	"fmt"
)

const HITTABLE_TYPE = "hittable_type"

func (list *HittableList) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	var rawMessagesForHittableList []*json.RawMessage
	err = json.Unmarshal(*objMap["List"], &rawMessagesForHittableList)
	if err != nil {
		return err
	}

	list.List = make([]Hittable, len(rawMessagesForHittableList))

	var m map[string]*json.RawMessage
	for index, rawMessage := range rawMessagesForHittableList {
		err = json.Unmarshal(*rawMessage, &m)
		if err != nil {
			return err
		}

		var hittableType string
		err = json.Unmarshal(*m[HITTABLE_TYPE], &hittableType)
		if err != nil {
			return err
		}

		switch hittableType {
		case SPHERE:
			var sphere Sphere
			err := json.Unmarshal(*m[SPHERE], &sphere)
			if err != nil {
				return err
			}

			list.List[index] = &sphere
		default:
			return errors.New(fmt.Sprintf("Unsupported hittable type found: %s", hittableType))
		}

	}

	return nil
}
