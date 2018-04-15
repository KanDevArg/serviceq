package model

type Config struct {
	ListenerPort          string
	Proto                 string
	Endpoints             []Endpoint
	CustomRequestHeaders  []string
	CustomResponseHeaders []string
	ConcurrencyPeak       int64
	RetryGap              int
	OutReqTimeout         int32
	EnableProfilingFor    string
}
