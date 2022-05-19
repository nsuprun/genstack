package stack_test

import (
	"testing"

	"github.com/magiconair/properties/assert"
	stack "github.com/nsuprun/genstack/src"
)

func TestPush(t *testing.T) {
	var stack stack.Stack[string]

	assert.Equal(t, stack.IsEmpty(), true)
	stack.Push("test")
	assert.Equal(t, stack.IsEmpty(), false)
	assert.Equal(t, stack.Size(), 1)
}

func TestPop(t *testing.T) {
	var stack stack.Stack[string]
	stack.Push("value1")
	stack.Push("value2")

	popped, found := stack.Pop()
	assert.Equal(t, popped, "value2")
	assert.Equal(t, found, true)
}

func TestTop(t *testing.T) {
	var stack stack.Stack[string]
	stack.Push("value1")
	stack.Push("value2")
	v, found := stack.Top()
	assert.Equal(t, "value2", v)
	assert.Equal(t, found, true)
}

func TestSize(t *testing.T) {
	var stack stack.Stack[string]
	assert.Equal(t, stack.Size(), 0)
	stack.Push("value1")
	stack.Push("value2")
	assert.Equal(t, stack.Size(), 2)
	stack.Pop()
	stack.Pop()
	assert.Equal(t, stack.Size(), 0)
	stack.Pop()
	assert.Equal(t, stack.Size(), 0)
}
