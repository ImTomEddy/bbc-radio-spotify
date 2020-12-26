package jobs

import "github.com/robfig/cron/v3"

//SetupJobs sets up any Cron jobs
func SetupJobs() {
	c := cron.New()
	c.AddFunc("* * * * *", UpdateInfo)
	c.Start()
}
