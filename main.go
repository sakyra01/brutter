package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	target := shell()
	file := "wordlist.txt"
	var av []string
	enumeration(target, file, av)
}

func shell() (accepting string) {
	arg := os.Args[1:]
	if len(arg) < 2 {
		fmt.Println("Bad usage format, try using flag -h for information how to use brutter")
		os.Exit(127)
	}
	if (arg[0] == "--help") || (arg[0] == "-h") {
		fmt.Print("Options:\n" +
			"-h, --help - show this help message and exit\n" +
			"-u, --url -  Target URL, you should input full format http(s)://target.com...\n\n " +
			"Usage Example:\n" +
			"./brutter -u <target url>\n\n")
		os.Exit(127)
	}
	if arg[0] != "-u" {
		fmt.Println("URL target is missing, try using flag -h for information how to use brutter\n")
		os.Exit(127)
	}
	if arg[0] == "-u" {
		if (strings.Contains(arg[1], "http://")) || (strings.Contains(arg[1], "https://")) {
			accepting = arg[1]
		} else {
			fmt.Println("Bad url format, try using -h for information how to use brutter\n")
			os.Exit(127)
		}
	}
	return accepting
}

func brute(target, line string) (avilation []string) {
	target = target + "/" + line
	resp, err := http.Get(target)
	if err != nil {
		log.Fatalln(err)
	}

	status := resp.StatusCode
	var list []string

	dt := time.Now()
	data := dt.Format("15:04:05") // Format hh:mm:ss

	format := fmt.Sprintf("[%s] %d - %s", data, status, target)

	if status == 200 {
		avilation := append(list, format)
		return avilation
	}
	fmt.Println(format)
	return
}

func enumeration(target, file string, av []string) {
	f, err := os.Open(file)

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		av = brute(target, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Accessible paths\n:")
	for _, val := range av {
		fmt.Println(val)
	}
}
