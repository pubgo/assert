package assert

import (
	"errors"
	"testing"
)

func a1() error {
	return _Try(func() {
		ErrWrap(errors.New("sbhbhbh"), func(m *M) {
			m.Msg("test shhh")
			m.M["ss"] = 1
			m.M["input"] = 1
		})

		T(true, func(m *M) {
			m.Msg("好东西%d", 1)
		})
	})
}

func TestName(t *testing.T) {
	P(IfEquals(1, 2, 34))
	P(IfEquals(nil, nil))
	P(IfEquals(0, ""))
	P(IfEquals("", ""))

	P(_Try(func() {
		Throw(_Try(func() {
			ErrWrap(_Try(func() {
				ErrWrap(a1(), func(m *M) {
					m.Msg("ok111")
					m.Tag("test tag")
				})
			}), func(m *M) {
				m.Msg("test 123")
			})
		}))
	}))
}
