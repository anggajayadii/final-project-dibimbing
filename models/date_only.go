package models

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type DateOnly struct {
	time.Time
}

// Untuk parsing dari JSON
func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := string(b)
	// Hapus tanda kutip
	s = s[1 : len(s)-1]

	if s == "" {
		// Jika kosong, set default time
		d.Time = time.Time{}
		return nil
	}

	// Parse pakai format "2006-01-02"
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return fmt.Errorf("invalid date format: should be YYYY-MM-DD")
	}
	d.Time = t
	return nil
}

// Untuk serialisasi ke JSON
func (d DateOnly) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", d.Time.Format("2006-01-02"))
	return []byte(formatted), nil
}

// Untuk GORM supaya tetap kompatibel
func (d DateOnly) Value() (driver.Value, error) {
	return d.Time, nil
}

func (d *DateOnly) Scan(value interface{}) error {
	if val, ok := value.(time.Time); ok {
		d.Time = val
		return nil
	}
	return fmt.Errorf("failed to scan DateOnly")
}
