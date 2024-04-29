package genfunc

// GetGenTableNameTemp get gen table template str
func GetGenTableNameTemp() string {
	return genTnf
}

// GetGenTableNameTemp get gen table template str
func GetGenTableExlTemp() string {
	return genExl
}

// GetGenTableNameTemp get gen table template str
func GetGenTableDataTemp() string {
	return genGetData
}

// GetGenTableNameTemp get gen table template str
func GetGenTableGetExlTemp() string {
	return genGetExl
}

// GetGenColumnNameTemp get gen column template str
func GetGenColumnNameTemp() string {
	return genColumn
}

// GetGenColumnNameTemp get gen column template str
func GetGenColumnCommentTemp() string {
	return genColumnComment
}

// GetGenBaseTemp get gen base template str
func GetGenBaseTemp() string {
	return genBase
}

// GetGenLogicTemp get gen logic template str
func GetGenLogicTemp() string {
	return genlogic
}

// GetGenPreloadTemp get gen preload template str
func GetGenPreloadTemp(multi bool) string {
	if multi {
		return genPreloadMulti
	}
	return genPreload
}

func GetGenPageTemp() string {
	return genPage
}
