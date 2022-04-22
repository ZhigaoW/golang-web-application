package main

import "net/http"

type MyHandler struct{}

func (m *MyHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("MyHandler Test!"))
}

func main() {
	mh := MyHandler{}
	http.ListenAndServe("localhost:9999", &mh)

}
