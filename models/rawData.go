package models

type RawData struct {
	Name        name        `json:"name"`
	PhoneNumber phoneNumber `json:"phoneNumber"`
	ZipAddress  zipAddress  `json:"zipAddress"`
	Address     address     `json:"address"`
	Products    []Products  `json:"products" swaggerignore:"true"`
	LogoData    LogoData    `json:"logoData"`
	SealData    SealData    `json:"sealData"`
	Template    int         `json:"template" binding:"required"`
}

type name struct {
	Name   string `json:"name"`
	Size   int    `json:"size"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	PageNo int    `json:"pageNo"`
}

type phoneNumber struct {
	PhoneNumber int `json:"phoneNumber"`
	Size        int `json:"size"`
	X           int `json:"x"`
	Y           int `json:"y"`
	PageNo      int `json:"pageNo"`
}

type zipAddress struct {
	ZipAddress int `json:"zipAddress"`
	Size       int `json:"size"`
	X          int `json:"x"`
	Y          int `json:"y"`
	PageNo     int `json:"pageNo"`
}

type address struct {
	Address string `json:"address"`
	Size    int    `json:"size"`
	X       int    `json:"x"`
	Y       int    `json:"y"`
	PageNo  int    `json:"pageNo"`
}

type Products struct {
	ProductName     productName `json:"productName"`
	ProductQuantity quantity    `json:"productQuantity"`
	ProductPrice    price       `json:"productPrice"`
}

type productName struct {
	Name   string `json:"name"`
	Size   int    `json:"size"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	PageNo int    `json:"pageNo"`
}

type quantity struct {
	Quantity int `json:"quantity"`
	Size     int `json:"size"`
	X        int `json:"x"`
	Y        int `json:"y"`
	PageNo   int `json:"pageNo"`
}

type price struct {
	Price  int `json:"price"`
	Size   int `json:"size"`
	X      int `json:"x"`
	Y      int `json:"y"`
	PageNo int `json:"pageNo"`
}

type LogoData struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
	PageNo int `json:"pageNo"`
}

type SealData struct {
	X      int `json:"x"`
	Y      int `json:"y"`
	Width  int `json:"width"`
	Height int `json:"height"`
	PageNo int `json:"pageNo"`
}
