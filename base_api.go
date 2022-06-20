package pay_lianlian

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"musenetwork.org/pay-channel-lianlian/xhttp"
)

// 收款-统一下单
//	文档地址：https://open.lianlianpay.com/apis/unified-payment.html

func (w *Client) UnifiedOrder(ctx context.Context, bm BodyMap) (orderRsp *UnifiedOrderResponse, err error) {
	err = bm.CheckEmptyError("no_order", "name_goods", "money_order", "notify_url", "user_id")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if w.IsProd {
		bs, err = w.doProdPost(ctx, bm, tradeCreate, nil)
	} else {
		bm.Set("money_order", 0.01)
		bs, err = w.doProdPost(ctx, bm, tradeCreate, nil)
	}
	if err != nil {
		return nil, err
	}
	orderRsp = new(UnifiedOrderResponse)
	if err = json.Unmarshal(bs, orderRsp); err != nil {
		return nil, fmt.Errorf("[%w]: %v, bytes: %s", xhttp.UnmarshalErr, err, string(bs))
	}
	if orderRsp.RetCode != RET_CODE_SUCCESS {
		return nil, errors.New(orderRsp.RetMsg)
	}
	return orderRsp, nil
}
