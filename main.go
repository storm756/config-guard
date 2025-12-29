package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {

	dirPtr := flag.String("dir", "./test_configs","Directory containing JSON configs")
	flag.Parse()

	files,err := os.ReadDir(*dirPtr)
	if err != nil {
		fmt.Println("Error in reading directory", err)
		return
	}
	var wg sync.WaitGroup
	resultsChannel := make(chan ValidationResult , len(files))

	rules := []Rule { PortRule{} , ReplicaRule{MinReplicas: 2}}
	fmt.Printf("Scanning Directory: %s\n", *dirPtr)

	for _, file := range files {
		if filepath.Ext(file.Name()) != ".json" {
		continue 
	}
	wg.Add(1)

	go func(fileName string) {
		defer wg.Done()

		path := filepath.Join(*dirPtr , fileName)
		data , err := os.ReadFile(path)
		if err != nil {
			resultsChannel <- ValidationResult{false , fmt.Sprintf("File Read error %s", fileName)}
			return
		}
		var config Config
		if err := json.Unmarshal(data, &config); err != nil {
			resultsChannel <- ValidationResult{false , fmt.Sprintf("JSON Error: %s %s",fileName , err)}
			 return
		}
		for _,rule := range rules {
			valid , msg := rule.Validate(config)
			resultsChannel <- ValidationResult{
				Success: valid,
				Message: fmt.Sprintf("[%s] %s", fileName , msg),
			}
		}


	} (file.Name())

	}

		go func() {
			wg.Wait()
			close(resultsChannel)
		} ()
		for res := range resultsChannel {
			if res.Success {
				fmt .Printf("%s\n" , res.Message)
			}else {
				fmt.Printf("%s\n",res.Message)
			}
		}
}