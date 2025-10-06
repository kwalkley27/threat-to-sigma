package main

import (
	"fmt"
	// "os"
	// "log"
	//"context"

	// "github.com/google/generative-ai-go/genai"
	// "google.golang.org/api/option"
	//"github.com/kwalkley27/threat-to-sigma/genai"
	"github.com/kwalkley27/threat-to-sigma/feeds"
)

const ModelName = "gemini-2.5-flash"

func main() {
	//ctx := context.Background()

	ipList := feeds.Retrieve()

	for _,ip := range ipList {
		fmt.Println(ip)
	}

	//genai.FormatSigma(ctx, []string{"1.2.3.4"})
}