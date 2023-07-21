package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gstore/model"
	"io"
	"os"
)

// 测试查询答案是否正确
func Train() {
	// 初始化请求数据
	// url := "http://10.112.41.37:12300/gstore"
	url := "http://localhost:12300/gstore"
	database := "CCKS"

	// 打开训练集文件
	// path := "./train.txt"
	path := "D:\\goprojects\\src\\gstore\\txt\\train.txt"
	fileHandle, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()
	reader := bufio.NewReader(fileHandle)

	// 查询的结果
	// filePath := "./myans"
	filePath := "D:\\goprojects\\src\\gstore\\txt\\myans.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)

	// 无法查询到结果的语句
	// filePath1 := "./sparql_not_train.txt"
	filePath1 := "D:\\goprojects\\src\\gstore\\txt\\sparql_not_train.txt"
	file1, err := os.OpenFile(filePath1, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	write1 := bufio.NewWriter(file1)

	// 无法查询到结果, 和答案完全一样, 总共查询数量
	null, right, all := 0, 0, 0
	// 按行读取文件
	for {
		// 读取文件, 直到读取到sparql
		line := []byte{}
		for {
			line, _, err = reader.ReadLine()
			if err == io.EOF {
				goto eof
			}
			t := string(line)
			if len(t) != 0 && t[0] == 's' {
				break
			}
			line = nil
		}

		if err == io.EOF {
			goto eof
		}

		// 发送请求
		sparql := string(line)
		gstoreRequest := model.GstoreRequest{database, sparql}
		payload, _ := json.Marshal(gstoreRequest)
		data := doHTTP(url, payload)

		// 读取答案
		ansBytes, _, _ := reader.ReadLine()
		ans := string(ansBytes)

		// 未查询到结果
		if data.AnsNum == 0 {
			write1.WriteString(sparql)
			write1.WriteString(ans)
			write1.WriteString("\n")
			null++
		}

		// 将数据写入文件
		vars := data.Head.Vars
		myans := getValues(&data.Results, &vars)
		write.WriteString(myans)
		write.WriteString("\n")

		// 答案相等
		if myans == ans {
			right++
		}
		all++
	}

eof:
	defer fmt.Println("right: ", right)
	defer fmt.Println("can't find the answer: ", null)
	defer fmt.Println("all queries: ", all)
	defer write.Flush()
	defer write1.Flush()
}
