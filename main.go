package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func calcFileMD5(filename string) (string, error) {
	f, err := os.Open(filename) //打开文件
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	defer f.Close()
	md5Handle := md5.New()         //创建 md5 句柄
	_, err = io.Copy(md5Handle, f) //将文件内容拷贝到 md5 句柄中
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	md := md5Handle.Sum(nil)        //计算 MD5 值，返回 []byte
	md5str := fmt.Sprintf("%x", md) //将 []byte 转为 string
	return md5str, nil
}

func calcFileSha1(filename string) (string, error) {
	f, err := os.Open(filename) //打开文件
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	defer f.Close()
	sha1Handle := sha1.New()        //创建 sha1 句柄
	_, err = io.Copy(sha1Handle, f) //将文件内容拷贝到 sha1 句柄中
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	sha := sha1Handle.Sum(nil)        //计算 SHA1 值，返回 []byte
	sha1str := fmt.Sprintf("%x", sha) //将 []byte 转为 string
	return sha1str, nil
}

func calcFileSha256(filename string) (string, error) {
	f, err := os.Open(filename) //打开文件
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	defer f.Close()
	sha256Handle := sha256.New()      //创建 sha256 句柄
	_, err = io.Copy(sha256Handle, f) //将文件内容拷贝到 sha256 句柄中
	if nil != err {
		fmt.Println(err)
		return "", err
	}
	sha := sha256Handle.Sum(nil)        //计算 SHA256 值，返回 []byte
	sha256str := fmt.Sprintf("%x", sha) //将 []byte 转为 string
	return sha256str, nil
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
		md5str, err := calcFileMD5(os.Args[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("MD5:%s\n", md5str)
		sha1str, err := calcFileSha1(os.Args[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("SHA1:%s\n", sha1str)
		sha256str, err := calcFileSha256(os.Args[i])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("SHA256:%s\n", sha256str)
	}
	fmt.Println("Press any key to exit...")
	var input string
	fmt.Scanln(&input)
}
