package main

import (
	"flag"
	"fmt"
	"time"
)

var dataMB int
var renewDate int
var downloadKb int
var usedDataMB int
var dontCarePercentage float64

func init() {
	//dataMB = 1000000
	flag.IntVar(&dataMB, "cap", 1000000, "Data cap in MB")
	//renewDate = 1
	flag.IntVar(&renewDate, "renew", 1, "Renew date (Default 1st)")
	//downloadKb = 10000
	flag.IntVar(&downloadKb, "download", 10000, "Download speed in Kbps")
	//usedDataMB = 0
	flag.IntVar(&usedDataMB, "used", 0, "Data used in MB")
	//dontCarePercentage = 50
	flag.Float64Var(&dontCarePercentage, "percentage", 50, "Percentage used before we actually care")
}

func main() {
	currentPercentageUsed := (float64(usedDataMB) / float64(dataMB)) * 100.0
	if currentPercentageUsed < dontCarePercentage {
		fmt.Printf("%d", downloadKb)
		return
	}

	nextRenew := getNextRenewDate()
	secondsLeft := int(time.Since(nextRenew)/time.Second) * -1

	dataLeftMB := dataMB - usedDataMB

	MBpsLeft := float64(dataLeftMB) / float64(secondsLeft)
	mbpsLeft := MBpsLeft * 8.0
	kbpsLeft := int(mbpsLeft * 1000)
	fmt.Printf("%d", min(kbpsLeft, downloadKb))
}

func getNextRenewDate() time.Time {
	t := time.Now()
	n := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	return n.AddDate(0, 1, renewDate-1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
