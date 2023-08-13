package crownd

import "encoding/json"

// MasternodeCount The current count of active masternodes
func (client *Client) MasternodeCount() (int, *CrownError) {
	resp, err := client.Request("masternode", "count")
	if resperr := parseErr(err, resp); resperr != nil {
		return -1, resperr
	}
	count := 0
	err = json.Unmarshal(resp.Result, &count)
	if err != nil {
		return -1, newCrownErrorFromError(err)
	}
	return count, nil
}

// SystemnodeCount The current count of active systemnodes
func (client *Client) SystemnodeCount() (int, *CrownError) {
	resp, err := client.Request("systemnode", "count")
	if resperr := parseErr(err, resp); resperr != nil {
		return -1, resperr
	}
	count := 0
	err = json.Unmarshal(resp.Result, &count)
	if err != nil {
		return -1, newCrownErrorFromError(err)
	}
	return count, nil
}
