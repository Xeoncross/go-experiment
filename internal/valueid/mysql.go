package valueid

import (
	"database/sql"
	"fmt"
)

// MySQLStore
type MySQLStore struct {
	Table  string
	Column string
	DB     *sql.DB
}

// insert inserts email into the database
func (m *MySQLStore) insert(value string) (ID int64, err error) {

	sqlstr := fmt.Sprintf("INSERT INTO %s (%s) VALUES (?)", m.Table, m.Column)

	var res sql.Result
	res, err = m.DB.Exec(sqlstr, value)
	if err != nil {
		return
	}

	ID, err = res.LastInsertId()
	return
}

// getID of value
func (m *MySQLStore) getID(value string) (ID int64, ok bool) {
	sqlstr := fmt.Sprintf("SELECT id FROM %s WHERE %s = ? LIMIT 1", m.Table, m.Column)
	err := m.DB.QueryRow(sqlstr, value).Scan(&ID)
	if err != nil {
		return 0, false
	}

	return ID, true
}

// GetByID of value
func (m *MySQLStore) GetByID(id int64) (value string, ok bool) {
	sqlstr := fmt.Sprintf("SELECT id FROM %s WHERE %s = ? LIMIT 1", m.Table, m.Column)
	err := m.DB.QueryRow(sqlstr, id).Scan(&value)
	if err != nil {
		return "", false
	}

	return value, true
}

// GetOrCreateValueID fetches or creates an entry for the given value
func (m *MySQLStore) GetOrCreateValueID(value string) (ID int64, err error) {

	var ok bool
	if ID, ok = m.getID(value); ok {
		return
	}

	ID, err = m.insert(value)
	if err == nil && ID > 0 {
		return
	}

	// Insert failed, try to read again
	ID, _ = m.getID(value)

	return
}
