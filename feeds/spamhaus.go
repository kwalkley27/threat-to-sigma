package feeds

import (
	"bufio"
	//"fmt"
	"log"
	"net/http"
	"strings"
)

func Retrieve() []string {
	
	ipList := []string{}
	
	// Fetch the DROP list
	resp, err := http.Get("https://www.spamhaus.org/drop/drop.txt")
	if err != nil {
		log.Fatalf("Failed to fetch DROP list: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch DROP list: %s", resp.Status)
	}

	// Create a scanner to read the response body line by line
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip comments, empty lines, and info lines
		if len(line) == 0 || line[0] == '#' || line[0] == ';' {
			continue
		}
		
		//fmt.Println(strings.Split(line, " ; ")[0])
		ipList = append(ipList, strings.Split(line, " ; ")[0])
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading DROP list: %v", err)
	}

	return ipList
}
