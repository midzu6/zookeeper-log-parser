package main

import (
	"fmt"
	"logParse/internal/models"
	"os"
	"time"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("Использование: go run cmd/parser/main.go <testdata/zookiper.log>")
		os.Exit(1)
	}

	//filePath := os.Args[1]

	//count, err := parser.ReadAndCountEntries(filePath)
	//if err != nil {
	//fmt.Fprintln(os.Stderr, "error:", err)
	//}
	//fmt.Printf("\nГотово! Найдено %d записей в логе.\n", count)

	entry := models.NewLogEntry(
		time.Date(2015, 7, 29, 17, 41, 44, 747000000, time.UTC),
		models.LogLevelInfo,
		"QuorumPeer[myid=1]/0:0:0:0:0:0:0:0:2181:FastLeaderElection@774",
		"FastLeaderElection@774",
		"Notification time out: 3200",
	)

	fmt.Println(entry)
}
