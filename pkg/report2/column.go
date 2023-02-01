package report2

type Column struct {
	SourceColumn   string
	Display        string
	Formula        string
	NumberFormat   int
	TransposeRange MonthRange
}

func (col *Column) Key() (key string) {
	if len(col.SourceColumn) > 0 {
		key = col.SourceColumn
	} else if len(col.Display) > 0 {
		key = col.Display
	} else {
		key = ""
	}
	return
}
