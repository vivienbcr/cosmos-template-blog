package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type LikePost struct {
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	ID      string         `json:"id" yaml:"id"`
	Like    string         `json:"like" yaml:"like"`
	PostID  string         `json:"postID" yaml:"postID"`
}
