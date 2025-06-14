package main

import (
	"shorturl/aserver/logger"
	"time"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

type JobConfig struct {
	service.ServiceConf
}

type JobIns struct {
	JobConfig
	Name string
	stop chan struct{}
}

func NewJobIns(name string, c JobConfig) *JobIns {
	return &JobIns{
		JobConfig: c,
		Name:      name,
		stop:      make(chan struct{}),
	}
}
func (j *JobIns) Run() {
	for {
		logx.Infof("job running %s", j.Name)
		logx.Infow("report ...", logx.LogField{Key: "report", Value: true})
		time.Sleep(1 * time.Second)
		select {
		case <-j.stop:
			logx.Infof("job %s stopped", j.Name)
			return
		default:
		}
	}
}

func main() {
	var c JobConfig
	conf.MustLoad("config.yaml", &c)

	c.MustSetUp()
	slackWriter := logger.NewSlackWriter("https://hooks.slack.com/services/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")
	logx.AddWriter(logx.NewWriter(slackWriter))
	logx.AddGlobalFields(logx.LogField{Key: "server", Value: c.Log.ServiceName})
	defer logx.Close()
	// do your job
	serviceGroup := service.NewServiceGroup()
	serviceGroup.Add(&AService{})
	serviceGroup.Add(&BService{})
	serviceGroup.Start()
}

type AService struct {
}

func (a *AService) Start() {
	for {
		logx.Info("A Service")
		time.Sleep(1 * time.Second)
	}
}

func (a *AService) Stop() {
	logx.Info("A stop recive")
}

type BService struct {
}

func (b *BService) Start() {
	for {
		logx.Info("B Service")
		time.Sleep(1 * time.Second)
	}
}

func (b *BService) Stop() {
	logx.Info("B stop recive")
}
