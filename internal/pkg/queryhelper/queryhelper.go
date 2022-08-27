package queryhelper

import (
	"fmt"
	"strings"
)

// Build the Sort Statement for List Features
// It will dynamically append the SQL string
// based on the argument being passed
func Sort(sortStr string, allowedCols map[string]string) (query string) {
	sorts := strings.Split(strings.ReplaceAll(string(sortStr), " ", ""), ",")
	stmts := []string{}
	for _, sortExpr := range sorts {
		col := strings.TrimPrefix(sortExpr, "-")
		isDesc := strings.HasPrefix(sortExpr, "-")

		if stmt, ok := allowedCols[col]; ok {
			if isDesc {
				stmt = fmt.Sprint(stmt, " DESC")
			}

			stmts = append(stmts, stmt)
		}
	}

	if len(stmts) == 0 {
		return ""
	}

	query = fmt.Sprintf("ORDER BY %s", strings.Join(stmts, ", "))

	return query
}

// Build the Pagination Statement for List Features
// It will dynamically append the SQL string
// based on the argument being passed
func Paginate(limit uint32, offset uint32) (query string, args []interface{}) {
	query = " LIMIT ? OFFSET ?"
	args = []interface{}{limit, offset}
	return query, args
}
