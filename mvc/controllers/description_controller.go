package controllers

import (
	"alterra/configs"
	"alterra/models/books"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddDescriptionController(c echo.Context) error {
	//Build the request
	apiKey := "?key=" + "AIzaSyCwSiNukiMozrgK3vFErCbIvDEp4Tn8PzU"
	id := "pgjmDAAAQBAJ"
	request := "https://www.googleapis.com/books/v1/volumes/" + id + apiKey
	req, err := http.NewRequest("GET", request, nil)
	if err != nil {
		fmt.Println("Error is req: ", err)
	}
	// create a Client
	client := &http.Client{}

	// Do sends an HTTP request and
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error in send req: ", err)
	}
	var data books.Getbook
	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println(err)
	}

	var desc books.Deskripsi
	desc.Description = data.VolumeInfo.Description
	defer resp.Body.Close()

	result := configs.DB.Create(&desc)
	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create the data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "Succes add book",
		"id":          desc.Id,
		"description": desc.Description,
	})
}
