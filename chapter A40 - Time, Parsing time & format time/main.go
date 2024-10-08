package main

import (
	"fmt"
	"time"
)

func main() {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)

	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
    fmt.Printf("time2 %v\n", time2)

	var now = time.Now()
	fmt.Println("year:", now.Year(), "month:", now.Month())


	//A40.3 Parsing dari string ke time.Time
	fmt.Println("//A40.3 Parsing dari string ke time.Time")
	var layoutFormat, value string
	var date time.Time

	layoutFormat = "2006-01-02 15:04:05"
	value = "2015-09-02 08:04:00"
	date, _ = time.Parse(layoutFormat, value)
	fmt.Println(value, "\t-v", date.String())

	//Buat ngubah time.Time ke string
	var date2, _ = time.Parse(time.RFC822, "02 Sep 15 08:00 WIB")
	var dateS1 = date2.Format("Monday 02, January 2006 15:04 MST")
	fmt.Printf("dateS1 = %T\n", date2)
	fmt.Println("dateS1", dateS1)
	// Wednesday 02, September 2015 08:00 WIB
	var dateS2 = date.Format(time.RFC3339)
	fmt.Printf("dateS2 = %T\n", dateS2)
	fmt.Println("dateS2", dateS2)

	//A40.6 Handle error parsing time.Time
	var date3, err = time.Parse("06 Jan 15", "02 Sep 15 08:00 WIB")
	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	fmt.Println(date3)
}