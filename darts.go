package darts

import (
	"sync"
)

/*Darts t*/
type Darts struct {
	lock   sync.RWMutex
	base   []uint64
	check  []uint64
	llTrie *linkedListTrie
}

/*BuildFromStrs is */
func (d *Darts) BuildFromStrs(words [][]rune) error {
	d.lock.Lock()
	defer d.lock.Unlock()

	d.llTrie = &linkedListTrie{}

	return nil
}

/*ExactMatchSearch xx*/
func (d *Darts) ExactMatchSearch(key string) bool {
	d.lock.RLock()
	defer d.lock.RUnlock()

	return false
}

/*CommonPrefixSearch xx*/
func (d *Darts) CommonPrefixSearch(size int64, key string) []string {
	d.lock.RLock()
	defer d.lock.RUnlock()

	return []string{}
}
