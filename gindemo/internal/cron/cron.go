package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func InitCron() {
	c := cron.New(cron.WithSeconds())
	// c.AddFunc("@every 5s", func() {
	// 	fmt.Println("tick every 5 second")
	// })

	// c.AddFunc("0/5 * * * * *", func() {
	// 	fmt.Println("tick every 5 second")
	// })
	c.AddFunc("0-15/5 * * * * *", func() {
		fmt.Println("tick every 5 second")
	})
	c.AddFunc("@every 1m", func() {
		fmt.Println("tick every 1 minute")
	})
	c.AddFunc("0 */5 * * * *", func() {
		fmt.Println("tick every 5 minute")
	})
	c.AddFunc("00 13 14 * * *", func() {
		fmt.Println("tick every day at 00 13 14")
	})
	c.Start()
}
