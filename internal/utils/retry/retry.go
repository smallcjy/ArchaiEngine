package retry

import (
	"github.com/avast/retry-go"
)

var (
	RetryStrategy = []retry.Option{}
)

// 重试配置
func NewRetryOption(retry []retry.Option) {
	if retry == nil {
		panic("data err")
	}
	RetryStrategy = retry
}

// 重试请求
func RetryFuncInterface(f func() error) error {
	err := retry.Do(f, RetryStrategy...)
	if err != nil {
		return err
	}
	return nil
}
