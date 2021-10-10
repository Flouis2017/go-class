package GoIO

import (
	"fmt"
	"os"
	"reflect"
)

func ReadTest(filePath string) string {
	bytes, _ := os.ReadFile(filePath)
	return string(bytes)
}

func WriteTest(fp string) {
	var content = []byte("穷人十字街头耍十把钢钩钩不着至亲骨肉\n富人在深山老林刀枪棍棒打不散无义宾朋")
	fmt.Println(reflect.TypeOf(content))
	os.WriteFile(fp, content, 0777)
}

