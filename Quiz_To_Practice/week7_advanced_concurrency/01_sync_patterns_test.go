package week7_advanced_concurrency

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"
	"github.com/stretchr/testify/assert"
)

func TestSemaphore(t *testing.T) {
	t.Skip("Remove this line after implementing Semaphore")

	// sem := NewSemaphore(2)

	var wg sync.WaitGroup
	results := make([]int, 0)
	mu := sync.Mutex{}

	// Start 5 goroutines, but only 2 can run at once
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// sem.Acquire()
			// defer sem.Release()

			mu.Lock()
			results = append(results, id)
			mu.Unlock()

			time.Sleep(50 * time.Millisecond)
		}(i)
	}

	// Test TryAcquire with timeout
	go func() {
		time.Sleep(10 * time.Millisecond)
		// acquired := sem.TryAcquire(10 * time.Millisecond)
		// assert.False(t, acquired) // Should fail, semaphore is full
	}()

	wg.Wait()
	assert.Len(t, results, 5)
}

func TestRWLock(t *testing.T) {
	t.Skip("Remove this line after implementing RWLock")

	// rwLock := NewRWLock(false) // Reader preference
	data := 0

	var wg sync.WaitGroup

	// Multiple readers
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// rwLock.RLock()
			// defer rwLock.RUnlock()
			_ = data // Read
			time.Sleep(50 * time.Millisecond)
		}()
	}

	// Writer
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		// rwLock.WLock()
		// defer rwLock.WUnlock()
		data = 42 // Write
	}()

	wg.Wait()
	assert.Equal(t, 42, data)
}

func TestBarrier(t *testing.T) {
	t.Skip("Remove this line after implementing Barrier")

	n := 3
	// barrier := NewBarrier(n)

	results := make([]int, 0)
	mu := sync.Mutex{}

	for i := 0; i < n; i++ {
		go func(id int) {
			time.Sleep(time.Duration(id*10) * time.Millisecond)

			mu.Lock()
			results = append(results, id)
			mu.Unlock()

			// barrier.Wait() // All goroutines wait here

			mu.Lock()
			results = append(results, id+10)
			mu.Unlock()
		}(i)
	}

	time.Sleep(200 * time.Millisecond)

	// All first phase should complete before any second phase
	assert.Len(t, results, 6)
	// First 3 should be 0,1,2 (any order)
	// Last 3 should be 10,11,12 (any order)
}

func TestFuture(t *testing.T) {
	t.Skip("Remove this line after implementing Future")

	// Test successful future
	// future := NewFuture(func() (interface{}, error) {
	// 	time.Sleep(50 * time.Millisecond)
	// 	return 42, nil
	// })

	// assert.False(t, future.IsReady())

	// result, err := future.Get()
	// assert.NoError(t, err)
	// assert.Equal(t, 42, result)
	// assert.True(t, future.IsReady())

	// Test future with error
	// errorFuture := NewFuture(func() (interface{}, error) {
	// 	return nil, errors.New("computation failed")
	// })

	// _, err = errorFuture.Get()
	// assert.Error(t, err)

	// Test timeout
	// slowFuture := NewFuture(func() (interface{}, error) {
	// 	time.Sleep(200 * time.Millisecond)
	// 	return "slow", nil
	// })

	// _, err = slowFuture.GetWithTimeout(50 * time.Millisecond)
	// assert.Error(t, err)
}

func TestCircuitBreaker(t *testing.T) {
	t.Skip("Remove this line after implementing CircuitBreaker")

	// cb := NewCircuitBreaker(3, 100*time.Millisecond)

	// Function that fails first 3 times
	callCount := 0
	fn := func() (interface{}, error) {
		callCount++
		if callCount <= 3 {
			return nil, errors.New("service error")
		}
		return "success", nil
	}

	// First 3 calls should fail and open circuit
	for i := 0; i < 3; i++ {
		// _, err := cb.Call(fn)
		// assert.Error(t, err)
	}

	// assert.Equal(t, Open, cb.GetState())

	// Circuit is open, call should fail immediately
	// _, err := cb.Call(fn)
	// assert.Error(t, err)
	// assert.Equal(t, 3, callCount) // Function not called

	// Wait for timeout
	time.Sleep(150 * time.Millisecond)

	// Circuit should be half-open, next call succeeds
	// result, err := cb.Call(fn)
	// assert.NoError(t, err)
	// assert.Equal(t, "success", result)
	// assert.Equal(t, Closed, cb.GetState())
}

func TestContextPipeline(t *testing.T) {
	t.Skip("Remove this line after implementing ContextPipeline")

	ctx, cancel := context.WithCancel(context.Background())
	input := make(chan int, 5)

	// Send values
	for i := 1; i <= 5; i++ {
		input <- i
	}
	close(input)

	// output := ContextPipeline(ctx, input)

	// Collect results
	var results []int
	// for val := range output {
	// 	results = append(results, val)
	// }

	// Pipeline: double -> add 10 -> square
	// 1: 2 -> 12 -> 144
	// 2: 4 -> 14 -> 196
	// 3: 6 -> 16 -> 256
	// 4: 8 -> 18 -> 324
	// 5: 10 -> 20 -> 400
	expected := []int{144, 196, 256, 324, 400}
	assert.Equal(t, expected, results)

	// Test cancellation
	ctx2, cancel2 := context.WithCancel(context.Background())
	input2 := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		input2 <- i
	}

	// output2 := ContextPipeline(ctx2, input2)

	// Cancel after reading first value
	// <-output2
	cancel2()

	// Should stop producing values
	time.Sleep(50 * time.Millisecond)
	select {
	// case <-output2:
		// Should not receive more values
	default:
		// Expected
	}

	cancel() // Clean up
}

func TestRetryWithBackoff(t *testing.T) {
	t.Skip("Remove this line after implementing RetryWithBackoff")

	config := RetryConfig{
		MaxAttempts:  4,
		InitialDelay: 10 * time.Millisecond,
		MaxDelay:     100 * time.Millisecond,
		Multiplier:   2.0,
	}

	// Test successful retry
	attempts := 0
	fn := func() error {
		attempts++
		if attempts < 3 {
			return errors.New("temporary error")
		}
		return nil
	}

	start := time.Now()
	// err := RetryWithBackoff(config, fn)
	elapsed := time.Since(start)

	// assert.NoError(t, err)
	assert.Equal(t, 3, attempts)
	// Should have delays: 10ms, 20ms (total ~30ms)
	assert.Greater(t, elapsed, 25*time.Millisecond)

	// Test max attempts exceeded
	failFn := func() error {
		return errors.New("persistent error")
	}

	// err = RetryWithBackoff(config, failFn)
	// assert.Error(t, err)
}

func TestObjectPool(t *testing.T) {
	t.Skip("Remove this line after implementing ObjectPool")

	created := 0
	factory := func() interface{} {
		created++
		return &struct{ ID int }{ID: created}
	}

	reset := func(obj interface{}) error {
		// Reset object state
		return nil
	}

	// pool := NewObjectPool(factory, reset, 3)

	// Get objects
	// obj1 := pool.Get()
	// obj2 := pool.Get()
	// obj3 := pool.Get()

	assert.Equal(t, 3, created)

	// Return one to pool
	// pool.Put(obj1)

	// available, total := pool.Stats()
	// assert.Equal(t, 1, available)
	// assert.Equal(t, 3, total)

	// Get should reuse from pool
	// obj4 := pool.Get()
	// assert.Equal(t, obj1, obj4) // Same object
	// assert.Equal(t, 3, created) // No new object created

	// Get new object when pool is empty
	// obj5 := pool.Get()
	// assert.Equal(t, 4, created) // New object created
}