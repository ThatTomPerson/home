package coinspot

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

func New(key, secret string) *Coinspot {
	return &Coinspot{
		key:    key,
		secret: []byte(secret),
	}
}

type Spot struct {
	LTC  decimal.Decimal `json:"ltcspot"`
	BTC  decimal.Decimal `json:"btcspot"`
	Doge decimal.Decimal `json:"dogespot"`
}

type Coinspot struct {
	key    string
	secret []byte
}

func (c *Coinspot) Spot() (*Spot, error) {
	res, err := c.request("/api/spot", nil)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var s struct {
		status string
		Spot   Spot
	}
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		return nil, err
	}

	return &s.Spot, nil
}

func addNonce(body []byte) ([]byte, error) {
	x := make(map[string]interface{})
	if body != nil {
		json.Unmarshal(body, &x)
	}
	x["nonce"] = time.Now().UnixNano()
	body, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("could not add nonce to body %v", err)
	}

	return body, nil
}

func (c *Coinspot) request(path string, body []byte) (*http.Response, error) {
	body, err := addNonce(body)
	if err != nil {
		return nil, err
	}

	signer := hmac.New(sha512.New, c.secret)
	signer.Write(body)
	mac := signer.Sum(nil)

	req, err := http.NewRequest("POST", fmt.Sprintf("https://www.coinspot.com.au%s", path), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}

	len := hex.EncodedLen(len(mac))

	sign := make([]byte, len)
	hex.Encode(sign, mac)

	req.Header.Add("sign", string(sign))
	req.Header.Add("key", c.key)
	req.Header.Add("Content-Type", "application/json")
	log.Println("doing request")
	return http.DefaultClient.Do(req)
}
