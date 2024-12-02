package keeper

import (
	"context"

	"hello-world/x/message/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateSupplychain(goCtx context.Context, msg *types.MsgCreateSupplychain) (*types.MsgCreateSupplychainResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgCreateSupplychainResponse{}, nil
}
