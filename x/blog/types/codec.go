package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on codec
func RegisterCodec(cdc *codec.Codec) {
	// this line is used by starport scaffolding
	cdc.RegisterConcrete(MsgCreateComment{}, "blog/CreateComment", nil)
	cdc.RegisterConcrete(MsgCreatePost{}, "blog/CreatePost", nil)
	cdc.RegisterConcrete(MsgCreateContact{}, "blog/CreateContact", nil)
	cdc.RegisterConcrete(MsgCreateLikeComment{}, "blog/CreateLikeComment", nil)
	cdc.RegisterConcrete(MsgCreateLikePost{}, "blog/CreateLikePost", nil)
}

// ModuleCdc defines the module codec
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
