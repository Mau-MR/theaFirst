package DB

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v7"
	"net"
	"net/http"
	"time"
)

func ElasticSearchClient(address, username, password string) (*elasticsearch.Client, error) {
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
