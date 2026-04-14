package proxmox

import (
	"crypto/tls"
	"net/http"
)

func NewClient() *http.Client {

	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: transport,
	}

	return client
}
