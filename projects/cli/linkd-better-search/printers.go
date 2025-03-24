package main

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

type HtmlNodePrinter struct{ *html.Node }

var _ fmt.Stringer = (*HtmlNodePrinter)(nil)

func (np *HtmlNodePrinter) String() string {
	if np.Node == nil {
		return fmt.Sprint("<nil>")
	}
	return fmt.Sprintf(
		"{type: %s tag: %s(%d) attr: %v}",
		HtmlNodeTypePrinter(np.Type),
		strings.TrimSpace(np.Data),
		np.DataAtom,
		np.Attr,
	)
}

type HtmlNodeTypePrinter html.NodeType

var _ fmt.Stringer = (*HtmlNodeTypePrinter)(nil)

func (ntp HtmlNodeTypePrinter) String() string {
	switch ntp {
	case HtmlNodeTypePrinter(html.ElementNode):
		return "Element"
	case HtmlNodeTypePrinter(html.TextNode):
		return "Text"
	case HtmlNodeTypePrinter(html.DocumentNode):
		return "Document"
	case HtmlNodeTypePrinter(html.ErrorNode):
		return "Error"
	case HtmlNodeTypePrinter(html.DoctypeNode):
		return "Doctype"
	case HtmlNodeTypePrinter(html.CommentNode):
		return "Comment"
	default:
		return "Unknown"
	}
}
