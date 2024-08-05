package admin

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"godemo.com/demo/internal/code"
	"godemo.com/demo/internal/model"
	"godemo.com/demo/internal/util"
)

type ReqTaskReward struct {
	Type   int `json:"type"     binding:"required"`
	Id     int `json:"id"       binding:"required"`
	Number int `json:"number"   binding:"required"`
}

func GetTaskList(c *gin.Context) {
	var taskList []model.Task
	model.DB.Find(&taskList)
	util.Success(c, gin.H{
		"items": taskList,
	})
}

type ReqCreateTask struct {
	Title  string        `json:"title"         binding:"required"`
	Desc   string        `json:"desc"          binding:"required"`
	Step   int           `json:"step"          binding:"required"`
	Type   int           `json:"type"          binding:"required"`
	Reward ReqTaskReward `json:"reward"        binding:"required"`
}

func CreateTaskItem(c *gin.Context) {
	var r ReqCreateTask
	err := c.ShouldBindBodyWithJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}

	reward, _ := json.Marshal(r.Reward)
	task := model.Task{
		Title:  r.Title,
		Desc:   r.Desc,
		Step:   r.Step,
		Type:   r.Type,
		Reward: string(reward),
	}
	model.DB.Create(&task)
	if task.Id == 0 {
		util.Error(c, code.TASK_CREATE_FAIL)
		return
	}
	util.Success(c, gin.H{
		"task": task,
	})
}

type ReqEditTask struct {
	Id     int           `json:"id"            binding:"required"`
	Title  string        `json:"title"         binding:"required"`
	Desc   string        `json:"desc"          binding:"required"`
	Step   int           `json:"step"          binding:"required"`
	Type   int           `json:"type"          binding:"required"`
	Reward ReqTaskReward `json:"reward"        binding:"required"`
}

func EditTaskItem(c *gin.Context) {
	var r ReqEditTask
	err := c.ShouldBindBodyWithJSON(&r)
	if err != nil {
		util.Error(c, code.PARAMS_ERROR)
		return
	}

	reward, _ := json.Marshal(r.Reward)
	task := model.Task{
		Id:     r.Id,
		Title:  r.Title,
		Desc:   r.Desc,
		Step:   r.Step,
		Type:   r.Type,
		Reward: string(reward),
	}
	model.DB.Save(&task)
	util.Success(c, gin.H{
		"task": task,
	})
}
