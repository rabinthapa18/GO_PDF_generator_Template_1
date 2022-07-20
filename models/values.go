package models

type Values struct {
	Items   []item          `json:"items"`
	Images  []images        `json:"images"`
	Details [][]fieldDetail `json:"details"`
}

type item struct {
	FieldName string `json:"fieldName"`
	Value     string `json:"value"`
}

type images struct {
	Name      string `json:"name"`
	ObjectKey string `json:"objectKey"`
}

type productDetail struct {
	Product []fieldDetail `json:"product"`
}

type fieldDetail struct {
	FieldName string `json:"fieldName"`
	Value     string `json:"value"`
}
