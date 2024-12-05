package helloworld_test

import (
	"testing"

	keepertest "hello-world/testutil/keeper"
	"hello-world/testutil/nullify"
	helloworld "hello-world/x/helloworld/module"
	"hello-world/x/helloworld/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HelloworldKeeper(t)
	helloworld.InitGenesis(ctx, k, genesisState)
	got := helloworld.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
