package exercise3

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddHandler(t *testing.T) {
	type requestStr struct {
		Values []int `json:"values"`
	}
	type responseStr struct {
		Result int `json:"result"`
	}
	type testCase struct {
		values []int
		want   int
	}
	cases := []testCase{
		{[]int{}, 0},
		{[]int{0, 0, 0}, 0},
		{[]int{1, 1}, 2},
		{[]int{1, -1}, 0},
		{[]int{123, 123}, 246},
	}
	for _, tC := range cases {
		t.Run(fmt.Sprintf("AddHandler(%v)", tC.values), func(t *testing.T) {
			recorder := httptest.NewRecorder()

			payload, _ := json.Marshal(requestStr{tC.values})
			request := httptest.NewRequest("GET", "/", bytes.NewReader(payload))
			AddHandler(recorder, request)
			if recorder.Code != 200 {
				t.Errorf("status code = %d, want %d", recorder.Code, http.StatusOK)
			}

			var got responseStr
			_ = json.Unmarshal(recorder.Body.Bytes(), &got)

			if got.Result != tC.want {
				t.Errorf("got = %d, want %d", got.Result, tC.want)
			}
		})
	}
}
