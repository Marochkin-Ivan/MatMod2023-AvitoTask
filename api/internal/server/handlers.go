package server

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func (s *Server) ping(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusOK)
}

func (s *Server) search(c *fiber.Ctx) error {
	const source = "search"

	queryParams := getQueryParams(c)

	searchReq := searchRequest{
		Query: Query{
			Bool: Bool{
				Should: Should{
					MultiMatch: MultiMatch{
						Query:     queryParams["q"],
						Fuzziness: "AUTO",
						Type:      "best_fields",
						Fields: []string{
							"title^3",
							"requirements^1.5",
							"companyName^1",
						},
						TieBreaker: 0.3,
					},
				},
			},
		},
	}

	for param, value := range queryParams {
		if filterFunc, exist := filtersTypeMap[param]; exist {
			searchReq.withFilter(filterFunc, param, value)
		}
	}

	b, _ := json.MarshalIndent(searchReq, "", "  ")

	log.Println(string(b))

	return c.SendStatus(http.StatusOK)
}
