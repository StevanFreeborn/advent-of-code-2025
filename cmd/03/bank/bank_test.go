package bank_test

import (
	"slices"
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/03/bank"
)

func TestFrom(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}
	input := "12345"

	result := bank.From(input)

	if slices.Equal(result.Cells(), expected) == false {
		t.Errorf("got %v but wanted %v", result.Cells(), expected)
	}
}

type testCase struct {
	numberOfBanksToTurnOn int
	bank                  bank.Bank
	expected              int64
}

func TestJoltage(t *testing.T) {
	testCases := []testCase{
		{
			numberOfBanksToTurnOn: 2,
			bank:                  bank.From("987654321111111"),
			expected:              98,
		},
		{
			numberOfBanksToTurnOn: 2,
			bank:                  bank.From("811111111111119"),
			expected:              89,
		},
		{
			numberOfBanksToTurnOn: 2,
			bank:                  bank.From("234234234234278"),
			expected:              78,
		},
		{
			numberOfBanksToTurnOn: 2,
			bank:                  bank.From("818181911112111"),
			expected:              92,
		},
		{
			numberOfBanksToTurnOn: 12,
			bank:                  bank.From("987654321111111"),
			expected:              987654321111,
		},
		{
			numberOfBanksToTurnOn: 12,
			bank:                  bank.From("811111111111119"),
			expected:              811111111119,
		},
		{
			numberOfBanksToTurnOn: 12,
			bank:                  bank.From("234234234234278"),
			expected:              434234234278,
		},
		{
			numberOfBanksToTurnOn: 12,
			bank:                  bank.From("818181911112111"),
			expected:              888911112111,
		},
	}

	for _, testCase := range testCases {
		result := testCase.bank.Joltage(testCase.numberOfBanksToTurnOn)

		if result != testCase.expected {
			t.Errorf("got %d but wanted %d", result, testCase.expected)
		}
	}
}
