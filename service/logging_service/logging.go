package logging_service

import (
	"bufio"
	"github.com/AzusaChino/ficus/pkg/kafka"
	"log"
	"os"
)

const FicusMessageTopic = "FICUS_MESSAGE"

func AsyncSend(f interface{}) {

	file, ok := f.(*os.File)
	if !ok {
		panic("wrong parameter type for AsyncSend")
	}
	// make sure file close when function ends
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64k
	const maxCap = 1024 * 64
	// new buffer (byte array)
	buf := make([]byte, maxCap)
	// set max capacity
	scanner.Buffer(buf, maxCap)
	for scanner.Scan() {
		kafka.SendMessage(FicusMessageTopic, "", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error when reading file, %v", err)
	}
}
