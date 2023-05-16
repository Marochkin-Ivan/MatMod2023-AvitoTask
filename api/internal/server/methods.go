package server

import (
	"strconv"
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
