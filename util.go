package svc

import (
	"crypto/md5"
	"encoding/hex"
	"time"

	"github.com/gofrs/uuid"
)

func MD5(text string) string {
	h := md5.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}

func MinMaxTime(u []uuid.UUID) (time.Time, time.Time) {
	min := time.Now()
	var max time.Time
	for _, v := range u {
		ts, _ := uuid.TimestampFromV1(v)
		t, _ := ts.Time()
		if t.Before(min) {
			min = t
		}
		if t.After(max) {
			max = t
		}
	}
	return min, max
}
