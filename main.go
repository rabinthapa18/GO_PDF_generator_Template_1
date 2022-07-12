// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

package main

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"grrow_pdf/api"
	"grrow_pdf/docs"
	"grrow_pdf/env"
)

func main() {
	env.Config()

	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "GROW PDF API"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Description = "Documentation for PDF API"
	docs.SwaggerInfo.Host = "localhost:3000"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	// controllers.GetS3()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router.POST("/addData", api.GenerateTemp1)
	router.POST("/addToTemplate", api.AddToTemplate)

	router.POST("/uploadTemplate", api.UploadTemplate)

	router.Run(os.Getenv("PORT"))

}
