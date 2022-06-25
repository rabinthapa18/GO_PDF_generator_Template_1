package models

import "mime/multipart"

type RawData struct {
	Name        name           `form:"name"`
	PhoneNumber phoneNumber    `form:"phoneNumber"`
	ZipAddress  zipAddress     `form:"zipAddress"`
	Address     address        `form:"address"`
	Products    []Products     `form:"products" binding:"requried" swaggerignore:"true"`
	LogoData    LogoData       `form:"logoData"`
	Logo        multipart.File `form:"logo" binding:"required" swaggerignore:"true"`
}

type name struct {
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

type phoneNumber struct {
	PhoneNumber int `json:"phoneNumber"`
	X           int `json:"x"`
	Y           int `json:"y"`
}

type zipAddress struct {
	ZipAddress int `json:"zipAddress"`
	X          int `json:"x"`
	Y          int `json:"y"`
}

type address struct {
	Address string `json:"address"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
}

type Products struct {
	ProductName     productName `json:"productName"`
	ProductQuantity quantity    `json:"productQuantity"`
	ProductPrice    price       `json:"productPrice"`
}

type productName struct {
	Name string `json:"name"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

type quantity struct {
	Quantity int `json:"quantity"`
	X        int `json:"x"`
	Y        int `json:"y"`
}

type price struct {
	Price int `json:"price"`
	X     int `json:"x"`
	Y     int `json:"y"`
}

type LogoData struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}
