package route

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func enableCORSSupport(router *httprouter.Router) {
	router.GlobalOPTIONS = http.HandlerFunc(func(resWriter http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Access-Control-Request-Method") != "" {
			// Set CORS headers
			header := resWriter.Header()
			header.Set("Access-Control-Allow-Methods", header.Get("Allow"))
			header.Set("Access-Control-Allow-Origin", "*")
			header.Set("Access-Control-Allow-Headers", "Access-Control-Allow-Headers, Origin, Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers, authorization")
		}

		resWriter.WriteHeader(http.StatusNoContent)
	})
}
