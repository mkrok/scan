package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int) {
	server := "scanme.nmap.org"
	// server := "localhost"
	for p := range ports {
		address := fmt.Sprintf(server+":%d", p)
		conn, err := net.Dial("tcp", address)
		if err == nil {
			results <- p
			conn.Close()
		} else {
			results <- 0
		}
	}
}

func main() {
	number := 10000
	ports := make(chan int, 100)
	results := make(chan int)
	var openports []int

	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i <= number; i++ {
			ports <- i
		}
	}()

	for i := 1; i <= number; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}

	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}

}
