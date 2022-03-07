package base

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	logger "github.com/fideism/golang-wechat/log"
	"github.com/fideism/golang-wechat/util"
	"github.com/sirupsen/logrus"
)

// Response 返回结构
type Response struct {
	ReturnCode string      `json:"return_code"`
	ReturnMsg  string      `json:"return_msg"`
	ErrCode    string      `json:"err_code"`
	ErrCodeDes string      `json:"err_code_des"`
	ResultCode string      `json:"result_code"`
	Data       util.Params `json:"data"`
	Detail     string      `json:"detail"`
}

const contentType = "application/xml; charset=utf-8"

// PostWithoutCert https no cert post
func PostWithoutCert(url string, params util.Params) (string, error) {
	logger.Entry().WithFields(logrus.Fields{
		"url":    url,
		"params": params,
	}).Debug("发起微信请求 without cert")
	h := &http.Client{}
	response, err := h.Post(url, contentType, strings.NewReader(MapToXML(params)))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

// PostWithTSL https need cert post
func PostWithTSL(url string, params util.Params, config *tls.Config) (string, error) {
	transport := &http.Transport{
		TLSClientConfig:    config,
		DisableCompression: true,
	}
	h := &http.Client{Transport: transport}
	response, err := h.Post(url, contentType, strings.NewReader(MapToXML(params)))
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	res, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

//ProcessResponseXML 处理 HTTPS API返回数据，转换成Map对象
func ProcessResponseXML(xmlStr string) (*Response, error) {
	params := XMLToMap(xmlStr)

	var response Response

	if params.Exists("return_code") {
		response.ReturnCode = params.GetString("return_code")
	} else {
		return nil, errors.New("no return_code in XML")
	}

	response.ReturnMsg = params.GetString("return_msg")
	response.ErrCode = params.GetString("err_code")
	response.ErrCodeDes = params.GetString("err_code_des")
	response.ResultCode = params.GetString("result_code")
	response.Data = params

	return &response, nil
}

// XMLToMap xml to map
func XMLToMap(xmlStr string) util.Params {
	params := make(util.Params)
	decoder := xml.NewDecoder(strings.NewReader(xmlStr))

	var (
		key   string
		value string
	)

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement: // 开始标签
			key = token.Name.Local
		case xml.CharData: // 标签内容
			content := string([]byte(token))
			value = content
		}
		if key != "xml" {
			if value != "\n" {
				params.Set(key, value)
			}
		}
	}

	return params
}

// MapToXML map to xml
func MapToXML(params util.Params) string {
	var buf bytes.Buffer
	buf.WriteString(`<xml>`)
	for k, v := range params {
		buf.WriteString(`<`)
		buf.WriteString(k)
		buf.WriteString(`><![CDATA[`)
		buf.WriteString(util.InterfaceToString(v))
		buf.WriteString(`]]></`)
		buf.WriteString(k)
		buf.WriteString(`>`)
	}
	buf.WriteString(`</xml>`)

	fmt.Println(buf.String())
	return buf.String()
}
