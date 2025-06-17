package category

import (
	"context"
	"super-catalog/internal/infra"
)

type Category struct {
	ID   string `bson:"_id,omitempty"`
	Name string `bson:"name"`
}

func InsertCategories(ctx context.Context, categories []Category) error {
	coll, err := infra.GetCollection("supercatalog", "categories")
	if err != nil {
		return err
	}
	insertDocs := make([]interface{}, len(categories))
	for i, c := range categories {
		insertDocs[i] = c
	}
	_, err = coll.InsertMany(ctx, insertDocs)
	return err
}
