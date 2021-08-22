package main

import (
	"context"
	"fmt"
	"time"

	"github.com/180909/task/db"
	"github.com/180909/task/models"
	"github.com/spf13/cast"
)

func main() {
	db.ConnDB()
	jobs, err := db.GetAllJobs(context.Background())
	if err != nil {
		panic("something error: " + err.Error())
	}
	for i := 0; i < len(jobs); i++ {
		go task(jobs[i])
	}
	time.Sleep(1 * time.Minute)
}

func task(job models.Job) {
	fmt.Sprint("%s, start time: %s \n", job.Name, time.Now().String())
	timer := time.NewTimer(time.Duration(cast.ToInt64(job.Interval)) * time.Second)
	fmt.Println(<-timer.C)
}
