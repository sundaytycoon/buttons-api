package retry

import "time"

func Retry(count int, dur time.Duration, f func() (interface{}, error)) (interface{}, error) {
	var value interface{}
	for i := 0; i <= count; i++ {
		v, err := f()
		if err != nil {
			<-time.After(dur)
			continue
		}
		value = v
		continue
	}
	return value, nil
}
