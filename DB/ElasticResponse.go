package DB

import (
	"encoding/json"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"io/ioutil"
)

type Hits struct {
	ID     string          `json:"_id"`
	Source json.RawMessage `json:"_source"`
	Sort   []interface{}   `json:"sort"`
	//Todo: add new  relevant fields
}

//ResponseWrapper is the format on which esapi.Response returns the data from the query on elasticsearch
type ResponseWrapper struct {
	Took int
	Hits struct {
		Total struct {
			Value int `json:"value"`
		} `json:"total"`
		Hits []Hits `json:"hits"`
	} `json:"hits"`
}

//NewResponseWrapper converts and esapi.Response into a ResponseWrapper
func NewResponseWrapper(res *esapi.Response) (*ResponseWrapper, error) {
	var r ResponseWrapper
	defer res.Body.Close()
	jsonByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonByte, &r) //NOTE:it was used json.Unmarshal since utils.FromJSON doesnt work for nested structs
	return &r, err
}
