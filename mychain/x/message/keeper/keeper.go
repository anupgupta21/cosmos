package keeper

import (
    "context"

    "github.com/cosmos/cosmos-sdk/store/prefix"
    sdk "github.com/cosmos/cosmos-sdk/types"

)

type Keeper struct {
    storeKey sdk.StoreKey
}

// NewKeeper creates a new Keeper instance
func NewKeeper(storeKey sdk.StoreKey) Keeper {
    return Keeper{
        storeKey: storeKey,
    }
}


