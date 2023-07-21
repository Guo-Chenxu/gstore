package service

import "gstore/model"

// 解析返回值
func getValues(results *model.Results, vars *[]string) string {
	myans := ""
	bindings := results.Bindings
	for _, b := range bindings {
		for _, k := range *vars {
			v := b[k]
			uri := (v["type"] == "uri")
			str := v["value"]
			if uri {
				myans += "<" + str + ">\t"
			} else {
				myans += "\"" + str + "\"\t"
			}
		}
	}
	return myans
}
