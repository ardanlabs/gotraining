package main

import (
	"context"
	"errors"
	"fmt"
	"time"
)

func main() {
	worker1 := func(ctx context.Context) (interface{}, error) {
		time.Sleep(time.Millisecond)
		return "test string", nil
	}
	v1, err := retry(context.Background(), time.Second, worker1)
	fmt.Println(v1.(string), err)

	worker2 := func(ctx context.Context) (interface{}, error) {
		time.Sleep(time.Millisecond)
		return 9999999, nil
	}
	v2, err := retry(context.Background(), time.Second, worker2)
	fmt.Println(v2.(int), err)

	worker3 := func(ctx context.Context) (interface{}, error) {
		time.Sleep(time.Millisecond)
		return &user{"bill", "b@email.com"}, nil
	}
	v3, err := retry(context.Background(), time.Second, worker3)
	fmt.Println(v3.(*user), err)
}

// =============================================================================

type Worker func(ctx context.Context) (interface{}, error)

func retry(ctx context.Context, retryInterval time.Duration, worker Worker) (interface{}, error) {
	var retry *time.Timer

	if ctx.Err() != nil {
		return nil, errors.New("error")
	}

	for {
		if value, err := worker(ctx); err == nil {
			return value, nil
		}

		if ctx.Err() != nil {
			return nil, errors.New("error")
		}

		if retry == nil {
			retry = time.NewTimer(retryInterval)
		}

		select {
		case <-ctx.Done():
			retry.Stop()
			return nil, errors.New("error")
		case <-retry.C:
			retry.Reset(retryInterval)
		}
	}
}

// =============================================================================

type user struct {
	name  string
	email string
}
