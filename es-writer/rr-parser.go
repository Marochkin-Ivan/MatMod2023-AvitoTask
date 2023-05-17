package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

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

type q struct {
	Vacancies []FileInfo `json:"vacancies"`
}

//type FileInfo struct {
//	ProfessionalSphereName string `json:"professionalSphereName"`
//}

const filename = "vacancy.json"

func getKeys(vs q) {
	vacancyMap := make(map[string]struct{})
	// Обработка каждой вакансии
	for _, vacancy := range vs.Vacancies {
		specialisations := strings.Split(vacancy.ProfessionalSphereName, ",")

		// Удаление начальных и конечных пробелов в каждой специализации
		for i := range specialisations {
			specialisations[i] = strings.TrimSpace(specialisations[i])
		}

		// Добавление специализаций в мапу (избежание повторений)
		for _, specialisation := range specialisations {
			if specialisation != "" {
				vacancyMap[specialisation] = struct{}{}
			}
		}
	}

	var res []string
	for v := range vacancyMap {
		res = append(res, v)
	}

	f, e := os.Create("rr_keywords.json")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	bs, _ := json.MarshalIndent(res, " ", " ")
	f.Write(bs)
}

func main() {
	bs, e := os.ReadFile(filename)
	if e != nil {
		log.Fatal(e)
	}

	var vs q
	e = json.Unmarshal(bs, &vs)
	if e != nil {
		log.Fatal(e)
	}
	bs = nil
	// getKeys(vs)

	// Обработка каждой вакансии
	f, e := os.Create("rr_vacancies_normal_format.json")
	if e != nil {
		log.Fatal(e)
	}
	defer f.Close()

	bs, _ = json.MarshalIndent(vs.Vacancies, " ", " ")
	f.Write(bs)
}
