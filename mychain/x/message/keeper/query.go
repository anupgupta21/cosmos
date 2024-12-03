package keeper

import (
	"github.com/username/mychain/x/message/types"
)

var _ types.QueryServer = Keeper{}
