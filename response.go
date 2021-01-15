package graphrpc

import (
	"encoding/json"
	"fmt"

	"github.com/aeramu/graphrpc/proto"
)

type Response struct {
	Data   []byte
	Errors []*proto.Error
}

func (r Response) Consume(v interface{}) error {
	if len(r.Errors) > 0 {
		return fmt.Errorf("%v", r.Errors)
	}
	if err := json.Unmarshal(r.Data, &v); err != nil {
		return err
	}
	return nil
}
