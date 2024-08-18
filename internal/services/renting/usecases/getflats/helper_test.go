//go:build unit

package getflats

import (
	"github.com/gojuno/minimock/v3"
	"testing"
)

type testHelper struct {
	flats   *RepositoryMock
	service *Service
}

func newTestHelper(t *testing.T) testHelper {
	mc := minimock.NewController(t)
	helper := testHelper{}
	helper.flats = NewRepositoryMock(mc)
	helper.service = NewService(helper.flats)
	return helper
}
