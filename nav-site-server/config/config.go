package conf

import (
	"bufio"
	_ "embed"
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"nav-site-server/extend/util"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"
)

//go:embed config.yaml
var DefaultConfigFile []byte

//go:embed nav-site.service
var NavSiteServerSystemCtl []byte

var HasLogo bool = false

type Program struct {
	// 程序所在目录
	ProgramDir string
	// 配置及数据目录
	ConfDir string
}

// Config server config
type Config struct {
	Server     Server  `yaml:"server"`
	Site       Site    `yaml:"site"`
	Store      Store   `yaml:"store"`
	GroupStore Store   `yaml:"GroupStore"`
	Static     Static  `yaml:"static"`
	Account    Account `yaml:"account"`
	Program    Program
}

// App application config
type Server struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
}

type Site struct {
	Title     string `yaml:"title"`
	Logo      string `yaml:"logo"`
	HasLogo   bool
	Url       string `yaml:"url"`
	Copyright string `yaml:"copyright"`
}

// Store store config
type Store struct {
	Drive      string `yaml:"drive"`
	Type       string `yaml:"type"`
	Path       string `yaml:"path"`
	Suffix     string `yaml:"suffix"`
	BackupsDir string `yaml:"backupsDir"`
	BackupsMax int    `yaml:"backupsMax"`
	FileSync   *util.FileSync
}

type Static struct {
	Root   string `yaml:"root"`
	Upload Upload `yaml:"upload"`
}

type Upload struct {
	Path    string `yaml:"path"`
	Maxsize int64  `yaml:"maxsize"`
	BaseUrl string `yaml:"baseUrl"`
}

type Account struct {
	Secret              string        `yaml:"secret"`
	Admin               string        `yaml:"admin"`
	Members             []User        `yaml:"members"`
	CookieExpireSeconds time.Duration `yaml:"cookieExpireSeconds"`
}

type User struct {
	Name     string `yaml:"name"`
	Rule     string `yaml:"rule"`
	Password string `yaml:"password"`
}

var App Config

const (
	StoreDriveFile = "file"
	// RuleAdd 添加权限
	RuleAdd = "add"
	// RuleEdit 编辑权限
	RuleEdit = "edit"
	// RuleDelete 删除权限
	RuleDelete = "delete"
)

func InitConfig(confDir string) (*Config, error) {
	release := util.IsRelease()
	var programPath = ""
	if !release {
		fmt.Println("非正式环境运行中..." + strconv.FormatBool(!release))
	} else {
		path, _ := util.GetExecPath()
		fmt.Println("程序所在目录:" + path)
		programPath = path
	}

	var conf Config
	conf.Program.ConfDir = confDir
	conf.Program.ProgramDir = programPath
	confDataDir := conf.GetConfDataDir()
	fmt.Println("数据配置目录:" + confDataDir)

	confFileFull := path.Join(confDataDir, "conf/config.yaml")

	err := createConfAuto(confFileFull)
	if err != nil {
		return nil, err
	}
	config, err := util.ParseYaml(confFileFull)
	if err != nil {
		errMsg := errors.New("parse config.yaml file error : " + err.Error())
		return nil, errMsg
	}

	if err := yaml.Unmarshal(config, &conf); err != nil {
		errMsg := errors.New("parse config []byte to struct error ：" + err.Error())
		return nil, errMsg
	}

	conf.initAccount()

	if err := conf.initStoreDrive(); err != nil {
		return &conf, nil
	}
	return &conf, nil
}

// 自动创建配置文件
func createConfAuto(confFile string) error {
	confExists, _ := util.FileExists(confFile)

	if !confExists {
		baseDir := filepath.Dir(confFile)
		os.MkdirAll(baseDir, os.ModePerm)
		file, err := os.OpenFile(confFile, os.O_WRONLY|os.O_CREATE, 0644)
		defer file.Close()

		if err != nil {
			fmt.Println("create conf err")
			return err
		}

		//file.WriteAt(DefaultConfigFile, 0)
		writer := bufio.NewWriter(file)
		writer.Write(DefaultConfigFile)
		writer.Flush()

	}
	// 写出linux-centos8-9配置
	linuxServerConf := "conf/nav-site.service"
	linuxServerFile, _ := os.OpenFile(linuxServerConf, os.O_WRONLY|os.O_CREATE, 0644)
	defer linuxServerFile.Close()

	writer := bufio.NewWriter(linuxServerFile)
	writer.Write(NavSiteServerSystemCtl)
	writer.Flush()

	return nil
}

func (c *Config) GetConfDataDir() string {
	var confDir string
	if c.Program.ConfDir != "" {
		confDir = c.Program.ConfDir
	} else {
		confDir = c.Program.ProgramDir
	}
	return confDir
}

// initStoreDrive 初始化存储驱动
func (c *Config) initStoreDrive() (err error) {
	switch c.Store.Drive {
	case StoreDriveFile:
		fallthrough
	default:
		confDataDir := c.GetConfDataDir()

		fileSync := util.FileSync{}
		webDataPath := path.Join(confDataDir, c.Store.Path)
		fileSync.FilePath = webDataPath
		if err = fileSync.InitStoreFile(webDataPath, 0755); err != nil {
			return err
		}
		c.Store.FileSync = &fileSync

		fileGroupSync := util.FileSync{}
		groupPath := path.Join(confDataDir, c.GroupStore.Path)
		fileGroupSync.FilePath = groupPath
		if err = fileGroupSync.InitStoreFile(groupPath, 0755); err != nil {
			return err
		}
		c.GroupStore.FileSync = &fileGroupSync

		if c.Store.BackupsMax <= 7 {
			c.Store.BackupsMax = 7
		}
	}
	return nil
}

func (c *Config) initAccount() {
	if c.Account.CookieExpireSeconds <= 7200 {
		c.Account.CookieExpireSeconds = 7200
	}
}
