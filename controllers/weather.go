package controllers

import (
	"encoding/json"
	"github.com/charlesbourget/WeatherFetcher/services"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"regexp"
)

type WeatherController struct{}

var re = regexp.MustCompile(`(?m)^-?(\d*\.\d+|\d+\.\d*|\d+)$`)

func (h WeatherController) Current(c *gin.Context) {
	lat := c.Query("lat")
	lon := c.Query("lon")
	if lat == "" || lon == "" || !re.MatchString(lat) || !re.MatchString(lon) {
		c.String(http.StatusBadRequest, "Takes two float parameters: lat and lon, ex. lat=45.34&lon=-73.45")
		return
	}
	response, err := services.GetWeather(lat, lon)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		log.Println(err)
		return
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		log.Println(err)
		return
	}

	c.Data(http.StatusOK, gin.MIMEJSON, jsonResponse)
}
