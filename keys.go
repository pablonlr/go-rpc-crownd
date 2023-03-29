package crownd

func (client *Client) Unlock(password string, seconds int) *CrownError {
	resp, err := client.Request("walletpassphrase", password, seconds)
	if resperr := parseErr(err, resp); resperr != nil {
		return resperr
	}
	return nil
}
