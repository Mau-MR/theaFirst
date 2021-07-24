package connection

import (
	"crypto/tls"
	"github.com/elastic/go-elasticsearch/v7"
	"net"
	"net/http"
	"os"
	"time"
)

//ElasticConnection manage the methods for matching the connection interface with elasticsearch
type ElasticConnection struct {
	Client *elasticsearch.Client
}

func (ec *ElasticConnection) Connect() error {
	//The retrieve of the credentials
	address := os.Getenv("ELASTICURI")
	username := os.Getenv("EUSER")
	password := os.Getenv("EPASSWORD")

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
	//TODO: MAKE A PING TO VERIFY THE CONNECTION
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}
	ec.Client = es
	return nil
}
func (ec *ElasticConnection) Close() error {
	//Nil implementation since elastic client doesnt  has a method for this
	return nil
}
