package testutil

import (
	"fmt"

	"github.com/gogo/protobuf/proto"
	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/testutil/rest"
	sdk "github.com/cosmos/cosmos-sdk/types"

	htlccli "github.com/irisnet/irismod/modules/htlc/client/cli"
	htlctypes "github.com/irisnet/irismod/modules/htlc/types"
)

func (s *IntegrationTestSuite) TestGRPC() {
	val := s.network.Validators[0]

	//------test GetCmdCreateHTLC()-------------
	baseURL := val.APIAddress
	from := val.Address
	to := sdk.AccAddress(crypto.AddressHash([]byte("dgsbl")))
	amount := "1000" + sdk.DefaultBondDenom
	receiverOnOtherChain := "0xcd2a3d9f938e13cd947ec05abc7fe734df8dd826"
	hashLock := "e8d4133e1a82c74e2746e78c19385706ea7958a0ca441a08dacfa10c48ce2561"
	timeLock := uint64(50)
	timestamp := uint64(1580000000)
	stateOpen := "HTLC_STATE_OPEN"

	args := []string{
		fmt.Sprintf("--%s=%s", htlccli.FlagTo, to),
		fmt.Sprintf("--%s=%s", htlccli.FlagAmount, amount),
		fmt.Sprintf("--%s=%s", htlccli.FlagReceiverOnOtherChain, receiverOnOtherChain),
		fmt.Sprintf("--%s=%s", htlccli.FlagHashLock, hashLock),
		fmt.Sprintf("--%s=%d", htlccli.FlagTimeLock, timeLock),
		fmt.Sprintf("--%s=%d", htlccli.FlagTimestamp, timestamp),

		fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
		fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
		fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(s.cfg.BondDenom, sdk.NewInt(10))).String()),
	}

	respType := proto.Message(&sdk.TxResponse{})
	expectedCode := uint32(0)

	bz, err := CreateHTLCExec(val.ClientCtx, from.String(), args...)
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(bz.Bytes(), respType), bz.String())
	txResp := respType.(*sdk.TxResponse)
	s.Require().Equal(expectedCode, txResp.Code)

	//------test GetCmdQueryHTLC()-------------
	url := fmt.Sprintf("%s/irismod/htlc/htlcs/%s", baseURL, hashLock)
	resp, err := rest.GetRequest(url)
	respType = proto.Message(&htlctypes.QueryHTLCResponse{})
	s.Require().NoError(err)
	s.Require().NoError(val.ClientCtx.Codec.UnmarshalJSON(resp, respType))
	htlcResp := respType.(*htlctypes.QueryHTLCResponse)
	s.Require().Equal(amount, htlcResp.Htlc.Amount.String())
	s.Require().Equal(from.String(), htlcResp.Htlc.Sender)
	s.Require().Equal(to.String(), htlcResp.Htlc.To)
	s.Require().Equal(receiverOnOtherChain, htlcResp.Htlc.ReceiverOnOtherChain)
	s.Require().Equal(timestamp, htlcResp.Htlc.Timestamp)
	s.Require().Equal(stateOpen, htlcResp.Htlc.State.String())
}
