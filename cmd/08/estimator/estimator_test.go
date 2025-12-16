package estimator_test

import (
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/08/box"
	"github.com/StevanFreeborn/advent-of-code-2025/cmd/08/estimator"
)

func TestFrom(t *testing.T) {
	result := estimator.From([]box.Box{})

	if result == nil {
		t.Errorf("estimator.From returned nil; expected non-nil value")
	}

	if len(result.PossibleConnections()) != 0 {
		t.Errorf("estimator.From returned estimator with PossibleConnections() = %d; expected 0", result.PossibleConnections())
	}
}

func TestPossibleConnections(t *testing.T) {
	expectedNumOfConnections := 1

	est := estimator.From([]box.Box{
		box.From("0,0,0"),
		box.From("1,1,1"),
	})

	result := est.PossibleConnections()

	if len(result) != expectedNumOfConnections {
		t.Errorf("PossibleConnections() returned %d connections; expected %d", len(result), expectedNumOfConnections)
	}
}
