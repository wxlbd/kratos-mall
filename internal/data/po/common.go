package po

import (
	"database/sql/driver"
	"encoding/json"
)

type JsonArray[T any] []T

func (j *JsonArray[T]) Scan(src any) error {
	return json.Unmarshal(src.([]byte), j)
}

func (j JsonArray[T]) Value() (driver.Value, error) {
	return json.Marshal(j)
}
