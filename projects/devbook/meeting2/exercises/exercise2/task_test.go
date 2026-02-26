package exercise2

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddHandler(t *testing.T) {
	type testCase struct {
		a, b int
		want int
	}
	cases := []testCase{
		{0, 0, 0},
		{1, 1, 2},
		{1, 2, 3},
		{2, 1, 3},
		{2, 3, 5},
		{10, 10, 20},
		{10, -10, 0},
		{-10, 10, 0},
	}
	for _, tC := range cases {
		t.Run(fmt.Sprintf("AddHandler(%d,%d)", tC.a, tC.b), func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", fmt.Sprintf("/add?a=%d&b=%d", tC.a, tC.b), nil)
			AddHandler(recorder, request)
			if recorder.Code != 200 {
				t.Errorf("status code = %d, want %d", recorder.Code, http.StatusOK)
			}
			if recorder.Body.String() != fmt.Sprintf("%d", tC.want) {
				t.Errorf("body = %s, want %d", recorder.Body.String(), tC.want)
			}
		})
	}
}

func TestSubHandler(t *testing.T) {
	type testCase struct {
		a, b int
		want int
	}
	cases := []testCase{
		{0, 0, 0},
		{1, 1, 0},
		{1, 2, -1},
		{2, 1, 1},
		{2, 3, -1},
		{10, 10, 0},
		{10, -10, 20},
		{-10, 10, -20},
	}
	for _, tC := range cases {
		t.Run(fmt.Sprintf("SubHandler(%d,%d)", tC.a, tC.b), func(t *testing.T) {
			recorder := httptest.NewRecorder()
			request := httptest.NewRequest("GET", fmt.Sprintf("/sub?a=%d&b=%d", tC.a, tC.b), nil)
			SubHandler(recorder, request)
			if recorder.Code != 200 {
				t.Errorf("status code = %d, want %d", recorder.Code, http.StatusOK)
			}
			if recorder.Body.String() != fmt.Sprintf("%d", tC.want) {
				t.Errorf("body = %s, want %d", recorder.Body.String(), tC.want)
			}
		})
	}
}
