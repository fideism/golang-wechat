package base

import (
	"crypto/tls"
	"encoding/pem"
	"io/ioutil"

	logger "github.com/fideism/golang-wechat/log"
	"github.com/fideism/golang-wechat/pay/config"
	"golang.org/x/crypto/pkcs12"
)

// CertTLSConfig 证书 tls
func CertTLSConfig(mchID string, certCfg config.Cert) (*tls.Config, error) {
	certData, err := getCertByte(certCfg)
	if err != nil {
		return nil, err
	}

	blocks, err := pkcs12.ToPEM(certData, mchID)
	if err != nil {
		return nil, err
	}

	defer func() {
		if x := recover(); x != nil {
			logger.Entry().WithField("recover err", x).Error("recover")
		}
	}()

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	keyPair, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{keyPair},
	}, nil
}

func getCertByte(certCfg config.Cert) ([]byte, error) {
	if len(certCfg.Content) > 0 {
		return certCfg.Content, nil
	}

	return ioutil.ReadFile(certCfg.Path)
}
