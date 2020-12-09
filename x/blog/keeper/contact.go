package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/sdk-tutorials/blog/blog/x/blog/types"
)

func (k Keeper) CreateContact(ctx sdk.Context, contact types.Contact) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.ContactPrefix + contact.Creator.String())
	value := k.cdc.MustMarshalBinaryLengthPrefixed(contact)
	store.Set(key, value)
}

func listContact(ctx sdk.Context, k Keeper) ([]byte, error) {
	var ContactList []types.Contact
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(types.ContactPrefix))
	for ; iterator.Valid(); iterator.Next() {
		var contact types.Contact
		k.cdc.MustUnmarshalBinaryLengthPrefixed(store.Get(iterator.Key()), &contact)
		ContactList = append(ContactList, contact)
	}
	res := codec.MustMarshalJSONIndent(k.cdc, ContactList)
	return res, nil
}
