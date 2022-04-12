package StorageLock

import (
	"sync"
)

// Intialize global variables to store the locks for the database in memory
var tableLocks map[string]*sync.Mutex
var catalogLock sync.Mutex

// IntiializeLocks is a function to initialize the global lock variables so
// they can then be used.
func InitializeLocks() {
	tableLocks = make(map[string]*sync.Mutex)
}

// AcquitreTableLock takes in a string to a table name and obtains the lock
// for the corresponding table associated with that name. If a lock by the 
// corresponding name does not exist, then create it.
func AcquireTableLock(table string) {
	_, exists := tableLocks[table]
	if(!exists) {
		tableLocks[table] = &sync.Mutex{}
	}
	tableLocks[table].Lock()
}

// ReleaseTableLock takes in a string to a table name and releases the lock
// for the corresponding table associated with that name.
func ReleaseTableLock(table string) {
	tableLocks[table].Unlock()
}

// AcquireCatalogLock obtains the global catalog lock.
func AcquireCatalogLock() {
	catalogLock.Lock()
}

// ReleaseCatalogLock releases the global catalog lock.
func ReleaseCatalogLock() {
	catalogLock.Unlock()
}