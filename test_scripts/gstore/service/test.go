package service

import (
	"bufio"
	"encoding/json"
	"fmt"
	"gstore/model"
	"io"
	"os"
)

// 测试生成的sparql是否能够查询到结果
func Test() {
	// 初始化请求数据
	// url := "http://10.112.41.37:12300/gstore"
	url := "http://localhost:12300/gstore"
	database := "CCKS"

	// 打开文件
	// path := "./generated_predictions.txt"
	path := "D:\\goprojects\\src\\gstore\\txt\\generated_predictions.txt"
	fileHandle, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer fileHandle.Close()
	reader := bufio.NewReader(fileHandle)

	// 可以查到结果的语句
	// filePath := "./my_pred.txt"
	filePath := "D:\\goprojects\\src\\gstore\\txt\\my_pred.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)

	// 无法查到结果的语句
	// filePath1 := "./sparql_not_test.txt"
	filePath1 := "D:\\goprojects\\src\\gstore\\txt\\sparql_not_test.txt"
	file1, err := os.OpenFile(filePath1, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	write1 := bufio.NewWriter(file1)

	// 查不到结果, 可以查询到结果, 总共查询数量
	no, ok, all := 0, 0, 0
	// 按行读取文件
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		// 发送请求
		test := &model.TestJson{}
		json.Unmarshal(line, &test)
		sparql := string(test.Predict)
		gstoreRequest := model.GstoreRequest{database, sparql}
		payload, _ := json.Marshal(gstoreRequest)
		data := doHTTP(url, payload)

		// 未查询到结果
		if data.AnsNum == 0 {
			write1.WriteString(sparql)
			write1.WriteString("\n")
			no++
		} else {
			ok++
		}

		// 将数据写入文件
		vars := data.Head.Vars
		myans := getValues(&data.Results, &vars)
		write.WriteString(myans)
		write.WriteString("\n")
		all++
	}

	defer fmt.Println("ok: ", ok)
	defer fmt.Println("no: ", no)
	defer fmt.Println("all queries: ", all)
	defer write.Flush()
	defer write1.Flush()
}
