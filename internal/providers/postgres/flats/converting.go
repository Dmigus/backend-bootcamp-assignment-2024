package flats

import "backend-bootcamp-assignment-2024/internal/models"

func flatDtoToService(dto Flat) models.Flat {
	return models.Flat{
		Id:      int(dto.ID),
		HouseId: int(dto.HouseID),
		Price:   int(dto.Price),
		Rooms:   int(dto.Rooms),
		Status:  strToFlatStatus(dto.Status),
	}
}

func strToFlatStatus(str string) models.FlatStatus {
	switch str {
	case "Created":
		return models.Created
	case "OnModerate":
		return models.OnModerate
	case "Approved":
		return models.Approved
	case "Declined":
		return models.Declined
	default:
		return 0
	}
}
