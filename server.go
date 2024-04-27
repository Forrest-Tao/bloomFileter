package bloomFileter

import (
	"github.com/demdxx/gocast"
)

type BloomServer struct {
	//m->len(bitMap) k->(the number of hash)
	m, k      int32
	encryptor *Encryptor
	client    *RedisClient
}

func NewBloomServer(m, k int32, e *Encryptor, client *RedisClient) *BloomServer {
	return &BloomServer{
		m:         m,
		k:         k,
		encryptor: e,
		client:    client,
	}
}

func (s *BloomServer) getEncrypted(val string) []int32 {
	encrypteds := make([]int32, 0, s.k)
	origin := val
	//前一次encrypt后的结果作为这次的 input，再次encrypt，循环（用这种方式取代 多个hash func）
	for i := 0; int32(i) < s.k; i++ {
		encypted := s.encryptor.Encrypt(origin)
		encrypteds = append(encrypteds, encypted)
		origin = gocast.ToString(encypted)
	}
	return encrypteds
}
