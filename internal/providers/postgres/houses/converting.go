package houses

import (
	"backend-bootcamp-assignment-2024/internal/models"
	"backend-bootcamp-assignment-2024/internal/services/renting/usecases/housecreate"

	"github.com/jackc/pgx/v5/pgtype"
)

func houseServiceToDto(req housecreate.HouseCreateRequest) createHouseParams {
	params := createHouseParams{Address: req.Address, Year: int32(req.Year)}
	if req.Developer != nil {
		params.Developer = pgtype.Text{String: *req.Developer, Valid: true}
	}
	return params
}

func houseDtoToService(dto House) models.House {
	result := models.House{
		Id:        int(dto.ID),
		Address:   dto.Address,
		Year:      int(dto.Year),
		CreatedAt: dto.CreatedAt.Time,
	}
	if dto.Developer.Valid {
		result.Developer = &dto.Developer.String
	}
	if dto.UpdateAt.Valid {
		result.UpdateAt = &dto.UpdateAt.Time
	}
	return result
}
