package keeper

import (
	"github.com/username/mychain/x/mychain/types"
)

var _ types.QueryServer = Keeper{}
