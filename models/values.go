package models

type Values struct {
	Items  []item   `json:"items"`
	Images []images `json:"images"`
}

type item struct {
	FieldName string `json:"fieldName"`
	Value     string `json:"value"`
}

type images struct {
	Name      string `json:"name"`
	ObjectKey string `json:"objectKey"`
}
