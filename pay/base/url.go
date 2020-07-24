package base

import "github.com/fideism/golang-wechat/pay/config"

const (
	//正常模式
	MicroPayURL         = "https://api.mch.weixin.qq.com/pay/micropay"
	UnifiedOrderURL     = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	OrderQueryURL       = "https://api.mch.weixin.qq.com/pay/orderquery"
	ReverseURL          = "https://api.mch.weixin.qq.com/secapi/pay/reverse"
	CloseOrderURL       = "https://api.mch.weixin.qq.com/pay/closeorder"
	RefundURL           = "https://api.mch.weixin.qq.com/secapi/pay/refund"
	RefundQueryURL      = "https://api.mch.weixin.qq.com/pay/refundquery"
	DownloadBillURL     = "https://api.mch.weixin.qq.com/pay/downloadbill"
	DownloadFundFlowURL = "https://api.mch.weixin.qq.com/pay/downloadfundflow"
	ReportURL           = "https://api.mch.weixin.qq.com/payitil/report"
	ShortURL            = "https://api.mch.weixin.qq.com/tools/shortURL"
	AuthCodeToOpenidURL = "https://api.mch.weixin.qq.com/tools/authcodetoopenid"

	//沙箱模式
	SandboxMicroPayURL         = "https://api.mch.weixin.qq.com/sandboxnew/pay/micropay"
	SandboxUnifiedOrderURL     = "https://api.mch.weixin.qq.com/sandboxnew/pay/unifiedorder"
	SandboxOrderQueryURL       = "https://api.mch.weixin.qq.com/sandboxnew/pay/orderquery"
	SandboxReverseURL          = "https://api.mch.weixin.qq.com/sandboxnew/secapi/pay/reverse"
	SandboxCloseOrderURL       = "https://api.mch.weixin.qq.com/sandboxnew/pay/closeorder"
	SandboxRefundURL           = "https://api.mch.weixin.qq.com/sandboxnew/secapi/pay/refund"
	SandboxRefundQueryURL      = "https://api.mch.weixin.qq.com/sandboxnew/pay/refundquery"
	SandboxDownloadBillURL     = "https://api.mch.weixin.qq.com/sandboxnew/pay/downloadbill"
	SandboxDownloadFundFlowURL = "https://api.mch.weixin.qq.com/sandboxnew/pay/downloadfundflow"
	SandboxReportURL           = "https://api.mch.weixin.qq.com/sandboxnew/payitil/report"
	SandboxShortURL            = "https://api.mch.weixin.qq.com/sandboxnew/tools/shortURL"
	SandboxAuthCodeToOpenidURL = "https://api.mch.weixin.qq.com/sandboxnew/tools/authcodetoopenid"
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
