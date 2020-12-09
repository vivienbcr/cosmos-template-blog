package keeper

import (
	// this line is used by starport scaffolding
	"github.com/cosmos/sdk-tutorials/blog/blog/x/blog/types"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewQuerier creates a new querier for blog clients.
func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryListComment:
			return listComment(ctx, k)
		case types.QueryListPost:
			return listPost(ctx, k)
		case types.QueryListLikeComment:
			return listLikeComment(ctx, k)
		case types.QueryListLikePost:
			return listLikePost(ctx, k)
		case types.QueryListContact:
			return listContact(ctx, k)
		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown blog query endpoint")
		}
	}
}