package clients

import (
	"fmt"

	"github.com/ethereum/go-ethereum/core/types"
)

// Describe a node setup, which consists of:

// - Bundler Client
type BundlerDefinition struct {
	BundlerClient string `json:"bundler_client"`

	// Bundler Config
	Chain []*types.Block `json:"chain,omitempty"`
}

func (n *BundlerDefinition) String() string {
	return fmt.Sprintf("%s", n.BundlerClient)
}
