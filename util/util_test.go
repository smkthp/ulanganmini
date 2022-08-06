package util

import (
	"reflect"
	"testing"

	"github.com/smkthp/ulanganmini/system"
)

func TestUnmarshalTasks(t *testing.T) {
	body := `[{"id":1,"name":"t1","desc":"t1"},{"id":2,"name":"t2","desc":"t2"}]`
	want := []system.Task{
		{ID: 1, Name: "t1", Desc: "t1"},
		{ID: 2, Name: "t2", Desc: "t2"},
	}

	tasks, err := UnmarshalTasks([]byte(body))
	if err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}

	if !reflect.DeepEqual(want, tasks) {
		t.Fatalf("Got %+v, want %+v", tasks, want)
	}
}
