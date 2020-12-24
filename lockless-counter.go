package lockless_counter

type LocklessCounter interface {
	Take() uint64
}

type locklessCounter struct {
	takeCH (chan bool)
	giveCH (chan uint64)
}

func NewLocklessCounter() LocklessCounter {
	llc := new(locklessCounter)
	llc.takeCH = make(chan bool, 1)
	llc.giveCH = make(chan uint64, 1)
	counter := uint64(0)
	go func() {
		for {
			select {
			case <-llc.takeCH:
				llc.giveCH <- counter
				counter++
			}
		}
	}()
	return llc
}

func (llc *locklessCounter) Take() (res uint64) {
	llc.takeCH <- true
	res = <-llc.giveCH
	return
}
