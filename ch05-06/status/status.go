package status

import "fmt"

// Status is to represent status of task
type Status int

//
const (
	Unknown Status = (iota + 1)
	WIP
	Done
)

var (
	mapping = map[Status]string{
		Unknown: `"Unknown"`,
		WIP:     `"WIP"`,
		Done:    `"Done"`,
	}
	revMapping = make(map[string]Status, len(mapping))
)

func (s Status) String() string {
	str, ok := mapping[s]
	if ok {
		return str
	}
	return ""
}

// MarshalJSON implements json.Marshaler
func (s Status) MarshalJSON() ([]byte, error) {
	str := s.String()
	if str == "" {
		return nil, fmt.Errorf("status.MarshalJSON: unknown value %+v", s)
	}
	return []byte(s.String()), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (s *Status) UnmarshalJSON(data []byte) error {
	status, ok := revMapping[string(data)]
	if ok {
		*s = status
		return nil
	}
	*s = 0
	return fmt.Errorf("status.UnmarshalJSON: unknown value %+v", string(data))
}

func init() {
	for status, str := range mapping {
		revMapping[str] = status
	}
}
