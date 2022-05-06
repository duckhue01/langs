package main

import "testing"

func TestStubObject_Stub(t *testing.T) {
	type fields struct {
		data string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StubObject{
				data: tt.fields.data,
			}
			if got := s.Stub(); got != tt.want {
				t.Errorf("StubObject.Stub() = %v, want %v", got, tt.want)
			}
		})
	}
}
