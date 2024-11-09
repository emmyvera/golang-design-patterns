package streamer

import "fmt"

type VideoDispatcher struct {
	WorkerPool chan chan VideoProcessingJob
	maxWorkers int
	jobQueue   chan VideoProcessingJob
	Processor  Processor
}

// type videoWorker
type videoWorker struct {
	id         int
	jobQueue   chan VideoProcessingJob
	workerPool chan chan VideoProcessingJob
}

// newVideoWorker
func newVideoWorker(id int, workerPool chan chan VideoProcessingJob) videoWorker {
	fmt.Println("newVideoWorker(): Creating video worker id", id)
	return videoWorker{
		id:         id,
		jobQueue:   make(chan VideoProcessingJob),
		workerPool: workerPool,
	}
}

// start()
func (w videoWorker) start() {
	fmt.Println("w.start(): Starting video worker id", w.id)
	go func() {
		for {
			// Add job queun to the worker pool
			w.workerPool <- w.jobQueue

			// wait for a job to come back
			job := <-w.jobQueue

			// process the job
			w.processVideoJob(job.Video)
		}
	}()
}

// Run()
func (vd *VideoDispatcher) Run() {
	fmt.Println("vd.Run(): Starting worker pool by running worker")
	for i := 0; i < vd.maxWorkers; i++ {
		fmt.Println("vd.Run(): Srarting worker id", i+1)
		worker := newVideoWorker(i+1, vd.WorkerPool)
		worker.start()
	}

	go vd.dispatch()
}

// dispatch()
func (vd *VideoDispatcher) dispatch() {
	for {
		// wait for a job to comes in
		job := <-vd.jobQueue
		fmt.Println("vd.dispatch(): sending job", job.Video.ID, "to worker job queue")
		go func() {
			workerJobQueue := <-vd.WorkerPool
			workerJobQueue <- job
		}()
	}
}

// processVideoJob
func (w videoWorker) processVideoJob(video Video) {
	fmt.Println("w.processVideoJob(): starting encode on video", video.ID)
	video.encode()
}
