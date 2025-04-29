package tronprotocol

import (
	"testing"

	pbtron "github.com/streamingfast/tron-protocol/pb/core"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func TestAnyContractFitsGRPC(t *testing.T) {
	testCases := []struct {
		name         string
		message      proto.Message
		expectedType string
	}{
		{
			name:         "TriggerSmartContract",
			message:      &pbtron.TriggerSmartContract{},
			expectedType: "type.googleapis.com/protocol.TriggerSmartContract",
		},
		{
			name:         "TransferContract",
			message:      &pbtron.TransferContract{},
			expectedType: "type.googleapis.com/protocol.TransferContract",
		},
		{
			name:         "TransferAssetContract",
			message:      &pbtron.TransferAssetContract{},
			expectedType: "type.googleapis.com/protocol.TransferAssetContract",
		},
		{
			name:         "VoteWitnessContract",
			message:      &pbtron.VoteWitnessContract{},
			expectedType: "type.googleapis.com/protocol.VoteWitnessContract",
		},
		{
			name:         "FreezeBalanceContract",
			message:      &pbtron.FreezeBalanceContract{},
			expectedType: "type.googleapis.com/protocol.FreezeBalanceContract",
		},
		{
			name:         "UnfreezeBalanceContract",
			message:      &pbtron.UnfreezeBalanceContract{},
			expectedType: "type.googleapis.com/protocol.UnfreezeBalanceContract",
		},
		{
			name:         "WithdrawBalanceContract",
			message:      &pbtron.WithdrawBalanceContract{},
			expectedType: "type.googleapis.com/protocol.WithdrawBalanceContract",
		},
		{
			name:         "UpdateAssetContract",
			message:      &pbtron.UpdateAssetContract{},
			expectedType: "type.googleapis.com/protocol.UpdateAssetContract",
		},
		// Add more contracts here if needed
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := anypb.New(tc.message)
			require.NoError(t, err)
			require.Equal(t, tc.expectedType, out.TypeUrl)
		})
	}
}
