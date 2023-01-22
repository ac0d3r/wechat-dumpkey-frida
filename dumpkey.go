package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/frida/frida-go/frida"
)

func main() {
	key, err := Key()
	if err != nil || key == "" {
		fmt.Println("not found wechat message key")
		return
	}
	fmt.Println(key)
}

var js = `
var key = ObjC.chooseSync(ObjC.classes.DBEncryptInfo)[0];
var data = key['- m_dbEncryptKey']();
console.log(hexdump(data.bytes(), { offset: 0, length: data.length(), header: false, ansi: false }));
`

type Log struct {
	Type    string `json:"type,omitempty"`
	Level   string `json:"level,omitempty"`
	Payload string `json:"payload,omitempty"`
}

func Key() (string, error) {
	var key string
	c := make(chan struct{}, 1)

	mgr := frida.NewDeviceManager()
	dev, err := mgr.LocalDevice()
	if err != nil {
		return "", err
	}

	session, err := dev.Attach("微信", nil)
	if err != nil {
		return "", err
	}

	script, err := session.CreateScript(js)
	if err != nil {
		return "", err
	}

	script.On("message", func(msg string) {
		defer func() {
			c <- struct{}{}
		}()

		m := Log{}
		err := json.Unmarshal([]byte(msg), &m)
		if err == nil {
			key = parse(m.Payload)
		}
	})

	if err := script.Load(); err != nil {
		return "", err
	}

	<-c
	return key, nil
}

func parse(payload string) string {
	var r strings.Builder
	r.WriteString("0x")

	data := strings.Split(payload, "\n")
	if len(data) == 0 {
		return ""
	}
	for i := range data {
		v := strings.Split(data[i], "  ")
		if len(v) != 3 {
			continue
		}
		key := strings.ReplaceAll(v[1], " ", "")
		r.WriteString(key)
	}
	if r.Len() == 2 {
		return ""
	}
	return r.String()
}
