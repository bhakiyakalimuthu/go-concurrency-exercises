package main

import (
	"bufio"
	"os"
)

func main() {

}

func reader(fileName string) {
	// if file path provided
	// open file
	file, err := os.Open(fileName)
	if err != nil {
		log.Errorf("failed to open file : %d", fileName)
	}
	// read file
	reader := bufio.NewReader(file)
	out := make([]byte, 1024*4)
	for {
		_, err := reader.Read(out)
		if err != nil {
			if err == os.EOF {
				log.Errorf("end of file")
				break
			}
			log.Errof("failed to read from file")
			break
		}

	}
	// close file
	// return the content
}

func writer() {

}
