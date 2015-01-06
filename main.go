package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("Usage: reload filename command")
	}

	filename := os.Args[1]
	// TODO: Handle escaped quotes (nested args)
	command_array := strings.Split(os.Args[2], " ")
	command := command_array[0]
	args := command_array[1:len(command_array)]

	stats, err := os.Stat(filename)
	if err != nil {
		log.Fatal("No such file: %s (%v)", filename, err)
	}

	timestamp := stats.ModTime()

	for {

		stats, err = os.Stat(filename)
		if err != nil {
			log.Fatal("No such file: %s (%v)", filename, err)
		}

		if timestamp != stats.ModTime() {
			log.Println("File changed. Restarting.")

			out, err := exec.Command(command, args...).Output()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%s", out)

			timestamp = stats.ModTime()
		}
		time.Sleep(1 * time.Second)
	}

}
