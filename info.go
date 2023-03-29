package crownd

import (
	"encoding/json"
)

type GetInfoResponse struct {
	Version         float64         `json:"version"`
	ProtocolVersion float64         `json:"protocolversion"`
	WalletVersion   float64         `json:"walletversion"`
	Balance         float64         `json:"balance"`
	Blocks          int             `json:"blocks"`
	TimeOffset      int             `json:"timeoffset"`
	Connections     int             `json:"connections"`
	Proxy           string          `json:"proxy"`
	Difficulty      float64         `json:"difficulty"`
	Tesnet          bool            `json:"testnet"`
	StakingActive   bool            `json:"staking_active"`
	KeyPoolOldest   float64         `json:"keypoololdest"`
	KeyPoolSize     float64         `json:"keypoolsize"`
	UnlockedUntil   float64         `json:"unlocked_until"`
	PayTxFee        float64         `json:"paytxfee"`
	RelayFee        float64         `json:"relayfee"`
	Errors          json.RawMessage `json:"errors"`
}

func (client *Client) GetInfo() (*GetInfoResponse, *CrownError) {
	resp, err := client.Request("getinfo")
	if resperr := parseErr(err, resp); resperr != nil {
		return nil, resperr
	}
	getinforesp := &GetInfoResponse{}
	err = json.Unmarshal(resp.Result, getinforesp)
	if err != nil {
		return nil, newCrownErrorFromError(err)
	}
	return getinforesp, nil
}
