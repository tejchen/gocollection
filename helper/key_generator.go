package helper

func StringKeyGenerator() func(item interface{}) string {
	return func(item interface{}) string {
		return item.(string)
	}
}
