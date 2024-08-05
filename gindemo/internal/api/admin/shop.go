package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"godemo.com/demo/internal/code"
	"godemo.com/demo/internal/model"
	"godemo.com/demo/internal/util"
)

func GetShopList(c *gin.Context) {
	var r ReqList
	err := c.ShouldBindJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}

	var list []model.Shop
	offset := (r.Page - 1) * r.PageSize
	model.DB.Limit(r.PageSize).Offset(offset).Find(&list)

	var count int64
	model.DB.Model(&model.Shop{}).Count(&count)

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

type ReqCreateShop struct {
	ItemType   int    `json:"item_type"   binding:"required"`
	ItemId     int    `json:"item_id"     binding:"required"`
	ItemTitle  string `json:"item_title"  binding:"required"`
	ItemImgurl string `json:"item_imgurl" binding:"required"`
	ItemDesc   string `json:"item_desc"   binding:"required"`
	ItemPrice  int    `json:"item_price"  binding:"required"`
}

func CreateShopItem(c *gin.Context) {
	var r ReqCreateShop
	err := c.ShouldBindBodyWithJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}

	shopItem := model.Shop{
		ItemType:   r.ItemType,
		ItemId:     r.ItemId,
		ItemTitle:  r.ItemTitle,
		ItemImgurl: r.ItemImgurl,
		ItemDesc:   r.ItemDesc,
		ItemPrice:  r.ItemPrice,
	}
	model.DB.Create(&shopItem)
	if shopItem.Id == 0 {
		util.Error(c, code.CREATE_SHOP_ITEM_ERROR)
		return
	}
	util.Success(c, gin.H{
		"item": shopItem,
	})
}

type ReqEditShop struct {
	Id         int    `json:"id"          binding:"required,gte=0"`
	ItemType   int    `json:"item_type"   binding:"required,gte=0,lte=20"`
	ItemId     int    `json:"item_id"     binding:"required,gte=0"`
	ItemTitle  string `json:"item_title"  binding:"required"`
	ItemImgurl string `json:"item_imgurl" binding:"required"`
	ItemDesc   string `json:"item_desc"   binding:"required"`
	ItemPrice  int    `json:"item_price"  binding:"required,gte=0"`
}

func EditShopItem(c *gin.Context) {
	var r ReqEditShop
	err := c.ShouldBindBodyWithJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}
	shopItem := model.Shop{
		Id:         r.Id,
		ItemType:   r.ItemType,
		ItemId:     r.ItemId,
		ItemTitle:  r.ItemTitle,
		ItemImgurl: r.ItemImgurl,
		ItemDesc:   r.ItemDesc,
		ItemPrice:  r.ItemPrice,
	}
	model.DB.Create(&shopItem)
	if shopItem.Id == 0 {
		util.Error(c, code.UPDATE_SHOP_ITEM_ERROR)
		return
	}
	util.Success(c, gin.H{
		"item": shopItem,
	})
}
