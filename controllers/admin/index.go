package admin

import (
	"github.com/astaxie/beego"
	"goblog/models"
	"os"
	"runtime"
	beeLogger "github.com/beego/bee/logger"
	"strings"
	"path/filepath"
	"log"
	"fmt"
)

type IndexController struct {
	baseController
}


func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func (this *IndexController) Index() {

	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["adminid"] = this.userid
	this.Data["adminname"] = this.username

	var str1, str2 string
	str1 = getCurrentDirectory()

	str2 = getParentDirectory(str1)
	fmt.Println(str2)

	this.TplName = this.moduleName + "/index/index.html"

	beeLogger.Log.Infof("IndexController Index %s", this.TplName)
}

func (this *IndexController) Main() {

	this.Data["hostname"], _ = os.Hostname()
	this.Data["version"] = beego.AppConfig.String("AppVer")
	this.Data["gover"] = runtime.Version()
	this.Data["os"] = runtime.GOOS
	this.Data["cpunum"] = runtime.NumCPU()
	this.Data["arch"] = runtime.GOARCH

	this.Data["postnum"], _ = new(models.Post).Query().Count()
	this.Data["tagnum"], _ = new(models.Tag).Query().Count()
	this.Data["usernum"], _ = new(models.User).Query().Count()

	beeLogger.Log.Infof("IndexController Main %s", this.TplName)
	this.display()
}
