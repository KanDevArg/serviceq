package main

import (
	"model"
	"testing"
	"time"
)

func TestWorkAssigment(t *testing.T) {

	// assumption -- all services are down

	sqp := model.ServiceQProperties{}
	sqp.ListenerPort = "5252"
	sqp.Proto = "http"
	sqp.ServiceList = []model.Endpoint{
		{RawUrl: "http://example.org:2001", Scheme: "http", QualifiedUrl: "http://example.org:2001", Host: "example.org:2001"},
		{RawUrl: "http://example.org:3001", Scheme: "http", QualifiedUrl: "http://example.org:3001", Host: "example.org:3001"},
		{RawUrl: "http://example.org:4001", Scheme: "http", QualifiedUrl: "http://example.org:4001", Host: "example.org:4001"},
		{RawUrl: "http://example.org:5001", Scheme: "http", QualifiedUrl: "http://example.org:5001", Host: "example.org:5001"},
	}
	sqp.MaxConcurrency = 8 // if changing, do check value of duplicateWork
	sqp.EnableDeferredQ = true
	sqp.DeferredQRequestFormats = []string{"ALL"}
	sqp.MaxRetries = 1  // we know it's down
	sqp.RetryGap = 0 // ms
	sqp.IdleGap = 500   // ms
	sqp.RequestErrorLog = make(map[string]int, 2)
	sqp.OutRequestTimeout = 300000

	cw := make(chan int, sqp.MaxConcurrency)
	cr := make(chan interface{}, sqp.MaxConcurrency)

	var reqParam model.RequestParam
	reqParam.Protocol = "HTTP/1.1"
	reqParam.Method = "GET"
	reqParam.RequestURI = "/getRefund"
	reqParam.Headers = make(map[string][]string, 1)
	reqParam.Headers["Content-Type"] = []string{"application/json"}
	reqParam.BodyBuff = nil

	cr <- reqParam
	cw <- 1

	go workBackground(cr, cw, &sqp) // this will start executing req

	// increment/decrement buffer (+1/-1) in creq, cwork and give time to orchestrate

	duplicateWork := int(sqp.MaxConcurrency/2) + 1

	time.Sleep(1000 * time.Millisecond)
	// add req and work again
	for i := 0; i < duplicateWork; i++ {
		cr <- reqParam
		cw <- 1
	}

	time.Sleep(1000 * time.Millisecond)
	// remove all work without removing req
	for i := 0; i < duplicateWork+1; i++ {
		<-cw
	}

	time.Sleep(1000 * time.Millisecond)
	// add half of works
	for i := 0; i < ((duplicateWork + 1) / 2); i++ {
		cw <- 1
	}

	time.Sleep(3000 * time.Millisecond)

	if len(cw) < ((duplicateWork+1)/2)-1 || len(cw) > ((duplicateWork+1)/2) {
		t.Errorf("Work not being orchestrated properly\n")
	}
}
