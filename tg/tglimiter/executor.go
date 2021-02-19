package tglimiter

import (
	"time"

	limiter "github.com/chatex-com/rate-limiter"
	"github.com/chatex-com/rate-limiter/pkg/config"
)

const maxRequestsInSecond = 25

type Executor struct {
	rateLimiter *limiter.RateLimiter
}

func newRateLimiter() *limiter.RateLimiter {
	cfg := config.NewConfigWithQuotas([]*config.Quota{
		config.NewQuota(maxRequestsInSecond, time.Second),
	})
	cfg.Concurrency = 1

	rateLimiter, _ := limiter.NewRateLimiter(cfg)
	rateLimiter.Start()

	return rateLimiter
}

func NewExecutor() *Executor {
	return &Executor{
		rateLimiter: newRateLimiter(),
	}
}

func (e *Executor) Execute(f func() (interface{}, error)) (interface{}, error) {
	response := <-e.rateLimiter.Execute(func() (interface{}, error) {
		return f()
	})

	return response.Result, response.Error
}
