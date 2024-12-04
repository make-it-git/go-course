package example

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

func TestWithErrGroup(t *testing.T) {
	g, ctx := errgroup.WithContext(context.Background())
	_ = ctx

	tasks := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	g.SetLimit(3)

	for _, task := range tasks {
		task := task // создаём копию для правильной работы с переменной цикла, актуально до версии 1.22
		g.Go(func() error {
			time.Sleep(time.Second)
			select {
			case <-ctx.Done():
				t.Logf(fmt.Sprintf("Context cancelled for %d", task))
				return errors.New("context cancelled")
			default:
				// do nothing
			}
			if task == 4 {
				return errors.New(fmt.Sprintf("error in task %d", task))
			}
			t.Logf("Task %d completed", task)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		t.Errorf("errgroup returned an error: %v", err)
	} else {
		t.Log("All tasks completed successfully")
	}
}
