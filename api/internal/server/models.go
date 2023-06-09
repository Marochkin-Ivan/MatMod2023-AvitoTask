package server

import (
	"api/internal/models"
	"api/internal/repo/es"
	"api/pkg/errs"
	"github.com/gofiber/fiber/v2"
)

type Getter interface {
	GetValue(dbID int, key string) (string, *errs.Error)
	GetArray(dbID int, key string) ([]string, *errs.Error)
}

type Server struct {
	cfg   Config
	logs  errs.LogChan
	a     *fiber.App
	es    *es.Client
	cache Getter
}

type Config struct {
	models.BaseConfig
	// special fields for server ...
}

var FiltersMap = map[string][]string{
	"keywords": {
		"Безопасность",
		"Юриспруденция",
		"Маркетинг",
		"Консалтинг",
		"Пищевая промышленность",
		"Логистика",
		"Управление персоналом",
		"Строительство",
		"Производство",
		"Сельское хозяйство",
		"Образование",
		"Торговля",
		"Информационные технологии",
		"Электроэнергетика",
		"Здравоохранение и социальное обеспечение",
		"Бухгалтерия",
		"Реклама",
		"Туризм",
	},
	"employment": {
		"Полная занятость",
		"Частичная занятость",
		"Проектная работа",
		"Стажировка",
		"Волонтёрство",
	},
	"schedule": {
		"Полный день",
		"Сменный график",
		"Гибкий график",
		"Удалённая работа",
		"Вахтовый метод",
	},
}

var filtersTypeMap = map[string]string{
	"keywords":   "match_phrase",
	"salaryMin":  "range",
	"salaryMax":  "range",
	"employment": "match_phrase",
	"schedule":   "match_phrase",
	"page":       "",
}

type searchRequest struct {
	Query Query `json:"query"`
	From  int   `json:"from,omitempty"`
	Size  int   `json:"size,omitempty"`
}

type Query struct {
	Bool Bool `json:"bool"`
}

type Bool struct {
	Should Should           `json:"should"`
	Filter []map[string]any `json:"filter"`
}

type Should struct {
	MultiMatch MultiMatch `json:"multi_match"`
}

type MultiMatch struct {
	Query      string   `json:"query"`
	Fuzziness  string   `json:"fuzziness"`
	Type       string   `json:"type"`
	Fields     []string `json:"fields"`
	TieBreaker float64  `json:"tie_breaker"`
}

type ElasticResponse struct {
	Took int      `json:"took"`
	Hits HitsWrap `json:"hits"`
}

type HitsWrap struct {
	Hits Hits `json:"hits"`
}

type Hits []Hit

type Hit struct {
	Index  string  `json:"_index"`
	Id     string  `json:"_id"`
	Score  float64 `json:"_score"`
	Source Source  `json:"_source"`
	Rating float64 `json:"-"`
}

type Source struct {
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

type Response []Vacancy

type Vacancy struct {
	ID           string  `json:"id"`
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

type RedisEvent struct {
	VacancyID  string  `json:"vacancy_id"`
	TypePoints float64 `json:"type_points"`
}
