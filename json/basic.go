package main

import "encoding/json"

type Valiable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func load(obj []byte) (*Valiable, error) {
	valiable := new(Valiable)
	err := json.Unmarshal(obj, valiable)

	return valiable, err
}

func dump(valiable *Valiable) ([]byte, error) {
	return json.Marshal(valiable)
}

func main() {
	object := []byte(`{"key": "egg", "value": "spam"}`)

	// to struct
	valiable, _ := load(object)

	// to []byte
	obj, _ := dump(valiable)

	_ = obj
}
