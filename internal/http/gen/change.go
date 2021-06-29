package gen

import (
	"encoding/json"
)

func (p Task) change() string {
	s, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	return string(s)
}
