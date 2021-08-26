package GoMap

/**
 * map三种声明方式如下
 */

var MapOne = map[string]string{}
var MapTwo = make(map[string]string)

var AnyMap = map[interface{}]interface{}{}

func Init() {
	MapOne["country"] = "China"
	MapOne["province"] = "Fujian"
	MapOne["city"] = "Xiamen"
	MapOne["district"] = "Jimei"

	AnyMap[1] = 123.23
	AnyMap["name"] = "hahaha"
	AnyMap["arr"] = []int{3, 3, 0, 6}
}

