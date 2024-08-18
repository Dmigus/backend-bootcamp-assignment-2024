//go:build unit

package createflat

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/gojuno/minimock/v3"
)

func TestService_createFlat_Positive(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)
	req := Request{HouseId: 5, Price: 10, Rooms: 10}
	createdFlat := models.Flat{Id: 10, HouseId: 5, Price: 10, Rooms: 10, Status: models.Created}
	h.flats.CreateFlatMock.Expect(minimock.AnyContext, req).Return(&createdFlat, nil)
	h.houses.HouseUpdatedMock.Expect(minimock.AnyContext, 5).Return(nil)
	returned, err := h.service.CreateFlat(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, createdFlat, *returned)
}

func TestService_createFlatAndUpdateHouse_Errors(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx  context.Context
		flat Request
	}
	errorToThrow := fmt.Errorf("oops error")
	tests := []struct {
		name      string
		mockSetup func(testHelper)
		args      args
		err       error
	}{
		{
			name: "error creating flat",
			mockSetup: func(helper testHelper) {
				req := Request{HouseId: 5, Price: 10, Rooms: 10}
				helper.flats.CreateFlatMock.Expect(minimock.AnyContext, req).Return(nil, errorToThrow)
			},
			args: args{ctx: context.Background(), flat: Request{HouseId: 5, Price: 10, Rooms: 10}},
			err:  errorToThrow,
		},
		{
			name: "error updating house",
			mockSetup: func(helper testHelper) {
				req := Request{HouseId: 5, Price: 10, Rooms: 10}
				createdFlat := models.Flat{Id: 10, HouseId: 5, Price: 10, Rooms: 10, Status: models.Created}
				helper.flats.CreateFlatMock.Expect(minimock.AnyContext, req).Return(&createdFlat, nil)
				helper.houses.HouseUpdatedMock.Expect(minimock.AnyContext, 5).Return(errorToThrow)
			},
			args: args{ctx: context.Background(), flat: Request{HouseId: 5, Price: 10, Rooms: 10}},
			err:  errorToThrow,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			helper := newTestHelper(t)
			tt.mockSetup(helper)
			_, err := helper.service.CreateFlat(tt.args.ctx, tt.args.flat)
			assert.ErrorIs(t, err, tt.err)
		})
	}
}
