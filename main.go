package main

import (
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	log "github.com/sirupsen/logrus"
)

var Logger *log.Logger

type ProblemResult struct {
	Problem        string `json:"problem"`
	NumberA        int    `json:"numberA"`
	ElapsedSeconds int    `json:"elapsedSeconds"`
	MovesA         int    `json:"movesA"`
	DifferencesA   int    `json:"differencesA"`
	Complete       int    `json:"complete"`
	AttemptsB      int    `json:"attemptsB"`
	SourceAddress  string `json:"sourceAddress"`
	StartTime      int64  `json:"startTime"`
}
type BCheck struct {
	Answer string `json:"answer"`
}

var ResultsA = []ProblemResult{}
var ResultsB = []ProblemResult{}
var timeLimit = 300
var bucketSize = 30

func init() {
	Logger = log.New()
	ResultsA = []ProblemResult{
		{Problem: "A", NumberA: 1, Complete: 1, ElapsedSeconds: 15, MovesA: 25, SourceAddress: "1.2.3.4", StartTime: 1640975675},
		{Problem: "A", NumberA: 2, Complete: 1, ElapsedSeconds: 125, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 2, Complete: 1, ElapsedSeconds: 120, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 2, Complete: 1, ElapsedSeconds: 149, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 2, Complete: 1, ElapsedSeconds: 150, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 1, Complete: 1, ElapsedSeconds: 20, MovesA: 14, SourceAddress: "1.2.3.4", StartTime: 1640975676},
		{Problem: "A", NumberA: 1, Complete: 1, ElapsedSeconds: 45, MovesA: 25, SourceAddress: "1.2.3.4", StartTime: 1640975675},
		{Problem: "A", NumberA: 2, Complete: 1, ElapsedSeconds: 300, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 2, Complete: 1, ElapsedSeconds: 300, MovesA: 14, SourceAddress: "1.2.3.4", StartTime: 1640975676},

		{Problem: "A", NumberA: 1, Complete: 1, ElapsedSeconds: 123, MovesA: 25, SourceAddress: "1.2.3.4", StartTime: 1640975675},
		{Problem: "A", NumberA: 2, Complete: 1, ElapsedSeconds: 300, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 1, Complete: 1, ElapsedSeconds: 111, MovesA: 14, SourceAddress: "1.2.3.4", StartTime: 1640975676},
	}
}

func main() {
	router := gin.Default()
	router.GET("/resultsA", getResultsA)
	router.GET("/resultsB", getResultsB)
	router.GET("/graphA", graphA)
	router.GET("/graphB", graphB)
	router.POST("/add", addResult)
	router.POST("/check", checkAnswer)
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

func checkAnswer(c *gin.Context) {
	var check BCheck
	err := c.BindJSON(&check)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	if check.Answer != "I-KLMN---J----OP" {
		c.AbortWithStatus(400)
		return
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Status(http.StatusOK)
	// IndentedJSON(http.StatusCreated, result)
}
func graphA(c *gin.Context) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Problem A",
		Subtitle: "Number of solutions in each 30 second interval",
	}))
	bar.SetXAxis([]string{"<30", "30", "60", "90", "120", "150", "180", "210", "240", "270", "300+"}).
		AddSeries("Number A1", buildSeriesProblemA(1)).
		AddSeries("Number A2", buildSeriesProblemA(2))
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
	bar.Render(c.Writer)
}
func graphB(c *gin.Context) {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "Problem B",
		Subtitle: "Number of solutions in each 30 second interval",
	}))
	bar.SetXAxis([]string{"<30", "30", "60", "90", "120", "150", "180", "210", "240", "270", "300+"}).
		AddSeries("Problem B", buildSeriesProblemB())
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")
	bar.Render(c.Writer)
}

func buildSeriesProblemA(number int) []opts.BarData {
	items := make([]opts.BarData, 0)
	var buckets int = (timeLimit / bucketSize) + 1
	for i := 0; i < buckets; i++ {
		items = append(items, opts.BarData{Value: 0})
	}
	var bucket int
	for _, result := range ResultsA {
		if result.NumberA == number && resultIncluded(&result) {
			bucket = int(math.Floor(float64(result.ElapsedSeconds) / float64(bucketSize)))
			items[bucket].Value = items[bucket].Value.(int) + 1
		}
	}
	return items
}
func buildSeriesProblemB() []opts.BarData {
	items := make([]opts.BarData, 0)
	var buckets int = (timeLimit / bucketSize) + 1
	for i := 0; i < buckets; i++ {
		items = append(items, opts.BarData{Value: 0})
	}
	var bucket int
	for _, result := range ResultsB {
		if resultIncluded(&result) {
			bucket = int(math.Floor(float64(result.ElapsedSeconds) / float64(bucketSize)))
			items[bucket].Value = items[bucket].Value.(int) + 1
		}
	}
	return items
}

func resultIncluded(result *ProblemResult) bool {
	return (result.Complete == 1) || (result.ElapsedSeconds == timeLimit)
}
