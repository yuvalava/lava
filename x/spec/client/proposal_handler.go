package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
	"github.com/lavanet/lava/x/spec/client/cli"
)

// ProposalHandler is the param change proposal handler.
var SpecAddProposalHandler = govclient.NewProposalHandler(cli.NewSubmitSpecAddProposalTxCmd)
var SpecModifyProposalHandler = govclient.NewProposalHandler(cli.NewSubmitSpecModifyProposalTxCmd)
