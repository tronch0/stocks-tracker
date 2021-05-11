package analytics

import (
	"sync"
)

type Analytics struct {
	avgProcessingTime int
	requestsCounter   int // 2^31 -1
	mutex sync.RWMutex
}

func New() *Analytics {
	return &Analytics{
		avgProcessingTime: 0,
		requestsCounter:   0,
		mutex:             sync.RWMutex{},
	}
}

func (a *Analytics) GetStats() (requestsCount, avgProcessTime int) {

	a.mutex.RLock()
	requestsCount = a.requestsCounter
	avgProcessTime = a.avgProcessingTime
	a.mutex.RUnlock()

	return
}

func (a *Analytics) AddRequest(processTime int) {
	a.mutex.Lock()

	currentSumAvg := a.requestsCounter * a.avgProcessingTime // lock

	currentSumAvg = currentSumAvg + processTime

	a.requestsCounter++
	newAvg := currentSumAvg / a.requestsCounter
	a.avgProcessingTime = newAvg

	a.mutex.Unlock()
}
