package apis

import (
	"encoding/json"
	"net/http"
)

type RequestHandler func(req *http.Request) *Response

func read(req *http.Request, body interface{}) error {
	if e := json.NewDecoder(req.Body).Decode(&body); e != nil {
		return e
	}
	return nil
}
