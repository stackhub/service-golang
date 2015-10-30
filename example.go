package main

import (
	"fmt"
	"time"
)

func timeUntilNewYear() {
	t := time.Now().Round(time.Second)
	endOfYear := time.Date(t.Year(), time.December, 31, 23, 59, 59, 59, time.UTC)
	endOfYear = endOfYear.Round(time.Second)

	timeToNewYear := endOfYear.Sub(t)
	hours := int(timeToNewYear.Hours())
	minutes := int(timeToNewYear.Minutes()) - hours*60
	seconds := int(timeToNewYear.Seconds()) - hours*60*60 - minutes*60

	fmt.Printf("We're only %d hours, %d minutes, %d seconds until the new year!\n", hours, minutes, seconds)
}

func main() {
	runs := 0
	for runs < 600 {
		timeUntilNewYear()
		time.Sleep(time.Second * 1)
		runs = runs + 1
	}
}
