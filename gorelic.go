package gorelic

import (
	"fmt"
	"github.com/go-martini/martini"
	metrics "github.com/yvasiyarov/go-metrics"
	"github.com/yvasiyarov/gorelic"
	"time"
)

var agent *gorelic.Agent

func Handler(c martini.Context) {
	startTime := time.Now()
	c.Next()
	agent.HTTPTimer.UpdateSince(startTime)
}

func InitNewrelicAgent(license string, appname string, verbose bool) error {

	if license == "" {
		return fmt.Errorf("Please specify NewRelic license")
	}

	agent = gorelic.NewAgent()
	agent.NewrelicLicense = license

	agent.HTTPTimer = metrics.NewTimer()
	agent.CollectHTTPStat = true
	agent.Verbose = verbose

	agent.NewrelicName = appname
	agent.Run()
	return nil
}
