package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type LikeComment struct {
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	ID        string         `json:"id" yaml:"id"`
	Like      string         `json:"like" yaml:"like"`
	CommentID string         `json:"commentId" yaml:"commentId"`
}
