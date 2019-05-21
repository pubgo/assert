package assert

func Marshal(v interface{}) []byte {
	_dt, err := json.Marshal(v)
	Throw(err)
	return _dt
}

func Unmarshal(data []byte, v interface{}) {
	Throw(json.Unmarshal(data, v))
}
