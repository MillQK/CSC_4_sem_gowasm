package hitable

import (
	"encoding/json"
	"errors"
	"fmt"
)

const TYPE = "type"

func (list *HitableList) UnmarshalJSON(b []byte) error {
	var objMap map[string]*json.RawMessage
	err := json.Unmarshal(b, &objMap)
	if err != nil {
		return err
	}

	var rawMessagesForHitableList []*json.RawMessage
	err = json.Unmarshal(*objMap["List"], &rawMessagesForHitableList)
	if err != nil {
		return err
	}

	list.List = make([]Hitable, len(rawMessagesForHitableList))

	var m map[string]*json.RawMessage
	for index, rawMessage := range rawMessagesForHitableList {
		err = json.Unmarshal(*rawMessage, &m)
		if err != nil {
			return err
		}

		var hitableType string
		err = json.Unmarshal(*m[TYPE], &hitableType)
		if err != nil {
			return err
		}

		switch hitableType {
		case SPHERE_TYPE:
			var sphere Sphere
			err := json.Unmarshal(*m[hitableType], &sphere)
			if err != nil {
				return err
			}

			list.List[index] = &sphere
		default:
			return errors.New(fmt.Sprintf("Unsupported type found: %s", hitableType))
		}

	}

	return nil
}
