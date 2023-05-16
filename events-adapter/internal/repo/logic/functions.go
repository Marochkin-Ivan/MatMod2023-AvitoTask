package logic

import (
	"events-adapter/internal/models"
)

func CreateRedisEvent(vacancyID, eventType string) models.RedisEvent {
	var points float64

	switch eventType {
	case "open_card":
		points = 0.3
	case "like":
		points = 0.4
	case "call":
		points = 0.5
	case "apply":
		points = 0.7
	}

	return models.RedisEvent{
		VacancyID:  vacancyID,
		TypePoints: points,
	}
}
