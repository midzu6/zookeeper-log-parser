package parser

import (
	"errors"
	"logParse/internal/utils"
)
import "bufio"
import "os"
import "fmt"

var ErrCannotOpenFile = errors.New("cannot open file")

func ReadAndCountEntries(path string) (int, error) {

	file, err := os.Open(path)
	if err != nil {
		return 0, fmt.Errorf("не удалось открыть файл %s: %w", path, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	entryCount := 0

	for scanner.Scan() {
		line := scanner.Text()

		if utils.IsNewEntry(line) {
			entryCount++
			fmt.Printf("[НОВАЯ ЗАПИСЬ #%d] %s\n\n", entryCount, line[:70])
		}
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("ошибка при чтении файла: %w", err)
	}

	return entryCount, nil
}
