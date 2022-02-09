// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/EldersJavas/SaponoAI/app/tool"
	"io/ioutil"
	"path/filepath"
)

// GetFileFromDir 通过文件扩展名筛选需要的文件
func GetFileFromDir(p string, t string) (n []string, err error) {
	filelist, err := ioutil.ReadDir(p)
	if err != nil {
		return nil, err
	}
	for _, file := range filelist {
		if filepath.Ext(file.Name()) == t {
			n = append(n, tool.PathAddGang(p)+file.Name())
		}
	}
	return n, nil
}
