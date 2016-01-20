package structures

import (
	//"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

//---------------------- Tests ------------------------
const limit = math.MaxUint8

// Send and check Int values
func TestQueueIntValues(t *testing.T) {

	q := NewMutexQueue(1024, 2)
	ch := make(chan bool, 2)
	go func() {
		for i := 0; i < limit; i++ {
			q.Enqueue(i)
		}
		ch <- true
	}()
	go func() {
		var r int
		for i := 0; i < limit; i++ {
			r = (<-q.Dequeue()).(int)
			assert.Equal(t, i, r)
		}
		ch <- true
	}()
	<-ch
	<-ch
	assert.Equal(t, 0, len(q.queue))
}

// Send and check Int pointers
func TestQueueIntPointers(t *testing.T) {
	q := NewMutexQueue(1024, 2)
	ch := make(chan bool, 2)
	go func() {
		for i := 0; i < limit; i++ {
			x := i
			q.Enqueue(&x)
		}
		ch <- true
	}()
	go func() {
		var r *int
		for i := 0; i < limit; i++ {
			r = (<-q.Dequeue()).(*int)
			assert.Equal(t, i, *r)
		}
		ch <- true
	}()
	<-ch
	<-ch
	assert.Equal(t, 0, len(q.queue))
}

// ------------------ Benchmark  -------------------

var queueBenchmarkElement interface{}

func BenchmarkQueue(b *testing.B) {

	var r interface{}
	q := NewMutexQueue(1024, 2)

	for n := 0; n < b.N; n++ {
		ch := make(chan bool, 2)
		go func() {
			for i := 0; i < math.MaxUint16; i++ {
				q.Enqueue(i)
			}
			ch <- true
		}()
		go func() {
			for i := 0; i < math.MaxUint16; i++ {
				r = <-q.Dequeue()
			}
			ch <- true
		}()
		<-ch
		<-ch
	}
	queueBenchmarkElement = r
}