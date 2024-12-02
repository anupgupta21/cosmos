package keeper

import (
	"github.com/username/helloapp/x/helloapp/types"
)

var _ types.QueryServer = Keeper{}
