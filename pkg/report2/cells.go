package report2

type Cells struct {
	cells []Cell
}

func (c *Cells) Append(cell Cell) {
	c.cells = append(c.cells, cell)
}

func (c *Cells) All() []Cell {
	return c.cells
}

func (c *Cells) Row(row int) (inRow *Cells) {
	inRow = &Cells{cells: []Cell{}}
	for _, cell := range c.cells {
		if cell.Location.Row == row {
			inRow.Append(cell)
		}
	}
	return
}

func (c *Cells) Col(colName string) (inCol *Cells) {
	inCol = &Cells{cells: []Cell{}}
	if len(c.cells) > 0 {
		col, _, _ := getDataMapHeading(c.cells[0].Location.Sheet, colName)
		for _, cell := range c.cells {
			if cell.Location.Col == col {
				inCol.Append(cell)
			}
		}
	}
	return
}
