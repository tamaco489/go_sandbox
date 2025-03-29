package main

import (
	// lib_mutex "github.com/tamaco489/go_sandbox/sync/lib/mutex"
	lib_rw_mutex "github.com/tamaco489/go_sandbox/sync/lib/rw_mutex"
)

func main() {
	// sync.Mutex
	// lib_mutex.MutexProcess()

	// sync.RWMutex
	lib_rw_mutex.RWMutexProcess()
}
