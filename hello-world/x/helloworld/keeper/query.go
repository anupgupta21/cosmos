package keeper

import (
	"hello-world/x/helloworld/types"
)

var _ types.QueryServer = Keeper{}
