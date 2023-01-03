package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's the weekday!")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	case t.Hour() > 12:
		fmt.Println("It is after noon")
	}

	whatAmI := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Println("Dont know this type")
		}
	}

	whatAmI(true)
	whatAmI("abracadabra")
	whatAmI(234)
}
