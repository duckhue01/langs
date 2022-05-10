package main

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestFoo(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := NewMockFoo(ctrl)

	m.EXPECT().Bar(10).Return(20)

	assert.Equal(t, m.Bar(10), 10)

}
