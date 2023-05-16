package server

import (
	"api/internal/repo/cache"
	"api/pkg/errs"
	"api/pkg/tools/den"
	"sort"
	"strconv"
	"strings"
)

func (req *searchRequest) withFilter(filterFunc, param, value string) {
	switch param {
	case "keywords", "employment", "schedule":
		req.Query.Bool.Filter = append(req.Query.Bool.Filter,
			map[string]any{
				filterFunc: map[string]string{
					param: value,
				},
			})

	case "salaryMin", "salaryMax":
		var p string
		if param == "salaryMin" {
			p = "gte"
		} else {
			p = "lte"
		}

		req.Query.Bool.Filter = append(req.Query.Bool.Filter,
			map[string]any{
				filterFunc: map[string]map[string]string{
					"salary": {
						p: value,
					},
				},
			})

	case "from":
		v, _ := strconv.Atoi(value)
		req.From = v

	case "size":
		v, _ := strconv.Atoi(value)
		req.Size = v
	}
}

func (esResp ElasticResponse) RankSort(userID string, c Getter) *errs.Error {
	const source = "RankSort"

	keywordsPoints := make(map[string]float64)

	events, err := c.GetArray(cache.EventsTable, userID)
	if err != nil {
		return err.Wrap(source)
	}

	for _, ev := range events {
		var event RedisEvent
		err = den.DecodeJson(&event, []byte(ev))
		if err != nil {
			return err.Wrap(source)
		}

		keywords, err := c.GetArray(cache.KeyWordsTable, event.VacancyID)
		if err != nil {
			return err.Wrap(source)
		}

		for _, kw := range keywords {
			keywordsPoints[kw] += event.TypePoints
		}
	}

	for _, vacancy := range esResp.Hits.Hits {
		vKeywords := strings.Split(vacancy.Source.Keywords, ",")
		for _, kw := range vKeywords {
			vacancy.Rating += keywordsPoints[kw]
		}
	}

	sort.Slice(esResp.Hits.Hits, func(i, j int) bool {
		return esResp.Hits.Hits[i].Rating > esResp.Hits.Hits[j].Rating
	})

	return nil
}

func (esResp ElasticResponse) ToResponse() Response {
	var res Response

	for _, v := range esResp.Hits.Hits {
		res = append(res, Vacancy{
			ID:           v.Id,
			Title:        v.Source.Title,
			Requirements: v.Source.Requirements,
			Keywords:     v.Source.Keywords,
			Salary:       v.Source.Salary,
			Region:       v.Source.Region,
			CompanyName:  v.Source.CompanyName,
			Schedule:     v.Source.Schedule,
			Experience:   v.Source.Experience,
			Employment:   v.Source.Employment,
			CreatedAt:    v.Source.CreatedAt,
		})
	}

	return res
}
