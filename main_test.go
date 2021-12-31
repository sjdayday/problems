package main

import (
	// "encoding/json"
	// "errors"
	"fmt"

	// "time"

	"testing"
	// "cdts.kp.org/metrics"
)

func TestMain(t *testing.T) {
	fmt.Println(t.Name())
	// metrics.Logger = log.New()
	// metrics.Logger.SetLevel(log.InfoLevel)
	// var hook *test.Hook
	// metrics.Logger, hook = test.NewNullLogger()
	// metrics.Logger.SetLevel(log.InfoLevel)

	// // Cntl = &Control{}
	// // Cntl.KpathsConfig = &KpathsConfig{}
	// var jsonStrings []*string
	// jsonStrings = append(jsonStrings, &metrics.TestServicesJson, &metrics.TestCountsJson, &metrics.TestErrorsJson,
	// 	&metrics.TestResponseTimeJson, &metrics.TestCountsJson2, &metrics.TestErrorsJson2, &metrics.TestResponseTimeJson2)
	// metrics.SendDynatraceFunc = func(request *http.Request) ([]byte, error) {
	// 	tempJson := *jsonStrings[0]
	// 	jsonStrings = jsonStrings[1:]
	// 	// fmt.Printf("printing: %v\n", tempJson)
	// 	return []byte(tempJson), nil
	// }
	// var jsonTsdbStrings []*string
	// jsonTsdbStrings = append(jsonTsdbStrings, &metrics.TestSuccessPutTsdbJsonFour, &metrics.TestSuccessPutTsdbJsonFour, &metrics.TestSuccessPutTsdbJson, &metrics.TestSuccessPutTsdbJsonFour, &metrics.TestSuccessPutTsdbJsonFour, &metrics.TestSuccessPutTsdbJson)
	// metrics.SendTsdbFunc = func(request *http.Request) ([]byte, error) {
	// 	tempTsdbJson := *jsonTsdbStrings[0]
	// 	jsonTsdbStrings = jsonTsdbStrings[1:]
	// 	// fmt.Printf("printing: %v\n", tempJson)
	// 	return []byte(tempTsdbJson), nil
	// }
	// hook.Reset()
	// //var successPutTsdbJson string = `{"success":3,"failed":0,"errors":[]}`
	// // err := Cntl.KpathsConfig.GetDynatraceMetricsPutIntoTsdb()
	// go main()
	// time.Sleep(time.Millisecond * 5)

	// assert.Equal(t, 6, len(hook.Entries))
	// assert.Equal(t, hook.Entries[0].Message, `Read 2 services from Dynatrace`)
	// assert.Equal(t, hook.Entries[1].Message, `Read 24 data points from Dynatrace`)
	// assert.Equal(t, hook.Entries[2].Message, `Created 22 data points for opentsdb`)
	// assert.Equal(t, hook.Entries[3].Message, `Data points not written to opentsdb because null/zero: 2`)
	// assert.Equal(t, hook.Entries[4].Message, `Successfully wrote 22 data points to opentsdb`)
	// assert.Equal(t, hook.Entries[5].Message, `Failed to write 0 data points to opentsdb`)

	// assert.Equal(t, metrics.Cntl.Stats.DynatracePoints, metrics.Cntl.Stats.TsdbPointsCreated+metrics.Cntl.Stats.TsdbPointsDropped)
	// assert.Equal(t, metrics.Cntl.Stats.TsdbPointsCreated, metrics.Cntl.Stats.TsdbPointsSuccess)
	// assert.Equal(t, metrics.Cntl.Stats.TsdbPointsFailed, 0)
	// hook.Reset()

}
