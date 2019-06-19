package assert

import (
	"encoding/json"
	"fmt"
)

type _KErr struct {
	Tag    string                 `json:"tag,omitempty"`
	M      map[string]interface{} `json:"m,omitempty"`
	Err    error                  `json:"err,omitempty"`
	Msg    string                 `json:"msg,omitempty"`
	Caller string                 `json:"caller,omitempty"`
	Sub    *_KErr                 `json:"sub,omitempty"`
}

type KErr struct {
	tag    string
	m      map[string]interface{}
	err    error
	msg    string
	caller string
	sub    *KErr
}

func (t *KErr) kerr() *_KErr {
	return &_KErr{
		Tag:    t.tag,
		M:      t.m,
		Err:    t.err,
		Msg:    t.msg,
		Caller: t.caller,
	}
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

func (t *KErr) Err() error {
	return t.err
}

func (t *KErr) Error() string {
	return t.err.Error()
}

func (t *KErr) Caller(caller string) {
	t.caller = caller
}

func (t *KErr) Tag() string {
	return t.tag
}

func (t *KErr) StackTrace() string {
	kerr := t.kerr()
	c := kerr
	for t.sub != nil {
		c.Sub = t.sub.kerr()
		t.sub = t.sub.sub
		c = c.Sub
	}

	_dt, _ := json.MarshalIndent(kerr, "", "\t")
	return string(_dt)
}

func (t *KErr) tErr() (err error) {
	err = t.err
	t.err = nil
	return
}

func (t *KErr) P() {
	fmt.Println(t.StackTrace())
}
