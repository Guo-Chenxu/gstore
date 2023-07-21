package service

import (
	"bytes"
	"encoding/json"
	"gstore/model"
	"io/ioutil"
	"net/http"
)

// 发送http请求
func doHTTP(url string, payload []byte) model.Data {
	// 发送请求
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json;charset=utf-8")
	if err != nil {
		panic(err)
	}

	// 处理请求结果
	res, _ := http.DefaultClient.Do(req)
	body, _ := ioutil.ReadAll(res.Body)
	r := model.R{}
	json.Unmarshal(body, &r)

	return r.Data
}
