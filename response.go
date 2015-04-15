package alfred

import (
	"encoding/xml"
	"fmt"
	"strings"
)

type Response struct {
	XMLName xml.Name `xml:"items"`
	Items   []ResponseItem
}

type ResponseItem struct {
	XMLName  xml.Name `xml:"item"`
	Valid    bool     `xml:"valid,attr"`
	Arg      string   `xml:"arg,attr"`
	UID      string   `xml:"uid,attr"`
	Title    string   `xml:"title"`
	Subtitle string   `xml:"subtitle"`
	Icon     string   `xml:"icon"`
	Extra    map[string]string
}

func NewResponse() *Response {
	r := new(Response)
	r.Items = []ResponseItem{}

	return r
}

func ErrorXML(title, subtitle, arg string) string {
	r := NewResponse()
	item := ResponseItem{
		Valid:    false,
		UID:      "error",
		Title:    title,
		Subtitle: subtitle,
		Arg:      arg,
		Icon:     "/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns",
	}
	r.AddItem(&item)
	str, _ := r.ToXML()

	return str
}

func (r *Response) AddItem(item *ResponseItem) *Response {
	r.Items = append(r.Items, *item)

	return r
}

func (r *Response) ToXML() (string, error) {
	var x, err = xml.Marshal(r)

	return xml.Header + string(x), err
}

func (ri ResponseItem) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Name.Local = "item"
	start.Attr = []xml.Attr{
		{Name: xml.Name{Local: "valid"}, Value: fmt.Sprint(ri.Valid)},
		{Name: xml.Name{Local: "arg"}, Value: ri.Arg},
		{Name: xml.Name{Local: "uid"}, Value: ri.UID},
	}
	e.EncodeToken(start)
	e.EncodeElement(ri.Title, xml.StartElement{Name: xml.Name{Local: "title"}})
	e.EncodeElement(ri.Subtitle, xml.StartElement{Name: xml.Name{Local: "subtitle"}})
	e.EncodeElement(ri.Icon, xml.StartElement{Name: xml.Name{Local: "icon"}})
	for k, v := range ri.Extra {
		l := strings.ToLower(k)
		e.EncodeElement(v, xml.StartElement{Name: xml.Name{Local: l}})
	}
	e.EncodeToken(start.End())
	return nil
}
