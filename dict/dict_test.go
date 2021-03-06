package main

import "testing"

func TestDict(t *testing.T) {
	t.Run("get a value with key", func(t *testing.T) {
		d := Dict{"k": "v"}
		v, err := d.Get("k")
		if err != nil {
			t.Errorf("unexpected error occur: %v", err)
		}
		want := "v"
		if v != want {
			t.Errorf("got: %s, want: %s", v, want)
		}
	})
	t.Run("get a value with key not existing", func(t *testing.T) {
		d := Dict{"k": "v"}
		v, err := d.Get("k1")
		if err != ErrNotFound {
			t.Error("get key which does not exist should get a ErrNotFound")
		}
		want := ""
		if v != want {
			t.Errorf("got: %s, want: %s", v, want)
		}
	})
	t.Run("put an entry into the map", func(t *testing.T) {
		d := Dict{}
		d.Put("k", "v")
		v, err := d.Get("k")
		if err != nil {
			t.Errorf("unexpected error occur: %v", err)
		}
		want := "v"
		if v != want {
			t.Errorf("got: %s, want: %s", v, want)
		}
	})
}
