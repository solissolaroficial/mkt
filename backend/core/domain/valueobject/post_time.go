package valueobject

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidTimeFormat = errors.New("invalid time format, expected HH:MM")

type PostTime struct {
	time string
}

func NewPostTime(timeStr string) (*PostTime, error) {
	if len(timeStr) != 5 {
		return nil, ErrInvalidTimeFormat
	}

	parts := strings.Split(timeStr, ":")
	if len(parts) != 2 {
		return nil, ErrInvalidTimeFormat
	}

	// Validar hora (00-23)
	hour, err := strconv.Atoi(parts[0])
	if err != nil || hour < 0 || hour > 23 {
		return nil, ErrInvalidTimeFormat
	}

	// Validar minuto (00-59)
	minute, err := strconv.Atoi(parts[1])
	if err != nil || minute < 0 || minute > 59 {
		return nil, ErrInvalidTimeFormat
	}

	return &PostTime{time: timeStr}, nil
}

func ReconstructPostTime(timeStr string) *PostTime {
	return &PostTime{time: timeStr}
}

func (t *PostTime) Value() string {
	return t.time
}

func (t *PostTime) String() string {
	return t.time
}
