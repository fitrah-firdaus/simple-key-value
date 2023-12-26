package keyvalue

import (
	"database/sql"
	"github.com/fitrah-firdaus/simple-key-value/pkg/entities"
	"github.com/gofiber/fiber/v2/log"
	"time"
)

type mysqlRepository struct {
	DB *sql.DB
}

func (m mysqlRepository) CreateOrUpdateKey(value *entities.KeyValue) (*entities.KeyValue, error) {
	value.CreatedAt = time.Now()
	value.UpdatedAt = time.Now()
	res, err := m.DB.Query("SELECT * FROM kv where keylog = ?", value.Key)
	defer res.Close()
	if err != nil {
		log.Error(err)
	}
	if res.Next() {
		stmt, err := m.DB.Prepare("UPDATE kv set value =?, updated_at = ? where keylog =?")
		if err != nil {
			log.Error(err)
		}
		resUpdate, err := stmt.Exec(value.Value, value.UpdatedAt, value.Key)
		if err != nil {
			log.Error(err)
		}
		affected, err := resUpdate.RowsAffected()
		if err != nil {
			log.Error(err)
		}
		if affected != 0 {
			return value, nil
		}
	}
	stmt, err := m.DB.Prepare("INSERT INTO kv (keylog, value, created_at, updated_at) VALUES (?,?,?,?)")
	if err != nil {
		log.Error(err)
	}
	_, err = stmt.Exec(value.Key, value.Value, value.CreatedAt, value.UpdatedAt)
	if err != nil {
		log.Error(err)
	}
	return value, nil
}

func (m mysqlRepository) GetKey(key string) (*entities.KeyValue, error) {
	var keyValue *entities.KeyValue
	res, err := m.DB.Query("SELECT * FROM kv where keylog = ?", key)
	defer res.Close()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if res.Next() {
		res.Scan(nil, &keyValue.Key, &keyValue.Value, &keyValue.CreatedAt, &keyValue.UpdatedAt)
	}
	return keyValue, nil
}

func (m mysqlRepository) DeleteKey(key string) error {
	stmt, err := m.DB.Prepare("DELETE FROM kv where keylog =?")
	if err != nil {
		return err
	}
	stmt.Exec(key)
	return nil
}

func NewMySQLRepo(db *sql.DB) Repository {
	return &mysqlRepository{
		DB: db,
	}
}
