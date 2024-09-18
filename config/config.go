package config

func GetSetting(key string) (interface{}, bool) {
	configMap := map[string]interface{}{
		"http_port": "8080",
	}
	if value, ok := configMap[key]; ok {
		return value, true
	} else {
		return nil, false
	}
}

func GetSettingString(key string) string {
	value, ok := GetSetting(key)
	if !ok {
		return ""
	}
	strValue, ok := value.(string)
	if !ok {
		return ""
	}
	return strValue
}

func GetSettingInt(key string) int {
	value, ok := GetSetting(key)
	if !ok {
		return 0
	}
	intValue, ok := value.(int)
	if !ok {
		return 0
	}
	return intValue
}
