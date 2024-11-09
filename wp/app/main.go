package main

import (
	"fmt"
	"streamer"
)

func main() {
	// Define number of worker and jobs
	const NUM_JOBS = 1
	const NUM_WORKERS = 2

	// Create channel for work and results
	notifyChan := make(chan streamer.ProcessingMessage, NUM_JOBS)
	defer close(notifyChan)

	videoQueue := make(chan streamer.VideoProcessingJob, NUM_JOBS)
	defer close(videoQueue)

	// Get a worker pool.
	wp := streamer.New(videoQueue, NUM_WORKERS)

	// Start the worker pool
	wp.Run()

	// Create a videos to send to the worker pool
	video := wp.NewVideo(1, "./input/puppy1.mp4", "./output", "mp4", notifyChan, nil)

	// Send the video to the worker pool
	videoQueue <- streamer.VideoProcessingJob{Video: video}

	// Print out results
	for i := 1; i <= NUM_JOBS; i++ {
		msg := <-notifyChan
		fmt.Println("i:", i, "msg:", msg)
	}

	fmt.Println("Done")

}
