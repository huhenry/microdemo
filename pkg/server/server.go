package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Start(port int) error {
	r := mux.NewRouter()
	r.HandleFunc("/{status_code}", returnStatusCode)
	r.HandleFunc("/", index)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

func index(w http.ResponseWriter, r *http.Request) {
	body := `
These web site returns HTTP status code.
Please access to /200, /404, and more.
`
	w.WriteHeader(200)
	w.Write([]byte(body))
}

func returnStatusCode(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	statusCode, ok := vars["status_code"]
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte(`Invalid path parameter`))
		return
	}
	statusCodeInt, err := strconv.Atoi(statusCode)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(`Invalid path parameter`))
		return
	}
	switch statusCodeInt {
	// Informational
	case 100, 101, 102:
		w.WriteHeader(statusCodeInt)
	// Success
	case 200, 201, 202, 203, 204, 205, 207, 208, 226:
		w.WriteHeader(statusCodeInt)
	// Redirection
	case 300, 301, 302, 303, 304, 305, 307, 308:
		w.WriteHeader(statusCodeInt)
	// Client Error
	case 400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418, 421, 422, 423, 424, 426, 428, 429, 431, 444, 451, 499:
		w.WriteHeader(statusCodeInt)
	// Server Error
	case 500, 501, 502, 503, 504, 505, 506, 507, 508, 510, 511, 599:
		w.WriteHeader(statusCodeInt)
	default:
		w.WriteHeader(400)
		w.Write([]byte(`Invalid path parameter`))
	}
}
