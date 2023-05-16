package logic

import (
	"events-adapter/internal/models"
)

func CreateRedisEvent(vacancyID, eventType string) models.RedisEvent {
	return models.RedisEvent{
		VacancyID: vacancyID,
		Type:      eventType,
	}
}
