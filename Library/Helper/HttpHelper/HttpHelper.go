package HttpHelper

import (
	"crypto/md5"
	"fmt"

	// "logkar-backend/Constant"
	"net/textproto"
	"path"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func StaticPath(Staticpath string) string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Join(path.Dir(filename), Staticpath)
}

// func GetTemplateEmail(TemplateFile string, templateData map[string]interface{}) (data string, err error) {
// 	path := StaticPath(Constant.EmailTemplatePath)
// 	t, err := template.ParseFiles(path + TemplateFile)
// 	if err != nil {
// 		log.Println(err.Error())
// 		return
// 	}
// 	var tpl bytes.Buffer
// 	err = t.Execute(&tpl, templateData)
// 	return tpl.String(), err
// }

func GenerateNameFile(params textproto.MIMEHeader) string {
	checkHeader := params.Get("Content-Type")
	currentTime := time.Now().Unix()
	convertionToString := strconv.Itoa(int(currentTime))
	getFormat := strings.TrimPrefix(checkHeader, "image/")

	returnGenerateNameFile := convertionToString + "." + getFormat
	return returnGenerateNameFile
}

func GenerateHashPass(secret string) string {
	data := []byte(secret)
	hash := fmt.Sprintf("%x", md5.Sum(data))
	return hash
}

func HashUserId() string {
	currentTime := time.Now().Unix()
	conversionToString := strconv.Itoa(int(currentTime))
	data := []byte(conversionToString)
	hash := fmt.Sprintf("%x", md5.Sum(data))
	return hash
}

// func HandlerParameterURL(param string) (value string, err error) {
// 	newValue := strings.ReplaceAll(param, Constant.SpecialCharPlus, Constant.UrlSpecialCharReplacement)
// 	value, err = url.QueryUnescape(newValue)
// 	value = strings.ReplaceAll(value, Constant.UrlSpecialCharReplacement, Constant.SpecialCharPlus)
// 	return
// }
