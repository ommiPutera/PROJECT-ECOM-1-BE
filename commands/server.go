package commands

import (
	"context"
	"ecomjc-be/config"
	server "ecomjc-be/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func InitServer() {
	/* Notify an Interruption Signal
	signal.Notify() will be listening the os.Interrupt (an interrupt such as
	typing ctrl + C in terminal) to gracefully shutdown the app by sending
	the signal to osSignal */
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, os.Interrupt)

	/* Handling an Interruption Signal
	signalCtx have a Done() method to stop the program. signalCtx.Done() is triggered
	by calling the sigCancel(). */
	signalCtx, sigCancel := context.WithCancel(context.Background())
	defer sigCancel()

	// go func() is running concurrently (non-blocking to its below codes)
	go func() {
		/* osCall is fulfilled when osSignals got interrupted. If it's not filled yet,
		it's blocking the codes below  */
		osCall := <-osSignals
		log.Printf("system call:%+v", osCall)
		sigCancel()
	}()

	server := server.ServeHTTP(config.GetConfig())

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Println("Starting http server...")

	<-signalCtx.Done()

	log.Println("Stopping http server...")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		shutdownCancel()
	}()

	// hover to the Shutdown() method for more information
	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	log.Printf("server exited properly")
}
