package admin

import "github.com/gin-gonic/gin"

type ReqCreateLandLevel struct {
	Level   int    `json:"level"         binding:"required"`
	Title   string `json:"title"         binding:"required"`
	Imgurl  string `json:"imgurl"        binding:"required"`
	Desc    string `json:"desc"          binding:"required"`
	PlantId int    `json:"plant_id"      binding:"required"`
}

func GetLandList(c *gin.Context) {

}
