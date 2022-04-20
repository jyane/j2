package util

import (
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
)

func ReadProtoTextFromLocalStorage(path string, m proto.Message) error {
	r, err := os.Open(path)
	if err != nil {
		return err
	}
	defer r.Close()
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	if err := proto.UnmarshalText(string(data), m); err != nil {
		return err
	}
	return nil
}
