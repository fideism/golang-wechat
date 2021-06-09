package base

import "github.com/fideism/golang-wechat/pay/config"

//正常模式
const (
	MicroPayURL          = "https://api.mch.weixin.qq.com/pay/micropay"
	UnifiedOrderURL      = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	OrderQueryURL        = "https://api.mch.weixin.qq.com/pay/orderquery"
	ReverseURL           = "https://api.mch.weixin.qq.com/secapi/pay/reverse"
	CloseOrderURL        = "https://api.mch.weixin.qq.com/pay/closeorder"
	RefundURL            = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	RefundQueryURL       = "https://api.mch.weixin.qq.com/pay/refundquery"
	DownloadBillURL      = "https://api.mch.weixin.qq.com/pay/downloadbill"
	DownloadFundFlowURL  = "https://api.mch.weixin.qq.com/pay/downloadfundflow"
	ReportURL            = "https://api.mch.weixin.qq.com/payitil/report"
	BatchQueryCommentURL = "https://api.mch.weixin.qq.com/billcommentsp/batchquerycomment"
	AuthCodeToOpenidURL  = "https://api.mch.weixin.qq.com/tools/authcodetoopenid"
)

//沙箱模式
const (
	SandboxMicroPayURL          = "https://api.mch.weixin.qq.com/sandboxnew/pay/micropay"
	SandboxUnifiedOrderURL      = "https://api.mch.weixin.qq.com/sandboxnew/pay/unifiedorder"
	SandboxOrderQueryURL        = "https://api.mch.weixin.qq.com/sandboxnew/pay/orderquery"
	SandboxReverseURL           = "https://api.mch.weixin.qq.com/sandboxnew/secapi/pay/reverse"
	SandboxCloseOrderURL        = "https://api.mch.weixin.qq.com/sandboxnew/pay/closeorder"
	SandboxRefundURL            = "https://api.mch.weixin.qq.com/sandboxnew/secapi/pay/refund"
	SandboxRefundQueryURL       = "https://api.mch.weixin.qq.com/sandboxnew/pay/refundquery"
	SandboxDownloadBillURL      = "https://api.mch.weixin.qq.com/sandboxnew/pay/downloadbill"
	SandboxDownloadFundFlowURL  = "https://api.mch.weixin.qq.com/sandboxnew/pay/downloadfundflow"
	SandboxReportURL            = "https://api.mch.weixin.qq.com/sandboxnew/payitil/report"
	SandboxBatchQueryCommentURL = "https://api.mch.weixin.qq.com/billcommentsp/batchquerycomment"
	SandboxAuthCodeToOpenidURL  = "https://api.mch.weixin.qq.com/sandboxnew/tools/authcodetoopenid"
)

// GetUnifyURL 统一下单
func GetUnifyURL(c *config.Config) string {
	if c.Sandbox {
		return SandboxUnifiedOrderURL
	}

	return UnifiedOrderURL
}

// GetMicroPayURL 付款码支付
func GetMicroPayURL(c *config.Config) string {
	if c.Sandbox {
		return SandboxMicroPayURL
	}

	return MicroPayURL
}

// GetOrderQueryURL 查询订单
func GetOrderQueryURL(c *config.Config) string {
	if c.Sandbox {
		return SandboxOrderQueryURL
	}

	return OrderQueryURL
}

// GetCloseOrderURL 关闭订单
func GetCloseOrderURL(c *config.Config) string {
	if c.Sandbox {
		return SandboxCloseOrderURL
	}

	return CloseOrderURL
}

// GetReverseURL 撤销订单
func GetReverseURL(c *config.Config) string {
	if c.Sandbox {
		return SandboxReverseURL
	}

	return ReverseURL
}

// GetRefundURL 退款接口
func GetRefundURL(c *config.Config) string {
	if c.Sandbox {
		return SandboxRefundURL
	}

	return RefundURL
}

// GetRefundQueryURL 查询退款接口
func GetRefundQueryURL(c *config.Config) string {
	if c.Sandbox {
		return SandboxRefundQueryURL
	}

	return RefundQueryURL
}

// GetDownloadBillURL 获取下载交易账单URL
func GetDownloadBillURL(c *config.Config) string {
	if c.Sandbox {
		return SandboxDownloadBillURL
	}

	return DownloadBillURL
}

// GetDownloadFundFlowURL  下载资金账单
func GetDownloadFundFlowURL(c *config.Config) string {
	if c.Sandbox {
		return SandboxDownloadFundFlowURL
	}

	return DownloadFundFlowURL
}

// GetReportURL 交易保障
func GetReportURL(c *config.Config) string {
	if c.Sandbox {
		return SandboxReportURL
	}

	return ReportURL
}

// GetAuthCodeToOpenidURL 付款码查询openid
func GetAuthCodeToOpenidURL(c *config.Config) string {
	if c.Sandbox {
		return SandboxAuthCodeToOpenidURL
	}

	return AuthCodeToOpenidURL
}

// GetBatchQueryCommentURL 拉取订单评价数据
func GetBatchQueryCommentURL(c *config.Config) string {
	if c.Sandbox {
		return SandboxBatchQueryCommentURL
	}

	return BatchQueryCommentURL
}
