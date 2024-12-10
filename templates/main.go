package templates

import "net/http"

func Apply(mux *http.ServeMux) {
    apply_card(mux)
}
