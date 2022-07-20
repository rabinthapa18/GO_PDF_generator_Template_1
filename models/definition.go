package models

type Definitions struct {
	Texts   []text  `json:"texts"`
	Images  []image `json:"images"`
	Details details `json:"details"`
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

type details struct {
	Schema     []schema `json:"schema"`
	IncrementY int      `json:"increment"`
	PageNo     int      `json:"pageNo"`
	Size       int      `json:"size"`
}

type schema struct {
	FieldName string `json:"fieldName"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
}
