package flowdiagrams

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Drawer struct {
	nodeNames []string
}

func NewDrawer() *Drawer {
	return &Drawer{}
}

func (d *Drawer) AddNode(name string) {
	d.nodeNames = append(d.nodeNames, name)
}

func (d *Drawer) SaveToXml(file string) error {
	type node struct {
		Name string `xml:",attr"`
	}
	type XmlData struct {
		Nodes []node `xml:"nodes>node"`
	}

	xmlData := &XmlData{}
	for _, nodeName := range d.nodeNames {
		xmlData.Nodes = append(xmlData.Nodes, node{
			Name: nodeName,
		})
	}

	xmlBytes, err := xml.MarshalIndent(xmlData, "", "  ")
	if err != nil {
		return fmt.Errorf("Unable to format data as XML, error: %s", err.Error())
	}

	return ioutil.WriteFile(file, xmlBytes, 0655)
}
