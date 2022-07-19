package models

type Definitions struct {
	Texts  []text  `json:"texts"`
	Images []image `json:"images"`
}

type text struct {
	FieldName string `json:"fieldName"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Size      int    `json:"size"`
	PageNo    int    `json:"pageNo"`
}

type image struct {
	Name   string `json:"name"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	PageNo int    `json:"pageNo"`
}
