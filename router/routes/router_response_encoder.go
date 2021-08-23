package route

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func EncodeResponse(w http.ResponseWriter, resp interface{}, code int) error {
	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	w.WriteHeader(code)

	_, err = fmt.Fprint(w, string(b))
	return err
}

func EncodeErrorResponse(w http.ResponseWriter, srvErr error) error {
	log.Println("srvErr:", srvErr)

	b, err := json.Marshal(srvErr)
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusInternalServerError)

	log.Println("err:", string(b))

	_, err = fmt.Fprint(w, string(b))
	return err
}
