package pay_lianlian

import (
	"context"
	"crypto/tls"
	"fmt"
	"musenetwork.org/pay-channel-lianlian/util"
	"musenetwork.org/pay-channel-lianlian/xhttp"
	"musenetwork.org/pay-channel-lianlian/xlog"
	"sync"
)

type Client struct {
	ApiVersion string
	OidPartner string //测试商户号
	PrivateKey []byte
	PublicKey  []byte
	BaseURL    string
	SignType   string //RSA
	IsProd     bool   //true 生产环境  false 测试环境
	Debug      string
	BaseDomain string
	mu         sync.RWMutex
}

func NewClient(oidPartner string, privateKey, publicKey []byte, isProd bool, debug string) (client *Client) {
	baseDomain := ""
	if isProd {
		baseDomain = baseUrlCh
	} else {
		baseDomain = baseUrlChTest
	}
	return &Client{
		OidPartner: oidPartner,
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		IsProd:     isProd,
		Debug:      debug,
		BaseDomain: baseDomain,
		ApiVersion: "1.0",
		SignType:   "RSA",
	}
}

// Post请求、正式
func (w *Client) doProdPost(ctx context.Context, bm BodyMap, path string, tlsConfig *tls.Config) (bs []byte, err error) {
	var url = w.BaseDomain + path
	if bm.GetString("sign_type") == util.NULL {
		bm.Set("sign_type", w.SignType)
	}
	if bm.GetString("oid_partner") == util.NULL {
		bm.Set("oid_partner", w.OidPartner)
	}
	signStr := GetReleaseSign(w.PrivateKey, bm.GetString("sign_type"), bm)
	httpClient := xhttp.NewClient()
	httpClient.Header.Add("Signature-Type", w.SignType)
	httpClient.Header.Add("Signature-Data", signStr)
	if w.IsProd && tlsConfig != nil {
		httpClient.SetTLSConfig(tlsConfig)
	}
	req := GenerateJson(bm)
	if w.Debug == DebugOn {
		xlog.Debugf("Lianlian_Request: %s", req)
	}
	res, bs, err := httpClient.Type(xhttp.TypeJSON).Post(url).SendString(req).EndBytes(ctx)
	if err != nil {
		return nil, err
	}
	if w.Debug == DebugOn {
		xlog.Debugf("Lianlian_Response: %s%d %s%s", xlog.Red, res.StatusCode, xlog.Reset, string(bs))
	}
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Request Error, StatusCode = %d", res.StatusCode)
	}
	return bs, nil
}
