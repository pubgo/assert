package assert

import (
	"encoding/json"
)

type KErr struct {
	Tag    string                 `json:"tag,omitempty"`
	M      map[string]interface{} `json:"m,omitempty"`
	Err    error                  `json:"err,omitempty"`
	Msg    string                 `json:"msg,omitempty"`
	Caller string                 `json:"caller,omitempty"`
	Sub    *KErr                  `json:"sub,omitempty"`
}

func (t *KErr) Error() string {
	return t.Err.Error()
}

func (t *KErr) StackTrace() interface{} {
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
