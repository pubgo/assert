package assert

import (
	"log"
	"runtime"
	"time"
)

type TaskFn func(args ...interface{}) *_task_fn

func NewTask(max int, maxDur time.Duration) *task {
	_t := &task{max: max, maxDur: maxDur, q: make(chan *_task_fn, max)}
	_t._handle()
	return _t
}

func FuncOf(fn interface{}, efn func(err error)) TaskFn {
	assertFn(fn)
	assertFn(efn)

	return func(args ...interface{}) *_task_fn {
		return &_task_fn{
			fn:   fn,
			args: args,
			efn:  efn,
		}
	}
}

type _task_fn struct {
	fn   interface{}
	args []interface{}
	efn  func(err error)
}

func (t *_task_fn) _do() {
	if err := KTry(t.fn, t.args...).(*KErr); err != nil {
		if t.efn != nil {
			t.efn(err)
		}
	}
}

type task struct {
	maxDur time.Duration
	curDur time.Duration
	max    int
	q      chan *_task_fn
}

func (t *task) Do(f TaskFn, args ...interface{}) {

	for {
		if len(t.q) < t.max && t.curDur < t.maxDur {
			t.q <- f(args...)
			break
		}

		if len(t.q) < runtime.NumCPU()*2 {
			t.curDur = 0
		}

		log.Printf("q_l:%d cur_dur:%s", len(t.q), t.curDur.String())
		time.Sleep(time.Millisecond * 200)
	}
}

func (t *task) _handle() {
	go func() {
		for {
			select {
			case _fn := <-t.q:
				go func() {
					t.curDur = FnCost(_fn._do)
				}()
			case <-time.NewTicker(time.Second * 2).C:

			}
		}
	}()
}
