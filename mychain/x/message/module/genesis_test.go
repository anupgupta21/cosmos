package message_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "github.com/username/mychain/testutil/keeper"
	"github.com/username/mychain/testutil/nullify"
	message "github.com/username/mychain/x/message/module"
	"github.com/username/mychain/x/message/types"
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
