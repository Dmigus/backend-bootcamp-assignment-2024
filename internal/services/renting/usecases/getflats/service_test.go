//go:build unit

package getflats

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"context"
	"fmt"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_GetFlats_PositiveModer(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)
	allFlats := []models.Flat{
		{Id: 10, HouseId: 1, Price: 10, Rooms: 10, Status: models.Created},
		{Id: 11, HouseId: 1, Price: 10, Rooms: 10, Status: models.Approved},
	}
	h.flats.GetFlatsMock.Expect(minimock.AnyContext, 1).Return(allFlats, nil)
	flatsForModer, err := h.service.GetFlats(context.Background(), 1, models.Moderator)
	require.NoError(t, err)
	assert.Equal(t, flatsForModer, allFlats)
}

func TestService_GetFlats_PositiveClient(t *testing.T) {
	t.Parallel()
	h := newTestHelper(t)
	approved := []models.Flat{{Id: 11, HouseId: 1, Price: 10, Rooms: 10, Status: models.Approved}}
	h.flats.GetApprovedFlatsMock.Expect(minimock.AnyContext, 1).Return(approved, nil)
	flatsForClient, err := h.service.GetFlats(context.Background(), 1, models.Client)
	require.NoError(t, err)
	assert.Equal(t, flatsForClient, approved)
}

func TestService_GetFlats_Errors(t *testing.T) {
	t.Parallel()
	type args struct {
		ctx  context.Context
		id   int
		role models.UserRole
	}
	errorToThrow := fmt.Errorf("oops error")
	tests := []struct {
		name      string
		mockSetup func(testHelper)
		args      args
		err       error
	}{
		{
			name:      "unknown role",
			mockSetup: func(helper testHelper) {},
			args:      args{context.Background(), 1, models.UserRole(10)},
			err:       models.ErrUnknownRole,
		},
		{
			name: "error from repo",
			mockSetup: func(helper testHelper) {
				helper.flats.GetFlatsMock.Expect(minimock.AnyContext, 1).Return(nil, errorToThrow)
			},
			args: args{context.Background(), 1, models.Moderator},
			err:  errorToThrow,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			helper := newTestHelper(t)
			tt.mockSetup(helper)
			_, err := helper.service.GetFlats(tt.args.ctx, tt.args.id, tt.args.role)
			assert.ErrorIs(t, err, tt.err)
		})
	}
}
