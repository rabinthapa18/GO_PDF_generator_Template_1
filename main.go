// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"grrow_pdf/api"
	"grrow_pdf/docs"
)

func main() {

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "GROW PDF API"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Description = "Documentation for PDF API"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()

	// v1 := router.Group("")
	// {
	// 	pdfData := v1.Group("/addData")
	// 	{
	// 		pdfData.POST("", api.AddData)
	// 	}
	// }

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router.GET("/getData", api.GetData)

	router.POST("/addData", api.AddData)

	router.Run("localhost:3000")
}
