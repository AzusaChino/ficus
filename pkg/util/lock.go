package util

import (
	"fmt"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

func getGID() int64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = b[:strings.Index(string(b), " ")]
	b = b[10:]
	var gid int64
	fmt.Sscanf(string(b), "%d", &gid)
	return gid
}

type RecursiveMutex struct {
	sync.Mutex
	owner     int64
	recursion uint32
}

func (rm *RecursiveMutex) Lock() {
	gid := getGID()
	if atomic.LoadInt64(&rm.owner) == gid {
		rm.recursion++
		return
	}
	rm.Mutex.Lock()
	atomic.StoreInt64(&rm.owner, gid)
	rm.recursion = 1
}

func (rm *RecursiveMutex) UnLock() {
	gid := getGID()
	if atomic.LoadInt64(&rm.owner) != gid {
		panic(fmt.Sprintf("wrong owner(%d): %d!", rm.owner, gid))
	}
	rm.recursion--
	if rm.recursion != 0 {
		return
	}
	atomic.StoreInt64(&rm.owner, -1)
	rm.Mutex.Unlock()
}
