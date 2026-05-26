package reservedkeywords

import (
	"embed"
)

var (
	//go:embed tablescolumns.json
	tablesColumnsFile embed.FS

	//go:embed namespaces.json
	namespacesFile embed.FS

	reservedTablesColumns, reservedNamespaces map[string]map[string]struct{}
)

func init() {
	reservedTablesColumns = load(tablesColumnsFile, "tablescolumns.json")
	reservedNamespaces = load(namespacesFile, "namespaces.json")
}

func load(file embed.FS, fileName string) map[string]map[string]struct{} {
	_ = "STUB: not implemented"
	return nil
}

// IsTableOrColumn checks if the given keyword is a reserved table/column keyword for the destination type.
func IsTableOrColumn(destType, keyword string) bool { _ = "STUB: not implemented"; return false }

// IsNamespace checks if the given keyword is a reserved namespace keyword for the destination type.
func IsNamespace(destType, keyword string) bool { _ = "STUB: not implemented"; return false }

func isKeywordReserved(keywords map[string]map[string]struct{}, destType, keyword string) bool {
	_ = "STUB: not implemented"
	return false
}
