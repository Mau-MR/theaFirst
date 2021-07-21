package DB

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"log"
	"net"
	"net/http"
	"strings"
	"time"
)

type ElasticWrapper struct {
	client *elasticsearch.Client
	l      *log.Logger
}
type Query string

func NewElasticWrapper(address, username, password string, logger *log.Logger) (*ElasticWrapper, error) {
	ElasticClient, err := elasticSearchClient(address, username, password)
	if err != nil {
		return nil, err
	}
	return &ElasticWrapper{
		client: ElasticClient,
		l:      logger,
	}, nil
}
func elasticSearchClient(address, username, password string) (*elasticsearch.Client, error) {
	//The retrieve of the credentials
	//The configuration  of the client
	cfg := elasticsearch.Config{
		Addresses: []string{
			address,
		},
		Username: username,
		Password: password,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: 10 * time.Second,
			DialContext:           (&net.Dialer{Timeout: 10 * time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MaxVersion:         tls.VersionTLS11,
				InsecureSkipVerify: true,
			},
		},
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return es, nil
}

//InsertStructTo inserts a struct to the specific index and with the given id, if not provided "" elasticsearch generates one id
func (e *ElasticWrapper) InsertStructTo(index, id string, Type interface{}) (*esapi.Response, error) {
	if id == "" { //if the id is not provided elastic generates it
		return e.client.Index(index, esutil.NewJSONReader(&Type))
	}
	return e.client.Index(
		index, esutil.NewJSONReader(Type),
		e.client.Index.WithDocumentID(id),
	)
}

//DeleteDocumentByID  receives the index and the ID of the document and deletes it, returns error in case of failure
func (e *ElasticWrapper) DeleteDocumentByID(index, ID string) (*esapi.Response, error) {
	res, err := e.client.Delete(
		index, ID,
		e.client.Delete.WithContext(context.Background()),
	)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf("ElasticError: %v", res.Status())
	}
	return res, err
}

func (e *ElasticWrapper) BuildSearchQueryByFields(SearchTerm string, fields []string) string {
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
	return fmt.Sprintf(query, SearchTerm, fields)
}

//ESearchWithDefault makes a default search with the specified query and index, returns the response as a esapi.Response, and and error if occurred
func (e *ElasticWrapper) ESearchWithDefault(index, query string) (*esapi.Response, error) {
	res, err := e.client.Search(
		e.client.Search.WithContext(context.Background()),
		e.client.Search.WithIndex(index),
		e.client.Search.WithBody(strings.NewReader(query)),
		e.client.Search.WithTrackTotalHits(true),
	)
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf("ElasticError: %s", res.Status())
	}
	return res, err
}
func (e *ElasticWrapper) GetResponseWrapper(res *esapi.Response) (*ResponseWrapper, error) {
	rw, err := NewResponseWrapper(res)
	if err != nil {
		return nil, err
	}
	return rw, err
}

func (e *ElasticWrapper) SearchIn(index string, query string) (*ResponseWrapper, error) {
	res, err := e.ESearchWithDefault(index, query)
	if err != nil {
		return nil, err
	}
	return e.GetResponseWrapper(res)
}
