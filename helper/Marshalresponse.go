package helper

import (
	"encoding/json"
)

func MarshalResponse(res interface{}) []byte {

	resp, er := json.Marshal(res)

	if er != nil {
		return []byte(er.Error())
	}
	return resp
}
