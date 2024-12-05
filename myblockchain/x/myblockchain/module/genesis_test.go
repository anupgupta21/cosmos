package myblockchain_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/username/myblockchain/testutil/keeper"
	"github.com/username/myblockchain/testutil/nullify"
	myblockchain "github.com/username/myblockchain/x/myblockchain/module"
	"github.com/username/myblockchain/x/myblockchain/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MyblockchainKeeper(t)
	myblockchain.InitGenesis(ctx, k, genesisState)
	got := myblockchain.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
