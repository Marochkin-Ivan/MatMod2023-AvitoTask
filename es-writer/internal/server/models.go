package server

import (
	"es-writer/internal/models"
	"es-writer/internal/repo/es"
	"es-writer/pkg/errs"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	cfg   Config
	logs  errs.LogChan
	a     *fiber.App
	es    *es.Client
	cache Setter
}

type Config struct {
	models.BaseConfig
	// special fields for server ...
}

var createIndexRequest = map[string]interface{}{
	"settings": map[string]interface{}{
		"number_of_shards":   1,
		"number_of_replicas": 0,
	},
	"mappings": map[string]interface{}{
		"properties": map[string]interface{}{
			"title":        map[string]interface{}{"type": "text"},    // vacancy_name
			"requirements": map[string]interface{}{"type": "text"},    // position_requirements
			"keywords":     map[string]interface{}{"type": "text"},    // professionalSphereName
			"salary":       map[string]interface{}{"type": "double"},  // salary
			"region":       map[string]interface{}{"type": "keyword"}, // regionName
			"companyName":  map[string]interface{}{"type": "text"},    // full_company_name
			"schedule":     map[string]interface{}{"type": "text"},    // schedule_type
			"experience":   map[string]interface{}{"type": "double"},  // required_experience
			"employment":   map[string]interface{}{"type": "text"},    // busy_type
			"createdAt":    map[string]interface{}{"type": "date"},    // date_create
		},
	},
}

type Setter interface {
	// AddValueInPipeline добавляет множество значений к массиву |ключ(string) : значение(string)|
	AddValueInPipeline(dbID int, KVMap map[string][]string) *errs.Error
	// SetValueInPipeline создает элемент в указанной таблице |ключ(string) : значение(string)|
	SetValueInPipeline(dbID int, KVMap map[string]string) *errs.Error

	FlushAll() *errs.Error
	Ping() *errs.Error
}
