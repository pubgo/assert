package assert

import (
	"log"
	"runtime"
	"time"
)

type TaskFn func(args ...interface{}) *_task_fn

func NewTask(max int, maxDur time.Duration) *task {
	_t := &task{
		max: max, maxDur:
		maxDur, q: make(chan *_task_fn, max),
		_curDur:   make(chan time.Duration, max),
		_done:     make(chan bool, max),
		_stop_q:   make(chan error),
	}
	go _t._handle()
	return _t
}

func TaskOf(fn interface{}, efn ...func(err error)) TaskFn {
	assertFn(fn)
	T(len(efn) != 0 && efn[0] == nil, "efn error")

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
	efn  []func(err error)
}

type task struct {
	maxDur time.Duration

	curDur  time.Duration
	_curDur chan time.Duration

	_done chan bool
	max   int

	q chan *_task_fn

	_stop_q chan error
	_stop   error
}

func (t *task) Wait() {
	for len(t._done) > 0 {
		if t._stop != nil {
			return
		}
		time.Sleep(time.Millisecond * 200)
	}
}

func (t *task) Do(f TaskFn, args ...interface{}) error {
	for {
		if t._stop != nil {
			return t._stop
		}

		if len(t._done) < t.max && t.curDur < t.maxDur {
			t.q <- f(args...)
			t._done <- true
			return nil
		}

		if len(t._done) < runtime.NumCPU()*2 {
			t.curDur = 0
		}

		log.Printf("q_l:%d cur_dur:%s max_q:%d max_dur:%s", len(t.q), t.curDur.String(), t.max, t.maxDur.String())
		time.Sleep(time.Millisecond * 200)
	}
}

func (t *task) _handle() {
	for t._stop == nil {
		select {
		case _fn := <-t.q:
			go func() {
				t._curDur <- FnCost(func() {
					if err := KTry(_fn.fn, _fn.args...); err != nil {
						if len(_fn.efn) != 0 && _fn.efn[0] != nil {
							if _err := KTry(_fn.efn[0], err); _err != nil {
								t._stop_q <- _err
							}
						}
					}
				})
				<-t._done
			}()
		case _c := <-t._curDur:
			t.curDur = _c
		case _e := <-t._stop_q:
			t._stop = _e
		}
	}
}
