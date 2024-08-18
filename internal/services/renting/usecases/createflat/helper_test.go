//go:build unit
// +build unit

package createflat

import (
	"context"
	"github.com/gojuno/minimock/v3"
	"testing"
)

type testHelper struct {
	flats   *FlatsRepoMock
	houses  *HousesRepoMock
	txM     *TxManagerMock
	service *Service
}

func newTestHelper(t *testing.T) testHelper {
	mc := minimock.NewController(t)
	helper := testHelper{}
	helper.flats = NewFlatsRepoMock(mc)
	helper.houses = NewHousesRepoMock(mc)
	helper.txM = NewTxManagerMock(mc)
	helper.txM.WithinTransactionMock.Set(func(ctx context.Context, f func(context.Context) bool) error {
		f(ctx)
		// ошибки фиксации транзакции тестировать не будем
		return nil
	})
	helper.service = NewService(helper.flats, helper.houses, helper.txM)
	return helper

}
