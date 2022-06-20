##支付SDK配置使用

### yaml 配置

```cgo
pay_lian_lian:
  oid_partner: "2020042200284052"
  public_key: "config/lianlian_cert/merchant_rsa_public_key.pem"
  private_key: "config/lianlian_cert/merchant_rsa_private_key.pem"
  notifyurl: https://test-api.xxx.vip/v1/lianlian_notify
  debug: on
  is_prod: false
  
```

### 获取配置实例化配置信息

```cgo
type Config struct {
	Yaml              *yamls
	LianLianPayClient *lianlian.Client
}

type yamls struct {
	PayLianLian payLianLian `yaml:"pay_lian_lian"`
}

type payLianLian struct {
	OidPartner string `yaml:"oid_partner"`
	PublicKey  string `yaml:"public_key"`
	PrivateKey string `yaml:"private_key"`
	Notifyurl  string `yaml:"notifyurl"`
	Debug      string `yaml:"debug"`
	IsProd     bool   `yaml:"is_prod"`
}

func (c *Config) initLianLianPay() {
	publickKey, err := ioutil.ReadFile(c.Yaml.PayLianLian.PublicKey)
	if err != nil {
		panic(err)
	}
	privateKey, err := ioutil.ReadFile(c.Yaml.PayLianLian.PrivateKey)
	if err != nil {
		panic(err)
	}
	client := lianlian.NewClient(c.Yaml.PayLianLian.OidPartner, privateKey, publickKey, c.Yaml.PayLianLian.IsProd, c.Yaml.PayLianLian.Debug)
	client.SetCountry()
	c.LianLianPayClient = client
}

//初始化支付系统
func ConfInstance() *Config {
	confOnce.Do(func() {
		//file := flag.String("conf", "", "请指定配置文件")
		//flag.Parse()
		Conf = new(Config)
		Conf.initLianLianPay()
	})
	return Conf
}

```

### 业务调取，参数组装

```cgo
//支付调取
func callLianLianPay(goodsName, tradeNum, expire string, userId int, payMoney float64) (map[string]string, error) {
	result := make(map[string]string)
	timeStamp := utils.TimestampToDatetimeNoStr(time.Now().Unix())
	//初始化参数Map
	bm := make(pay_lianlian.BodyMap)
	bm.Set("api_version", "1.0").
		Set("time_stamp", timeStamp).
		Set("platform", "").
		Set("user_id", userId).
		Set("busi_partner", pay_lianlian.BusiPartnerVirtual).
		Set("no_order", tradeNum).
		Set("dt_order", timeStamp).
		Set("name_goods", goodsName).
		Set("money_order", payMoney).
		Set("notify_url", config.Conf.Yaml.PayLianLian.Notifyurl).
		Set("risk_item", fmt.Sprintf("{\"user_info_bind_phone\":\"19933610000\",\"user_info_dt_register\":\"%s\",\"frms_ware_category \":\"1009\"}", timeStamp)).
		Set("flag_pay_product", "0").
		Set("flag_chnl", pay_lianlian.FlagChnlH5).
		Set("bank_code", "0100000")
	//Set("card_type", "")
	_, err := config.Conf.LianLianPayClient.UnifiedOrder(context.Background(), bm)
	if err != nil {
		return result, err
	}
	return result, nil
}

```

### RSAWithMd5 签名方式

+参考文档
https://github.com/wenzhenxi/gorsa/blob/15feec0f05a6feb896e30fbf3c9df5d77fa1a421/gorsaSign.go