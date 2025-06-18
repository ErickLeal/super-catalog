package category

type Size struct {
	ID   string `bson:"id" json:"id"`
	Name string `bson:"name" json:"name"`
}

type SizeFlavor struct {
	ID          string `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	MaxFlavours int64  `bson:"max_flavours" json:"max_flavours"`
}

type AskGroupOption struct {
	ID          string `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Description string `bson:"description" json:"description"`
	Value       int64  `bson:"value" json:"value"`
}

type AskGroup struct {
	ID           string           `bson:"id" json:"id"`
	Group        string           `bson:"group" json:"group"`
	MinimunLimit int              `bson:"minimun_limit" json:"minimun_limit"`
	MaximunLimit int              `bson:"maximun_limit" json:"maximun_limit"`
	Options      []AskGroupOption `bson:"options" json:"options"`
}

type Schedul struct {
	Day   string `bson:"day" json:"day"`
	Hours string `bson:"hours" json:"hours"`
}

type FoodsCategory struct {
	Type        string     `bson:"type" json:"type"`
	StoreId     string     `bson:"store_id" json:"store_id"`
	ID          string     `bson:"id" json:"id"`
	Name        string     `bson:"name" json:"name"`
	Description string     `bson:"description" json:"description"`
	Culinary    string     `bson:"culinary" json:"culinary"`
	Sizes       []Size     `bson:"sizes" json:"sizes"`
	AskGroups   []AskGroup `bson:"ask_groups" json:"ask_groups"`
}

type SlicedFoodsCategory struct {
	Type        string       `bson:"type" json:"type"`
	StoreId     string       `bson:"store_id" json:"store_id"`
	ID          string       `bson:"id" json:"id"`
	Name        string       `bson:"name" json:"name"`
	Description string       `bson:"description" json:"description"`
	Sizes       []SizeFlavor `bson:"sizes" json:"sizes"`
	AskGroups   []AskGroup   `bson:"ask_groups" json:"ask_groups"`
}

type MaketCategory struct {
	Type        string `bson:"type" json:"type"`
	ID          string `bson:"id" json:"id"`
	Name        string `bson:"name" json:"name"`
	Section     string `bson:"section" json:"section"`
	Description string `bson:"description" json:"description"`
}

type SchedulCategory struct {
	Type        string    `bson:"type" json:"type"`
	ID          string    `bson:"id" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	Schedul     []Schedul `bson:"schedul" json:"schedul"`
}

type OpenCategory struct {
	Type        string     `bson:"type" json:"type"`
	StoreId     string     `bson:"store_id" json:"store_id"`
	ID          string     `bson:"id" json:"id"`
	Name        string     `bson:"name" json:"name"`
	Section     string     `bson:"section" json:"section"`
	Description string     `bson:"description" json:"description"`
	Culinary    string     `bson:"culinary" json:"culinary"`
	Sizes       []Size     `bson:"sizes" json:"sizes"`
	AskGroups   []AskGroup `bson:"ask_groups" json:"ask_groups"`
	Schedul     []Schedul  `bson:"schedul" json:"schedul"`
}
