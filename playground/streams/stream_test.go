package streams

import (
	"fmt"
	"gitlab.myteksi.net/snd/ironbank/common/util"
	"io/ioutil"
	"strings"
	"testing"
)

func TestGetRemain(t *testing.T) {
	exist := util.ReadFile("exsiting")
	existStreams := strings.Fields(exist)
	
	existMap := map[string]bool{}
	
	for _, s := range existStreams {
		existMap[s] = true
	}
	
	full := util.ReadFile("full")
	fullStreams := strings.Fields(full)
	
	remain := []string{}
	
	for _, s := range fullStreams {
		if !existMap[s] {
			remain = append(remain, s)
		}
	}
	for _, s :=range remain {
		fmt.Printf("%v\n",s)
	}

	util.WriteFile("remain", strings.Join(remain, "\n"))
}

func TestProcessRemain(t *testing.T) {
	path := "/Users/yuguang.xiao/go/src/gitlab.myteksi.net/gophers/go/streams"
	ownerMap := map[string]string{}
	GetOwnerRecursive(path, "", ownerMap)
	
	content := util.ReadFile("remain")
	remain := strings.Fields(content)
	
	tfs := []string{}
	
	for _, s := range remain {
		tfs = append(tfs, ownerMap[s])
	}

	util.WriteFile("tfs", strings.Join(tfs, "\n"))
}

func GetOwnerRecursive(dir string, parent string, owners map[string]string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("err getting stream folders for dir %v:%v\n", dir, err)
		return
	}
	for _, file := range files {
		//filepath := dir + "/" + file.Name()
		if file.IsDir() {
			if file.Name() == "apis" {
				continue
			}

			if strings.HasSuffix(file.Name(), "stream") {
				owners[file.Name()] = strings.Trim(parent,"/")
			} else {
				GetOwnerRecursive(dir + "/" + file.Name(), parent+"/"+file.Name(), owners)
			}
		}
	}
	return
}