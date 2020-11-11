package crypt

import (
	"crypto/cipher"
	"crypto/des"
	"cryptoForGo/utils"
	"fmt"
)



//3DES加密
func TripleDESEnCrypt(data, key []byte) ([]byte, error) {
	//分组
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	fmt.Println("分组的长度：", block.BlockSize())
	//对明文分组进行填充
	originData := utils.PKCS5Padding(data, block.BlockSize())
	//实例化加密模式
	mode := cipher.NewCBCEncrypter(block, key[:8])
	//创建一个容器装加密后的数据
	dst := make([]byte, len(originData))
	//对明文进行加密
	mode.CryptBlocks(dst, originData)
	return dst, nil
}

//3DES解密
func TripleDESDeCrypt(data, key []byte) ([]byte, error) {
	//分组
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	//实例化一个加密模式
	mode := cipher.NewCBCDecrypter(block, key[:8])
	//创建一个容器装解密后的数据
	dst := make([]byte, len(data))
	//解密
	mode.CryptBlocks(dst, data)
	//去除填充的数据
	dst=utils.ClearPKCS5Padding(dst)
	return dst, nil
}
