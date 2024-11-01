package replacer

var (
	useDeclarationStart    = []byte(`use (`)
	multiImportPrefixStart = []byte(`import (`)
	multiImportPrefixEnd   = []byte(`)`)
	singleImportPrefix     = []byte(`import `)
	moduleDeclaration      = []byte(`module `)
	replaceDeclaration     = []byte(`//+replace:`)
	currentFolderSegment   = []byte(`./`)
	endLine                = byte('\n')
)

func addLine(data, line []byte) []byte {
	data = append(data, line...)
	return append(data, endLine)
}
