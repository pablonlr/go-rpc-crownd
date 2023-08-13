package crownd

import (
	"encoding/json"
	"strconv"
)

// NFToken represents the information of a Non-Fungible Token in a Crown logic
type NFToken struct {
	ID                 string `json:"nftId"`
	NFTProtocolID      string `json:"nftProtocolId"`
	OwnerKeyID         string `json:"nftOwnerKeyId"`
	MetadataAdminKeyId string `json:"metadataAdminKeyId"`
	Metadata           string `json:"metadata"`
	BlockHash          string `json:"blockHash"`
	RegistrationTxHash string `json:"registrationTxHash"`
	Height             int    `json:"height"`
	Timestamp          uint64 `json:"timestamp"`
}

// RegisterNFToken Register a new NFT in a given protocol.
// Creates and sends a new non-fungible token transaction
func (client *Client) RegisterNFToken(nftProtoID string, ID string, nfTokenOwnerAddr string, metadatAdminAddr string, metadata string) (string, *CrownError) {
	resp, err := client.Request("nftoken", "register", nftProtoID, ID, nfTokenOwnerAddr, metadatAdminAddr, metadata)
	if resperr := parseErr(err, resp); resperr != nil {
		return "", resperr
	}
	return string(resp.Result), nil
}

// ListNFTokensExplicit 	List tokens registered on the blockchain by:
// <protocol> <owner address>, <number of elements to list>, <numb of txs to skip from tip>,<to a specific block height or "*" for current height>
// Return a slice of NFToken struct
func (client *Client) ListNFTokensExplicit(nftproto string, nftOwnerAddr string, count, skipFromTip int, toHeight string) ([]NFToken, *CrownError) {
	resp, err := client.Request("nftoken", "list", nftproto, nftOwnerAddr, strconv.Itoa(count), strconv.Itoa(skipFromTip), toHeight)
	if resperr := parseErr(err, resp); resperr != nil {
		return nil, resperr
	}
	tokens := []NFToken{}
	err = json.Unmarshal(resp.Result, &tokens)
	if err != nil {
		return nil, newCrownErrorFromError(err)
	}
	return tokens, nil
}

// ListLastNFTokens List the last registered tokens on the Crown blockchain.
// Receives the number of tokens to list.
// Return a slice of NFToken struct
func (client *Client) ListLastNFTokens(count int) ([]NFToken, *CrownError) {
	return client.ListLastNFTokensInProtocol("*", count)
}

// ListLastNFTokensInProtocol List the last registered tokens on a given protocol.
// Receives the protocol unique ID, and the number of tokens to list.
// Return a slice of NFToken struct
func (client *Client) ListLastNFTokensInProtocol(nftproto string, count int) ([]NFToken, *CrownError) {
	return client.ListLastNFTokensInProtocolOwnerBy(nftproto, "*", count)
}

// ListLastNFTokensOwnerBy List the last registered tokens owner by a given CRW address.
// Receives the CRW owner address, and the number of tokens to list.
// Return a slice of NFToken struct
func (client *Client) ListLastNFTokensOwnerBy(ownerAddr string, count int) ([]NFToken, *CrownError) {
	return client.ListLastNFTokensInProtocolOwnerBy("*", ownerAddr, count)
}

// ListLastNFTokensInProtocolOwnerBy List the last registered tokens on a given protocol, owner by a CRW address.
// Receives the protocol unique ID, the CRW owner address and the number of tokens to list
// Return a slice of NFToken struct
func (client *Client) ListLastNFTokensInProtocolOwnerBy(nftproto, ownerAddr string, count int) ([]NFToken, *CrownError) {
	return client.ListNFTokensExplicit(nftproto, ownerAddr, count, 0, "*")
}

// GetNFToken Obtain a registered NFT by a given protocol and unique NFToken ID.
// Return a pointer to the NFToken representation.
func (client *Client) GetNFToken(protocolID, tokenID string) (*NFToken, *CrownError) {
	resp, err := client.Request("nftoken", "get", protocolID, tokenID)
	if resperr := parseErr(err, resp); resperr != nil {
		return nil, resperr
	}
	token := NFToken{}
	err = json.Unmarshal(resp.Result, &token)
	if err != nil {
		return nil, newCrownErrorFromError(err)
	}
	return &token, nil
}

// GetNFTokenByTxID Obtain a registered NFT by the registration transaction ID.
// Return a pointer to the NFToken representation.
func (client *Client) GetNFTokenByTxID(txID string) (*NFToken, *CrownError) {
	resp, err := client.Request("nftoken", "getbytxid", txID)
	if resperr := parseErr(err, resp); resperr != nil {
		return nil, resperr
	}
	token := NFToken{}
	err = json.Unmarshal(resp.Result, &token)
	if err != nil {
		return nil, newCrownErrorFromError(err)
	}
	return &token, nil
}

// TotalSupply Number of NFTokens registered on a protocol. Require the unique NFT protocol ID.
func (client *Client) TotalSupply(protocolID string) (int, *CrownError) {
	resp, err := client.Request("nftoken", "totalsupply", protocolID)
	if resperr := parseErr(err, resp); resperr != nil {
		return -1, resperr
	}
	supply := 0
	err = json.Unmarshal(resp.Result, &supply)
	if err != nil {
		return -1, newCrownErrorFromError(err)
	}
	return supply, nil
}

// GlobalSupply Total of NFTokens registered on Crown NFT Framework
func (client *Client) GlobalSupply() (int, *CrownError) {
	resp, err := client.Request("nftoken", "totalsupply")
	if resperr := parseErr(err, resp); resperr != nil {
		return -1, resperr
	}
	supply := 0
	err = json.Unmarshal(resp.Result, &supply)
	if err != nil {
		return -1, newCrownErrorFromError(err)
	}
	return supply, nil
}

// BalanceOf Number of NFTokens owned by a given CRW address.
func (client *Client) BalanceOf(address string) (int, *CrownError) {
	return client.balanceOf(address)
}

// BalanceOf Number of NFTokens owned by a given CRW address in a specific protocol.
// Receives the  the CRW owner address and the protocol unique ID.
func (client *Client) BalanceOfInProto(address, nfproto string) (int, *CrownError) {
	return client.balanceOf(address, nfproto)
}

func (client *Client) balanceOf(params ...string) (int, *CrownError) {
	resp, err := client.Request("nftoken", "balanceof", params)
	if resperr := parseErr(err, resp); resperr != nil {
		return -1, resperr
	}
	balance := 0
	err = json.Unmarshal(resp.Result, &balance)
	if err != nil {
		return -1, newCrownErrorFromError(err)
	}
	return balance, nil
}

// OwnerOfToken Get the CRW address owner of a given NFT.
// Receives protocol ID and the NFToken ID.
func (client *Client) OwnerOfToken(nfprotoID, nftokenID string) (string, *CrownError) {
	resp, err := client.Request("nftoken", "ownerof", nfprotoID, nftokenID)
	if resperr := parseErr(err, resp); resperr != nil {
		return "", resperr
	}
	owner := ""
	err = json.Unmarshal(resp.Result, &owner)
	if err != nil {
		return "", newCrownErrorFromError(err)
	}
	return owner, nil
}
