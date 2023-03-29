package crownd

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// NFTProtocol represents the information of a Non-Fungible Token Protocol in a Crown logic
type NFTProtocol struct {
	ID                     string `json:"nftProtocolId"`
	Name                   string `json:"tokenProtocolName"`
	TokenProtocolOwnerAddr string `json:"tokenProtocolOwnerId"`
	NFTRegSign             string `json:"nftRegSign"`
	MetadataMimeType       string `json:"tokenMetadataMimeType"`
	MetadataSchemaUri      string `json:"tokenMetadataSchemaUri"`
	IsTokenTransferible    bool   `json:"isTokenTransferable"`
	IsMetadataEmbedded     bool   `json:"isMetadataEmbedded"`
	MaxMetadataSize        uint8  `json:"maxMetadataSize"`
	BlockHash              string `json:"blockHash"`
	RegistrationTxHash     string `json:"registrationTxHash"`
	Height                 int    `json:"height"`
	Timestamp              uint64 `json:"timestamp"`
	nftRegCode             NFTSignCode
}

func (proto *NFTProtocol) NFTRegCode() NFTSignCode {
	if proto.nftRegCode != NoValue {
		return proto.nftRegCode
	}
	switch proto.NFTRegSign {
	case fmt.Sprint(SelfSign):
		proto.nftRegCode = SelfSign
	case fmt.Sprint(SignByCreator):
		proto.nftRegCode = SignByCreator
	case fmt.Sprint(SignPayer):
		proto.nftRegCode = SignPayer
	}
	return proto.nftRegCode
}

type NFTSignCode int

const (
	NoValue NFTSignCode = iota
	SelfSign
	SignByCreator
	SignPayer
)

func (client *Client) RegisterNFTProtocol(ID, name, tokenProtoOwnerAddr string, nftRegSign NFTSignCode, metadataMimeType, metadataSchemaUri string, isTokenTransferible, isMetadataEmbedded bool, maxMetadataSize uint8) (string, *CrownError) {
	resp, err := client.Request("nftproto", "register", ID, name, tokenProtoOwnerAddr, strconv.Itoa(int(nftRegSign)), metadataMimeType, metadataSchemaUri, isTokenTransferible, isMetadataEmbedded, strconv.Itoa(int(maxMetadataSize)))
	if resperr := parseErr(err, resp); resperr != nil {
		return "", resperr
	}
	return string(resp.Result), nil
}

// GET NFT Protocol

func (client *Client) GetNFTProtocol(ID string) (*NFTProtocol, *CrownError) {
	return client.getNFTProtocol("get", ID)
}

func (client *Client) GetNFTProtocolByTxID(TxID string) (*NFTProtocol, *CrownError) {
	return client.getNFTProtocol("getbytxid", TxID)
}
func (client *Client) getNFTProtocol(getmethod, ID string) (*NFTProtocol, *CrownError) {
	resp, err := client.Request("nftproto", getmethod, ID)
	if resperr := parseErr(err, resp); resperr != nil {
		return nil, resperr
	}
	getproto := &NFTProtocol{}
	err = json.Unmarshal(resp.Result, getproto)
	if err != nil {
		return nil, newCrownErrorFromError(err)
	}
	return getproto, nil
}

//

//LIST NFT Protocols

func (client *Client) ListNFTProtocolsExplicit(count int, skiptxs int, toHeight string) ([]NFTProtocol, *CrownError) {
	resp, err := client.Request("nftproto", "list", strconv.Itoa(count), strconv.Itoa(skiptxs), toHeight)
	if resperr := parseErr(err, resp); resperr != nil {
		return nil, resperr
	}
	protos := []NFTProtocol{}
	err = json.Unmarshal(resp.Result, &protos)
	if err != nil {
		return nil, newCrownErrorFromError(err)
	}
	return protos, nil

}
func (client *Client) ListNFTProtocolsToBlockN(count int, toBlockHeight int) ([]NFTProtocol, *CrownError) {
	return client.ListNFTProtocolsExplicit(count, 0, strconv.Itoa(toBlockHeight))
}

func (client *Client) ListLastNFTProtocols(count int) ([]NFTProtocol, *CrownError) {
	return client.ListNFTProtocolsExplicit(count, 0, "*")
}

//

func (client *Client) OwnerOfProto(nftProtoID string) (string, *CrownError) {
	resp, err := client.Request("nftproto", "ownerof", nftProtoID)
	if resperr := parseErr(err, resp); resperr != nil {
		return "", resperr
	}
	var s string
	err = json.Unmarshal(resp.Result, &s)
	return s, nil
}
