package utils

import "bytes"

//使用PKCS5对明文进行填充
func PKCS5Padding(data []byte,blocksize int)[]byte{
	//计算填充的个数
	size:=blocksize-len(data)%blocksize
	//准备填充的内容
	padtext:=bytes.Repeat([]byte{byte(size)},size)
	//填充并返回
	return append(data,padtext...)
}

//使用Zeros对明文进行填充
func ZerosPadding(data []byte,blocksize int)[]byte{
	//计算填充的个数
	size:=blocksize-len(data)%blocksize
	//准备填充的内容
	padtext:=bytes.Repeat([]byte{byte(0)},size)
	//填充并返回
	return append(data,padtext...)
}

//去除PKCS5尾部填充的内容
func ClearPKCS5Padding(data []byte)[]byte{
	lenth:=len(data)
	clear:=int(data[lenth-1])
	return data[:(lenth-clear)]
}

//去除Zeros尾部填充的数据
func ClearZerosPadding(data []byte,blockSize int)[]byte{
	size:=blockSize-len(data)%blockSize
	return data[:len(data)-size]
}