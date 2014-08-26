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
		avg int64
	}

	val := Vals{count: 0, avg: 0}
	data := make(map [string] Vals)
	var eml string
	var du int64
	//var date string

	scanner := bufio.NewScanner(inpfl)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ",")

		eml = line[0]
		du, _ = strconv.ParseInt(line[1], 10, 32)
		//date = line[2]

		if data[eml].count == 0 {
			val.count = 1
			val.avg = du 
		} else {
			val.count = data[eml].count + 1
			val.avg = (data[eml].avg + du) / int64(2)
		}
		data[eml] = val
		//fmt.Println(eml)
		//fmt.Println(data[eml].count)
		//fmt.Println(data[eml].avg)
	}
	err = scanner.Err()
	check_err(err)

	fmt.Printf("%20s %5s %10s\n", "Email Address", "Days", "Average")
	fmt.Printf("%20s %5s %10s\n", "---", "---", "---")
	for eml, val := range data {
		fmt.Printf("%s: %d, %f\n", eml, val.count, val.avg/1024)
	}

	return
}
