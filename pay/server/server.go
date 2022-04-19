package server

import (
	"strings"

	"github.com/fideism/golang-wechat/pay/base"
	"github.com/fideism/golang-wechat/pay/config"
	"github.com/fideism/golang-wechat/util"
)

// Server server
type Server struct {
	*config.Config
}

// NewServer NewServer
func NewServer(c *config.Config) *Server {
	return &Server{c}
}

// DownloadBill 下载交易账单
func (server *Server) DownloadBill(params util.Params) (*base.Response, error) {
	params = base.Sign(params, server.Config)

	xmlStr, err := base.PostWithoutCert(base.GetDownloadBillURL(server.Config), params)
	if err != nil {
		return nil, err
	}
	// 如果出现错误，返回XML数据
	if strings.Index(xmlStr, "<") == 0 {
		return base.ProcessResponseXML(xmlStr)
	}

	return &base.Response{
		ReturnCode: "SUCCESS",
		ReturnMsg:  "OK",
		Detail:     xmlStr,
	}, nil
}

// DownloadFundFlow 下载资金账单
func (server *Server) DownloadFundFlow(params util.Params, certCfg config.Cert) (*base.Response, error) {
	params = base.Sign(params, server.Config)

	tsl, err := base.CertTLSConfig(server.Config.MchID, certCfg)
	if err != nil {
		return nil, err
	}

	xmlStr, err := base.PostWithTSL(base.GetDownloadFundFlowURL(server.Config), params, tsl)
	if err != nil {
		return nil, err
	}
	// 如果出现错误，返回XML数据
	if strings.Index(xmlStr, "<") == 0 {
		return base.ProcessResponseXML(xmlStr)
	}
	return &base.Response{
		ReturnCode: "SUCCESS",
		ReturnMsg:  "OK",
		Detail:     xmlStr,
	}, nil
}

// Report 交易保障
func (server *Server) Report(params util.Params) (*base.Response, error) {
	params = base.Sign(params, server.Config)

	xmlStr, err := base.PostWithoutCert(base.GetReportURL(server.Config), params)
	if err != nil {
		return nil, err
	}

	return base.ProcessResponseXML(xmlStr)
}

// AuthCodeToOpenid 付款码查询openid
func (server *Server) AuthCodeToOpenid(params util.Params) (*base.Response, error) {
	params = base.Sign(params, server.Config)

	xmlStr, err := base.PostWithoutCert(base.GetAuthCodeToOpenidURL(server.Config), params)
	if err != nil {
		return nil, err
	}

	return base.ProcessResponseXML(xmlStr)
}

// BatchQueryComment 拉取订单评价数据
func (server *Server) BatchQueryComment(params util.Params) (*base.Response, error) {
	params = base.Sign(params, server.Config)

	xmlStr, err := base.PostWithoutCert(base.GetBatchQueryCommentURL(server.Config), params)
	if err != nil {
		return nil, err
	}

	// 如果出现错误，返回XML数据
	if strings.Index(xmlStr, "<") == 0 {
		return base.ProcessResponseXML(xmlStr)
	}
	// 正常返回csv数据
	return &base.Response{
		ReturnCode: "SUCCESS",
		ReturnMsg:  "OK",
		Detail:     xmlStr,
	}, nil
}
