package zkmanager

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type XmlConfig struct {
	Zookeeper   string   `xml:"Zookeeper"`
}

var BasicConfig = &XmlConfig{}

func init() {
	file, err := os.Open("/config/zk.config")
	if err != nil {
		log.Printf("error: %v", err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("error: %v", err)
	}
	err = xml.Unmarshal(data, BasicConfig)
	if err != nil {
		log.Printf("error: %v", err)
	}
}

func (c *XmlConfig)Servers()[]string  {
	ss := strings.Split(BasicConfig.Zookeeper,",")
	return ss
}

