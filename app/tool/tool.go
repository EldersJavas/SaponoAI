package tool

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

//获取应用根目录
func GetAppRootDir() string {
	if rootDir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		return ""
	} else {
		return WinDir(rootDir)
	}
}

//验证 s 是否存在 slice 中
func InSliceString(s string, slices []string) bool {
	for _, v := range slices {
		if v == s {
			return true
		}
	}
	return false
}

//Windows下Dir路径统一Linux格式
func WinDir(dir string) string {
	return strings.Replace(dir, "\\", "/", -1)
}

//自动判断系统并返回使用的格式
func DirFormat() string {
	if os.Getenv("GOOS") == "windows" {
		return "\\"
	} else {
		return "/"
	}
}

//获取文件名称（不带后缀）
func GetFileBaseName(filepath string) string {
	basefile := path.Base(filepath)
	ext := path.Ext(filepath)

	return strings.Replace(basefile, ext, "", 1)
}

//检验目录是否存在
func DirExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

//创建目录
func CreateDir(path string, all bool) error {
	var err error
	if all {
		err = os.Mkdir(path, os.ModePerm)
	} else {
		err = os.MkdirAll(path, os.ModePerm)
	}
	if err != nil {
		return err
	}
	return nil
}

//获取随机字符串
func GetRandomCodeString(len int) string {
	seed := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seedArr := strings.Split(seed, "")

	result := []string{}
	index := 0
	for index < len {
		s := GetIntRandomNumber(0, 61)
		result = append(result, seedArr[s])

		index++
	}

	return strings.Join(result, "")
}

//设置随机种子
func SetRandomSeed() {
	rand.Seed(time.Now().Unix()) //设置随机种子
}

//获取某范围的随机整数
func GetIntRandomNumber(min int64, max int64) int64 {
	return rand.Int63n(max-min) + min
}

func RepeatStr(str string, s string, length int, before bool) string {
	ln := len(str)

	if ln >= length {
		return str
	}

	if before {
		return strings.Repeat(s, (length-ln)) + str
	} else {
		return str + strings.Repeat(s, (length-ln))
	}
}

//校验文件是否存在
func VaildFile(file string) bool {
	_, err := os.Stat(file) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//文本md5
func Md5String(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetNowDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetTimestampDateTimeString(timeStamp int64) string {
	//返回time对象
	t := time.Unix(timeStamp, 0)
	//返回string
	return t.Format("2006-01-02 15:04:05")
}
