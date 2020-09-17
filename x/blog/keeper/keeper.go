package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/linus1501/blog/x/blog/types"
)

// Keeper of the blog store
type Keeper struct {
	CoinKeeper bank.Keeper
	storeKey   sdk.StoreKey
	cdc        *codec.Codec
	// paramspace types.ParamSubspace
}

// NewKeeper creates a blog keeper
func NewKeeper(coinKeeper bank.Keeper, cdc *codec.Codec, key sdk.StoreKey) Keeper {
	keeper := Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   key,
		cdc:        cdc,
		// paramspace: paramspace.WithKeyTable(types.ParamKeyTable()),
	}
	return keeper
}

// Logger returns a module-specific logger.
func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) CreatePost(ctx sdk.Context, post types.Post) {
	store := ctx.KVStore(k.storeKey)
	key := []byte(types.PostPrefix + post.ID)
	value := k.cdc.MustMarshalBinaryLengthPrefixed(post)
	store.Set(key, value)
}
