package domain

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID          uuid.UUID  `json:"id"`
	ServiceName string     `json:"service_name"`
	Price       uint32     `json:"price"`
	UserID      uuid.UUID  `json:"user_id"`
	StartDate   MonthYear  `json:"start_date"`
	EndDate     *MonthYear `json:"end_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type MonthYear struct {
	Month int
	Year  int
}

func (my MonthYear) String() string {
	return time.Date(my.Year, time.Month(my.Month), 1, 0, 0, 0, 0, time.UTC).Format("01-2006")
}

func (my MonthYear) Time() time.Time {
	return time.Date(my.Year, time.Month(my.Month), 1, 0, 0, 0, 0, time.UTC)
}

func ParseMonthYear(s string) (MonthYear, error) {
	t, err := time.Parse("01-2006", s)
	if err != nil {
		return MonthYear{}, err
	}
	return MonthYear{
		Month: int(t.Month()),
		Year:  t.Year(),
	}, nil
}

func (my *MonthYear) UnmarshalJSON(data []byte) error {
	s := string(data)
	if len(s) < 2 {
		return nil
	}
	s = s[1 : len(s)-1]

	parsed, err := ParseMonthYear(s)
	if err != nil {
		return err
	}

	*my = parsed
	return nil
}

func (my MonthYear) MarshalJSON() ([]byte, error) {
	return []byte(`"` + my.String() + `"`), nil
}
