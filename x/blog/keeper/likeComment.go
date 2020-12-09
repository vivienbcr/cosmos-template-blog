package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/sdk-tutorials/blog/blog/x/blog/types"
)

func (k Keeper) CreateLikeComment(ctx sdk.Context, likeComment types.LikeComment) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.LikePostPrefix + likeComment.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(likeComment)
	store.Set(key, value)
}

func listLikeComment(ctx sdk.Context, k Keeper) ([]byte, error) {
	var likeCommentList []types.LikePost
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.LikePostPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var likeComment types.LikePost
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &likeComment)
		likeCommentList = append(likeCommentList, likeComment)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, likeCommentList)
	return res, nil
}
