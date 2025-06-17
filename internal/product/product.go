package product

import (
	"context"
	"fmt"
	"super-catalog/internal/infra"
)

func InsertProducts(ctx context.Context, products []interface{}) error {
	coll, err := infra.GetCollection("supercatalog", "products")
	if err != nil {
		return err
	}

	if len(products) == 0 {
		return nil
	}

	insertDocs := make([]interface{}, len(products))
	for i, c := range products {
		insertDocs[i] = c
		fmt.Printf("[InsertProducts] Item %d tipo: %T\n", i, c)
	}
	_, err = coll.InsertMany(ctx, insertDocs)
	return err
}
