package models

import "database/sql"

//delete deletes a group from database
// db the DB connexion
//sid the id of the record to be deleted
//
//return:
//int64 the number of record that has been deleted (0 or 1 if all is OK, -1 if an error is thrown)
//error the error that has been raised during the delete action (nil if all is OK)
func delete(db *sql.DB, query string, sid int) (int64, error) {
	res, err := db.Exec(query, sid)
	if err != nil {
		return -1, err
	}
	count, err2 := res.RowsAffected()
	if err2 != nil {
		return -1, err2
	}
	return count, nil
}
