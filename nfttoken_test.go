package crownd

import (
	"errors"
	"os"
	"testing"
)

func loadTestEnvVars() (string, string, error) {
	ruser := os.Getenv("RPC_USER")
	if ruser == "" {
		return "", "", errors.New("error reading enviroment var RPC_USER")
	}

	rpass := os.Getenv("RPC_PASS")
	if rpass == "" {
		return "", "", errors.New("error reading enviroment var RPC_PASS")
	}
	return ruser, rpass, nil

}

func TestGetNFToken(t *testing.T) {
	ruser, rpass, err := loadTestEnvVars()
	if err != nil {
		t.Fatal(err)
	}
	c, err := NewClient("localhost", 9341, ruser, rpass, 5000)
	if err != nil {
		t.Fatal(err)
	}
	protocolId := "testperiod"
	tokenId := "612d637bddcdd31467b5c44812f053d7faa30c8f740912fe5d02ff784ef1a9ff"
	nft, err2 := c.GetNFToken(protocolId, tokenId)
	if err2 != nil {
		t.Fatal(err)
	}
	if nft == nil {
		t.Fatal("nft is nil")
	}
	if nft.ID != tokenId {
		t.Fatal("nft.ID != tokenId")
	}
	if nft.NFTProtocolID != protocolId {
		t.Fatal("nft.NFTProtocolID != protocolId")
	}
}
