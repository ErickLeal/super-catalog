package product

import (
	"context"
	"super-catalog/internal/infra"
)

type Product struct {
	ID    string  `bson:"_id,omitempty"`
	Name  string  `bson:"name"`
	Price float64 `bson:"price"`
}

func InsertProducts(ctx context.Context, products []Product) error {
	coll, err := infra.GetCollection("supercatalog", "products")
	if err != nil {
		return err
	}
	insertDocs := make([]interface{}, len(products))
	for i, p := range products {
		insertDocs[i] = p
	}
	_, err = coll.InsertMany(ctx, insertDocs)
	return err
}
