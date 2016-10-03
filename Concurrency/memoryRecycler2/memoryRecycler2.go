//Using a buffered channel to get buffers from a pool.
package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func makeBuffer() []byte {
	return make([]byte, rand.Intn(5000000)+5000000)
}

func main() {
	// our big pool
	pool := make([][]byte, 20)
	// pool of buffers
	buffer := make(chan []byte, 5)

	var m runtime.MemStats
	makes := 0
	for {
		var b []byte
		select {
		// retrieve a buffer if pool has previously released buffers
		case b = <-buffer:
		default:
			// or create a new buffer
			makes++
			b = makeBuffer()
		}
		i := rand.Intn(len(pool))
		if pool[i] != nil {
			select {
			// release buffers to the buffer pool
			case buffer <- pool[i]:
				pool[i] = nil
			default:
			}
		}

		// store created buffer in the big pool
		pool[i] = b

		time.Sleep(time.Second)
		bytes := 0
		for i := 0; i < len(pool); i++ {
			if pool[i] != nil {
				bytes += len(pool[i])
			}
		}

		runtime.ReadMemStats(&m)
		fmt.Printf("%d,%d,%d,%d,%d,%d\n", m.HeapSys, bytes, m.HeapAlloc,
			m.HeapIdle, m.HeapReleased, makes)
	}
}
