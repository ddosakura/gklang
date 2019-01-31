package gklang

// Er exit with error log
func Er(e error) {
	Log(LCrash, e)
}
