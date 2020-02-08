package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "conf-path", "configs/apiserver.json", "path to config json file")
}

func main() {
	flag.Parse()

	file, err := ioutil.ReadFile(configPath)

	if err != nil {
		log.Fatal(err)
	}

	config := apiserver.Config{}

	bytes := []byte(file)

	json.Unmarshal(bytes, &config)

	log.Println(config.LogLevel)

	s := apiserver.New(&config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
