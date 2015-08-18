package rapi

import (
	"log"
	"time"
)

type APIDate struct {
	original []byte
	time     *time.Time
}

const (
	DATE_LONG_1 string = `"2006-01-02T15:04:05+00:00"`
	DATE_LONG_2 string = `"2006-01-02T15:04:05Z00:00"`
	DATE_SHORT  string = `"2006-01-02"`
)

var layouts = []string{
	DATE_LONG_1,
	DATE_LONG_2,
	DATE_SHORT,
}

func (ad *APIDate) UnmarshalJSON(data []byte) error {
	var (
		t   time.Time
		err error
	)

	ad.original = data

	for _, layout := range layouts {
		t, err = time.Parse(layout, string(data))
		if err != nil {
			log.Printf("Error when parsing date: %s", err)
		} else {
			ad.time = &t
			return nil
		}
	}

	return err
}

func (ad *APIDate) MarshalJSON() ([]byte, error) {
	d := ad.time.Format(DATE_SHORT)
	return []byte(d), nil
}
