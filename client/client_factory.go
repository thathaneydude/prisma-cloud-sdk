package client

import (
	"crypto/tls"
	"crypto/x509"
	"github.com/thathaneydude/prisma-cloud-sdk/constants"
	"log"
	"net/http"
)

const (
	ContentTypeHeader = "Content-Type"
	UserAgentHeader   = "UserAgent"
	ApplicationJSON   = "application/json"
)

func NewBaseClient(sslVerify bool, maxRetries int, schema string) *BaseClientImpl {
	headers := &http.Header{}
	headers.Set(ContentTypeHeader, ApplicationJSON)
	headers.Set(UserAgentHeader, constants.UserAgent)
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
	return &BaseClientImpl{
		httpClient,
		headers,
		schema,
		maxRetries,
	}
}
