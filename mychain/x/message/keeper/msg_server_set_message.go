package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/username/mychain/x/message/types"
)

func (k msgServer) SetMessage(goCtx context.Context, msg *types.MsgSetMessage) (*types.MsgSetMessageResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)


	_ = ctx

	return &types.MsgSetMessageResponse{}, nil
}
