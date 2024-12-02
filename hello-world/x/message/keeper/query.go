package keeper

import (
	"hello-world/x/message/types"
)

var _ types.QueryServer = Keeper{}
