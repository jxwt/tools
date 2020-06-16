package tools

import (
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
	"strings"
	"time"
)

func GetUuid() uuid.UUID {
	uuid, _ := uuid.NewV4()
	return uuid
}

/**
获取uuid的字符串
*/
func GetUuidString() string {
	return GetUuid().String()
}

/**
获取随机字符串
*/
func GetUuidRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

func GetRandIntn(n int, min int) int {
	rand.Seed(time.Now().UnixNano())
	if min > n {
		return min
	}
	x := rand.Intn(n)
	for index := 0; index < 100; index++ {
		if x >= min {
			break
		}
		x = rand.Intn(n)
	}
	return x
}

func GetRandNanoNumber() string {
	rand.Seed(time.Now().UnixNano())
	return fmt.Sprintf("%d", rand.Intn(9999)+1000)
}

/**
获取随机字符串
*/
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := make([]byte, 0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// 根据主键生成邀请码

func GetInviteCode(id uint) string {
	var source = "HVE8S2DZX9C7P5K3MJUAR4WYLTN6BGQ"

	id = id*121 + 1000000
	mod := uint(0)
	res := ""
	for id != 0 {
		mod = id % 31
		id = id / 31
		res += string(source[mod])
	}

	resLen := len(res)
	if resLen < 6 {
		res += "F"
		for i := 0; i < 6-resLen-1; i++ {
			rand.Seed(time.Now().UnixNano())
			res += string(source[rand.Intn(31)])
		}
	}

	if res[1] == res[2] && res[2] == res[3] && res[3] == res[4] && res[4] == res[5] {
		res += "T"
	}
	return res
}

func ParseInviteCode(code string) uint {
	res := uint(0)
	lenCode := len(code)
	var source = "HVE8S2DZX9C7P5K3MJUAR4WYLTN6BGQ"
	baseArr := []byte(source)     // 字符串进制转换为byte数组
	baseRev := make(map[byte]int) // 进制数据键值转换为map
	for k, v := range baseArr {
		baseRev[v] = k
	}

	// 查找补位字符的位置
	isPad := strings.Index(code, "F")
	if isPad != -1 {
		lenCode = isPad
	}

	r := 0
	for i := 0; i < lenCode; i++ {
		// 补充字符直接跳过
		if string(code[i]) == "F" {
			continue
		}
		index := baseRev[code[i]]
		b := uint(1)
		for j := 0; j < r; j++ {
			b *= 31
		}
		// pow 类型为 float64 , 类型转换太麻烦, 所以自己循环实现pow的功能
		//res += float64(index) * math.Pow(float64(32), float64(2))
		res += uint(index) * b
		r++
	}

	return (res - 1000000) / 121
}
