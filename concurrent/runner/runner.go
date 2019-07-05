package runner

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	// interrupt 通道报告从操作系统发送信号
	interrupt chan os.Signal

	// complete 通道报告处理任务已经完成
	complete chan error

	// timeout 报告处理任务已经超时
	// 接收一个 time.Time 的值，该程序会试图清理状态并停止工作
	timeout <-chan time.Time

	// tasks 持有一组以索引顺序依次执行的函数
	tasks []func(int)
}

// 接收到任务执行超时时返回
var ErrTimeout = errors.New("received timeout")

// 接收到操作系统的事件时返回
var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

func (r *Runner) Add(tasks ...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start() error {
	// 我们希望接收所有中断的信号
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	// select 是一个控制结构，每个 case 必须是一个通信操作，要么是发送要么是接收
	// select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。
	select {
	case err := <-r.complete:
		return err
	case <-r.timeout:
		return ErrTimeout
	}
}

func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}

		// 执行已注册的任务
		task(id)
	}

	return nil
}

// gotInterrupt 验证是否接收到了中断信号
func (r *Runner) gotInterrupt() bool {
	select {
	// 当中断事件被触发时发出的信号
	case <-r.interrupt:
		// 停止接收后续的任何信号
		signal.Stop(r.interrupt)
		return true
	// 继续正常运行
	default:
		return false
	}
}
