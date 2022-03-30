package httpOperations

import (
	"fmt"
	"net/http"
)

func call(res http.ResponseWriter, req *http.Request) {

	name := req.URL.Query().Get("name")
	res.Write([]byte(fmt.Sprintf("Hello: %s", name)))
}

func serve() {
	http.HandleFunc("/hello", call)
	http.ListenAndServe(":5000", nil)
}
