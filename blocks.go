package crownd

import "encoding/json"

type GetBlockResponse struct {
	Hash              string   `json:"hash"`
	Confirmations     int      `json:"confirmations"`
	Size              int      `json:"size"`
	Height            int      `json:"height"`
	Version           int      `json:"version"`
	Merkleroot        string   `json:"merkleroot"`
	TX                []string `json:"tx"`
	Time              int64    `json:"time"`
	Nonce             int      `json:"nonce"`
	Bits              string   `json:"bits"`
	Difficulty        float64  `json:"difficulty"`
	PreviousBlockHash string   `json:"previousblockhash"`
	NextBlockHash     string   `json:"nextblockhash"`
}

func (client *Client) GetBlock(hash string) (*GetBlockResponse, *CrownError) {
	resp, err := client.Request("getblock", hash)
	if resperr := parseErr(err, resp); resperr != nil {
		return nil, resperr
	}
	getblockresp := &GetBlockResponse{}
	err = json.Unmarshal(resp.Result, getblockresp)
	if err != nil {
		return nil, newCrownErrorFromError(err)
	}
	return getblockresp, nil
}

func (client *Client) GetBlockCount() (int, *CrownError) {
	resp, err := client.Request("getblockcount")
	if resperr := parseErr(err, resp); resperr != nil {
		return 0, resperr
	}
	blockcount := 0
	err = json.Unmarshal(resp.Result, &blockcount)
	if err != nil {
		return 0, newCrownErrorFromError(err)
	}
	return blockcount, nil
}
