package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
)

// для запроса через апи (отдает максимум 10к записей)

//type Response struct {
//	Status  string `json:"status"`
//	Results struct {
//		Vacancies []struct {
//			Vacancy struct {
//				Id     string `json:"id"`
//				Source string `json:"source"`
//				Region struct {
//					RegionCode string `json:"region_code"`
//					Name       string `json:"name"`
//				} `json:"region"`
//				Category struct {
//					Specialisation string `json:"specialisation"`
//				} `json:"category"`
//			} `json:"vacancy"`
//		} `json:"vacancies"`
//	} `json:"results"`
//}
//
//func main() {
//	// Параметры запроса
//	perPage := 100 // Количество вакансий на странице
//	page := 1
//
//	vacancyMap := make(map[string]bool)
//
//	for {
//		// Создание URL для запроса
//		url := fmt.Sprintf("https://opendata.trudvsem.ru/api/v1/vacancies/region/6100000000000?limit=%d&offset=%d", perPage, page)
//		//log.Println("Запрос:", url)
//		//url := fmt.Sprintf("http://opendata.trudvsem.ru/api/v1/vacancies")
//
//		// Выполнение запроса к API
//		resp, err := http.Get(url)
//		if err != nil {
//			log.Fatal(err)
//		}
//		defer resp.Body.Close()
//
//		// Чтение ответа API
//		var response Response
//		err = json.NewDecoder(resp.Body).Decode(&response)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		if response.Status != "200" {
//			log.Println(response)
//			log.Println("Запрос:", url)
//			break
//		}
//
//		// Обработка каждой вакансии
//		for _, vacancy := range response.Results.Vacancies {
//			specialisations := strings.Split(vacancy.Vacancy.Category.Specialisation, ",")
//
//			// Удаление начальных и конечных пробелов в каждой специализации
//			for i := range specialisations {
//				specialisations[i] = strings.TrimSpace(specialisations[i])
//			}
//
//			// Добавление специализаций в мапу (избежание повторений)
//			for _, specialisation := range specialisations {
//				if specialisation != "" {
//					vacancyMap[specialisation] = true
//				}
//			}
//		}
//
//		page++
//		log.Println("Обработано вакансий:", page*perPage)
//		//break
//	}
//
//	// Вывод всех уникальных специализаций
//	//log.Println(vacancyMap)
//	var res []string
//	for specialisation := range vacancyMap {
//		res = append(res, specialisation)
//	}
//
//	f, e := os.Create("rr_keywords.json")
//	log.Println(e)
//	defer f.Close()
//
//	bs, _ := json.MarshalIndent(res, " ", " ")
//	f.Write(bs)
//}

type q struct {
	Vacancies []FileInfo `json:"vacancies"`
}

type FileInfo struct {
	ProfessionalSphereName string `json:"professionalSphereName"`
}

const filename = "vacancy.json"

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

	bs, _ = json.MarshalIndent(res, " ", " ")
	f.Write(bs)
}
