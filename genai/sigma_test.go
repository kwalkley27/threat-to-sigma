package genai

import (
	"strings"
	"testing"
)

func TestGetPrompt(t *testing.T) {
	ioc := "1.2.3.4/24"
	prompt := getPrompt(ioc)

	if !strings.Contains(prompt, ioc) {
		t.Errorf("prompt does not contain the IOC: got %s", prompt)
	}

	if !strings.Contains(prompt, "sigma rule") {
		t.Errorf("prompt does not ask for a sigma rule: got %s", prompt)
	}
}
