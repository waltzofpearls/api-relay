package rapi

import "time"

type APIDate struct {
	original []byte
	time     *time.Time
}

func (ad *APIDate) Unmarshal(data []byte) error {
	return nil
}

func (ad *APIDate) Marshal() ([]byte, error) {
	return nil, nil
}
