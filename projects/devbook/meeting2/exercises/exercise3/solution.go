package exercise3

import (
	"encoding/json"
	"net/http"
)

type Solution_Request struct {
	Values []int `json:"values"`
}

type Solution_Response struct {
	Result int `json:"result"`
}

func Solution_AddHandler(w http.ResponseWriter, r *http.Request) {
	var req Solution_Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var out Solution_Response
	for _, v := range req.Values {
		out.Result += v
	}

	if err := json.NewEncoder(w).Encode(out); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
