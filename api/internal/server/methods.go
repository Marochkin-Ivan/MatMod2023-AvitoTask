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
