// Package goDOM provides methodes similar to the Javascript Document interface.
//
// It can be used to extract data from HTML documents.
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Document
package goDOM

import (
	"bytes"
	"fmt"
	"io"
	"slices"
	"strings"

	"golang.org/x/net/html"
)

// New returns the parsed tree for the HTML from the given Reader as a DOM object.
func New(r io.Reader) (*DOM, error) {
	node, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return newDOM(node), nil
}

func newDOM(node *html.Node) *DOM {
	return &DOM{node, nil, nil}
}

// A DOM represents a parsed HTML document.
// It implements methodes to extract data from the document.
type DOM struct {
	node            *html.Node
	flatElementList []*DOM
	flatNodeList    []*DOM
}

// TagName returns a string representation of the nodes tag.
func (d *DOM) TagName() string {
	nodeType := d.node.Type
	if nodeType == html.TextNode {
		return "text"
	}
	if nodeType == html.DocumentNode {
		return "document"
	}
	if nodeType == html.DoctypeNode {
		return "doctype"
	}
	return d.node.Data
}

// FirstElementChild returns the document's first child Element, or nil if there are no child elements.
//
// FirstElementChild includes only element nodes. Other node types like text or comment are ignored.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Document/firstElementChild
func (d *DOM) FirstElementChild() *DOM {
	if d.node == nil || d.node.FirstChild == nil {
		return newDOM(nil)
	}
	node := d.node.FirstChild
	for node != nil {
		if node.Type == html.ElementNode {
			return newDOM(node)
		}
		node = node.NextSibling
	}
	return newDOM(nil)
}

// LastElementChild returns the document's last child Element, or nil if there are no child elements.
//
// LastElementChild includes only element nodes. Other node types like text or comment are ignored.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Document/lastElementChild
func (d *DOM) LastElementChild() *DOM {
	if d.node == nil || d.node.LastChild == nil {
		return newDOM(nil)
	}
	node := d.node.LastChild
	for node != nil {
		if node.Type == html.ElementNode {
			return newDOM(node)
		}
		node = node.PrevSibling
	}
	return newDOM(nil)
}

// NextElementSibling returns the element immediately following the specified one in its parent's children list,
// or null if the specified element is the last one in the list.
//
// NextElementSibling includes only element nodes. Other node types like text or comment are ignored.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Element/nextElementSibling
func (d *DOM) NextElementSibling() *DOM {
	if d.node == nil {
		return newDOM(nil)
	}
	node := d.node.NextSibling
	for node != nil {
		if node.Type == html.ElementNode {
			return newDOM(node)
		}
		node = node.NextSibling
	}
	return newDOM(nil)
}

// PreviousElementSibling returns the element immediately prior the specified one in its parent's children list,
// or null if the specified element is the first one in the list
//
// PreviousElementSibling includes only element nodes. Other node types like text or comment are ignored.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Element/previousElementSibling
func (d *DOM) PreviousElementSibling() *DOM {
	if d.node == nil {
		return newDOM(nil)
	}
	node := d.node.PrevSibling
	for node != nil {
		if node.Type == html.ElementNode {
			return newDOM(node)
		}
		node = node.PrevSibling
	}
	return newDOM(nil)
}

// Children returns a slice which contains all of the child elements of the element upon which it was called.
//
// The Children slice includes only element nodes. Other node types like text or comment are ignored.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Element/children
func (d *DOM) Children() []*DOM {
	children := make([]*DOM, 0)
	nodes := d.childNodes()
	for _, node := range nodes {
		if node.isElementNode() {
			children = append(children, node)
		}
	}
	return children
}

// ChildElementCount returns the number of child elements of this element.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Element/childElementCount
func (d *DOM) ChildElementCount() int {
	return len(d.Children())
}

// Attributes returns a map of all attribute nodes registered to the specified node
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Element/attributes
func (d *DOM) Attributes() map[string]string {
	var attr = make(map[string]string)
	for _, a := range d.node.Attr {
		attr[a.Key] = a.Val
	}
	return attr
}

// ClassName returns a string representing the class or space-separated classes of the current element.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Element/className
func (d *DOM) ClassName() string {
	return d.getAttribute("class")
}

// ClassList returns a slice containing all the classes of the current element.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Element/classList
func (d *DOM) ClassList() []string {
	return strings.Split(d.ClassName(), " ")
}

// Id returns returns a string representing the id of the current element.
//
// If the id value is not the empty string, it must be unique in a document.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Element/id
func (d *DOM) Id() string {
	return d.getAttribute("id")
}

// HasAttribute returns a Boolean value indicating whether the specified element has the specified attribute or not.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Element/hasAttribute
func (d *DOM) HasAttribute(key string) bool {
	attributes := d.Attributes()
	_, ok := attributes[key]
	return ok
}

// HasAttribute returns a boolean value indicating whether the current element has any attributes or not.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Element/hasAttributes
func (d *DOM) HasAttributes() bool {
	return len(d.Attributes()) > 0
}

// Render returns a string representation of the DOM.
func (d *DOM) Render() (string, error) {
	var buffer bytes.Buffer
	err := html.Render(&buffer, d.node)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}

// GetElementById returns a DOM object representing the element whose id property matches the specified string.
//
// Since element IDs are required to be unique if specified,
// they're a useful way to get access to a specific element quickly.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Document/getElementById
func (d *DOM) GetElementById(id string) *DOM {
	nodes := d.getFlatElementList(true)
	for _, node := range nodes {
		if node.Id() == id {
			return node
		}
	}
	return nil
}

// GetElementsByTagName returns a slice of elements with the given tag name.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByTagName
func (d *DOM) GetElementsByTagName(tag string) []*DOM {
	elements := make([]*DOM, 0)
	nodes := d.getFlatElementList(true)
	for _, node := range nodes {
		if node.TagName() == tag {
			elements = append(elements, node)
		}
	}
	return elements
}

// GetElementsByClassName returns a slice of elements with the given class name.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Element/getElementsByClassName
func (d *DOM) GetElementsByClassName(class string) []*DOM {
	elements := make([]*DOM, 0)
	nodes := d.getFlatElementList(true)
	for _, node := range nodes {
		if slices.Contains(node.ClassList(), class) {
			elements = append(elements, node)
		}
	}
	return elements
}

// Text returns the text content of the node.
//
// If full is set to true, Text returns the text content of the node and its descendants.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Node/textContent
func (d *DOM) Text(full bool) string {
	var sb strings.Builder
	nodes := make([]*DOM, 0)
	if !full {
		nodes = d.getTextNodes()
	} else {
		for _, node := range d.getFlatNodeList(true) {
			if node.TagName() == "text" {
				nodes = append(nodes, node)
			}
		}
	}
	for _, node := range nodes {
		sb.WriteString(fmt.Sprintf("%s ", node.node.Data))
	}
	return strings.TrimSpace(sb.String())
}

// Parent returns the parent of the specified node in the DOM tree.
//
// See Javascript equivalent: https://developer.mozilla.org/en-US/docs/Web/API/Node/parentNode
func (d *DOM) Parent() *DOM {
	return newDOM(d.node.Parent)
}

// getTextNodes returns all text node children of the given node.
func (d *DOM) getTextNodes() []*DOM {
	children := make([]*DOM, 0)
	nodes := d.childNodes()
	if len(nodes) == 0 {
		return children
	}
	for _, node := range nodes {
		if node.TagName() == "text" {
			children = append(children, node)
		}
	}
	return children
}

// getFlatElementList returns all element nodes in the DOM.
func (d *DOM) getFlatElementList(setCache bool) []*DOM {
	if d.flatElementList != nil {
		return d.flatElementList
	}
	elements := make([]*DOM, 0)
	nodes := d.getFlatNodeList(true)
	for _, node := range nodes {
		if node.isElementNode() {
			elements = append(elements, node)
		}
	}
	if setCache {
		d.flatElementList = elements
	}
	return elements
}

// getFlatNodeList returns all nodes in the DOM.
func (d *DOM) getFlatNodeList(setCache bool) []*DOM {
	if d.flatNodeList != nil {
		return d.flatNodeList
	}
	flatNodeList := make([]*DOM, 0)
	flatNodeList = append(flatNodeList, d)
	if d.node.FirstChild != nil {
		children := d.childNodes()
		for _, child := range children {
			flatNodeList = append(flatNodeList, child.getFlatNodeList(false)...)
		}
	}
	if setCache {
		d.flatNodeList = flatNodeList
	}
	return flatNodeList
}

// childNodes returns all child nodes of a given node.
func (d *DOM) childNodes() []*DOM {
	children := make([]*DOM, 0)
	child := d.node.FirstChild
	if child == nil {
		return children
	}
	children = append(children, newDOM(child))
	for child.NextSibling != nil {
		next := child.NextSibling
		if next != nil {
			children = append(children, newDOM(next))
		}
		child = next
	}
	return children
}

// getAttribute returns the string value of a given attribute.
func (d *DOM) getAttribute(key string) string {
	attributes := d.Attributes()
	val, ok := attributes[key]
	if !ok {
		return ""
	}
	return val
}

// isElementNode returns a Boolean value indicating whether the specified node is an element node or not.
func (d *DOM) isElementNode() bool {
	if d.node == nil {
		return false
	}
	return d.node.Type == html.ElementNode
}
