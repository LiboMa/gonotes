package main

import (
	//	"backend/support/libraries/loggers"
	"bytes"
	"crypto"
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"time"
)

const (
	ApiUrl = "https://devapi.bbx.com/v1/cloud/gateway"
	//ApiUrl             = "http://127.0.0.1:9100/gateway"
	AppId              = "2017184040" // 测试的券商uid
	OriginSubAccountId = "10006430"   //券商自己账户体系下的某个用户的id

)

// 券商自己的私钥，每个券商自己生成
var privateKey = []byte(`
-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQDZsfv1qscqYdy4vY+P4e3cAtmvppXQcRvrF1cB4drkv0haU24Y
7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0DgacdwYWd/7PeCELyEipZJL07Vro7
Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NLAUeJ6PeW+DAkmJWF6QIDAQAB
AoGBAJlNxenTQj6OfCl9FMR2jlMJjtMrtQT9InQEE7m3m7bLHeC+MCJOhmNVBjaM
ZpthDORdxIZ6oCuOf6Z2+Dl35lntGFh5J7S34UP2BWzF1IyyQfySCNexGNHKT1G1
XKQtHmtc2gWWthEg+S6ciIyw2IGrrP2Rke81vYHExPrexf0hAkEA9Izb0MiYsMCB
/jemLJB0Lb3Y/B8xjGjQFFBQT7bmwBVjvZWZVpnMnXi9sWGdgUpxsCuAIROXjZ40
IRZ2C9EouwJBAOPjPvV8Sgw4vaseOqlJvSq/C/pIFx6RVznDGlc8bRg7SgTPpjHG
4G+M3mVgpCX1a/EU1mB+fhiJ2LAZ/pTtY6sCQGaW9NwIWu3DRIVGCSMm0mYh/3X9
DAcwLSJoctiODQ1Fq9rreDE5QfpJnaJdJfsIJNtX1F+L3YceeBXtW0Ynz2MCQBI8
9KP274Is5FkWkUFNKnuKUK4WKOuEXEO+LpR+vIhs7k6WQ8nGDd4/mujoJBr5mkrw
DPwqA3N5TMNDQVGv8gMCQQCaKGJgWYgvo3/milFfImbp+m7/Y3vCptarldXrYQWO
AQjxwc71ZGBFDITYvdgJM1MTqc8xQek1FXn1vfpy2c6O
-----END RSA PRIVATE KEY-----
`)

// 合约云的公钥（测试用）
var publicKey = []byte(`
-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDZsfv1qscqYdy4vY+P4e3cAtmv
ppXQcRvrF1cB4drkv0haU24Y7m5qYtT52Kr539RdbKKdLAM6s20lWy7+5C0Dgacd
wYWd/7PeCELyEipZJL07Vro7Ate8Bfjya+wltGK9+XNUIHiumUKULW4KDx21+1NL
AUeJ6PeW+DAkmJWF6QIDAQAB
-----END PUBLIC KEY-----
`)

//////////////////////////////////////

func originStr(params url.Values) string {
	var (
		keys   []string
		oriStr string
	)
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, k := range keys {
		oriStr = fmt.Sprintf("%s%s", oriStr, params.Get(k))
	}
	fmt.Println("Ori Str: ", oriStr)
	return oriStr
}

func genSignature(params url.Values) (string, error) {
	var (
		keys []string
	)
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "", errors.New("public key error")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	buff := bytes.NewBuffer(nil)
	for _, k := range keys {
		buff.WriteString(params.Get(k))
	}
	hashed := sha256.Sum256(buff.Bytes())
	encrypted, err := rsa.SignPKCS1v15(crand.Reader, priv, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func checkSign(r url.Values) error {
	var (
		keys []string
	)
	fmt.Println(r.Get("signature"))
	signature, err := base64.StdEncoding.DecodeString(r.Get("signature"))
	if err != nil {
		fmt.Println(err)
		return err
	}

	for key := range r {
		if key == "signature" {
			continue
		}
		keys = append(keys, key)
	}
	sort.Strings(keys)
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	buff := bytes.NewBuffer(nil)
	pub := pubInterface.(*rsa.PublicKey)
	for _, k := range keys {
		buff.WriteString(r.Get(k))
	}
	hashed := sha256.Sum256(buff.Bytes())
	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], []byte(signature))
}

func verifySignature(origData, sign []byte) error {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)
	hashed := sha256.Sum256(origData)

	return rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed[:], sign)
}

func callCloudApi(params url.Values) error {
	var (
		trans = &http.Transport{
			MaxIdleConns:       4,
			IdleConnTimeout:    time.Second * 30,
			DisableCompression: true,
		}
		client = http.Client{
			Transport: trans,
			Timeout:   time.Second * 360,
		}
	)
	params.Add("timestamp", fmt.Sprintf("%d", time.Now().Unix()))
	params.Add("nonce", fmt.Sprintf("%d", 100000000000+rand.Int63n(89999999999)))
	params.Add("version", "v1")

	signature, err := genSignature(params)
	params.Add("signature", signature)

	if err := checkSign(params); err != nil {
		fmt.Println(err)
		return err
	}
	if err != nil {
		return err
	}

	fmt.Println(params)
	req, err := http.NewRequest(http.MethodPost, ApiUrl, bytes.NewReader([]byte(params.Encode())))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		err = resp.Body.Close()
	}()

	buff := bytes.NewBuffer(nil)
	if _, err = buff.ReadFrom(resp.Body); err != nil {
		return err
	}

	fmt.Println(buff.String())
	buff.WriteString(resp.Header.Get("Bbx-Nonce"))
	buff.WriteString(resp.Header.Get("Bbx-Ts"))
	sign, err := base64.StdEncoding.DecodeString(resp.Header.Get("Bbx-Signature"))
	if err != nil {
		return nil
	}

	if err := verifySignature(buff.Bytes(), sign); err != nil {
		fmt.Println("Verify sign error:")
	} else {
		fmt.Println("Verify sign OK")
	}

	return nil
}

//////////////////////////////////////

func CreateCloudAccount() {
	params := make(url.Values)
	rand.Seed(time.Now().UnixNano())
	params.Add("app_id", AppId)
	params.Add("method", "account.create")
	params.Add("api_key_life_span", fmt.Sprintf("%d", 24*3600*30))
	params.Add("origin_uid", fmt.Sprintf("%d", 10000000+rand.Int63n(10000)))
	if err := callCloudApi(params); err != nil {
		//loggers.Error.Println(err)
	}
}

func AccountFreeze() {
	params := make(url.Values)
	rand.Seed(time.Now().UnixNano())
	params.Add("app_id", AppId)
	params.Add("method", "account.freeze")
	params.Add("origin_uid", OriginSubAccountId)
	if err := callCloudApi(params); err != nil {
		//loggers.Error.Println(err)
	}
}

func AccountUnfreeze() {
	params := make(url.Values)
	rand.Seed(time.Now().UnixNano())
	params.Add("app_id", AppId)
	params.Add("method", "account.unfreeze")
	params.Add("origin_uid", OriginSubAccountId)
	if err := callCloudApi(params); err != nil {
		//loggers.Error.Println(err)
	}
}

func AccountApiKeyUpdate() {
	params := make(url.Values)
	rand.Seed(time.Now().UnixNano())
	params.Add("app_id", AppId)
	params.Add("method", "account.api_key.update")
	params.Add("origin_uid", OriginSubAccountId)
	params.Add("api_key_life_span", fmt.Sprintf("%d", 24*3600*30))
	if err := callCloudApi(params); err != nil {
		//loggers.Error.Println(err)
	}
}

func AccountApiKeyQuery() {
	params := make(url.Values)
	rand.Seed(time.Now().UnixNano())
	params.Add("app_id", AppId)
	params.Add("method", "account.api_key.query")
	params.Add("origin_uid", OriginSubAccountId)
	if err := callCloudApi(params); err != nil {
		//loggers.Error.Println(err)
	}
}

func AccountAssetTransfer() {
	params := make(url.Values)
	rand.Seed(time.Now().UnixNano())
	params.Add("app_id", AppId)
	params.Add("method", "account.asset.transfer")
	params.Add("origin_uid", OriginSubAccountId)
	params.Add("vol", "100")
	params.Add("coin_code", "BTC")
	params.Add("out_trade_no", "btctest001")
	if err := callCloudApi(params); err != nil {
		//loggers.Error.Println(err)
	}
}

func AccountAssetTransferOut() {
	params := make(url.Values)
	rand.Seed(time.Now().UnixNano())
	params.Add("app_id", AppId)
	params.Add("method", "account.asset.transferout")
	params.Add("origin_uid", OriginSubAccountId)
	params.Add("vol", "100")
	params.Add("coin_code", "BTC")
	params.Add("out_trade_no", "btctest002")
	if err := callCloudApi(params); err != nil {
		//loggers.Error.Println(err)
	}
}

func AccountAssetQuery() {
	params := make(url.Values)
	rand.Seed(time.Now().UnixNano())
	params.Add("app_id", AppId)
	params.Add("method", "account.asset.query")
	params.Add("origin_uid", OriginSubAccountId)
	params.Add("coin_code", "")
	if err := callCloudApi(params); err != nil {
		//loggers.Error.Println(err)
	}
}

func AccountQueryTradeNo() {
	params := make(url.Values)
	rand.Seed(time.Now().UnixNano())
	params.Add("app_id", AppId)
	params.Add("method", "account.tradeno.query")
	params.Add("out_trade_no", "btctest001")
	if err := callCloudApi(params); err != nil {
		//	loggers.Error.Println(err)
	}
}

func AccountQueryOrders() {
	params := make(url.Values)
	rand.Seed(time.Now().UnixNano())
	params.Add("app_id", AppId)
	params.Add("method", "account.orders.query")
	params.Add("origin_uid", OriginSubAccountId)
	params.Add("status", "0") //0: All; 1: 审批中; 2: 委托中; 4: 结束
	params.Add("offset", "0")
	params.Add("size", "100")
	params.Add("contract_id", "6")
	if err := callCloudApi(params); err != nil {
		//loggers.Error.Println(err)
	}
}

func AccountQueryPositions() {
	params := make(url.Values)
	rand.Seed(time.Now().UnixNano())
	params.Add("app_id", AppId)
	params.Add("method", "account.positions.query")
	params.Add("origin_uid", OriginSubAccountId)
	params.Add("status", "0") //0: All; 1: 持仓中; 2: 系统代持中; 4: 已平仓
	params.Add("offset", "0")
	params.Add("size", "100")
	params.Add("coin_code", "")
	if err := callCloudApi(params); err != nil {
		//loggers.Error.Println(err)
	}
}

func main() {
	fmt.Println("Start --------")
	//CreateCloudAccount()
	//AccountFreeze()
	//AccountUnfreeze()
	//AccountApiKeyUpdate()
	//AccountApiKeyQuery()
	//AccountAssetTransfer()
	//AccountAssetTransferOut()
	//AccountQueryTradeNo()
	//AccountAssetQuery()
	AccountQueryOrders()
	//AccountQueryPositions()
	//signTest()
	fmt.Println("End --------")
}

func signTest() {
	params := make(url.Values)
	params.Add("origin_uid", "1")
	params.Add("apikeylife_span", "1000")
	params.Add("method", "account.create")
	params.Add("app_id", "2244248428")
	params.Add("nonce", "eqfdsqmk7qtsf42k")
	params.Add("version", "v1")
	params.Add("timestamp", "1566466346")
	params.Add("signature", "PpJ9tuKSmIbsurGbUnTWs5ORxalggD9mtQXLrgxrfiNCULhurE7C1UlGWXd2YNr1ciNsL+fHrSzTkK4cOb3adTrxUgnPkFMSK6NlgPp16bPJeN++/4pcWvILOfWb8yZ42prFIKkaPJeREowwf3Qt6f+B8ugjrIZTKnBdg4aHcjs=")

	//resp, _ := http.PostForm("http://localhost:9100/gateway", params)
	resp, e := http.PostForm("https://devapi.bbx.com/v1/cloud/gateway", params)
	if e != nil {
		fmt.Println(e)
		return
	}

	all, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		fmt.Println(e)
		return
	}
	fmt.Println(string(all))
}
