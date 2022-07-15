package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func getBatch(n int64, pool int64) []user {
	ch := make(chan user, pool)
	for i := int64(0); i < n; i++ {
		go func(id int64) {
			ch <- getOne(id)
		}(i)
	}
	res := make([]user, 0, n)
	for i := int64(0); i < n; i++ {

		res = append(res, <-ch)
	}
	return res
}
