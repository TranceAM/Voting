package voting_test

import (
	"testing"

	keepertest "github.com/TranceAM/voting/testutil/keeper"
	"github.com/TranceAM/voting/testutil/nullify"
	"github.com/TranceAM/voting/x/voting"
	"github.com/TranceAM/voting/x/voting/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VotingKeeper(t)
	voting.InitGenesis(ctx, *k, genesisState)
	got := voting.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
