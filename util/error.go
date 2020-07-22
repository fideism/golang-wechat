package util

import (
	"encoding/json"
	"fmt"
	logger "github.com/fideism/golang-wechat/log"
	"github.com/sirupsen/logrus"
	"reflect"
)

// CommonError 微信返回的通用错误json
type CommonError struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

// DecodeWithCommonError 将返回值按照CommonError解析
func DecodeWithCommonError(response []byte, apiName string) (err error) {
	var commError CommonError
	logger.Entry().WithFields(logrus.Fields{
		"api_name": apiName,
		"data":     commError,
	}).Debug("解析微信返回信息")
	err = json.Unmarshal(response, &commError)
	if err != nil {
		return
	}
	if commError.ErrCode != 0 {
		return fmt.Errorf("%s Error , errcode=%d , errmsg=%s", apiName, commError.ErrCode, commError.ErrMsg)
	}
	return nil
}

// DecodeWithCustomerStruct 将返回值按照CommonError解析
func DecodeWithCustomerStruct(response []byte, obj interface{}, apiName string) error {
	err := json.Unmarshal(response, obj)
	logger.Entry().WithFields(logrus.Fields{
		"api_name": apiName,
		"data":     obj,
	}).Debug("解析微信返回信息")
	if err != nil {
		return fmt.Errorf("%s json Unmarshal Error, err=%v", apiName, err)
	}
	responseObj := reflect.ValueOf(obj)
	if !responseObj.IsValid() {
		return fmt.Errorf("%s obj is invalid", apiName)
	}

	return nil
}

// DecodeWithError 将返回值按照解析
func DecodeWithError(response []byte, obj interface{}, apiName string) error {
	err := json.Unmarshal(response, obj)
	logger.Entry().WithFields(logrus.Fields{
		"api_name": apiName,
		"data":     obj,
	}).Debug("解析微信返回信息")
	if err != nil {
		return fmt.Errorf("json Unmarshal Error, err=%v", err)
	}
	responseObj := reflect.ValueOf(obj)
	if !responseObj.IsValid() {
		return fmt.Errorf("obj is invalid")
	}
	commonError := responseObj.Elem().FieldByName("CommonError")
	if !commonError.IsValid() || commonError.Kind() != reflect.Struct {
		return fmt.Errorf("commonError is invalid or not struct")
	}
	errCode := commonError.FieldByName("ErrCode")
	errMsg := commonError.FieldByName("ErrMsg")
	if !errCode.IsValid() || !errMsg.IsValid() {
		return fmt.Errorf("errcode or errmsg is invalid")
	}
	if errCode.Int() != 0 {
		return fmt.Errorf("%s Error , errcode=%d , errmsg=%s", apiName, errCode.Int(), errMsg.String())
	}
	return nil
}
