package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"

	"github.com/charlesbourget/WeatherFetcher/services"
	"github.com/gin-gonic/gin"
)

type WeatherController struct{}

var reFloat = regexp.MustCompile(`(?m)^-?(\d*\.\d+|\d+\.\d*|\d+)$`)
var reInt = regexp.MustCompile(`^\d+$`)

func (h WeatherController) Current(c *gin.Context) {
	lat := c.Query("lat")
	lon := c.Query("lon")
	if lat == "" || lon == "" || !reFloat.MatchString(lat) || !reFloat.MatchString(lon) {
		c.String(http.StatusBadRequest, "Takes two float parameters: lat and lon, ex. lat=45.34&lon=-73.45")
		return
	}
	response, err := services.GetCurrent(lat, lon)
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

func (h WeatherController) Forecast(c *gin.Context) {
	lat := c.Query("lat")
	lon := c.Query("lon")
	cnt := c.Query("cnt")
	if lat == "" || lon == "" || cnt == "" || !reInt.MatchString(cnt) || !reFloat.MatchString(lat) || !reFloat.MatchString(lon) {
		c.String(http.StatusBadRequest, "Takes three parameters: lat, lon and cnt, ex. lat=45.34&lon=-73.45&cnt=3")
		return
	}
	response, err := services.GetForecast(lat, lon, cnt)
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

func (h WeatherController) Weather(c *gin.Context) {
	lat := c.Query("lat")
	lon := c.Query("lon")
	cnt := c.Query("cnt")
	if lat == "" || lon == "" || cnt == "" || !reInt.MatchString(cnt) || !reFloat.MatchString(lat) || !reFloat.MatchString(lon) {
		c.String(http.StatusBadRequest, "Takes three parameters: lat, lon and cnt, ex. lat=45.34&lon=-73.45&cnt=3")
		return
	}
	response, err := services.GetWeather(lat, lon, cnt)
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
