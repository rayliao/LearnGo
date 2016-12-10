package modals

import (
	"encoding/json"
	"fmt"

	"github.com/bitly/go-simplejson"
)

// Server struct
type Server struct {
	ServerName string `json:"serverName"`
	ServerIP   string `json:"serverIP"`
}

// Serverslice struct
type Serverslice struct {
	Servers []Server `json:"servers"`
}

// ParseJSON func 解析json
func ParseJSON() {
	var s Serverslice
	str := `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`

	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)

	// 使用simplejson解析未知结构的json
	js, err := simplejson.NewJson([]byte(`{
        "test": {
        "array": [1, "2", 3],
        "int": 10,
        "float": 5.150,
        "bignum": 9223372036854775807,
        "string": "simplejson",
        "bool": true
    }
    }`))

	if err != nil {
		fmt.Println(err)
	}

	arr, _ := js.Get("test").Get("array").Array()
	i, _ := js.Get("test").Get("int").Int()
	ms := js.Get("test").Get("string").MustString()

	fmt.Printf("arr %v", arr)
	fmt.Printf("i %v", i)
	fmt.Printf("ms %v", ms)
}

// GenerateJSON func 生成json
func GenerateJSON() {
	var s Serverslice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "172.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "172.0.0.2"})

	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}
	fmt.Println(string(b))
}
