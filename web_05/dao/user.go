package dao

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var database = map[string]string{}

func Initdatabase() {
	file, err := os.Open("info.data")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var (
		username string
		password string
	)
	for i := 0; ; i++ {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		if len(line) != 0 && i%2 == 0 {
			username = line
			s_username := []byte(username)
			s_username = s_username[:len(s_username)-1]
			username = string(s_username)
		} else if len(line) != 0 && i%2 == 1 {
			password = line
			s_password := []byte(password)
			s_password = s_password[:len(s_password)-1]
			password = string(s_password)
			database[username] = password
		}
	}
	fmt.Println(database["immhj"])
}

func AddUSer(username, password string) {
	file1, err1 := os.OpenFile("info.data", os.O_APPEND|os.O_WRONLY, 0666)
	if err1 != nil {
		fmt.Println("open file failed, err:", err1)
		return
	}
	defer file1.Close()
	writer := bufio.NewWriter(file1)
	writer.WriteString("\n" + username + "\n") //将数据先写入缓存
	writer.WriteString(password + "\n")
	writer.Flush() //将缓存中的内容写入文件
	database[username] = password
}

func SelectUser(username string) bool {
	if database[username] == "" {
		return false
	}
	return true
}

func SelectPasswordFromUsername(username string) string {
	return database[username]
}

func Updatepassword(username, password string) {
	var o int
	file1, err1 := os.OpenFile("info.data", os.O_CREATE|os.O_RDWR, 0666)
	if err1 != nil {
		fmt.Println("open file failed, err:", err1)
		return
	}
	defer file1.Close()
	reader := bufio.NewReader(file1)
	for i := 0; ; i++ {
		line, err := reader.ReadString('\n') //注意是字符
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		if line == username {
			o = i
		}
		if i == o+1 {
			str := []byte(line)
			str = str[:len(str)-1]
			line = string(str)
		}
	}
	database[username] = password
}
