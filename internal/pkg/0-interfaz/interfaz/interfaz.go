package interfaz

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Web struct {
}

func (this *Web) Serve() {
	http.HandleFunc("/", serves)

	http.ListenAndServe("localhost:8080", nil)
}

func serves(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		file, err := ioutil.ReadFile("./web/index.html")
		if err != nil {
			notFound(w, r)
			return
		}
		fmt.Fprintf(w, string(file))
		return
	}

	file, err := ioutil.ReadFile("./web" + r.URL.Path)
	if err != nil {
		notFound(w, r)
		return
	}
	fmt.Fprintf(w, string(file))
}

func notFound(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusNotFound)

	file, err := ioutil.ReadFile("./web/notfound.html")
	if err != nil {
		return
	}

	fmt.Fprintf(w, string(file))
}
