package report

func (c *Column) Key() (k string) {
	if len(c.MapKey) > 0 {
		k = c.MapKey
	} else if len(c.Display) > 0 {
		k = c.Display
	} else {
		k = ""
	}
	return
}
