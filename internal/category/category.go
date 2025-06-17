package category

import (
	"context"
	"fmt"
	"super-catalog/internal/infra"
)

func InsertCategories(ctx context.Context, categories []interface{}) error {
	coll, err := infra.GetCollection("supercatalog", "categories")
	if err != nil {
		return err
	}

	fmt.Printf("[InsertCategories] Tipo recebido: %T\n", categories)

	if len(categories) == 0 {
		return nil
	}

	insertDocs := make([]interface{}, len(categories))
	for i, c := range categories {
		insertDocs[i] = c
		fmt.Printf("[InsertCategories] Item %d tipo: %T\n", i, c)
	}
	_, err = coll.InsertMany(ctx, insertDocs)
	return err
}

func GetCategoryByID(ctx context.Context, id string) (map[string]interface{}, error) {
	coll, err := infra.GetCollection("supercatalog", "categories")
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	err = coll.FindOne(ctx, map[string]interface{}{"id": id}).Decode(&result)
	return result, err
}
