package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"shift-image/shifter"
	"time"
)

const (
	COUNT      = 100 // COUNT is the total number of frames to create
	BATCH_SIZE = 20  // BATCH_SIZE is the number of frames inside each chunk, i.e., inside each goroutine.
)

// Its important to note that the code below is conderating the calculation of the frame creation in parallel based
// on the result of the division of the COUNT by the BATCH_SIZE. So if the COUNT is 100 and the BATCH_SIZE is 10, the
// code will spawn 10 goroutines to create 10 frames inside each goroutine.

func main() {
	// Setup environment
	os.RemoveAll("/output/")
	os.MkdirAll("/output/", 755)
	start := time.Now()

	// Load the source image
	src, err := shifter.GetImageFromFile("image.jpg")
	if err != nil {
		log.Fatalf("failed to load the image: %v", err)
	}

	// Create the actions
	actions := makeActions(COUNT)
	chunks := chunkActions(actions, 2)
	controlChan := make(chan int, len(actions))

	// Create the frames
	for _, chunk := range chunks {
		go handleChunk(controlChan, src, chunk)
	}

	// Wait for all frames to be created
	for i := 0; i < COUNT; i++ {
		processedAction := <-controlChan
		log.Println("Processed frame", processedAction)
	}

	log.Println(fmt.Sprintf("Finished creating %d frames", COUNT))
	log.Println("Time taken:", time.Since(start))
}

func handleChunk(ch chan int, src image.Image, chunk []int) {
	for _, value := range chunk {
		err := shifter.CreateFrame(src, value*10)
		if err != nil {
			log.Fatalf("failed to create frame: %v", err)
		}
		ch <- value
	}
}

func makeActions(size int) []int {
	actions := make([]int, size)
	for i := 0; i < size; i++ {
		actions[i] = i + 1
	}
	return actions
}

func chunkActions(actions []int, batchSize int) [][]int {
	batches := make([][]int, 0, (len(actions)+batchSize-1)/batchSize)

	for batchSize < len(actions) {
		actions, batches = actions[batchSize:], append(batches, actions[0:batchSize:batchSize])
	}
	batches = append(batches, actions)
	return batches
}
