package server

import (
	"github.com/charlesbourget/WeatherFetcher/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		weatherGroup := v1.Group("weather")
		{
			weather := new(controllers.WeatherController)
			weatherGroup.GET("/current", weather.Current)
			weatherGroup.GET("/forecast", weather.Forecast)
			weatherGroup.GET("/", weather.Weather)
		}
	}
	return router

}
