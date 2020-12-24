package brokerrepo

import (
	"database/sql"
	"errors"
	"os"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
	"github.com/na7r1x/acectl/internal/core/domain"
)

type sqliteRepo struct {
	DbLocation string
}

var database *sql.DB

// Instantiate an SQLite BrokerRepo
func NewSqliteRepo(loc string) *sqliteRepo {
	s := new(sqliteRepo)
	s.DbLocation = loc
	s.Init()
	return s
}

func (s sqliteRepo) Init() error {
	var err error
	database, err = sql.Open("sqlite3", s.DbLocation)
	if err != nil {
		return errors.New("failed opening database connection; " + err.Error())
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS broker (id TEXT PRIMARY KEY, created TIMESTAMP, host TEXT, port INT, username TEXT, password TEXT)")
	if err != nil {
		return errors.New("failed executing create statement; " + err.Error())
	}
	statement.Exec()

	return nil
}

func (s sqliteRepo) Destroy() error {
	err := database.Close()
	if err != nil {
		return errors.New("could not close database; " + err.Error())
	}

	errDel := os.Remove(s.DbLocation)
	if errDel != nil {
		return errors.New("failed deleting sqlite database; " + errDel.Error())
	}
	return nil
}

func (s sqliteRepo) GetAll() ([]domain.Broker, error) {
	records := []domain.Broker{}

	rows, err := database.Query("SELECT * FROM broker")
	if err != nil {
		return records, errors.New("failed retrieving brokers from storage; " + err.Error())
	}

	for rows.Next() {
		var record domain.Broker
		err := rows.Scan(&record.Id, &record.Created, &record.Host, &record.Port, &record.Username, &record.Password)
		if err != nil {
			return records, errors.New("failed mapping brokers from storage; " + err.Error())
		}

		records = append(records, record)
	}

	return records, nil
}

func (s sqliteRepo) Set(r domain.Broker) error {
	statement, err := database.Prepare("INSERT OR REPLACE INTO broker (id,created,host,port,username,password) values(?,?,?,?,?,?)")
	if err != nil {
		return errors.New("failed database upsert; " + err.Error())
	}
	statement.Exec(r.Id, r.Created, r.Host, r.Port, r.Username, r.Password)
	return nil
}

func (s sqliteRepo) Delete(id string) error {
	statement, err := database.Prepare("DELETE FROM broker WHERE id = ?")
	if err != nil {
		return errors.New("failed to prepare delete statement; " + err.Error())
	}
	result, err := statement.Exec(id)
	if err != nil {
		return errors.New("failed to delete record [" + id + "]; " + err.Error())
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return errors.New("failed to check affected rows; " + err.Error())
	}
	if affected == 0 {
		return errors.New("record [" + id + "] does not exist; ")
	}
	return nil
}

func (s sqliteRepo) Get(id string) (domain.Broker, error) {
	row := database.QueryRow("SELECT * FROM broker WHERE id = ?", id)

	var record domain.Broker

	err := row.Scan(&record.Id, &record.Created, &record.Host, &record.Port, &record.Username, &record.Password)
	if err != nil {
		return record, errors.New("failed retrieving broker from storage; " + err.Error())
	}

	return record, nil
}
