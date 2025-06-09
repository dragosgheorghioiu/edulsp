package main

import (
	"bufio"
	"log"
	"os"

	"github.com/dragosgheorghioiu/edulsp/src/rpc"
)

func main() {
	logger := getLogger("../log.txt")
	logger.Println("Started")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(rpc.Split)

	for scanner.Scan() {
		msg := scanner.Text()
		handleMessage(logger, msg)
	}
}

func handleMessage(logger *log.Logger, msg any) {
	logger.Println(msg)
}

func getLogger(filename string) *log.Logger {
	logfile, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		panic("please provide a file")
	}

	return log.New(logfile, "[edulsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}

