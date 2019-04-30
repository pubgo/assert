package assert

import (
	"encoding/json"
)

type KErr struct {
	Msg        string `json:"msg,omitempty"`
	FuncCaller string `json:"funcCaller,omitempty"`
	Sub        error  `json:"sub,omitempty"`
}

func (t *KErr) Error() string {
	_dt, _ := json.Marshal(t)
	return string(_dt)
}

func (t *KErr) Panic() {
	panic(t)
}

func (t *KErr) P() {
	P(t)
}
