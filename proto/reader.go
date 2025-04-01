package reader

import (
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
)

func ReadProto(data []byte, msgType proto.Message) (interface{}, error) {
	val := reflect.ValueOf(msgType)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return nil, fmt.Errorf("msgType must be a pointer to a struct")
	}

	t := val.Elem().Type()
	instance := reflect.New(t).Interface()

	if err := proto.Unmarshal(data, instance.(proto.Message)); err != nil {
		return nil, err
	}

	return instance, nil
}