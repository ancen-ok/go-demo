package utils

// Types 类型接口
type Types interface{}

func In[T comparable](target T, arr []T) (bool, int) {
	for i, v := range arr {
		if v == target {
			return true, i
		}
	}
	return false, -1
}

func InStr(val string, list []string) (bool, error) {
	for _, item := range list {
		if item == val {
			return true, nil
		}
	}
	return false, nil
}
