package models

type Values struct {
	Items  []item   `json:"items"`
	Images []images `json:"images"`
	Detail []detail `json:"details"`
}

type item struct {
	FieldName string `json:"fieldName"`
	Value     string `json:"value"`
}

type images struct {
	Name      string `json:"name"`
	ObjectKey string `json:"objectKey"`
}

type detail struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Price    int    `json:"price"`
}
