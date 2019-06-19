package assert

import (
	"encoding/json"
)

type KErr struct {
	tag    string                 `json:"tag,omitempty"`
	m      map[string]interface{} `json:"m,omitempty"`
	err    error                  `json:"err,omitempty"`
	msg    string                 `json:"msg,omitempty"`
	caller string                 `json:"caller,omitempty"`
	sub    *KErr                  `json:"sub,omitempty"`
}

func (t *KErr) copy() *KErr {
	return &KErr{
		tag:    t.tag,
		m:      t.m,
		err:    t.err,
		msg:    t.msg,
		caller: t.caller,
		sub:    t.sub,
	}
}

func (t *KErr) Error() string {
	return t.err.Error()
}

func (t *KErr) Caller(caller string) {
	t.caller = caller
}

func (t *KErr) StackTrace() string {
	_dt, _ := json.MarshalIndent(t, "", "\t")
	return string(_dt)
}

func (t *KErr) tErr() (err error) {
	err = t.err
	t.err = nil
	return
}

func (t *KErr) Panic() {
	panic(t)
}

func (t *KErr) P() {
	P(t)
}
