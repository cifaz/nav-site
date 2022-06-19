package conf

import (
	"errors"
	"gopkg.in/yaml.v2"
	"nav-site-server/extend/util"
	"time"
)

// Config server config
type Config struct {
	Server  Server  `yaml:"server"`
	Store   Store   `yaml:"store"`
	Static  Static  `yaml:"static"`
	Account Account `yaml:"account"`
}

// App application config
type Server struct {
	Name string `yaml:"name"`
	Port string `yaml:"port"`
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
	Static string `yaml:"static"`
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

func InitConfig() (*Config, error) {
	config, err := util.ParseYaml("config/config.yaml")
	if err != nil {
		errMsg := errors.New("parse config.yaml file error : " + err.Error())
		return nil, errMsg
	}
	var c Config
	if err := yaml.Unmarshal(config, &c); err != nil {
		errMsg := errors.New("parse config []byte to struct error ：" + err.Error())
		return nil, errMsg
	}

	c.initAccount()

	if err := c.initStoreDrive(); err != nil {
		return &c, nil
	}
	return &c, nil
}

// initStoreDrive 初始化存储驱动
func (c *Config) initStoreDrive() (err error) {
	switch c.Store.Drive {
	case StoreDriveFile:
		fallthrough
	default:
		fileSync := util.FileSync{}
		fileSync.FilePath = c.Store.Path
		if err = fileSync.InitStoreFile(c.Store.Path, 0755); err != nil {
			return err
		}
		c.Store.FileSync = &fileSync
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
