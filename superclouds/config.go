package superclouds

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"
)

// Config contains the configuration settings for connecting to the Superclouds API.
type Config struct {
	SuperURL   string
	CertPath   string
	KeyPath    string
	SuperToken string
	Client     *http.Client
}

// NewConfig creates a new Config instance using environment variables for cert and key paths, and token.
// The environment variables that need to be set are:
// - SUPER_CERT: The path to the SSL certificate file.
// - SUPER_KEY: The path to the SSL key file.
// - SUPER_TOKEN: The bearer token for API authorization.
//
// Example usage:
//
//	cfg, err := superclouds.NewConfig()
//	if err != nil {
//	    log.Fatalf("Failed to create config: %v", err)
//	}
func NewConfig() (*Config, error) {
	certPath := os.Getenv("SUPER_CERT")
	if certPath == "" {
		return nil, fmt.Errorf("missing SUPER_CERT environment variable")
	}

	keyPath := os.Getenv("SUPER_KEY")
	if keyPath == "" {
		return nil, fmt.Errorf("missing SUPER_KEY environment variable")
	}

	superToken := os.Getenv("SUPER_TOKEN")
	if superToken == "" {
		return nil, fmt.Errorf("missing SUPER_TOKEN environment variable")
	}

	client, err := setupClient(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	return &Config{
		SuperURL:   apiBaseURL,
		CertPath:   certPath,
		KeyPath:    keyPath,
		SuperToken: superToken,
		Client:     client,
	}, nil
}

// NewConfigWithParams creates a new Config instance using provided parameters for cert and key paths, and token.
//
// Parameters:
// - certPath: The path to the SSL certificate file.
// - keyPath: The path to the SSL key file.
// - token: The bearer token for API authorization.
//
// Example usage:
//
//	cfg, err := superclouds.NewConfigWithParams(certPath, keyPath, superToken)
//	if err != nil {
//	    log.Fatalf("Failed to create config: %v", err)
//	}
func NewConfigWithParams(certPath, keyPath, token string) (*Config, error) {
	client, err := setupClient(certPath, keyPath)
	if err != nil {
		return nil, err
	}

	return &Config{
		SuperURL:   apiBaseURL,
		CertPath:   certPath,
		KeyPath:    keyPath,
		SuperToken: token,
		Client:     client,
	}, nil
}

func setupClient(certPath, keyPath string) (*http.Client, error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load key pair: %v", err)
	}

	caCertPool := x509.NewCertPool()
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				Certificates:       []tls.Certificate{cert},
				RootCAs:            caCertPool,
			},
		},
	}
	return client, nil
}
