package rate

import (
	"log"
	"testing"
	"time"
)

func TestRateLimiter_Wait_noblock(t *testing.T) {
	start := time.Now()
	limit := 5
	interval := time.Second * 3
	limiter := New(limit, interval)
	for i := 0; i < limit; i++ {
		limiter.Wait()
	}
	if time.Now().Sub(start) >= interval {
		t.Error("The limiter blocked when it shouldn't have")
	}
}

func TestRateLimiter_Wait_block(t *testing.T) {
	start := time.Now()
	limit := 5
	interval := time.Second * 3
	limiter := New(limit, interval)
	for i := 0; i < limit+1; i++ {
		limiter.Wait()
	}
	if time.Now().Sub(start) < interval {
		t.Error("The limiter didn't block when it should have")
	}
}

func TestRateLimiter_Try(t *testing.T) {
	limit := 5
	interval := time.Second * 3
	limiter := New(limit, interval)
	for i := 0; i < limit; i++ {
		if ok, _, _ := limiter.Try(); !ok {
			t.Fatalf("Should have allowed try on attempt %d", i)
		}
	}
	if ok, _, _ := limiter.Try(); ok {
		t.Fatal("Should have not allowed try on final attempt")
	}
}

func TestRateLimiter_Remaining(t *testing.T) {
	limit := 5
	interval := time.Second * 3
	limiter := New(limit, interval)
	for i := 0; i < 3*limit; {
		if ok, _, _ := limiter.Try(); ok {
			i++
		} else {
			log.Println("wait...")
		}
		log.Println(limiter.Remaining())
		time.Sleep(time.Millisecond * 100)
	}
}

func TestRateLimiter_SetRemaining(t *testing.T) {
	limit := 5
	interval := time.Second * 3
	limiter := New(limit, interval)
	for i := 0; i < 2*limit; i++ {
		if ok, _, remain := limiter.Try(); !ok {
			log.Println("remain plus 2")
			limiter.SetRemaining(remain + 2)
		}
		log.Println(i)
	}
}

func TestRateLimiter_Reverse(t *testing.T) {
	limit := 5
	interval := time.Second * 3
	limiter := New(limit, interval)
	limiter.Reverse()
}
