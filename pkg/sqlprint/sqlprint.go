package sqlprint

import (
	"database/sql"
	"errors"
	"fmt"
)

func DumpTable(rows *sql.Rows) error {
	cols, err := rows.Columns()
	if err != nil {
		return errors.New("Failed to get columns; " + err.Error())
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))
	result := make([]string, len(cols))

	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i, _ := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return errors.New("Failed to scan row; " + err.Error())
		}

		for i, raw := range rawResult {
			if raw == nil {
				result[i] = "\\N"
			} else {
				result[i] = string(raw)
			}
		}

		fmt.Printf("%#v\n", result)
	}

	return nil
}
