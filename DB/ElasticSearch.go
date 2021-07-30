package DB

import (
	"context"
	"fmt"
	"github.com/Mau-MR/theaFirst/connection"
	"github.com/Mau-MR/theaFirst/data/types"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"strings"
)

type ElasticModifier struct {
	client *connection.ElasticConnection
	db     string
	index  string
}
type Query string

func NewElasticModifier(connection *connection.ElasticConnection, db, index string) *ElasticModifier {
	return &ElasticModifier{
		client: connection,
		db:     db,
		index:  index,
	}
}

func (em *ElasticModifier) Push(data types.Type) error {
	return nil
}

//Insert inserts a struct to the specific index and with the given id, if not provided "" elasticsearch generates one id
func (em *ElasticModifier) Insert(data types.Type) error {
	id := data.StringID()
	data.SetID("") //The ID is errased since elasticsearch doesnt allow the keyword _id into the body of a document
	res, err := em.client.Client.Index(
		em.index, esutil.NewJSONReader(data),
		em.client.Client.Index.WithDocumentID(id), //Assigning the specific id
	)
	if err != nil {
		return err
	}
	if res.IsError() {
		resInfo := res.String()
		return fmt.Errorf("Insert(): elasticError: %v:", resInfo)
	}
	return err
}

//Delete  receives the index and the PrimitiveID of the document and deletes it, returns error in case of failure
func (em *ElasticModifier) Delete(data types.Type) error {
	res, err := em.client.Client.Delete(
		em.index, data.StringID(),
		em.client.Client.Delete.WithContext(context.Background()),
	)
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("ElasticError: %v", res.Status())
	}
	return err
}
func (em *ElasticModifier) Update(data types.Type) error {
	return nil
}

func (em *ElasticModifier) BuildSearchQueryByFields(termAndFields *map[string]string) string {
	searchTerm := "" // todo Change to a string.Builder implementation
	var fields []string
	for field, term := range *termAndFields {
		searchTerm += term
		fields = append(fields, field)
	}
	query := `{
	"query": {
		"multi_match"  : {
			"query": %+q,
			"type": "best_fields",
			"fields": %+q
		}
	},
	"sort": [
	  {"_score": {"order" : "asc"}}
	],
	"size": 5
	}`
	return fmt.Sprintf(query, searchTerm, fields)
}

func (em *ElasticModifier) SearchID(data types.Type) (types.Type, error) {
	return nil, nil
}

//SearchFields makes a default search with the specified query and index, returns the response as a esapi.Response, and and error if occurred
func (em *ElasticModifier) SearchFields(data types.Type) ([]types.Type, error) {
	query := em.BuildSearchQueryByFields(data.SearchFields())
	res, err := em.client.Client.Search(
		em.client.Client.Search.WithContext(context.Background()),
		em.client.Client.Search.WithIndex(em.index),
		em.client.Client.Search.WithBody(strings.NewReader(query)),
		em.client.Client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, err
	}
	//TODO: Change return to return the type
	if res.IsError() {
		return nil, fmt.Errorf("ElasticError: %s", res.Status())
	}
	rw, err := em.GetResponseWrapper(res)
	if err != nil {
		return nil, err
	}
	var searchedTypes []types.Type
	for _, hit := range rw.Hits.Hits {
		tmpType := data.EmptyClone()
		err := tmpType.FromJSON(hit.Source)
		if err != nil {
			return searchedTypes, err
		}
		tmpType.SetID(hit.ID)
		searchedTypes = append(searchedTypes, tmpType)
	}
	return searchedTypes, nil
}

func (em *ElasticModifier) GetResponseWrapper(res *esapi.Response) (*ResponseWrapper, error) {
	rw, err := NewResponseWrapper(res)
	if err != nil {
		return nil, err
	}
	return rw, err
}
