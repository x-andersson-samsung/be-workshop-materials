package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"time"
)

func getNumFileDesciptors() (int, error) {
	pid := os.Getpid()
	fds, err := os.Open(fmt.Sprintf("/proc/%d/fd", pid))

	if err != nil {
		return 0, err
	}
	defer fds.Close()

	files, err := fds.Readdirnames(-1)
	if err != nil {
		return 0, err
	}

	return len(files), nil
}

func sendRequest() {

	payload := `{"key": "value"}`
	req, err := http.Post(
		"https://echo.free.beeceptor.com/sample-request",
		"application/json",
		strings.NewReader(payload),
	)
	if err != nil {
		panic(err)
	}

	_, err = io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	//req.Body.Close()

	//fmt.Println(string(body))
}

func main() {
	r, err := http.NewRequest()

	r.Header.Set("Content-Type", "application/json")

	http.DefaultClient.Do()

	go func() {

		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	fd, _ := getNumFileDesciptors()
	fmt.Printf("File descriptors: %d\n", fd)

	time.Sleep(time.Second * 10)

	sendRequest()
	sendRequest()
	sendRequest()
	sendRequest()
	sendRequest()

	fd, _ = getNumFileDesciptors()
	fmt.Printf("File descriptors 2: %d\n", fd)

	time.Sleep(time.Second * 10)

}
