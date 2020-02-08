package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"

	"api/internal/app/apiserver"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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

	db, err := sql.Open("mysql", "root:12345678@/mydb")

	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Begin()

	if err != nil {
		log.Fatal(err)
	}

	log.Println("db work")
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

	db.Close()

}
