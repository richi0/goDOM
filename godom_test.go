package goDOM_test

import (
	"os"
	"strings"
	"testing"

	"github.com/richi0/goDOM"
)

func createTestDOM() *goDOM.DOM {
	indexHTML, err := os.Open("test_data/index.html")
	if err != nil {
		panic("Cannot read file: test_data/index.html")
	}
	dom, err := goDOM.New(indexHTML)
	if err != nil {
		panic("Cannot create test dome object")
	}
	return dom
}

func TestAttributes(t *testing.T) {
	dom := createTestDOM()
	if len(dom.Attributes()) != 0 {
		t.Error("Expected no attributes on document")
	}
	meta := dom.LastElementChild().FirstElementChild().FirstElementChild()
	if meta.Attributes()["charset"] != "UTF-8" {
		t.Error("Expected no attributes on document", meta.Attributes())
	}
}

func TestTagName(t *testing.T) {
	dom := createTestDOM()
	tag := dom.TagName()
	if dom.TagName() != "document" {
		t.Error("Expected tag to be:", tag)
	}
}

func TestFirstElemetChild(t *testing.T) {
	dom := createTestDOM()
	first := dom.FirstElementChild()
	expectedTag := "html"
	tag := first.TagName()
	if tag != expectedTag {
		t.Error("Expected first element child tag to be:", expectedTag, ", got:", tag)
	}
}

func TestLastElemetChild(t *testing.T) {
	dom := createTestDOM()
	last := dom.LastElementChild()
	expectedTag := "html"
	tag := last.TagName()
	if tag != expectedTag {
		t.Error("Expected last element child tag to be:", expectedTag, ", got:", tag)
	}
}

func TestNextElemetSibling(t *testing.T) {
	dom := createTestDOM()
	meta := dom.LastElementChild().
		FirstElementChild().
		FirstElementChild()
	expectedTag := "meta"
	tag := meta.TagName()
	if tag != expectedTag {
		t.Error("Expected next element tag to be:", expectedTag, ", got:", tag)
	}
}

func TestPreviousElemetSibling(t *testing.T) {
	dom := createTestDOM()
	text := dom.
		LastElementChild().
		FirstElementChild().
		LastElementChild().
		PreviousElementSibling()
	expectedTag := "link"
	tag := text.TagName()
	if tag != expectedTag {
		t.Error("Expected previous element tag to be:", expectedTag, ", got:", tag)
	}
}

func TestChildren(t *testing.T) {
	dom := createTestDOM()
	children := dom.Children()
	expectedTag := "html"
	tag := children[0].TagName()
	if tag != expectedTag {
		t.Error("Expected element tag to be:", expectedTag, ", got:", tag)
	}
	expectedHeadTag := "meta"
	headTag := dom.LastElementChild().FirstElementChild().Children()[0].TagName()
	if headTag != expectedHeadTag {
		t.Error("Expected element tag to be:", expectedHeadTag, ", got:", headTag)
	}
}

func TestChildElementCount(t *testing.T) {
	dom := createTestDOM()
	count := dom.ChildElementCount()
	if count != 1 {
		t.Error("Expected child element count to be:", 1)
	}
	headChildCount := dom.LastElementChild().FirstElementChild().ChildElementCount()
	if headChildCount != 46 {
		t.Error("Expected child element count to be:", headChildCount)
	}
}

func TestClassName(t *testing.T) {
	dom := createTestDOM()
	expectedClass := ""
	class := dom.ClassName()
	if class != expectedClass {
		t.Error("Expected class to be:", expectedClass, ", got:", class)
	}
	head := dom.LastElementChild().FirstElementChild()
	body := head.NextElementSibling()
	expectedBodyClass := "skin-vector skin-vector-search-vue mediawiki"
	bodyClass := body.ClassName()
	if bodyClass != expectedBodyClass {
		t.Error("Expected class to be:", expectedBodyClass, ", got:", bodyClass)
	}
}

func TestClassList(t *testing.T) {
	dom := createTestDOM()
	classList := dom.ClassList()
	if len(classList) == 0 {
		t.Error("Expected class list len to be:", 0, ", got:", len(classList))
	}
	head := dom.LastElementChild().FirstElementChild()
	body := head.NextElementSibling()
	bodyClassList := body.ClassName()
	if len(bodyClassList) == 3 {
		t.Error("Expected class list len to be:", 3, ", got:", len(bodyClassList))
	}
}

func TestHasAttribute(t *testing.T) {
	dom := createTestDOM()
	hasAttribute := dom.HasAttribute("class")
	if hasAttribute != false {
		t.Error("Expected element to not have attribute 'class'")
	}
	head := dom.LastElementChild().FirstElementChild()
	body := head.NextElementSibling()
	hasAttribute = body.HasAttribute("class")
	if hasAttribute != true {
		t.Error("Expected element to have attribute 'class'")
	}
}

func TestHasAttributes(t *testing.T) {
	dom := createTestDOM()
	hasAttribute := dom.HasAttributes()
	if hasAttribute != false {
		t.Error("Expected element to not have attributes")
	}
	head := dom.LastElementChild().FirstElementChild()
	body := head.NextElementSibling()
	hasAttribute = body.HasAttributes()
	if hasAttribute != true {
		t.Error("Expected element to have attributes")
	}
}

func TestId(t *testing.T) {
	dom := createTestDOM()
	id := dom.Id()
	if id != "" {
		t.Error("Expected element id to be ''")
	}
	head := dom.LastElementChild().FirstElementChild()
	body := head.
		NextElementSibling().
		FirstElementChild().
		NextElementSibling().
		FirstElementChild().
		FirstElementChild().
		FirstElementChild().
		FirstElementChild()
	expectedId := "vector-main-menu-dropdown"
	id = body.Id()
	if id != expectedId {
		t.Error("Expected element id to be: ", expectedId)
	}
}

func TestGetElementById(t *testing.T) {
	dom := createTestDOM()
	element := dom.GetElementById("abc")
	if element != nil {
		t.Error("Expected element to be: nil")
	}
	element = dom.GetElementById("vector-main-menu-dropdown-checkbox")
	if element.TagName() != "input" {
		t.Error("Expected element to be: input")
	}
}

func TestGetElementsByTagName(t *testing.T) {
	dom := createTestDOM()
	elements := dom.GetElementsByTagName("abc")
	if len(elements) != 0 {
		t.Error("Expected no elements but found", len(elements))
	}
	elements = dom.GetElementsByTagName("div")
	if len(elements) != 209 {
		t.Error("Expected 209 elements but found", len(elements))
	}
	elements = dom.GetElementsByTagName("img")
	if len(elements) != 14 {
		t.Error("Expected 14 elements but found", len(elements))
	}
	elements = dom.GetElementsByTagName("a")
	if len(elements) != 1182 {
		t.Error("Expected 1182 elements but found", len(elements))
	}
}

func TestGetElementsByClassName(t *testing.T) {
	dom := createTestDOM()
	elements := dom.GetElementsByClassName("abc")
	if len(elements) != 0 {
		t.Error("Expected no elements but found", len(elements))
	}
	elements = dom.GetElementsByClassName("mw-default-size")
	if len(elements) != 1 {
		t.Error("Expected 1 elements but found", len(elements))
	}
	elements = dom.GetElementsByClassName("mw-editsection")
	if len(elements) != 31 {
		t.Error("Expected 31 elements but found", len(elements))
	}
}

func TestText(t *testing.T) {
	dom := createTestDOM()
	element := dom.GetElementById("toc-Enumerated_types")
	expected := ""
	text := element.Text(false)
	if text != expected {
		t.Errorf("Expected text to be %s but go %s", expected, text)
	}
	element = dom.GetElementById("Enumerated_types")
	expected = "Enumerated types"
	text = element.Text(true)
	if text != expected {
		t.Errorf("Expected text to be %s but go %s", expected, text)
	}
	element = dom.GetElementById("Package_system").Parent().NextElementSibling()
	expected = "Enumerated types"
	text = element.Text(true)
	if !strings.Contains(text, "other packages are accessible") {
		t.Errorf("Expected text to be %s but go %s", expected, text)
	}
}
