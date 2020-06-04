package connpool

import "time"

type Options struct {
	initialCap  int           // 初始化容量
	maxCap      int           // 最大容量
	idleTimeout time.Duration // 空闲超时时长
	maxIdle     int           // 最大空闲连接数
	dialTimeout time.Duration // 连接超时时长
}

type Option func(*Options)

func WithInitialCap (initialCap int) Option {
	return func(o *Options) {
		o.initialCap = initialCap
	}
}

func WithMaxCap (maxCap int) Option {
	return func(o *Options) {
		o.maxCap = maxCap
	}
}


func WithMaxIdle (maxIdle int) Option {
	return func(o *Options) {
		o.maxIdle = maxIdle
	}
}

func WithIdleTimeout(idleTimeout time.Duration) Option {
	return func(o *Options) {
		o.idleTimeout = idleTimeout
	}
}

func WithDialTimeout(dialTimeout time.Duration) Option {
	return func(o *Options) {
		o.dialTimeout = dialTimeout
	}
}
