package status

import (
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	dict := map[Status]string{
		Unknown: `"Unknown"`,
		WIP:     `"WIP"`,
		Done:    `"Done"`,
	}

	for status, str := range dict {
		b, err := status.MarshalJSON()
		if err != nil || string(b) != str {
			t.Errorf("Expected %+v but got %+v", str, string(b))
		}
	}
}

func TestUnmarshalJSON(t *testing.T) {
	dict := map[string]Status{
		`"Unknown"`: Unknown,
		`"WIP"`:     WIP,
		`"Done"`:    Done,
	}

	for str, status := range dict {
		var s Status
		(&s).UnmarshalJSON([]byte(str))
		if s != status {
			t.Errorf("Expected %+v but got %+v", status, s)
		}
	}
}
