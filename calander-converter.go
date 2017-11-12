package main

import (
	"fmt"
	"github.com/lestrrat/go-ical"
	"time"
)

func run() {
	f := GetSettings().IcalLocation
	p := ical.NewParser()
	c, err := p.ParseFile(f)
	if err != nil {
		fmt.Println("Unable to parse file")
	}

	var dbpath string = GetSettings().DatabaseLocation
	db := InitDb(dbpath)
	defer db.Close()
	CreateTable(db)
	// snip
	for e := range c.Entries() {
		ev, ok := e.(*ical.Event)
		if !ok {
			continue
		}
		summary, err := ev.GetProperty("summary")
		if err == false {
			fmt.Println("Error getting summary")
		}

		date, err := ev.GetProperty("dtstart")
		if err == false {
			fmt.Println("Error getting dtstart")
		}

		layout := "20060102T150405Z"
		t, err2 := time.Parse(layout, date.RawValue())

		if err2 != nil {
			fmt.Println(err)
		}
		//if t.Year() == 2017 || t.Year() == 2018 {
		//    fmt.Println(summary.RawValue())
		//    fmt.Println(date.RawValue())
		//}
		StoreItem(db, Match{summary.RawValue(), t})
		for _, i := range ReadItem(db) {
			fmt.Println(t)
			fmt.Println(i.date)
		}

	}

}
