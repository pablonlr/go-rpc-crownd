package crownd

import "encoding/json"

type GetRawTxResponse struct {
	TxID     string              `json:"txid"`
	Version  int                 `json:"version"`
	Locktime int                 `json:"locktime"`
	Vin      []*TxInputResponse  `json:"vin"`
	Vout     []*TxOutputResponse `json:"vout"`
}

type TxInputResponse struct {
	TxID      string             `json:"txid"`
	Vout      int                `json:"vout"`
	ScriptSig *ScriptSigResponse `json:"scriptSig"`
	Sequence  int                `json:"sequence"`
}

type ScriptSigResponse struct {
	Asm string `json:"asm"`
	Hex string `json:"hex"`
}

type TxOutputResponse struct {
	Value       float64         `json:"value"`
	N           int             `json:"n"`
	SriptPubKey *TxOutputScript `json:"scriptPubKey"`
}

type TxOutputScript struct {
	Asm       string   `json:"asm"`
	Hex       string   `json:"hex"`
	ReqSigs   int      `json:"reqSigs"`
	Type      string   `json:"type"`
	Addresses []string `json:"addresses"`
}

func (client *Client) SendRawTx(txHex string) (string, error) {
	resp, err := client.Request("sendrawtransaction", txHex)
	if resperr := parseErr(err, resp); resperr != nil {
		return "", resperr
	}
	var s string
	err = json.Unmarshal(resp.Result, &s)
	return s, nil
}

func (client *Client) GetRawTransaction(txid string) (string, error) {
	resp, err := client.Request("getrawtransaction", txid)
	if resperr := parseErr(err, resp); resperr != nil {
		return "", resperr
	}
	var hex string
	err = json.Unmarshal(resp.Result, &hex)
	return hex, nil
}

func (client *Client) GetRawTransactionDecoded(txid string) (*GetRawTxResponse, error) {
	resp, err := client.Request("getrawtransaction", txid, 1)
	if resperr := parseErr(err, resp); resperr != nil {
		return nil, resperr
	}
	tx := GetRawTxResponse{}
	err = json.Unmarshal(resp.Result, &tx)
	return &tx, nil
}

func (client *Client) DecodeRawTransaction(hex string) (*GetRawTxResponse, error) {
	resp, err := client.Request("decoderawtransaction", hex)
	if resperr := parseErr(err, resp); resperr != nil {
		return nil, resperr
	}
	tx := GetRawTxResponse{}
	err = json.Unmarshal(resp.Result, &tx)
	return &tx, nil

}

func (client *Client) SignRawTx(hex string) (string, error) {
	//in development :)
	var err error
	return "", err
}
