package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

type UserOp struct {
	Sender               common.Address `json:"sender"               mapstructure:"sender"               validate:"required"`
	Nonce                string         `json:"nonce"                mapstructure:"nonce"                validate:"required"`
	InitCode             string         `json:"initCode"             mapstructure:"initCode"             validate:"required"`
	CallData             string         `json:"callData"             mapstructure:"callData"             validate:"required"`
	CallGasLimit         string         `json:"callGasLimit"         mapstructure:"callGasLimit"         validate:"required"`
	VerificationGasLimit string         `json:"verificationGasLimit" mapstructure:"verificationGasLimit" validate:"required"`
	PreVerificationGas   string         `json:"preVerificationGas"   mapstructure:"preVerificationGas"   validate:"required"`
	MaxFeePerGas         string         `json:"maxFeePerGas"         mapstructure:"maxFeePerGas"         validate:"required"`
	MaxPriorityFeePerGas string         `json:"maxPriorityFeePerGas" mapstructure:"maxPriorityFeePerGas" validate:"required"`
	PaymasterAndData     string         `json:"paymasterAndData"     mapstructure:"paymasterAndData"     validate:"required"`
	Signature            string         `json:"signature"            mapstructure:"signature"            validate:"required"`
}

func loadUserOps() ([]UserOp, error) {
	contents, err := os.ReadFile("userops/userops.json")
	if err != nil {
		panic(fmt.Errorf("can't read userops file: %v", err))
	}
	var userOps []UserOp
	fmt.Println("Contents ", contents)
	if err := json.Unmarshal(contents, &userOps); err != nil {
		panic(fmt.Errorf("can't parse userops JSON: %v", err))
	}

	return userOps, nil
}
