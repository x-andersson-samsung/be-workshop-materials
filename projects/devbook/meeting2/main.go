package main

import (
	"net/http"
)

// Handling function for std must have the signature:
// func(http.ResponseWriter, *http.Request)
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
	r.PathValue("id")
}

func main() {
	srv := &http.Server{
		Addr:                         "",
		Handler:                      nil,
		DisableGeneralOptionsHandler: false,
		TLSConfig:                    nil,
		ReadTimeout:                  0,
		ReadHeaderTimeout:            0,
		WriteTimeout:                 0,
		IdleTimeout:                  0,
		MaxHeaderBytes:               0,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
		HTTP2:                        nil,
		Protocols:                    nil,
	}

	srv.ListenAndServe()
	srv.ListenAndServeTLS("cert.pem", "key.pem")
}
