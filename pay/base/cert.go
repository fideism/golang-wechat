package base

import (
	"crypto/tls"
	"encoding/pem"
	logger "github.com/fideism/golang-wechat/log"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
)

// CertTLSConfig 证书 tls
func CertTLSConfig(mchID, path string) (*tls.Config, error) {
	certData, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	blocks, err := pkcs12.ToPEM(certData, mchID)

	defer func() {
		if x := recover(); x != nil {
			logger.Entry().WithField("recover err", x).Error("recover")
		}
	}()

	if err != nil {
		return nil, err
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	pem, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		return nil, err
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{pem},
	}

	return config, nil
}
