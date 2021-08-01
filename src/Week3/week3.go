package Week3

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hi, the time is "+time.Now().Format("2006-01-02 15:04:05")+"\n")
	fmt.Printf("Responsed at %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

func runApp(server *http.Server) error {
	http.HandleFunc("/", handler)
	fmt.Println("App start")
	err := server.ListenAndServe()
	return err
}

func Do() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	g, errCtx := errgroup.WithContext(ctx)
	server := &http.Server{Addr: ":8080"}
	g.Go(func() error {
		return runApp(server)
	})
	g.Go(func() error {
		<-errCtx.Done()
		fmt.Println("App stop")
		return server.Shutdown(errCtx)
	})

	ch := make(chan os.Signal, 1)
	signal.Notify(ch)

	g.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-ch:
				cancel()
			}
		}
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Println("g error:", err)
	}
	fmt.Println("Done!")
}
