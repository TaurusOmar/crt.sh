package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type CrtshResult struct {
	NameValue string `json:"name_value"`
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	fmt.Println(`
                    __                    
    _____   _____  / /_      ____ _  ____ 
   / ___/  / ___/ / __/     / __  / / __ \
  / /__   / /    / /_   _  / /_/ / / /_/ /
  \___/  /_/     \__/  (_) \__, /  \\___/ 
                          /____/ @TaurusOmar_
	`)

	if len(os.Args) != 2 {
		fmt.Printf("Use: %s <domain>\n", os.Args[0])
		os.Exit(1)
	}

	domain := os.Args[1]
	resultDir := "result_directory"
	crtURL := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", domain)

	fmt.Printf("Scanning for domain: %s...\n", domain)

	go spinner(100 * time.Millisecond)

	resp, err := http.Get(crtURL)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var crtshResults []CrtshResult
	err = json.Unmarshal(body, &crtshResults)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	subdomains := make(map[string]struct{})
	for _, result := range crtshResults {
		subdomain := strings.TrimPrefix(result.NameValue, "*.")
		subdomains[subdomain] = struct{}{}
	}

	uniqueSubdomains := make([]string, 0, len(subdomains))
	for subdomain := range subdomains {
		uniqueSubdomains = append(uniqueSubdomains, subdomain)
	}

	sort.Strings(uniqueSubdomains)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	resultDir = filepath.Join(homeDir, resultDir)
	if _, err := os.Stat(resultDir); os.IsNotExist(err) {
		err = os.Mkdir(resultDir, 0755)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
	}

	crtOutput := filepath.Join(resultDir, fmt.Sprintf("%s.crt.txt", domain))
	err = ioutil.WriteFile(crtOutput, []byte(strings.Join(uniqueSubdomains, "\n")), 0644)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	fmt.Printf("\rScan completed. Results saved in %s\n", crtOutput)
}
