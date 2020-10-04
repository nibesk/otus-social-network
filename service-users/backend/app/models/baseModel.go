package models

import "database/sql"

func checkIsExecDidUpdate(res sql.Result) error {
	rows, err := res.RowsAffected()
	if nil != err {
		return err
	}

	if 0 == rows {
		return sql.ErrNoRows
	}

	return nil
}
