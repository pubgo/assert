package assert

import (
	"log"
	"sync"
)

var _kerr = &sync.Pool{
	New: func() interface{} {
		return &KErr{}
	},
}

func kerrGet() *KErr {
	defer func() {
		if err := recover(); err != nil {
			log.Println(funcCaller(3))
			log.Fatalln(err)
		}
	}()

	return _kerr.Get().(*KErr)
}

func kerrPut(m *KErr) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(funcCaller(3))
			log.Fatalln(err)
		}
	}()

	_kerr.Put(m)
}
