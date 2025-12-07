package grid

type Grid interface {
	Positions() map[Position]string
	Walk()
}

type grid struct {
	positions map[Position]string
}

func NewGrid(input []string) Grid {
	numberOfRows := len(input)
	numberOfColumns := len(input[0])
	positions := map[Position]string{}

	for row := range numberOfRows {
		for column := range numberOfColumns {
			value := string(input[row][column])
			pos := NewPosition(row, column)
			positions[pos] = value
		}
	}

	return grid{
		positions: positions,
	}
}

func (g grid) Positions() map[Position]string {
	positions := map[Position]string{}

	for key, value := range g.positions {
		newPosition := NewPosition(key.Row(), key.Column())
		positions[newPosition] = value
	}

	return positions
}

type Position interface {
	Column() int
	Row() int
}

type position struct {
	column int
	row    int
}

func NewPosition(row int, column int) Position {
	return position{
		row:    row,
		column: column,
	}
}

func (p position) Row() int {
	return p.row
}

func (p position) Column() int {
	return p.column
}

type Move interface {
	NumberOfRows() int
	NumberOfColumns() int
}

type move struct {
	numberOfRows    int
	numberOfColumns int
}

func (m move) NumberOfRows() int {
	return m.numberOfColumns
}

func (m move) NumberOfColumns() int {
	return m.numberOfColumns
}

var (
	Up        = move{numberOfRows: -1, numberOfColumns: 0}
	Down      = move{numberOfRows: 1, numberOfColumns: 0}
	Right     = move{numberOfRows: 0, numberOfColumns: 1}
	Left      = move{numberOfRows: 0, numberOfColumns: -1}
	UpRight   = move{numberOfRows: -1, numberOfColumns: 1}
	UpLeft    = move{numberOfRows: -1, numberOfColumns: -1}
	DownRight = move{numberOfRows: 1, numberOfColumns: 1}
	DownLeft  = move{numberOfRows: 1, numberOfColumns: -1}
)

var AllMoves = []Move{
	Up,
	Down,
	Right,
	Left,
	UpRight,
	UpLeft,
	DownRight,
	DownLeft,
}
