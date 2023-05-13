package logic

import (
	"es-writer/pkg/errs"
	"es-writer/pkg/tools/den"
)

func InsertDocuments(docs []any, indexName string, insert func(name string, req []byte) *errs.Error) *errs.Error {
	const source = "logic.InsertDocuments"

	for _, doc := range docs {
		req, err := den.EncodeJson(doc)
		if err != nil {
			return err.Wrap(source)
		}

		if err := insert(indexName, req.Bytes()); err != nil {
			return err.Wrap(source)
		}
	}

	return nil
}
