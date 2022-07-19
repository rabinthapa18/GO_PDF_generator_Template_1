package api

import (
	"encoding/json"
	"fmt"
	"grrow_pdf/controllers"
	"grrow_pdf/models"
	"net/http"
)

// request body for adding data to template
// @Summary      Add data to template
// @Description  Add data to template
// @Tags         Add Data to Template
// @Accept       application/json
// @Produce      json
// @Param        body body models.RequestBody true "body"
// @Success      200  {file}    file
// @Failure      400  {string}  error
// @Failure      404  {string}  error
// @Failure      500  {string}  error
// @Router       /addDataToTemplate [POST]
func AddDataToTemplate(res http.ResponseWriter, req *http.Request) {

	//cors
	res.Header().Set("Access-Control-Allow-Origin", "*")
	res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	res.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	//end cors

	var newData models.RequestBody
	err := json.NewDecoder(req.Body).Decode(&newData)
	if err != nil {
		fmt.Println(err.Error())
	}

	// get template from body
	template := newData.Template

	// get difinitions from body
	definitions := newData.Definitions

	// get value from body
	values := newData.Values

	byteData := controllers.AddToTemplate(template, definitions, values)

	res.Write(byteData)

}
