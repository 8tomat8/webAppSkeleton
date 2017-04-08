package serverHTTP

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Msg    json.RawMessage
	Error  string
	Status int `json:"-"`
}

func (h *Handlers) Send(w http.ResponseWriter, resp *Response) {
	w.WriteHeader(resp.Status)

	enc := json.NewEncoder(w)
	err := enc.Encode(resp)
	if err != nil {
		h.env.Log.Error(err)
	} else {
		h.env.Log.Debugf("%+v", resp)
	}
}
