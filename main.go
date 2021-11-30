package main

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/spf13/cast"
	"gopkg.in/yaml.v2"
)

type Jobs struct {
	Job []*Job `json:"job"`
}

type Job struct {
	Name     string `json:"name"`
	Interval int    `json:"interval"`
}

func getConfig() *Jobs {
	c := &Jobs{}
	yamlFile, err := ioutil.ReadFile("./jobs/jobs.yml")
	if err != nil {
		panic("Can't read file")
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		panic("Parse yaml file error")
	}
	return c
}

func main() {
	jobs := getConfig().Job
	for i := 0; i < len(jobs); i++ {
		go task(*jobs[i])
	}
	time.Sleep(1 * time.Minute)
}

func task(job Job) {
	for {
		fmt.Println(fmt.Sprintf("%s, start time: %s", job.Name, time.Now().String()))
		timer := time.NewTimer(time.Duration(cast.ToInt64(job.Interval)) * time.Second)
		fmt.Println(<-timer.C)
	}
}
