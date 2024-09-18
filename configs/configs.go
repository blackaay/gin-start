package configs

func GetSetting(key string) (interface{}, bool) {
	configMap := map[string]interface{}{
		"HTTP_PORT":     "8080",
		"DB_CONNECTION": "mysql",
		"DB_HOST":       "0.0.0.0",
		"DB_PORT":       "3306",
		"DB_DATABASE":   "admin_local",
		"DB_USERNAME":   "root",
		"DB_PASSWORD":   "rootdev",
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

func GetDbDsn() string {
	dbUser := GetSettingString("DB_USERNAME")
	dbHost := GetSettingString("DB_HOST")
	dbPort := GetSettingString("DB_PORT")
	dbName := GetSettingString("DB_DATABASE")
	dbPwd := GetSettingString("DB_PASSWORD")
	dsn := dbUser + ":" + dbPwd + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	return dsn
}
