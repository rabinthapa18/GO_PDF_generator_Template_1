package models

import "mime/multipart"

type PdfData struct {
	Name        string         `form:"name"`
	PhoneNumber int            `form:"phoneNumber"`
	ZipAddress  int            `form:"zipAddress"`
	Address     string         `form:"address"`
	Products    []ProductData  `form:"products" swaggerignore:"true"`
	Logo        multipart.File `form:"logo" binding:"required" swaggerignore:"true"`
}
