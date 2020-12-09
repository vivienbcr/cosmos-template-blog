package blog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/sdk-tutorials/blog/blog/x/blog/keeper"
	"github.com/cosmos/sdk-tutorials/blog/blog/x/blog/types"
)

func handleMsgCreateLikeComment(ctx sdk.Context, k keeper.Keeper, msg types.MsgCreateLikeComment) (*sdk.Result, error) {
	var likeComment = types.LikeComment{
		Creator:   msg.Creator,
		ID:        msg.ID,
		Like:      msg.Like,
		CommentID: msg.CommentID,
	}
	k.CreateLikeComment(ctx, likeComment)

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
