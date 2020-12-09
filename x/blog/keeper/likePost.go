package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/sdk-tutorials/blog/blog/x/blog/types"
)

func (k Keeper) CreateLikePost(ctx sdk.Context, likePost types.LikePost) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.LikePostPrefix + likePost.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(likePost)
	store.Set(key, value)
}

func listLikePost(ctx sdk.Context, k Keeper) ([]byte, error) {
	var likePostList []types.LikePost
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.LikePostPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var likePost types.LikePost
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &likePost)
		likePostList = append(likePostList, likePost)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, likePostList)
	return res, nil
}
