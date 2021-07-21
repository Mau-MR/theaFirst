package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//ToJSON converts the specified type to a json string
func ToJSON(i interface{}, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(i)
}

//FromJSON converts a json string on the specified type
func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(&i) //TODO CHECK IF THIS & AFFECTS OTHER PARTS OF THE CODE
}

//TODO: ADD  COMMENT
func ParseRequest(i interface{}, r io.Reader, rw http.ResponseWriter) error {
	err := FromJSON(i, r)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		ToJSON(GenericError{
			Message: fmt.Sprintf("Bad Format for type: %T", i),
		}, rw)
		return err
	}
	return nil
}
