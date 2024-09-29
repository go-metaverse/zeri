package graceful

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/go-metaverse/zeri/logger"
	"go.uber.org/zap"
)

// WaitGroup for managing shutdown processes.
var (
	stop, shutdown sync.WaitGroup
	log            *zap.SugaredLogger
)

func init() {
	log = logger.NewLoggerWithAttributes(logger.Attributes{
		"prefix": "graceful",
	})

	// Register signal notifications for graceful shutdown.
	stop.Add(1)
	go func() {
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP, syscall.SIGQUIT)
		log.Debug("Received signal: ", <-signalChan)
		stop.Done()
	}()
}

// AddProcess increments the stop WaitGroup counter for a new process.
func AddProcess() {
	stop.Add(1)
}

// DoneProcess decrements the stop WaitGroup counter for a completed process.
func DoneProcess() {
	stop.Done()
}

// Stop initiates a shutdown sequence after waiting for all processes to complete.
// It executes the provided cleanup function after a specified duration.
//
// Parameters:
//   - delay: Duration to wait before executing the cleanup function.
//   - cleanupFunc: The function to execute after the delay.
func Stop(delay time.Duration, cleanupFunc func()) {
	shutdown.Add(1)
	go func() {
		stop.Wait()
		log.Debug("Stopping")
		time.Sleep(delay)
		cleanupFunc()
		shutdown.Done()
	}()
}

// ShutDown waits for all processes to complete and then terminates the application.
func ShutDown() {
	stop.Wait()
	shutdown.Wait()
	log.Debug("Shutdown initiated")
	os.Exit(0)
}

// ShutDownSlowly waits for all processes, introduces a delay, and then terminates the application.
//
// Parameters:
//   - delay: Duration to wait before terminating the application.
func ShutDownSlowly(delay time.Duration) {
	stop.Wait()
	shutdown.Wait()
	log.Debug("ShutDownSlowly initiated")
	time.Sleep(delay)
	os.Exit(0)
}

// WaitShutdown triggers the Stop process and waits for it to complete,
// then performs a shutdown.
func WaitShutdown() {
	stopChan := make(chan bool, 1)
	Stop(0, func() {
		stopChan <- true // Signal completion to the channel.
	})
	<-stopChan // Wait for the completion signal.
	ShutDown()
}
