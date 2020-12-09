package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/sdk-tutorials/blog/blog/x/blog/keeper"
	"github.com/cosmos/sdk-tutorials/blog/blog/x/blog/types"
)

func handleMsgCreateLikePost(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateLikePost) (*sdk.Result, error) {
	var likePost = types.LikePost{
		Creator: msg.Creator,
		ID:      msg.ID,
		Like:    msg.Like,
		PostID:  msg.PostID,
	}
	k.CreateLikePost(ctx, likePost)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
