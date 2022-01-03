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

var ResultsA = []ProblemResult{}
var ResultsB = []ProblemResult{}

func init() {
	Logger = log.New()

	// metrics.Logger.SetLevel(log.DebugLevel)

	// metrics.SendDynatraceFunc = metrics.SendRequest
	// metrics.SendTsdbFunc = metrics.SendRequest
}

func main() {
	Cntl = &Control{}
	router := gin.Default()
	router.GET("/resultsA", getResultsA)
	router.GET("/resultsB", getResultsB)
	router.POST("/add", addResult)
	router.StaticFS("/problem", http.Dir("./problem"))
	router.Run(":8080")
}

func getResultsA(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ResultsA)
}
func getResultsB(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, ResultsB)
}

func addResult(c *gin.Context) {
	var result ProblemResult
	err := c.BindJSON(&result)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	if result.Problem == "A" {
		ResultsA = append(ResultsA, result)
	} else if result.Problem == "B" {
		ResultsB = append(ResultsB, result)
	} else {
		c.AbortWithStatus(400)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.IndentedJSON(http.StatusCreated, result)
}
