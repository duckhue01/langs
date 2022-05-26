package main

import (
	"testing"

	"github.com/duckhue01/lang/go/go-mock/mock"
	"github.com/golang/mock/gomock"
	is "github.com/stretchr/testify/assert"
)

func TestService_GetUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	per := mock.NewMockPersistence(ctrl)

	per.EXPECT().GetUser("duckhue01").Return("duckhue02")

	type fields struct {
		per Persistence
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "",
			fields: fields{
				per: per,
			},
			args: args{
				name: "duckhue01",
			},
			want: "duckhue01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				per: tt.fields.per,
			}
			got := s.GetUser(tt.args.name)
			is.Equal(t, tt.want, got, "dasdasdasd")
		})
	}

}
