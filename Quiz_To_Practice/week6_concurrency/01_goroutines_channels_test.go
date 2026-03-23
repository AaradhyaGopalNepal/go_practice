package week6_concurrency

import (
	"testing"
	"time"
	"sort"
	"github.com/stretchr/testify/assert"
)

func TestParallelProcessor(t *testing.T) {
	t.Skip("Remove this line after implementing ParallelProcessor")

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	processor := func(n int) int {
		time.Sleep(10 * time.Millisecond) // Simulate work
		return n * n
	}

	start := time.Now()
	// results := ParallelProcessor(items, 3, processor)
	elapsed := time.Since(start)

	// Check results are correct and in order
	expected := []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100}
	// assert.Equal(t, expected, results)

	// Check parallelism (should be faster than sequential)
	// Sequential would take ~100ms, parallel should be ~40ms
	assert.Less(t, elapsed, 60*time.Millisecond)
}

func TestPipelineProcessor(t *testing.T) {
	t.Skip("Remove this line after implementing PipelineProcessor")

	numbers := []int{1, 2, 3, 4, 5}

	// Pipeline: multiply by 2, add 10, divide by 3
	// 1: 1*2=2, 2+10=12, 12/3=4
	// 2: 2*2=4, 4+10=14, 14/3=4
	// 3: 3*2=6, 6+10=16, 16/3=5
	// 4: 4*2=8, 8+10=18, 18/3=6
	// 5: 5*2=10, 10+10=20, 20/3=6

	expected := []int{4, 4, 5, 6, 6}
	// results := PipelineProcessor(numbers)
	// assert.Equal(t, expected, results)
}

func TestFanInFanOut(t *testing.T) {
	t.Skip("Remove this line after implementing FanInFanOut")

	// Create multiple source channels
	ch1 := make(chan int, 3)
	ch2 := make(chan int, 3)
	ch3 := make(chan int, 3)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)

	ch2 <- 4
	ch2 <- 5
	close(ch2)

	ch3 <- 6
	close(ch3)

	sources := []<-chan int{ch1, ch2, ch3}
	// merged := FanInFanOut(sources)

	// Collect all values
	var results []int
	// for val := range merged {
	// 	results = append(results, val)
	// }

	// Check we got all values (order doesn't matter)
	sort.Ints(results)
	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, results)
}

func TestRateLimiter(t *testing.T) {
	t.Skip("Remove this line after implementing RateLimiter")

	// limiter := NewRateLimiter(3, 100*time.Millisecond)

	var executed []time.Time
	items := make([]func(), 6)
	for i := range items {
		items[i] = func() {
			executed = append(executed, time.Now())
		}
	}

	start := time.Now()
	// limiter.Process(items)

	// Should process 3 items immediately, then 3 more after 100ms
	assert.Len(t, executed, 6)

	// First 3 should be immediate
	for i := 0; i < 3; i++ {
		assert.Less(t, executed[i].Sub(start), 50*time.Millisecond)
	}

	// Next 3 should be after rate limit duration
	for i := 3; i < 6; i++ {
		assert.Greater(t, executed[i].Sub(start), 90*time.Millisecond)
	}
}

func TestProducerConsumer(t *testing.T) {
	t.Skip("Remove this line after implementing ProducerConsumer")

	// results := ProducerConsumer(3, 2, 5, 20)

	// Should have processed 20 items
	assert.Len(t, results, 20)

	// All numbers from 0-19 should be present (order doesn't matter)
	sort.Ints(results)
	expected := make([]int, 20)
	for i := range expected {
		expected[i] = i
	}
	assert.Equal(t, expected, results)
}

func TestTimeoutProcessor(t *testing.T) {
	t.Skip("Remove this line after implementing TimeoutProcessor")

	items := []int{1, 2, 3, 4, 5}
	processor := func(n int) (int, error) {
		if n == 3 {
			time.Sleep(200 * time.Millisecond) // This will timeout
		} else {
			time.Sleep(50 * time.Millisecond)
		}
		return n * 2, nil
	}

	// results := TimeoutProcessor(items, processor, 100*time.Millisecond)

	// Should have 4 results (item 3 timed out)
	assert.Len(t, results, 4)
	assert.Contains(t, results, 2)  // 1*2
	assert.Contains(t, results, 4)  // 2*2
	assert.NotContains(t, results, 6) // 3*2 (timed out)
	assert.Contains(t, results, 8)  // 4*2
	assert.Contains(t, results, 10) // 5*2
}

func TestBroadcaster(t *testing.T) {
	t.Skip("Remove this line after implementing Broadcaster")

	// broadcaster := NewBroadcaster()

	// Create subscribers
	// sub1 := broadcaster.Subscribe()
	// sub2 := broadcaster.Subscribe()
	// sub3 := broadcaster.Subscribe()

	// Collect results
	var results1, results2, results3 []interface{}
	done := make(chan bool, 3)

	go func() {
		// for val := range sub1 {
		// 	results1 = append(results1, val)
		// }
		done <- true
	}()

	go func() {
		// for val := range sub2 {
		// 	results2 = append(results2, val)
		// }
		done <- true
	}()

	go func() {
		// for val := range sub3 {
		// 	results3 = append(results3, val)
		// }
		done <- true
	}()

	// Broadcast values
	// broadcaster.Broadcast("hello")
	// broadcaster.Broadcast(42)
	// broadcaster.Broadcast(true)

	time.Sleep(100 * time.Millisecond)

	// All subscribers should receive all values
	expected := []interface{}{"hello", 42, true}
	assert.Equal(t, expected, results1)
	assert.Equal(t, expected, results2)
	assert.Equal(t, expected, results3)
}

func TestWorkerPool(t *testing.T) {
	t.Skip("Remove this line after implementing WorkerPool")

	// pool := NewWorkerPool(3)

	processor := func(job Job) Result {
		time.Sleep(10 * time.Millisecond)
		n := job.Data.(int)
		return Result{
			JobID: job.ID,
			Data:  n * 2,
			Error: nil,
		}
	}

	// pool.Start(processor)

	// Submit jobs
	for i := 0; i < 10; i++ {
		// pool.Submit(Job{ID: i, Data: i})
	}

	// Collect results
	var results []Result
	// resultChan := pool.Results()
	timeout := time.After(200 * time.Millisecond)

	// for i := 0; i < 10; i++ {
	// 	select {
	// 	case result := <-resultChan:
	// 		results = append(results, result)
	// 	case <-timeout:
	// 		t.Fatal("Timeout waiting for results")
	// 	}
	// }

	// pool.Stop()

	// Check all jobs were processed
	assert.Len(t, results, 10)

	// Check results are correct
	for _, result := range results {
		expected := result.JobID * 2
		assert.Equal(t, expected, result.Data)
	}
}