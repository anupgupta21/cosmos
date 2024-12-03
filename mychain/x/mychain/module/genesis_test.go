package mychain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/username/mychain/testutil/keeper"
	"github.com/username/mychain/testutil/nullify"
	mychain "github.com/username/mychain/x/mychain/module"
	"github.com/username/mychain/x/mychain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MychainKeeper(t)
	mychain.InitGenesis(ctx, k, genesisState)
	got := mychain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
