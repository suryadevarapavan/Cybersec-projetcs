package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// Take website input from the user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter website URL (e.g., scanme.nmap.org): ")
	website, _ := reader.ReadString('\n')
	website = strings.TrimSpace(website)

	// Scan ports from 1 to 1024
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("%s:%d", website, i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("Port %d open\n", i)
	}
}
