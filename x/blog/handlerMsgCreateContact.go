package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/sdk-tutorials/blog/blog/x/blog/keeper"
	"github.com/cosmos/sdk-tutorials/blog/blog/x/blog/types"
)

func handleMsgCreateContact(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateContact) (*sdk.Result, error) {
	var Contact = types.Contact{
		ID:      msg.ID,
		Creator: msg.Creator,
		Pseudo:  msg.Pseudo,
	}
	k.CreateContact(ctx, Contact)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
