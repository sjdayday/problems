package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var Cntl *Control
var Logger *log.Logger

type Control struct {
}
type ProblemResult struct {
	Problem        string `json:"problem"`
	NumberA        int    `json:"numberA"`
	ElapsedSeconds int    `json:"elapsedSeconds"`
	MovesA         int    `json:"movesA"`
	DifferencesA   int    `json:"differencesA"`
	AttemptsB      int    `json:"attemptsB"`
	SourceAddress  string `json:"sourceAddress"`
	StartTime      int64  `json:"startTime"`
}

var Results []ProblemResult

func init() {
	Logger = log.New()
	Results = []ProblemResult{
		{Problem: "A", NumberA: 1, ElapsedSeconds: 123, MovesA: 25, SourceAddress: "1.2.3.4", StartTime: 1640975675},
		{Problem: "A", NumberA: 2, ElapsedSeconds: 300, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 1, ElapsedSeconds: 111, MovesA: 14, SourceAddress: "1.2.3.4", StartTime: 1640975676},
	}

	// metrics.Logger.SetLevel(log.DebugLevel)

	// metrics.SendDynatraceFunc = metrics.SendRequest
	// metrics.SendTsdbFunc = metrics.SendRequest
}

func main() {
	Cntl = &Control{}
	router := gin.Default()
	router.GET("/results", getResults)
	router.POST("/add", addResult)
	// router.Run("localhost:8080")
	router.Run(":8080")
	// metrics.Cntl.KpathsConfig = &metrics.KpathsConfig{}
	// err := metrics.Cntl.KpathsConfig.GetDynatraceMetricsPutIntoTsdb()

	// if err != nil {
	// 	metrics.Logger.Debugf("Received error message: %v", err.Error())
	// }
}

func getResults(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Results)
}
func addResult(c *gin.Context) {
	var result ProblemResult
	if err := c.BindJSON(&result); err != nil {
		return
	}
	Results = append(Results, result)
	c.IndentedJSON(http.StatusCreated, result)
}
