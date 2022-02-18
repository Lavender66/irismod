package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (
	ProposalTypeCreateFarmPool string = "CommunityPoolCreateFarm"
	FarPoolPrefix              string = "SYS-"
)

// Implements Proposal Interface
var _ govtypes.Content = &CommunityPoolCreateFarmProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeCreateFarmPool)
	govtypes.RegisterProposalTypeCodec(&CommunityPoolCreateFarmProposal{}, "irismod/CommunityPoolCreateFarmProposal")
}

func GenSyspoolId(name string) string {
	return fmt.Sprintf(name, FarPoolPrefix)
}

func (cfp *CommunityPoolCreateFarmProposal) GetTitle() string       { return cfp.Title }
func (cfp *CommunityPoolCreateFarmProposal) GetDescription() string { return cfp.Description }
func (cfp *CommunityPoolCreateFarmProposal) ProposalRoute() string  { return RouterKey }
func (cfp *CommunityPoolCreateFarmProposal) ProposalType() string   { return ProposalTypeCreateFarmPool }
func (cfp *CommunityPoolCreateFarmProposal) ValidateBasic() error {
	if err := ValidateDescription(cfp.PoolDescription); err != nil {
		return err
	}

	if err := ValidateLpTokenDenom(cfp.LptDenom); err != nil {
		return err
	}

	if err := ValidateCoins("RewardsPerBlock", cfp.RewardPerBlock...); err != nil {
		return err
	}

	if err := ValidateCoins("TotalRewards", cfp.TotalReward...); err != nil {
		return err
	}

	if err := ValidateReward(cfp.RewardPerBlock, cfp.TotalReward); err != nil {
		return err
	}

	return govtypes.ValidateAbstract(cfp)
}

func (cfp CommunityPoolCreateFarmProposal) String() string {
	return fmt.Sprintf(`Community Pool Create Farm Proposal:
  Title:       %s
  Description: %s
  PoolDescription: %s
  LpTokenDenom: %s
  RewardPerBlock: %s
  TotalReward: %s
`, cfp.Title, cfp.Description, cfp.PoolDescription, cfp.LptDenom, sdk.Coins(cfp.RewardPerBlock), sdk.Coins(cfp.TotalReward))
}