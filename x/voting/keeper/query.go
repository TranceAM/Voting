package keeper

import (
	"github.com/TranceAM/voting/x/voting/types"
)

var _ types.QueryServer = Keeper{}
