package credential

//WxCardTicketHandle wx_card ticket获取
type WxCardTicketHandle interface {
	//GetTicket 获取ticket
	GetTicket(accessToken string) (ticket string, err error)
}
