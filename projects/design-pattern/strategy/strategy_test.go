package strategy

import "testing"

func TestStrategyA(t *testing.T) {
	ctx := NewContext(&StrategyA{})
	ctx.Execute()
}

func TestStrategyB(t *testing.T) {
	ctx := NewContext(&StrategyB{})
	ctx.Execute()
}
