package models

import "os"

var JWTSalt = "salt"

func init() {
	salt, ok := os.LookupEnv("JWT_SALT")
	if ok {
		JWTSalt = salt
	}
}

type BaseConfig struct {
	Host       string `env:"HOST" envDefault:":1234"`
	ServerName string `env:"SERVER_NAME" envDefault:"TestServer"`
	SentryDSN  string `env:"SENTRY_DSN" envDefault:""`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"trace"`
	TLSKey     string `env:"TLS_KEY" envDefault:""`
	TLSPem     string `env:"TLS_PEM" envDefault:""`
}

const IndexName = "vacancy"

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
