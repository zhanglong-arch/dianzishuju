package util

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/ioutil"
)

/**
 * 对一个字符串进行MD5哈希计算，并返回hash值
 */
func MD5HashString (data string) string{
	hashMd5 := md5.New()
	hashMd5.Write([]byte(data))
	passwordBytes := hashMd5.Sum(nil)
	return hex.EncodeToString(passwordBytes)
}

/**
 * 对一个io操作的reader （通常为文件）进行数据读取，并计算
 */
func MD5HashFile (reader io.Reader) (string,error){
	bytes, err := ioutil.ReadAll(reader)
	if err != nil{
		return "", err
	}
	md5Hash := md5.New()
	md5Hash.Write(bytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

func SHA256Hash(data []byte) ([]byte) {
	//1、对block字段进行拼接

	//2、对拼接后的数据进行sha256
	sha256Hash := sha256.New()
	sha256Hash.Write([]byte(""))
	return sha256Hash.Sum(nil)
}

