package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Contact struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
	Pseudo  string         `json:"pseudo" yaml:"pseudo"`
}
