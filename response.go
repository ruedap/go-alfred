package alfred

import "encoding/xml"

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
}

func NewResponse() *Response {
	r := new(Response)
	r.Items = []ResponseItem{}

	return r
}

func (r *Response) AddItem(item *ResponseItem) *Response {
	r.Items = append(r.Items, *item)

	return r
}

func (r *Response) ToXML() (string, error) {
	var x, err = xml.Marshal(r)

	return xml.Header + string(x), err
}
