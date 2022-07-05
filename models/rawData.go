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
	Template    int            `form:"template" binding:"required"`
}

type name struct {
	Name string `json:"name"`
	Size int    `json:"size"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

type phoneNumber struct {
	PhoneNumber int `json:"phoneNumber"`
	Size        int `json:"size"`
	X           int `json:"x"`
	Y           int `json:"y"`
}

type zipAddress struct {
	ZipAddress int `json:"zipAddress"`
	Size       int `json:"size"`
	X          int `json:"x"`
	Y          int `json:"y"`
}

type address struct {
	Address string `json:"address"`
	Size    int    `json:"size"`
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
	Size int    `json:"size"`
	X    int    `json:"x"`
	Y    int    `json:"y"`
}

type quantity struct {
	Quantity int `json:"quantity"`
	Size     int `json:"size"`
	X        int `json:"x"`
	Y        int `json:"y"`
}

type price struct {
	Price int `json:"price"`
	Size  int `json:"size"`
	X     int `json:"x"`
	Y     int `json:"y"`
}

type LogoData struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
}
