package postgres

import "go.trueardian.com/chronica"

func ExportBuildActaQuery(chronicumID string, q chronica.ActaQuery) (string, []any) {
	return buildActaQuery(chronicumID, q)
}
