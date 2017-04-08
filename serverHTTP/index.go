package serverHTTP

import (
	"net/http"
	"encoding/json"
)

type helloMsg struct {
	Question string
	Answer   int
}

func (h *Handlers) Index(w http.ResponseWriter, r *http.Request) {
	resp := Response{Status:http.StatusOK}
	msg, err := json.Marshal(helloMsg{
		"Answer to the Ultimate Question of Life, the Universe, and Everything",
		42,
	})
	if h.env.Check(err) {
		resp.Error = err.Error()
		resp.Status = http.StatusInternalServerError
		h.Send(w, &resp)
		return
	}

	resp.Msg = msg
	h.Send(w, &resp)
}
