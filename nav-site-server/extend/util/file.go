package util

import (
	"encoding/json"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"
)

type FileUtil struct {
}

// CreateFolderIfNotExist 文件夹不不存在则创建
func (f *FileUtil) CreateFolderIfNotExist(dir string, mode os.FileMode) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		// 创建文件夹
		if err := os.MkdirAll(dir, mode); err != nil {
			return err
		}
	}
	return nil
}

// CreateFileIfNotExist 文件不存在则创建后打开
func (f *FileUtil) CreateFileIfNotExist(path string, mode os.FileMode) (*os.File, error) {
	dir := filepath.Dir(path)
	if err := f.CreateFolderIfNotExist(dir, mode); err != nil {
		return nil, err
	}
	_, err := os.Stat(path)
	if err != nil {
		// 文件不存在
		file, err := os.Create(path)
		if err != nil {
			return nil, err
		}
		return file, nil
	}
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE, mode)
}

// RemoveFileIfExist 文件存在则删除
func (f *FileUtil) RemoveFileIfExist(path string) error {
	_, err := os.Stat(path)
	if os.IsExist(err) {
		return os.Remove(path)
	}
	return nil
}

type FileSync struct {
	File         *os.File
	FilePath     string
	FileFullPath string
	Type         string
	FileInfo     os.FileInfo
	rw           sync.RWMutex //读写锁
}

// InitStoreFile 初始化文件存储
func (f *FileSync) InitStoreFile(path string, mode os.FileMode) error {
	path, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	f.FileFullPath = path
	fileUtil := FileUtil{}
	f.File, err = fileUtil.CreateFileIfNotExist(path, mode)
	if err != nil {
		return err
	}
	info, err := f.File.Stat()
	if err != nil {
		return err
	}
	f.FileInfo = info
	return nil
}

// CloseStoreFile 退出程序时关闭文件资源
func (f *FileSync) CloseStoreFile() error {
	if f.File != nil {
		return f.File.Close()
	}
	return nil
}

// Read 文件安全读
func (f *FileSync) ReadJSON() ([]byte, error) {
	//添加读写锁
	f.rw.Lock()
	defer func() {
		// 解锁
		f.rw.Unlock()
	}()
	if f.FileInfo != nil && f.FileInfo.Size() == 0 {
		return nil, nil
	}
	var content interface{}
	_, err := f.File.Seek(0, 0)
	if err != nil {
		return nil, err
	}
	err = json.NewDecoder(f.File).Decode(&content)
	if err != nil {
		return nil, err
	}
	return json.Marshal(content)
}

// Write 文件安全写
func (f *FileSync) CoverJSON(content []byte) (err error) {
	//添加读写锁
	f.rw.Lock()
	defer func() {
		// 解锁
		f.rw.Unlock()
	}()
	err = f.File.Truncate(0)
	if err != nil {
		return err
	}
	f.File.Sync()
	_, err = f.File.WriteAt(content, 0)
	if err != nil {
		return err
	}
	info, err := f.File.Stat()
	if err != nil {
		return err
	}
	f.FileInfo = info
	return nil
}

// BackupsFileIfExist 文件备份
func (f *FileSync) Backups(backupsDir string) (err error) {
	if f.File == nil {
		return errors.New("file target is nil")
	}
	// 备份文件名称
	newFileName := backupsDir + string(os.PathSeparator) + time.Now().Format("2006.01.02.15.04.05.999.") + f.Type
	newFileName, err = filepath.Abs(newFileName)
	if err != nil {
		return err
	}
	content, err := f.ReadJSON()
	if err != nil {
		return err
	}
	fileUtil := FileUtil{}
	if err = fileUtil.RemoveFileIfExist(newFileName); err != nil {
		return err
	}
	file, err := fileUtil.CreateFileIfNotExist(newFileName, 0755)
	if err != nil {
		return err
	}
	defer func() {
		_ = file.Close()
	}()
	fileModel := FileSync{}
	fileModel.File = file
	return fileModel.CoverJSON(content)
}

// FileExists 判断所给路径文件/文件夹是否存在(返回true是存在)
func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, err
	}

	return false, err
}

func IsRelease() bool {
	arg1 := strings.ToLower(os.Args[0])
	return strings.Index(arg1, "go_build") < 0 && strings.Index(arg1, "go-build") < 0
}

func CurrentFile() string {
	_, file, _, ok := runtime.Caller(2)
	if !ok {
		panic(errors.New("Can not get current file info"))
	}
	return file
}

func GetExecPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}
