package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateLikePost{}

type MsgCreateLikePost struct {
	ID      string
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Like    string         `json:"like" yaml:"like"`
	PostID  string         `json:"postID" yaml:"postID"`
}

func NewMsgCreateLikePost(creator sdk.AccAddress, like string, postID string) MsgCreateLikePost {
	return MsgCreateLikePost{
		ID:      uuid.New().String(),
		Creator: creator,
		PostID:  postID,
		Like:    like,
	}
}

func (msg MsgCreateLikePost) Route() string {
	return RouterKey
}

func (msg MsgCreateLikePost) Type() string {
	return "CreateLikePost"
}

func (msg MsgCreateLikePost) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateLikePost) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateLikePost) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
