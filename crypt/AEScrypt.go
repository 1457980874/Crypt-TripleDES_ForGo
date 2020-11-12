package crypt

import (
	"crypto/cipher"
	"crypto/des"
	"cryptoForGo/utils"
)

//AES加密
func AESEnCrypter(data, key []byte)([]byte, error){
	//分组
	block,err:=des.NewCipher(key)
	if err != nil {
		return nil,err
	}
	//对明文进行填充
	padtext:=utils.PKCS5Padding(data,block.BlockSize())
	//实例化一个加密模式
	mode:=cipher.NewCBCEncrypter(block,padtext)
	//创建一个容器装加密后的密文
	dst:=make([]byte,(len(padtext)))
	//加密
	mode.CryptBlocks(dst,padtext)
	return dst,nil
}

//AES解密
func AESDeCrypter(data,key []byte)([]byte,error){
	//分组
	block,err:=des.NewCipher(key)
	if err != nil {
		return nil,err
	}
	//实例化一个加密模式
	mode:=cipher.NewCBCDecrypter(block,key)
	//创建一个装解密后的明文的容器
	dst:=make([]byte,len(data))
	//解密
	mode.CryptBlocks(dst,data)
	//去除解密后的明文尾部填充的数据
	dst=utils.ClearPKCS5Padding(dst)
	return dst,nil
}
