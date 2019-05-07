package assert

import (
	"encoding/json"
)

type KErr struct {
	Msg        string `json:"msg,omitempty"`
	FuncCaller string `json:"funcCaller,omitempty"`
	Sub        *KErr  `json:"sub,omitempty"`
	Err        error  `json:"err,omitempty"`
}

func (t *KErr) Error() string {
	return t.Err.Error()
}

func (t *KErr) StackTrace() string {
	_dt, _ := json.Marshal(t)
	return string(_dt)
}

func (t *KErr) tErr() (err error) {
	err = t.Err
	t.Err = nil
	return
}

func (t *KErr) Panic() {
	panic(t)
}

func (t *KErr) P() {
	P(t)
}
