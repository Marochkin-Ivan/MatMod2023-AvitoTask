package server

import (
	"es-writer/internal/models"
	"es-writer/internal/repo/es"
	"es-writer/pkg/errs"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	cfg  Config
	logs errs.LogChan
	a    *fiber.App
	es   *es.Client
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

type ReindexRequest struct {
	Source struct {
		Index string `json:"index"`
	} `json:"source"`
	Dest struct {
		Index string `json:"index"`
	} `json:"dest"`
	Script struct {
		Source string `json:"source"`
		Lang   string `json:"lang"`
		Params struct {
			Keywords string `json:"keywords"`
		} `json:"params"`
	} `json:"script"`
}

type ESDoc struct {
	Title        string  `json:"title"`
	Requirements string  `json:"requirements"`
	Keywords     string  `json:"keywords"`
	Salary       float64 `json:"salary"`
	Region       string  `json:"region"`
	CompanyName  string  `json:"companyName"`
	Schedule     string  `json:"schedule"`
	Experience   int     `json:"experience"`
	Employment   string  `json:"employment"`
	CreatedAt    string  `json:"createdAt"`
}

type FileInfo struct {
	AdditionalRequirements   string `json:"additional_requirements"`
	BusyType                 string `json:"busy_type"`
	ContactSource            string `json:"contact_source"`
	DateCreate               string `json:"date_create"`
	Education                string `json:"education"`
	FullCompanyName          string `json:"full_company_name"`
	ID                       string `json:"id"`
	OtherVacancyBenefit      string `json:"other_vacancy_benefit"`     // Плюшки
	PositionRequirements     string `json:"position_requirements"`     // Требования
	PositionResponsibilities string `json:"position_responsibilities"` // Обязанности
	ProfessionalSphereName   string `json:"professionalSphereName"`
	RegionName               string `json:"regionName"`
	RequiredExperience       int    `json:"required_experience"`
	Salary                   string `json:"salary"`
	SalaryMax                int    `json:"salary_max"`
	SalaryMin                int    `json:"salary_min"`
	ScheduleType             string `json:"schedule_type"`
	Address                  string `json:"vacancy_address"`
	Name                     string `json:"vacancy_name"`
	WorkPlaces               int    `json:"work_places"`
}
