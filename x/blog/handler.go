package blog

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/example/blog/x/blog/types"
	"github.com/linus1501/blog/x/blog/keeper"
	"github.com/linus1501/blog/x/blog/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgCreatePost:
			return handleMsgCreatePost(ctx, k, msg)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleMsgCreatePost(ctx sdk.Context, k Keeper, msg types.MsgCreatePost) (*sdk.Result, error) {
	var post = types.Post{
		Creator: msg.Creator,
		ID:      msg.ID,
		Title:   msg.Title,
	}
	k.CreatePost(ctx, post)
	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
