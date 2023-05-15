package server

import (
	"api/internal/models"
	"api/pkg/errs"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	cfg  Config
	logs errs.LogChan
	a    *fiber.App
}

type Config struct {
	models.BaseConfig
	// special fields for server ...
}

var filtersTypeMap = map[string]string{
	"keywords":   "match_phrase",
	"salaryMin":  "range",
	"salaryMax":  "range",
	"employment": "match_phrase",
	"schedule":   "match_phrase",
}

type searchRequest struct {
	Query Query `json:"query"`
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
