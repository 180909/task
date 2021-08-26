package main

import (
	"fmt"
	"github.com/spf13/cast"
	"time"

	"github.com/180909/task/db"
	"github.com/180909/task/models"
)

func main() {
	db.ConnDB()
	jobs, err := db.GetAllJobs()
	if err != nil {
		panic("something error: " + err.Error())
	}
	for i := 0; i < len(jobs); i++ {
		go task(jobs[i])
	}
	time.Sleep(1 * time.Minute)
}

func task(job models.Job) {
	for {
		fmt.Sprintf("%s, start time: %s \n", job.Name, time.Now().String())
		timer := time.NewTimer(time.Duration(cast.ToInt64(job.Interval)) * time.Second)
		fmt.Println(<-timer.C)
	}
}
