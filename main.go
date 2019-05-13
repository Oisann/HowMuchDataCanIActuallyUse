package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var dataMB int
var renewDate int
var downloadKb int
var usedDataMB int
var dontCarePercentage float32

func init() {
	cap, err := strconv.ParseInt(os.Getenv("CAP"), 10, 32)
	if err != nil {
		log.Fatal(err.Error())
	}
	dataMB = int(cap)

	day, err := strconv.ParseInt(os.Getenv("RENEW"), 10, 32)
	if err != nil {
		log.Fatal(err.Error())
	}
	renewDate = int(day)

	download, err := strconv.ParseInt(os.Getenv("DOWNLOAD"), 10, 32)
	if err != nil {
		log.Fatal(err.Error())
	}
	downloadKb = int(download)

	used, err := strconv.ParseInt(os.Getenv("USED"), 10, 32)
	if err != nil {
		log.Fatal(err.Error())
	}
	usedDataMB = int(used)

	percentage, err := strconv.ParseFloat(os.Getenv("PERCENTAGE"), 32)
	if err != nil {
		log.Fatal(err.Error())
	}
	dontCarePercentage = float32(percentage)
}

func main() {
	currentPercentageUsed := (float32(usedDataMB) / float32(dataMB)) * 100.0
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
