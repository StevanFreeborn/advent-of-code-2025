package dial_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/01/dial"
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/01/instruction"
)

func TestNew(t *testing.T) {
	expectedStartingValue := 50
	result := dial.New()

	if result.Value() != expectedStartingValue {
		t.Errorf("got unexpected beginning value. expected %d but got %d", expectedStartingValue, result.Value())
	}
}

type TurnTestCase struct {
	Instruction       instruction.Instruction
	ExpectedDialValue int
}

func TestTurn(t *testing.T) {
	testCases := []TurnTestCase{
		{
			Instruction:       instruction.FromParts("L", 68),
			ExpectedDialValue: 82,
		},
		{
			Instruction:       instruction.FromParts("L", 30),
			ExpectedDialValue: 52,
		},
		{
			Instruction:       instruction.FromParts("R", 48),
			ExpectedDialValue: 0,
		},
		{
			Instruction:       instruction.FromParts("L", 5),
			ExpectedDialValue: 95,
		},
		{
			Instruction:       instruction.FromParts("R", 60),
			ExpectedDialValue: 55,
		},
		{
			Instruction:       instruction.FromParts("L", 55),
			ExpectedDialValue: 0,
		},
		{
			Instruction:       instruction.FromParts("L", 1),
			ExpectedDialValue: 99,
		},
		{
			Instruction:       instruction.FromParts("L", 99),
			ExpectedDialValue: 0,
		},
		{
			Instruction:       instruction.FromParts("R", 14),
			ExpectedDialValue: 14,
		},
		{
			Instruction:       instruction.FromParts("L", 82),
			ExpectedDialValue: 32,
		},
	}

	dial := dial.New()

	for i, testCase := range testCases {

		dial.Turn(testCase.Instruction)

		if dial.Value() != testCase.ExpectedDialValue {
			t.Errorf("%d: got dial value %d, but expected dial value %d", i, dial.Value(), testCase.ExpectedDialValue)
		}
	}
}
