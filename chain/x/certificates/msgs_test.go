package certificates

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"

	"github.com/EmpowerPlastic/empowerchain/testutil/sample"
)

type validateTest struct {
	msgUnderTest  sdk.Msg
	expectedError error
}

func TestMsgUpdateParams_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgUpdateParams{
				Params: Params{
					AllowedIssuer: []string{sample.AccAddress()},
				},
			},
			expectedError: nil,
		},
		"invalid issuers params": {
			msgUnderTest: &MsgUpdateParams{
				Params: Params{
					AllowedIssuer: []string{"invalid"},
				},
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()

			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}

func TestMsgCreateCertificate_ValidateBasic(t *testing.T) {
	testCases := map[string]validateTest{
		"happy path": {
			msgUnderTest: &MsgCreateCertificate{
				Owner:  sample.AccAddress(),
				Issuer: sample.AccAddress(),
				Type:   0,
				Data:   map[string]string{},
			},
			expectedError: nil,
		},
		"invalid owner": {
			msgUnderTest: &MsgCreateCertificate{
				Owner:  "invalid",
				Issuer: sample.AccAddress(),
				Type:   0,
				Data:   map[string]string{},
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
		"invalid issuer": {
			msgUnderTest: &MsgCreateCertificate{
				Owner:  sample.AccAddress(),
				Issuer: "invalid",
				Type:   0,
				Data:   map[string]string{},
			},
			expectedError: sdkerrors.ErrInvalidAddress,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			err := tc.msgUnderTest.ValidateBasic()

			require.ErrorIs(t, err, tc.expectedError)
		})
	}
}
