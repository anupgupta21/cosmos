package helloapp_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/username/helloapp/testutil/keeper"
	"github.com/username/helloapp/testutil/nullify"
	helloapp "github.com/username/helloapp/x/helloapp/module"
	"github.com/username/helloapp/x/helloapp/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.HelloappKeeper(t)
	helloapp.InitGenesis(ctx, k, genesisState)
	got := helloapp.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
