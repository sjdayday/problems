package main

import (
	// "encoding/json"
	// "errors"
	// "fmt"q

	// "time"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	// "cdts.kp.org/metrics"
)

func TestAddProblemResultA(t *testing.T) {
	fmt.Println(t.Name())
	gin.SetMode(gin.TestMode)
	ResultsA = []ProblemResult{
		{Problem: "A", NumberA: 1, ElapsedSeconds: 123, MovesA: 25, SourceAddress: "1.2.3.4", StartTime: 1640975675},
		{Problem: "A", NumberA: 2, ElapsedSeconds: 300, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 1, ElapsedSeconds: 111, MovesA: 14, SourceAddress: "1.2.3.4", StartTime: 1640975676},
	}
	assert.Equal(t, 3, len(ResultsA))
	var jsonStr string = `{"problem": "A", "numberA": 1, "elapsedSeconds": 125, "movesA": 18, "sourceAddress": "1.2.3.4", "startTime": 1640975680}`
	// req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	io.WriteString(w, jsonStr)

	context, _ := gin.CreateTestContext(w)
	// from: https://stackoverflow.com/questions/67508787/how-to-mock-a-gin-context
	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST" // or PUT
	context.Request.Header.Set("Content-Type", "application/json")

	// jsonbytes, err := json.Marshal(ProblemResult{})
	// if err != nil {
	//     panic(err)
	// }

	context.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(jsonStr)))
	addResult(context)
	assert.Equal(t, 4, len(ResultsA))
	assert.Equal(t, 0, len(ResultsB))
	assert.Equal(t, 125, ResultsA[3].ElapsedSeconds)

	resp := w.Result()
	// body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)

	// fmt.Println(resp.Header.Get("Content-Type"))
	// fmt.Println(string(body))
}
func TestAddProblemResultB(t *testing.T) {
	fmt.Println(t.Name())
	gin.SetMode(gin.TestMode)
	ResultsA = []ProblemResult{
		{Problem: "A", NumberA: 1, ElapsedSeconds: 123, MovesA: 25, SourceAddress: "1.2.3.4", StartTime: 1640975675},
		{Problem: "A", NumberA: 2, ElapsedSeconds: 300, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 1, ElapsedSeconds: 111, MovesA: 14, SourceAddress: "1.2.3.4", StartTime: 1640975676},
	}
	assert.Equal(t, 3, len(ResultsA))
	assert.Equal(t, 0, len(ResultsB))
	var jsonStr string = `{"problem": "B", "numberA": 0, "elapsedSeconds": 30, "movesA": 0, "sourceAddress": "1.2.3.4", "startTime": 1640975680}`
	// req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	io.WriteString(w, jsonStr)

	context, _ := gin.CreateTestContext(w)
	// from: https://stackoverflow.com/questions/67508787/how-to-mock-a-gin-context
	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST" // or PUT
	context.Request.Header.Set("Content-Type", "application/json")

	// jsonbytes, err := json.Marshal(ProblemResult{})
	// if err != nil {
	//     panic(err)
	// }

	context.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(jsonStr)))
	addResult(context)
	assert.Equal(t, 3, len(ResultsA))
	assert.Equal(t, 1, len(ResultsB))
	assert.Equal(t, 30, ResultsB[0].ElapsedSeconds)

	resp := w.Result()
	// body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, 200, resp.StatusCode)

	// fmt.Println(resp.Header.Get("Content-Type"))
	// fmt.Println(string(body))
}

func TestInvalidProblemResultReturns400(t *testing.T) {
	fmt.Println(t.Name())
	gin.SetMode(gin.TestMode)
	ResultsA = []ProblemResult{
		{Problem: "A", NumberA: 1, ElapsedSeconds: 123, MovesA: 25, SourceAddress: "1.2.3.4", StartTime: 1640975675},
		{Problem: "A", NumberA: 2, ElapsedSeconds: 300, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 1, ElapsedSeconds: 111, MovesA: 14, SourceAddress: "1.2.3.4", StartTime: 1640975676},
	}
	assert.Equal(t, 3, len(ResultsA))
	var jsonStr string = `some not very JSON-like string`
	// req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	io.WriteString(w, jsonStr)

	context, _ := gin.CreateTestContext(w)
	// from: https://stackoverflow.com/questions/67508787/how-to-mock-a-gin-context
	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST" // or PUT
	context.Request.Header.Set("Content-Type", "application/json")

	// jsonbytes, err := json.Marshal(ProblemResult{})
	// if err != nil {
	//     panic(err)
	// }

	context.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(jsonStr)))
	addResult(context)
	assert.Equal(t, 3, len(ResultsA))

	// per: https://github.com/gin-gonic/gin/issues/1120
	// the resp.StatusCode is not updated consistently during testing,
	// so use context.Writer.Status() instead
	// resp := w.Result()
	// body, _ := io.ReadAll(resp.Body)
	// assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, 400, context.Writer.Status())
}
func TestInvalidProblemValueReturns400(t *testing.T) {
	fmt.Println(t.Name())
	gin.SetMode(gin.TestMode)
	ResultsA = []ProblemResult{
		{Problem: "A", NumberA: 1, ElapsedSeconds: 123, MovesA: 25, SourceAddress: "1.2.3.4", StartTime: 1640975675},
		{Problem: "A", NumberA: 2, ElapsedSeconds: 300, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 1, ElapsedSeconds: 111, MovesA: 14, SourceAddress: "1.2.3.4", StartTime: 1640975676},
	}
	assert.Equal(t, 3, len(ResultsA))
	var jsonStr string = `{"problem": "NotANorB", "numberA": 1, "elapsedSeconds": 125, "movesA": 18, "sourceAddress": "1.2.3.4", "startTime": 1640975680}`
	// req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	io.WriteString(w, jsonStr)

	context, _ := gin.CreateTestContext(w)
	// from: https://stackoverflow.com/questions/67508787/how-to-mock-a-gin-context
	context.Request = &http.Request{
		Header: make(http.Header),
	}
	context.Request.Method = "POST" // or PUT
	context.Request.Header.Set("Content-Type", "application/json")

	// jsonbytes, err := json.Marshal(ProblemResult{})
	// if err != nil {
	//     panic(err)
	// }

	context.Request.Body = io.NopCloser(bytes.NewBuffer([]byte(jsonStr)))
	addResult(context)
	assert.Equal(t, 3, len(ResultsA))

	// per: https://github.com/gin-gonic/gin/issues/1120
	// the resp.StatusCode is not updated consistently during testing,
	// so use context.Writer.Status() instead
	// resp := w.Result()
	// body, _ := io.ReadAll(resp.Body)
	// assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, 400, context.Writer.Status())
}
func TestBuildChartItemsForProblemA(t *testing.T) {
	fmt.Println(t.Name())
	ResultsA = []ProblemResult{
		{Problem: "A", NumberA: 1, ElapsedSeconds: 15, MovesA: 25, SourceAddress: "1.2.3.4", StartTime: 1640975675},
		{Problem: "A", NumberA: 2, ElapsedSeconds: 125, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 2, ElapsedSeconds: 120, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 2, ElapsedSeconds: 149, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 2, ElapsedSeconds: 150, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 1, ElapsedSeconds: 20, MovesA: 14, SourceAddress: "1.2.3.4", StartTime: 1640975676},
		{Problem: "A", NumberA: 1, ElapsedSeconds: 45, MovesA: 25, SourceAddress: "1.2.3.4", StartTime: 1640975675},
		{Problem: "A", NumberA: 2, ElapsedSeconds: 300, MovesA: 51, SourceAddress: "1.2.3.5", StartTime: 1640975670},
		{Problem: "A", NumberA: 2, ElapsedSeconds: 300, MovesA: 14, SourceAddress: "1.2.3.4", StartTime: 1640975676},
	}
	items := buildSeriesProblemA(1)
	assert.Equal(t, 2, items[0].Value)
	assert.Equal(t, 1, items[1].Value)
	assert.Equal(t, 11, len(items))
	items = buildSeriesProblemA(2)
	assert.Equal(t, 0, items[0].Value, "0-29")
	assert.Equal(t, 0, items[1].Value, "30-59")
	assert.Equal(t, 0, items[2].Value, "60-89")
	assert.Equal(t, 0, items[3].Value, "90-119")
	assert.Equal(t, 3, items[4].Value, "120-149")
	assert.Equal(t, 1, items[5].Value, "150-179")
	assert.Equal(t, 0, items[6].Value, "180-209")
	assert.Equal(t, 0, items[7].Value, "210-239")
	assert.Equal(t, 0, items[8].Value, "240-269")
	assert.Equal(t, 0, items[9].Value, "270-299")
	assert.Equal(t, 2, items[10].Value, "300 (time limit)")

	assert.Equal(t, 11, len(items))
}
