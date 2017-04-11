package main

import (
	"flag"
	"fmt"
	"github.com/fiam/gounidecode/unidecode"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var oldDirName, url string
	flag.StringVar(&oldDirName, "d", "", "dir")
	flag.StringVar(&url, "u", "", "url")
	flag.Parse()
	if oldDirName == "" {
		log.Fatal("Set flag dir please")
	}
	if url == "" {
		log.Fatal("Set flag url please")
	}
	newDirName := fmt.Sprintf("%s_new/", oldDirName)
	fileList, _ := ioutil.ReadDir(oldDirName)
	html := "\n"
	os.Mkdir(newDirName, 0777)
	for _, file := range fileList {
		if file.Name() != ".DS_Store" {
			fileName := file.Name()
			clearName := clearName(fileName)
			html += "<a href=\"" + url + "/" + clearName + "\" target=\"_blank\">" + fileName + "</a><br />\n"
			err := os.Link(filepath.Join(oldDirName, fileName), newDirName+clearName)
			if err != nil {
				log.Fatal(err.Error())
			}
		}
	}

	log.Printf("%+v", html)
}

func clearName(name string) string {
	tmpName := unidecode.Unidecode(name)
	withoutSpace := strings.Replace(tmpName, " ", "_", 100)
	withoutComa := strings.Replace(withoutSpace, ",", "", 100)
	return withoutComa
}
