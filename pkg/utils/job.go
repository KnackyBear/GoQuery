package utils

type Job struct {
	result interface{}
}

func (self *Job) execute() {
	// do some network activity
}
  
  // to be executed parallelly
doParallel := func(ctx context.Context, inputs <-chan Job, output chan<- Job) {
	for {
		select {
		case tag, ok := <-inputs:
		if !ok {
			return
		}
		job.result := job.execute()
		output <- job
		case <-ctx.Done():
		return
		}
	}
}
  
  // an array of jobs
  jobs := make([]Job)
  jobs.append(jobs,...)
  
  ctx := context.Background()
  max := runtime.NumCPU()
  queue := make(chan Job, max)
  output := make(chan Job)
  defer close(output)
  
  // spin up workers
  for i := 0; i < max; i++ {
	go doParallel(ctx, queue, output)
  }
  
  // passing jobs to workers though a queue, idle workers will pick the job and execute
  go func() {
	for _, job := range jobs {
	  queue <- job
	}
	close(queue)
  }()
  
  // collecting results
  results := make([]Job, len(jobs))
  for i := 0; i < len(jobs); i++ {
	select {
	case executedJob := <-output:
	  results[i] = executedJob
	case <-ctx.Done():
	  break
	}
  }