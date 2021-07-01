package server

type Semaphore struct {
	sm chan struct{}
}

func NewSemaphore(size int) *Semaphore {
	return &Semaphore{make(chan struct{}, size)}
}

func (s *Semaphore) Acquire(n int) {
	for i := 0; i < n; i++ {
		s.sm <- struct{}{}
	}
}

func (s *Semaphore) Release(n int) {
	for i := 0; i < n; i++ {
		<- s.sm
	}
}
