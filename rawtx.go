package crownd

import "encoding/json"

func (client *Client) SendRawTx(txHex string) (string, error) {
	resp, err := client.Request("sendrawtransaction", txHex)
	if resperr := parseErr(err, resp); resperr != nil {
		return "", resperr
	}
	var s string
	err = json.Unmarshal(resp.Result, &s)
	return s, nil
}
