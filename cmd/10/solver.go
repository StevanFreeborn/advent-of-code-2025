package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	White  = "\033[37m"
	Gray   = "\033[90m"
	Clear  = "\033[H\033[2J"
)

type Solution struct {
	a, b, c, d, e, f int
	sum              int
	found            bool
}

func main() {
	best := Solution{sum: 99999}
	delay := 150 * time.Millisecond

	for f := 0; f <= 5; f++ {
		for d := 0; d <= 7; d++ {

			a := 2 - d + f
			b := 5 - f
			c := 4 - d - f
			e := 3 - f

			currentSum := a + b + c + d + e + f

			validA := a >= 0
			validB := b >= 0
			validC := c >= 0
			validE := e >= 0

			isSolution := validA && validB && validC && validE

			isNewBest := false

			if isSolution {
				if currentSum < best.sum {
					best = Solution{a, b, c, d, e, f, currentSum, true}
					isNewBest = true
				}
			}

			printDashboard(d, f, a, b, c, e, currentSum, best, isSolution)

			if isNewBest {
				time.Sleep(1 * time.Second)
			} else {
				time.Sleep(delay)
			}
		}
	}

	printFinalResult(best)
}

func printDashboard(d, f, a, b, c, e, sum int, best Solution, valid bool) {
	var sb strings.Builder

	sb.WriteString(Clear)

	sb.WriteString(fmt.Sprintf(" %sFREE VARIABLES%s\n", Yellow, Reset))
	sb.WriteString(" ────────────────────────────────────────\n")
	sb.WriteString(fmt.Sprintf("  INPUT d: %s%-3d%s %s\n", White, d, Reset, bar(d, 7)))
	sb.WriteString(fmt.Sprintf("  INPUT f: %s%-3d%s %s\n\n", White, f, Reset, bar(f, 7)))

	sb.WriteString(fmt.Sprintf(" %sDEPENDENT VARIABLES%s\n", Yellow, Reset))
	sb.WriteString(" ────────────────────────────────────────\n")
	sb.WriteString(formatVar("a", "2 - d + f", a))
	sb.WriteString(formatVar("b", "5 - f", b))
	sb.WriteString(formatVar("c", "4 - d - f", c))
	sb.WriteString(formatVar("e", "3 - f", e))

	sb.WriteString("\n")

	sb.WriteString(fmt.Sprintf(" %sSTATUS%s\n", Yellow, Reset))
	sb.WriteString(" ────────────────────────────────────────\n")

	statusColor := Red
	statusText := "INVALID (Constraints Failed)"

	if valid {
		statusColor = Green
		statusText = "VALID SOLUTION"
	}

	sb.WriteString(fmt.Sprintf("  Current Sum:    %d\n", sum))
	sb.WriteString(fmt.Sprintf("  Constraint Check: %s%s%s\n\n", statusColor, statusText, Reset))

	sb.WriteString(fmt.Sprintf(" %sBEST MINIMUM FOUND SO FAR%s\n", Yellow, Reset))
	sb.WriteString(" ────────────────────────────────────────\n")

	if best.found {
		sb.WriteString(fmt.Sprintf("  Total Sum: %s%d%s\n", Green, best.sum, Reset))
		sb.WriteString(fmt.Sprintf("  Values:    a=%d, b=%d, c=%d, d=%d, e=%d, f=%d\n", best.a, best.b, best.c, best.d, best.e, best.f))
	} else {
		sb.WriteString("  Searching...\n")
	}

	fmt.Print(sb.String())
}

func formatVar(name, eq string, val int) string {
	color := Green
	check := "OK"

	if val < 0 {
		color = Red
		check = "FAIL (< 0)"
	}

	displayVal := max(val, 0)
	visual := bar(displayVal, 10)

	return fmt.Sprintf("  %s = %-10s = %s%-3d%s [%-10s] %s\n", name, eq, color, val, Reset, check, visual)
}

func bar(val, max int) string {
	if val < 0 {
		val = 0
	}
	if val > max {
		val = max
	}
	return "[" + strings.Repeat("█", val) + strings.Repeat("░", max-val) + "]"
}

func printFinalResult(s Solution) {
	fmt.Printf("Smallest Total Possible: %s%d%s\n", Green, s.sum, Reset)
	fmt.Printf("Configuration: a=%d, b=%d, c=%d, d=%d, e=%d, f=%d\n", s.a, s.b, s.c, s.d, s.e, s.f)
}
