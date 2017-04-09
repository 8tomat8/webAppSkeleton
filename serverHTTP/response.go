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
	logger := h.env.Log.WithField("type", "send")
	w.WriteHeader(resp.Status)

	if resp.Error == "" && resp.Status >= 300 {
		resp.Error = http.StatusText(resp.Status)
	}

	enc := json.NewEncoder(w)
	err := enc.Encode(resp)
	if err != nil {
		logger.Error(err)
	} else {
		logger.Debugf("Sent: %+v", resp)
	}
}
