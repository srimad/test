package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func check_err(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: " + os.Args[0] + " <file_path>")
		return
	}

	inpfl, err := os.Open(os.Args[1])
	check_err(err)
	defer inpfl.Close()

	type Vals struct {
		count int
		avg float64
	}

	val := Vals{count: 0, avg: 0}
	data := make(map [string] Vals)
	var eml string
	var du float64
	//var date string

	scanner := bufio.NewScanner(inpfl)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		eml = line[0]
		du, _ = strconv.ParseFloat(line[1], 10)
		//date = line[2]

		if data[eml].count == 0 {
			val.count = 1
			val.avg = du 
		} else {
			val.count = data[eml].count + 1
			val.avg = (data[eml].avg + du) / 2.00
		}
		data[eml] = val
		//fmt.Println(eml)
		//fmt.Println(data[eml].count)
		//fmt.Println(data[eml].avg)
	}
	err = scanner.Err()
	check_err(err)

	fmt.Printf("%-25s %5s %10s %10s\n", "Email Address", "Days", "Avg_gB", "Price_USD")
	for eml, val := range data {
		gb := val.avg / (1024 * 1024 * 1024)
		if gb > 512 {
			fmt.Printf("%-25.25s %5d %10.2f %10.2f\n", eml, val.count, gb, (gb - 512) * 22.5 / 1024)
		}
	}

	return
}
