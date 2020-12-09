package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateLikeComment{}

type MsgCreateLikeComment struct {
	ID        string
	Creator   sdk.AccAddress `json:"creator" yaml:"creator"`
	Like      string         `json:"like" yaml:"like"`
	CommentID string         `json:"postID" yaml:"postID"`
}

func NewMsgCreateLikeComment(creator sdk.AccAddress, like string, commentID string) MsgCreateLikeComment {
	return MsgCreateLikeComment{
		ID:        uuid.New().String(),
		Creator:   creator,
		CommentID: commentID,
		Like:      like,
	}
}

func (msg MsgCreateLikeComment) Route() string {
	return RouterKey
}

func (msg MsgCreateLikeComment) Type() string {
	return "CreateLikeComment"
}

func (msg MsgCreateLikeComment) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateLikeComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateLikeComment) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
