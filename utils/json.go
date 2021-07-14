package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}

func ParseRequest(i interface{}, r io.Reader, rw http.ResponseWriter) error{
	err :=FromJSON(i,r)
	if err !=nil  {
		rw.WriteHeader(http.StatusBadRequest)
		ToJSON(GenericError{
			Message: fmt.Sprintf("Bad Format for type: %T", i),
		},rw)
	}
	return nil
}
