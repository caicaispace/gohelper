package goqueue

var Queue chan interface{}

type GoQueue struct {
	queue chan interface{}
}

func NewGoQueue(len int) *GoQueue {
	queue := make(chan interface{}, len)

	goqueue := &GoQueue{
		queue: queue,
	}

	return goqueue
}

func (g *GoQueue) Push(data interface{}) {
	g.queue <- data
}

func (g *GoQueue) Pop() interface{} {
	return <-g.queue
}

func (g *GoQueue) Close() {
	close(g.queue)
}

func (g *GoQueue) Len() int {
	return len(g.queue)
}

func (g *GoQueue) Cap() int {
	return cap(g.queue)
}

func (g *GoQueue) IsEmpty() bool {
	return len(g.queue) == 0
}

func (g *GoQueue) IsFull() bool {
	return len(g.queue) == cap(g.queue)
}

func (g *GoQueue) Range(f func(interface{})) {
	for data := range g.queue {
		f(data)
	}
}

func (g *GoQueue) RangeWithBreak(f func(interface{}) bool) {
	for data := range g.queue {
		if f(data) {
			break
		}
	}
}

func (g *GoQueue) RangeWithContinue(f func(interface{}) bool) {
	for data := range g.queue {
		if f(data) {
			continue
		}
	}
}

func (g *GoQueue) RangeWithReturn(f func(interface{}) bool) {
	for data := range g.queue {
		if f(data) {
			return
		}
	}
}
