package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
)

var _ sdk.Msg = &MsgCreateContact{}

type MsgCreateContact struct {
	ID      string
	Creator sdk.AccAddress `json:"creator" yaml:"creator"`
	Pseudo  string         `json:"pseudo" yaml:"pseudo"`
}

func NewMsgCreateContact(creator sdk.AccAddress, pseudo string) MsgCreateContact {
	return MsgCreateContact{
		ID:      uuid.New().String(),
		Creator: creator,
		Pseudo:  pseudo,
	}
}

func (msg MsgCreateContact) Route() string {
	return RouterKey
}

func (msg MsgCreateContact) Type() string {
	return "CreateContact"
}

func (msg MsgCreateContact) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Creator)}
}

func (msg MsgCreateContact) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateContact) ValidateBasic() error {
	if msg.Creator.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator can't be empty")
	}
	return nil
}
