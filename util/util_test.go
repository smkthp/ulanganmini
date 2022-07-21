package util

import (
	"reflect"
	"testing"

	"github.com/smkthp/ulanganmini/system"
)

func TestUnmarshalTasks(t *testing.T) {
	body := `[{"taskID":1,"title":"t1"},{"taskID":2,"title":"t2"}]`
	want := []system.Task{
		{ID: 1, Title: "t1"},
		{ID: 2, Title: "t2"},
	}

	tasks, err := UnmarshalTasks([]byte(body))
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(want, tasks) {
		t.Fatalf("Got %+v, want %+v", tasks, want)
	}
}
