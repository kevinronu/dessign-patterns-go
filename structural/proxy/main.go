package main

import (
	"fmt"
	"time"

	"github.com/kevinronu/dessign-patterns-go/structural/proxy/proxy"
	"github.com/kevinronu/dessign-patterns-go/structural/proxy/real"
	"github.com/kevinronu/dessign-patterns-go/structural/proxy/service"
)

// timeDownload prints what a call cost, because the cost shows who answered: 0ms is the proxy,
// anything slower is the service.
//
// It takes the contract, so it runs the same against the real service and against the proxy.
func timeDownload(library service.VideoLibrary, id string) {
	start := time.Now()

	video, err := library.Download(id)
	if err != nil {
		fmt.Printf("  %-5s  failed: %v\n", id, err)

		return
	}

	fmt.Printf("  %-5s %4dms  %s\n", id, time.Since(start).Milliseconds(), video)
}

func main() {
	const ttl = time.Second

	cache := proxy.NewCache(real.Remote{}, ttl)
	defer cache.Stop()

	fmt.Printf("a video library behind a caching proxy, ttl %v\n\n", ttl)

	timeDownload(cache, "cats")
	timeDownload(cache, "dogs")
	timeDownload(cache, "cats")
	timeDownload(cache, "nope")

	time.Sleep(ttl)

	fmt.Println("\nthe cleanup ran, so nothing is stored any more")
	timeDownload(cache, "cats")
}
