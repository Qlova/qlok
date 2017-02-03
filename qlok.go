package main

import "time"
import "os"
import "fmt"
import "github.com/Splizard/ctrls"

type Qlok struct {
	Hours	time.Time
	Start 	time.Time
}

func main() {

	var qlok Qlok
	ctrls.Load(&qlok, "~/qlok/qlockfile")

	switch len(os.Args) {
		case 1:
			fmt.Println("Usage: qlock [start/stop/reset/total/add/remove]")
		
		case 3:
		
			d, err := time.ParseDuration(os.Args[2])
			if err != nil {
				fmt.Println(err)
				return
			}
		
			switch os.Args[1] {
				case "add":
			
					qlok.Hours = qlok.Hours.Add(d)
					
				case "remove":
				
					qlok.Hours = qlok.Hours.Add(-d)			
			}
		
		case 2:
	switch os.Args[1] {
	
		case "total":
			fmt.Println(qlok.Hours.Sub(time.Time{}))
		
		case "start":
			qlok.Start = time.Now()
		
		case "stop":
			qlok.Hours = qlok.Hours.Add(time.Since(qlok.Start))
		
		case "reset":
			qlok.Hours = time.Time{}
			
	}}
	
	ctrls.Save(qlok, "~/qlok/qlockfile")
}
