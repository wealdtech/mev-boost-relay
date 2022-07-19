package common

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/flashbots/go-boost-utils/types"
)

// BuilderEntry represents a builder that is allowed to send blocks
// Address will be schema://hostname:port
type BuilderEntry struct {
	Address string
	Pubkey  hexutil.Bytes
	URL     *url.URL
}

// NewBuilderEntry creates a new instance based on an input string
// builderURL can be IP@PORT, PUBKEY@IP:PORT, https://IP, etc.
func NewBuilderEntry(builderURL string) (entry *BuilderEntry, err error) {
	if !strings.HasPrefix(builderURL, "http") {
		builderURL = "http://" + builderURL
	}

	url, err := url.Parse(builderURL)
	if err != nil {
		return entry, err
	}

	entry = &BuilderEntry{
		URL:     url,
		Address: entry.URL.Scheme + "://" + entry.URL.Host,
	}
	err = entry.Pubkey.UnmarshalText([]byte(entry.URL.User.Username()))
	return entry, err
}

type EthNetworkDetails struct {
	Name                     string
	GenesisForkVersionHex    string
	GenesisValidatorsRootHex string
	BellatrixForkVersionHex  string
}

func NewEthNetworkDetails(networkName string) (ret *EthNetworkDetails, err error) {
	ret = &EthNetworkDetails{
		Name: networkName,
	}
	switch networkName {
	case "kiln":
		ret.GenesisForkVersionHex = types.GenesisForkVersionKiln
		ret.GenesisValidatorsRootHex = types.GenesisValidatorsRootKiln
		ret.BellatrixForkVersionHex = types.BellatrixForkVersionKiln
	case "ropsten":
		ret.GenesisForkVersionHex = types.GenesisForkVersionRopsten
		ret.GenesisValidatorsRootHex = types.GenesisValidatorsRootRopsten
		ret.BellatrixForkVersionHex = types.BellatrixForkVersionRopsten
	case "sepolia":
		ret.GenesisForkVersionHex = types.GenesisForkVersionSepolia
		ret.GenesisValidatorsRootHex = types.GenesisValidatorsRootSepolia
		ret.BellatrixForkVersionHex = types.BellatrixForkVersionSepolia
	default:
		return nil, fmt.Errorf("unknown network: %s", networkName)
	}
	return ret, nil
}
