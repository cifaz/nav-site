package conf

import (
	"bufio"
	"os"
	"testing"
)

func TestCreateConfAuto(t *testing.T) {
	confFile := "conf/config.yaml"
	os.MkdirAll("conf", os.ModePerm)
	file, _ := os.OpenFile(confFile, os.O_WRONLY|os.O_CREATE, 0644)
	defer file.Close()

	//file.WriteAt(DefaultConfigFile, 0)
	abc := "sdggggggggggggggggggggg"
	writer := bufio.NewWriter(file)
	//writer.Write(DefaultConfigFile)
	writer.Write([]byte(abc))
	writer.Flush()
}
