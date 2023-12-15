package dev01

import (
	"log"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func GetExactTime(timeServer string) time.Time {
	response, err := ntp.Query(timeServer)
	if err != nil {
		log.Printf("Error querying NTP server: %v\n", err)
		os.Exit(1)
	}

	// Adding this value to subsequent local system time measurements in order to obtain a more accurate time
	exactTime := time.Now().Add(response.ClockOffset)

	return exactTime
}
