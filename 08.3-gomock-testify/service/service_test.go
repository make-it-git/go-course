package service

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestService(t *testing.T) {
	testCases := map[string]struct {
		username           string
		amount             int
		balance            int
		initPaymentGateway func(m *MockPaymentGateway)
		initLogger         func(m *MockLogger)
		err                error
	}{
		"success": {
			username: "Jack",
			amount:   1,
			balance:  10,
			initPaymentGateway: func(m *MockPaymentGateway) {
				m.EXPECT().SendMoneyAndGetCurrentBalance("Jack", 1).Return(10, nil)
			},
			initLogger: func(m *MockLogger) {
				m.EXPECT().Info("send money ok")
			},
		},
		"error": {
			username: "John",
			amount:   10,
			balance:  0,
			initPaymentGateway: func(m *MockPaymentGateway) {
				m.EXPECT().SendMoneyAndGetCurrentBalance("John", 10).Return(0, ErrConnection)
			},
			initLogger: func(m *MockLogger) {
				m.EXPECT().Error("send money error")
			},
			err: ErrConnection,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			extDep := NewMockPaymentGateway(ctrl)
			if tc.initPaymentGateway != nil {
				tc.initPaymentGateway(extDep)
			}
			logger := NewMockLogger(ctrl)
			if tc.initLogger != nil {
				tc.initLogger(logger)
			}
			balance, err := TransferMoney(extDep, logger, tc.username, tc.amount)
			if err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
			require.Equal(t, balance, tc.balance)
		})
	}
}
