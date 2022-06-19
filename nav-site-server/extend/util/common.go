package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

// ParseYaml 配置文件转[]byte
func ParseYaml(filepath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// CreateMD5 create a md5 string
func CreateMD5(str string, long bool) string {
	h := md5.New()
	h.Write([]byte(str))
	pwd := hex.EncodeToString(h.Sum(nil))
	if long {
		return pwd
	}
	return pwd[8:24]
}

// CreateRandom @desc: CreateRandom crate a random string
// @params l: length of the string
// @paramss: is include special charters
// @paramst：is include timestamp s
// @return a random string
func CreateRandom(l int, s, t bool) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if s {
		str += "-+=@#$%^*!."
	}
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	if t {
		return string(result) + fmt.Sprintf("%d", time.Now().Unix())
	}
	return string(result)
}

func CreateRandomUNID(l int) string {
	str := "123456789123456789123456789123456789"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

// CreateUniqueID create a unique id
func CreateUniqueID(prefix string) string {
	str := CreateMD5(prefix+CreateRandom(32, true, true), true)
	return CreateMD5(str, true)
}

// UintInArray 数组中是否存在该该值
func UintInArray(id uint, arr []uint) bool {
	for _, val := range arr {
		if id == val {
			return true
		}
	}
	return false
}

// StringInArray 数组中是否存在该该值
func StringInArray(id string, arr []string) bool {
	for _, val := range arr {
		if id == val {
			return true
		}
	}
	return false
}
