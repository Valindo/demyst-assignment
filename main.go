package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/valindo/demyst-assignment/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
		log.Fatal(err)
	}
	baseUrl := os.Getenv("URL")
	iterations, err := strconv.Atoi(os.Getenv("ITERATIONS"))
	if err != nil {
		log.Fatal("Failed to retrive ITERATIONS env variable")
		log.Fatal(err)
	}
	fmt.Println("Application loading todos....")
	utils.Display(baseUrl, iterations)

}
