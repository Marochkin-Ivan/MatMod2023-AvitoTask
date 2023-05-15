package logic

import (
	"events-adapter/internal/models"
	"regexp"
	"strings"
)

func CreateRedisEvent(req models.BaseEventRequest) models.RedisEvent {
	return models.RedisEvent{
		VacancyID: req.VacancyID,
		Type:      req.Type,
	}
}

func CreateRedisSearch(req models.SearchEventRequest) models.RedisSearch {
	onlyWords := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(req.SearchPhrases, "")
	keywords := strings.Split(onlyWords, " ")

	return models.RedisSearch{
		Phrases: keywords,
	}
}
