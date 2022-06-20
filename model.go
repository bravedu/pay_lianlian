package pay_lianlian

const (

	// URL
	baseUrlCh     = "https://accpapi.lianlianpay.com/"
	baseUrlChTest = "https://accpapi-ste.lianlianpay-inc.com/"

	//业务接口
	//统一支付创单API
	payCreateBill = "v1/paycreatebill"
	//收款接口查询
	orderQuery = "orderquery.htm"

	//支付统一创单
	tradeCreate = "v1/txn/tradecreate"
	//网关类支付
	paymentGw = "v1/txn/payment-gw"

	// 签名方式
	SignType_MD5         = "MD5"
	SignType_HMAC_SHA256 = "HMAC-SHA256"
	SignType_MD5WithRSA  = "MD5WithRSA"
	SignType_RSA         = "RSA"

	RET_CODE_SUCCESS = "0000"
)

//统一下单返回
type UnifiedOrderResponse struct {
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	SignType   string `json:"sign_type"`
	Sign       string `json:"sign"`
	NoOrder    string `json:"no_order"`
	DtOrder    string `json:"dt_order"`
	MoneyOrder string `json:"money_order"`
	OidPaybill string `json:"oid_paybill"`
	UserID     string `json:"user_id"`
	OidPartner string `json:"oid_partner"`
	GatewayURL string `json:"gateway_url"`
	Token      string `json:"token"`
}

//订单收款结果查询
type QueryOrderResponse struct {
	DtOrder    string `json:"dt_order"`
	InfoOrder  string `json:"info_order"`
	MoneyOrder string `json:"money_order"`
	NoOrder    string `json:"no_order"`
	OidPartner string `json:"oid_partner"`
	OidPaybill string `json:"oid_paybill"`
	ResultPay  string `json:"result_pay"`
	SettleDate string `json:"settle_date"`
	RetCode    string `json:"ret_code"`
	RetMsg     string `json:"ret_msg"`
	PayType    string `json:"pay_type"`
	BankCode   string `json:"bank_code"`
	BankName   string `json:"bank_name"`
	CardNo     string `json:"card_no"`
	Sign       string `json:"sign"`
	SignType   string `json:"sign_type"`
}

//支付类回调通知
type PayNotifyResponse struct {
	OidPartner string `json:"oid_partner"`
	DtOrder    string `json:"dt_order"`
	NoOrder    string `json:"no_order"`
	OidPaybill string `json:"oid_paybill"`
	MoneyOrder string `json:"money_order"`
	ResultPay  string `json:"result_pay"`
	SettleDate string `json:"settle_date"`
	InfoOrder  string `json:"info_order"`
	PayType    string `json:"pay_type"`
	BankCode   string `json:"bank_code"`
	SignType   string `json:"sign_type"`
	Sign       string `json:"sign"`
}
