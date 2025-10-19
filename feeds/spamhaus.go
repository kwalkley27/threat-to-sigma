package feeds

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"github.com/kwalkley27/threat-to-sigma/config"
)

func Retrieve(cfg *config.Config) ([]string, error) {
	cidrList := []string{}

	// Fetch the DROP list
	resp, err := http.Get(cfg.SpamhausFeedURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch DROP list: %w", err)
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch DROP list: %s", resp.Status)
	}

	// Create a scanner to read the response body line by line
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()
		// Skip comments, empty lines, and info lines
		if len(line) == 0 || line[0] == '#' || line[0] == ';' {
			continue
		}

		//strip extra line details and add cidrs to list
		cidrList = append(cidrList, strings.Split(line, " ; ")[0])

		//stop processing cidrs when the limit is reached
		if len(cidrList) >= cfg.FeedLimit {
			break
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading DROP list: %w", err)
	}

	return cidrList, nil
}
