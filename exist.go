package bloomFileter

import (
	"context"
	"github.com/demdxx/gocast"
)

func (s *BloomServer) Exist(ctx context.Context, key, val string) (bool, error) {
	keysAndArgs := make([]interface{}, 0, 2+s.k)
	//其中key为 bitmap的 key，val是想要查询是否存在的值
	keysAndArgs = append(keysAndArgs, key)
	for _, encrypted := range s.getEncrypted(val) {
		keysAndArgs = append(keysAndArgs, encrypted)
	}
	//使用lua脚本，保证原子性
	rawResp, err := s.client.Eval(ctx, LuaBloomBatchGetBits, 1, keysAndArgs)
	if err != nil {
		return false, err
	}

	resp := gocast.ToInt(rawResp)
	if resp == 1 {
		return true, nil
	}
	return false, nil
}
