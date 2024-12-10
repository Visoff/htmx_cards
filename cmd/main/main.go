package main

import (
	"net/http"
	"strings"

	templs "github.com/Visoff/uni_cards/templates"
)

func main() {
    mux := http.NewServeMux()
    templs.Apply(mux)
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        path := strings.Split(r.URL.Path, "/")
        if path[len(path)-1] == "" {
            path[len(path)-1] = "index.html"
        }
        http.ServeFile(w, r, "dist/"+strings.Join(path, "/"))
    })
    
    println("server started on port 8080")
    err := http.ListenAndServe(":8080", mux)
    if err != nil {
        panic(err)
    }
}
