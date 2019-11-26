package internal

type RetryOptions struct {
	WaitTime int64
	Attempt  int
}

func MakeDefaultRetryOptions() *RetryOptions {
	return &RetryOptions{
		Attempt:  5,
		WaitTime: 3,
	}
}

var DefaultRetryOptions = MakeDefaultRetryOptions()
