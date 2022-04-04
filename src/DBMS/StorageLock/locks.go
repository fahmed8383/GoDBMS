package StorageLock

import (
	"sync"
)

var tableLocks map[string]*sync.Mutex
var catalogLock sync.Mutex

func InitializeLocks() {
	tableLocks = make(map[string]*sync.Mutex)
}

func AcquireTableLock(table string) {
	_, exists := tableLocks[table]
	if(!exists) {
		tableLocks[table] = &sync.Mutex{}
	}
	tableLocks[table].Lock()
}

func ReleaseTableLock(table string) {
	tableLocks[table].Unlock()
}

func AcquireCatalogLock() {
	catalogLock.Lock()
}

func ReleaseCatalogLock() {
	catalogLock.Unlock()
}