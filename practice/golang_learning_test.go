package practice

import (
	"context"
	"github.com/stretchr/testify/assert"
	"strings"
	"sync"
	"testing"
	"time"
)

// golang 학습 테스트
func TestGolang(t *testing.T) {
	t.Run("string test", func(t *testing.T) {
		str := "Ann,Jenny,Tom,Zico"
		actual := strings.Split(str, ",")
		expected := []string{"Ann", "Jenny", "Tom", "Zico"}
		assert.Equal(t, expected, actual)
	})

	t.Run("goroutine에서 slice에 값 추가해보기", func(t *testing.T) {
		const SIZE int = 100
		var numbers []int
		var wg sync.WaitGroup
		ch := make(chan int)

		wg.Add(SIZE)

		for i := 0; i < 100; i++ {
			i := i
			go func(i int) {
				ch <- i
			}(i)
		}

		go func() {
			for number := range ch {
				numbers = append(numbers, number)
				wg.Done()
			}
		}()

		var expected []int
		for i := 0; i < SIZE; i++ {
			expected = append(expected, i)
		}
		wg.Wait()
		assert.ElementsMatch(t, expected, numbers)
	})

	t.Run("fan out, fan in", func(t *testing.T) {

		inputCh := generate()
		outputCh := make(chan int)
		go func() {
			defer close(outputCh)
			for {
				select {
				case value, ok := <-inputCh:
					if !ok {
						return
					}
					outputCh <- value * 10
				}
			}
		}()

		var actual []int
		for value := range outputCh {
			actual = append(actual, value)
		}
		expected := []int{10, 20, 30}
		assert.Equal(t, expected, actual)
	})

	t.Run("context timeout", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		//goland:noinspection GoVetLostCancel
		ctx, _ := context.WithTimeout(context.Background(), add)

		var endTime time.Time
		select {
		case <-ctx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context deadline", func(t *testing.T) {
		startTime := time.Now()
		add := time.Second * 3
		//goland:noinspection GoVetLostCancel
		ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(add))

		var endTime time.Time
		select {
		case <-ctx.Done():
			endTime = time.Now()
			break
		}

		assert.True(t, endTime.After(startTime.Add(add)))
	})

	t.Run("context value", func(t *testing.T) {
		// context에 key, value를 추가해보세요.
		// 추가된 key, value를 호출하여 assert로 값을 검증해보세요.
		// 추가되지 않은 key에 대한 value를 assert로 검증해보세요.
	})
}

func generate() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 1; i <= 3; i++ {
			ch <- i
		}
	}()
	return ch
}
