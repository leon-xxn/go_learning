package leaky_bucket

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sync"
	"testing"
	"time"
	"unsafe"
)

type RateLimiterManager struct {
	limiterMap map[string]*UserRateLimiter
	mu         sync.Mutex
}

func NewRateLimiterManager() *RateLimiterManager {
	return &RateLimiterManager{
		limiterMap: make(map[string]*UserRateLimiter),
	}
}

func (r *RateLimiterManager) GetLimiter(userid string) *UserRateLimiter {
	r.mu.Lock()
	defer r.mu.Unlock()
	if limiter, ok := r.limiterMap[userid]; ok {
		return limiter
	}
	newLimiter := NewUserRateLimiter(10, 10)
	r.limiterMap[userid] = newLimiter
	return newLimiter
}

func TestManager(t *testing.T) {
	manager := NewRateLimiterManager()

	// 模拟多个用户的请求
	users := []string{"user1", "user2", "user1", "user3", "user2"}

	for i, userID := range users {
		fmt.Printf("User %s makes request %d\n", userID, i+1)
		limiter := manager.GetLimiter(userID)

		// 检查用户的请求是否被允许
		if limiter.AllowRequest() {
			fmt.Printf("Request %d by %s processed\n", i+1, userID)
		} else {
			fmt.Printf("Request %d by %s denied\n", i+1, userID)
		}
		time.Sleep(4 * time.Second) // 每 4 秒发送一次请求
	}
}

func StringToBytes(str string) []byte {
	strHeader := (*[2]uintptr)(unsafe.Pointer(&str))
	byteSliceHeader := [3]uintptr{strHeader[0], strHeader[1], strHeader[1]}
	return *(*[]byte)(unsafe.Pointer(&byteSliceHeader))

}

//

func reverseString(s string) string {
	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

type MyType struct {
	A int32
	B float64
}

func MyTypeSliceToByteSlice(data []MyType) ([]byte, error) {
	buf := new(bytes.Buffer)
	for _, value := range data {
		if err := binary.Write(buf, binary.LittleEndian, value); err != nil {
			return nil, err
		}
	}
	return buf.Bytes(), nil
}

func ByteSliceToMyTypeSlice(data []byte) ([]MyType, error) {
	buf := bytes.NewReader(data)
	var result []MyType
	for buf.Len() > 0 {
		var value MyType
		if err := binary.Write(buf, binary.LittleEndian, &value); err != nil {
			return nil, err
		}
	}
	return result, nil

}

//new 是内建函数，返回一个指向零值的指针，new(int)返回指向 0的*int
//make 是内建函数，返回一个初始化并且准备使用的值而不是指针，并且初始化数据结构
