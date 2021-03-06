package xhttp

import "errors"

var (
	MissInitParamErr   = errors.New("missing init parameter")
	MissParamErr       = errors.New("missing required parameter")
	MarshalErr         = errors.New("marshal error")
	UnmarshalErr       = errors.New("unmarshal error")
	SignatureErr       = errors.New("signature error")
	VerifySignatureErr = errors.New("verify signature error")
	CertNotMatchErr    = errors.New("cert not match error")
	GetSignDataErr     = errors.New("get signature data error")
)
