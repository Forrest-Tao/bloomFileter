package bloomFileter

import (
	"context"
	"fmt"
	"github.com/demdxx/gocast"
)

func (s *BloomServer) Set(ctx context.Context, key, val string) error {
	keysAndArgs := make([]interface{}, 0, 2+s.k)
	keysAndArgs = append(keysAndArgs, key, val)
	for _, encrypted := range s.getEncrypted(val) {
		keysAndArgs = append(keysAndArgs, encrypted)
	}

	rawResp, err := s.client.Eval(ctx, LuaBloomBatchSetBits, 1, keysAndArgs)
	if err != nil {
		return err
	}
	resp := gocast.ToInt(rawResp)
	if resp != 1 {
		return fmt.Errorf("resp: %d", resp)
	}
	return nil
}
