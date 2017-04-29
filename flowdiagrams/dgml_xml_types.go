package flowdiagrams

type DirectedGraph struct {
	//https://msdn.microsoft.com/en-us/library/dn966108.aspx
	Nodes      []*Node     `xml:"Nodes>Node"`
	Links      []*Link     `xml:"Links>Link"`
	Categories []*Category `xml:"Categories>Category"`
}

type Node struct {
	Id       string `xml:"Id,attr"`
	Label    string `xml:"Label,attr"`
	Category string `xml:"Category,attr"`
}

type Link struct {
	Source   string `xml:"Source,attr"`
	Target   string `xml:"Target,attr"`
	Category string `xml:"Category,attr"`
	Label    string `xml:"Label,attr"`
}

type Category struct {
	Id         string `xml:"Id,attr"`
	Background string `xml:"Background,attr"`
	Stroke     string `xml:"Stroke,attr"`
}
