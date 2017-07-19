package parser

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"github.com/parserlog/model"
)

func Process() (result map[string]*model.BodyCount, err error) {

	// get running dir
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))

	if err != nil {
		return result, err
	}

	// get log files
	files, err := ioutil.ReadDir(dir + "/logs")

	if err != nil {
		return result, err
	}

	// make result
	result = make(map[string]*model.BodyCount, len(files))

	// set the number of channels
	channels := len(files)

	// set the number of simultaneous processes
	simultaneous := runtime.NumCPU()

	// is there less processes to run than CPUs?
	if channels < simultaneous {
		simultaneous = channels
	}

	// init channels
	queue := make(chan *model.BodyCount, channels)

	// init wait group
	wg := &sync.WaitGroup{}

	for _, file := range files {

		// init body count
		result[file.Name()] = &model.BodyCount{
			FilePath: fmt.Sprintf("%s/logs/%s", dir, file.Name()),
		}

		// generate task
		queue <- result[file.Name()]

	}

	// start workers
	for i := 0; i < simultaneous; i++ {

		// add worker
		wg.Add(1)

		// go concurrency
		go func(wg *sync.WaitGroup) {

			// make concurrency safe (finish worker)
			defer wg.Done()

			// parse file
			for q := range queue {
				if err := Parser(q); err != nil {
					q.Error = err
				}
			}

		}(wg)

	}

	// close queue of tasks
	close(queue)

	// wait for all the workers to finish
	wg.Wait()

	return

}

