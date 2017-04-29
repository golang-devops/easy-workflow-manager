package flowdiagrams

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
)

type Drawer struct {
	idCounter         int
	diagramBackground string
	graph             *DirectedGraph
}

func NewDrawer(diagramBackground string) *Drawer {
	return &Drawer{
		diagramBackground: diagramBackground,
		graph:             &DirectedGraph{},
	}
}

func (d *Drawer) AddCategory(categoryProvider CategoryProvider) {
	categoryAlreadyAdded := false
	for _, existingCategory := range d.graph.Categories {
		if existingCategory.Id == categoryProvider.Id() {
			categoryAlreadyAdded = true
			break
		}
	}
	if !categoryAlreadyAdded {
		d.graph.Categories = append(d.graph.Categories, &Category{
			Id:         categoryProvider.Id(),
			Background: categoryProvider.Background(),
			Stroke:     categoryProvider.TextColor(),
		})
	}
}

func (d *Drawer) AddNode(label string, categoryProvider CategoryProvider) (id string) {
	d.idCounter++
	node := &Node{
		Id:       fmt.Sprintf("node-%d", d.idCounter),
		Label:    label,
		Category: categoryProvider.Id(),
	}
	d.graph.Nodes = append(d.graph.Nodes, node)

	d.AddCategory(categoryProvider)

	return node.Id
}

func (d *Drawer) AddLink(fromID, toID, linkName string, categoryProvider CategoryProvider) {
	d.graph.Links = append(d.graph.Links, &Link{
		Source:   fromID,
		Target:   toID,
		Category: categoryProvider.Id(),
		Label:    linkName,
	})

	d.AddCategory(categoryProvider)
}

func (d *Drawer) SaveToDGML(file string) error {
	xmlBytes, err := xml.MarshalIndent(d.graph, "", "  ")
	if err != nil {
		return fmt.Errorf("Unable to format data as XML, error: %s", err.Error())
	}

	xmlString := string(xmlBytes)
	xmlString = strings.Replace(xmlString,
		`<DirectedGraph>`,
		strings.Join([]string{
			`<?xml version="1.0" encoding="utf-8"?>`,
			fmt.Sprintf(`<DirectedGraph Background="%s" xmlns="http://schemas.microsoft.com/vs/2009/dgml">`, d.diagramBackground),
		}, "\n"),
		-1)

	return ioutil.WriteFile(file, []byte(xmlString), 0655)
}
