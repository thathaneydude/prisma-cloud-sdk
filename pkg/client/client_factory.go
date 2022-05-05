package client

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
)

const (
	ContentTypeHeader = "Content-Type"
	ApplicationJSON   = "application/json"
)

func NewBaseClient(baseUrl string, sslVerify bool, maxRetries int) *BaseClientImpl {
	headers := &http.Header{}
	headers.Set(ContentTypeHeader, ApplicationJSON)
	var tlsConfig tls.Config
	if !sslVerify {
		tlsConfig = tls.Config{
			InsecureSkipVerify: true,
		}
	} else {
		// On Unix systems other than macOS the environment variables SSL_CERT_FILE and SSL_CERT_DIR can be used to
		//override the system default locations for the SSL certificate file and SSL certificate
		//files directory, respectively.
		systemCertPool, err := x509.SystemCertPool()
		if err != nil {
			log.Printf("Failed to determine system level certificates to verify TLS connections: %v", err)
			return nil
		}
		tlsConfig.RootCAs = systemCertPool
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tlsConfig,
		},
	}
	return &BaseClientImpl{httpClient, headers, baseUrl, maxRetries}
}
