package api

import "time"

const TimeFormat = "2006-01-02"

type RequestDate time.Time

func (t *RequestDate) UnmarshalJSON(data []byte) (err error) {
	if len(data) == 2 {
		*t = RequestDate(time.Time{})
		return
	}

	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = RequestDate(now)
	return
}

func (t RequestDate) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

func (t RequestDate) String() string {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return string(b)
}
