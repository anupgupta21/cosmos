package message_test

import (
	"testing"

	keepertest "hello-world/testutil/keeper"
	"hello-world/testutil/nullify"
	message "hello-world/x/message/module"
	"hello-world/x/message/types"

	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MessageKeeper(t)
	message.InitGenesis(ctx, k, genesisState)
	got := message.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
