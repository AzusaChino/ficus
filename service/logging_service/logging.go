package logging_service

import (
	"bufio"
	"fmt"
	"github.com/AzusaChino/ficus/pkg/kafka"
	"io/fs"
	"log"
	"os"
	"syscall"
)

const FicusMessageTopic = "FICUS_MESSAGE"

func AsyncSend(f interface{}) {

	targetSrc, ok := f.(string)
	if !ok {
		panic("wrong parameter type for AsyncSend")
	}
	// make sure file close when function ends
	file, err := os.OpenFile(targetSrc, syscall.O_RDONLY, fs.ModeExclusive)
	if err != nil {
		log.Fatalf(`failed to open file %s: %v`, targetSrc, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64k
	const maxCap = 1024 * 64
	// new buffer (byte array)
	buf := make([]byte, maxCap)
	// set max capacity
	scanner.Buffer(buf, maxCap)
	for scanner.Scan() {
		s := scanner.Text()
		kafka.SendMessage(FicusMessageTopic, "", s)
		fmt.Printf("read message: %s\n", s)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error when reading file, %v", err)
	}
}
