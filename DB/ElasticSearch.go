package DB

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/elastic/go-elasticsearch/v7/esutil"
	"log"
	"net"
	"net/http"
	"time"
)

type ElasticWrapper struct {
	client *elasticsearch.Client
	l      *log.Logger
}

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

//InsertStructTo inserts a struct to the specific index and with the given id,
func (e *ElasticWrapper) InsertStructTo(Type interface{}, index, id string) (*esapi.Response, error) {
	if id == "" { //if the id is not provided elastic generates it
		return e.client.Index(index, esutil.NewJSONReader(&Type))
	}
	return e.client.Index(
		index, esutil.NewJSONReader(Type),
		e.client.Index.WithDocumentID(id),
	)
}
