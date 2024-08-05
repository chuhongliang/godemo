package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"godemo.com/demo/internal/code"
	"godemo.com/demo/internal/model"
	"godemo.com/demo/internal/util"
)

type ReqList struct {
	Page     int `json:"page" binding:"required,gte=0,lte=120"`
	PageSize int `json:"page_size" binding:"required,gte=0,lte=120"`
}

func GetSeedList(c *gin.Context) {
	var r ReqList
	err := c.ShouldBindJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}

	var list []model.Seed
	offset := (r.Page - 1) * r.PageSize
	model.DB.Limit(r.PageSize).Offset(offset).Find(&list)

	var count int64
	model.DB.Model(&model.Seed{}).Count(&count)

	totalPage := 1
	if count != 0 {
		fmt.Println("count:", count)
		totalPage = int(count / int64(r.PageSize))
	}
	fmt.Println("totalPage:", totalPage)
	if totalPage <= 0 {
		totalPage = 1
	}
	util.Success(c, gin.H{
		"items":      list,
		"page":       r.Page,
		"page_size":  r.PageSize,
		"total_page": totalPage,
	})
}

type ReqCreateSeed struct {
	Title       string `json:"title"         binding:"required"`
	Level       int    `json:"level"         binding:"required"`
	Imgurl      string `json:"imgurl"        binding:"required"`
	GrowTime    int    `json:"grow_time"     binding:"required"`
	Desc        string `json:"desc"          binding:"required"`
	PlantId     int    `json:"plant_id"      binding:"required"`
	PlantNumber int    `json:"plant_number"  binding:"required"`
}

func CreateSeedItem(c *gin.Context) {
	var r ReqCreateSeed
	err := c.ShouldBindBodyWithJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}
	seed := model.Seed{
		Title:       r.Title,
		Level:       r.Level,
		Imgurl:      r.Imgurl,
		GrowTime:    r.GrowTime,
		Desc:        r.Desc,
		PlantId:     r.PlantId,
		PlantNumber: r.PlantNumber,
	}
	model.DB.Create(&seed)
	if seed.Id == 0 {
		util.Error(c, code.PARAMS_ERROR)
		return
	}
	util.Success(c, gin.H{
		"seed": seed,
	})
}

type ReqEditSeed struct {
	Id          int    `json:"id"            binding:"required"`
	Title       string `json:"title"         binding:"required"`
	Level       int    `json:"level"         binding:"required"`
	Imgurl      string `json:"imgurl"        binding:"required"`
	GrowTime    int    `json:"grow_time"     binding:"required"`
	Desc        string `json:"desc"          binding:"required"`
	PlantId     int    `json:"plant_id"      binding:"required"`
	PlantNumber int    `json:"plant_number"  binding:"required"`
}

func EditSeedItem(c *gin.Context) {
	var r ReqEditSeed
	err := c.ShouldBindBodyWithJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}
	seed := model.Seed{
		Id:          r.Id,
		Title:       r.Title,
		Level:       r.Level,
		Imgurl:      r.Imgurl,
		GrowTime:    r.GrowTime,
		Desc:        r.Desc,
		PlantId:     r.PlantId,
		PlantNumber: r.PlantNumber,
	}
	model.DB.Save(&seed)
	util.Success(c, gin.H{
		"seed": seed,
	})
}
