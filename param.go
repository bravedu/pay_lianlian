package pay_lianlian

import (
	"crypto"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"hash"
	"musenetwork.org/pay-channel-lianlian/util"
	"musenetwork.org/pay-channel-lianlian/xlog"
	"strings"
)

func (w *Client) SetCountry() (client *Client) {
	w.mu.Lock()
	switch w.IsProd {
	case true:
		w.BaseURL = baseUrlCh
	case false:
		w.BaseURL = baseUrlCh //测试环境暂时未看到区分url,统一正式暂时
		w.OidPartner = "2020042200284052"
	default:
		w.BaseURL = baseUrlCh
	}
	w.mu.Unlock()
	return w
}

//签名方式
func GetReleaseSign(privateKey []byte, signType string, bm BodyMap) (sign string) {
	var h hash.Hash
	if signType == SignType_HMAC_SHA256 {
		h = hmac.New(sha256.New, privateKey)
	} else {
		h = md5.New()
	}
	sortStr := strings.ToLower(bm.EncodeLianLianSignParams())
	h.Write([]byte(sortStr))
	hashMd5 := h.Sum([]byte(nil))
	block, _ := pem.Decode(privateKey)
	if block == nil {
		xlog.Error("GetReleaseSign privateKey empty: %s", "")
		return ""
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		xlog.Error("ParsePKCS1PrivateKey  error: %s", err.Error())
		return ""
	}
	resEnc, err := rsa.SignPKCS1v15(nil, priv, crypto.MD5, hashMd5[:])
	if err != nil {
		xlog.Error("SignPKCS1v15  error: %s", err.Error())
		return ""
	}
	return base64.StdEncoding.EncodeToString(resEnc)
	//return hex.EncodeToString(resEnc)
}

// 生成请求Json的Body体
func GenerateJson(bm BodyMap) string {
	bs, err := json.Marshal(bm)
	if err != nil {
		return util.NULL
	}
	return string(bs)
}
