package exercise2

import (
	"net/http"
	"strconv"
)

func Solution_AddHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	aStr, bStr := r.Form.Get("a"), r.Form.Get("b")
	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)

	w.Write([]byte(strconv.Itoa(a + b)))
}

func Solution_SubHandler(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	aStr, bStr := r.Form.Get("a"), r.Form.Get("b")
	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)

	w.Write([]byte(strconv.Itoa(a - b)))
}
