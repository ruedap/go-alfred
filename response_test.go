package alfred

import (
	"reflect"
	"testing"
)

func TestResponse_NewResponse(t *testing.T) {
	actual := NewResponse()
	expected := new(Response)
	expected.Items = []ResponseItem{}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestResponse_ErrorXML(t *testing.T) {
	actual := ErrorXML("foo", "bar", "baz")
	expected := `<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="false" arg="baz" uid="error"><title>foo</title><subtitle>bar</subtitle><icon>/System/Library/CoreServices/CoreTypes.bundle/Contents/Resources/AlertStopIcon.icns</icon></item></items>`

	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestResponse_AddItem(t *testing.T) {
	r := NewResponse()

	actual := r.AddItem(&ResponseItem{Title: "title-foo"}).Items
	expected := []ResponseItem{ResponseItem{Title: "title-foo"}}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestResponse_ToXML(t *testing.T) {
	r := NewResponse()
	item := ResponseItem{
		Valid:    true,
		UID:      "uid-foo",
		Title:    "title-foo",
		Subtitle: "Subtitle foo.",
		Arg:      "arg-foo",
		Icon:     "./icons/title-foo.png",
	}
	r.AddItem(&item)

	actual, err := r.ToXML()
	if err != nil {
		t.Error("failed to convert to XML")
	}

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<items><item valid="true" arg="arg-foo" uid="uid-foo"><title>title-foo</title><subtitle>Subtitle foo.</subtitle><icon>./icons/title-foo.png</icon></item></items>`
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}

func TestResponse_ToXML_Blank(t *testing.T) {
	r := NewResponse()

	actual, err := r.ToXML()
	if err != nil {
		t.Error("failed to convert to XML")
	}

	expected := `<?xml version="1.0" encoding="UTF-8"?>
<items></items>`
	if actual != expected {
		t.Errorf("expected %v to eq %v", actual, expected)
	}
}
