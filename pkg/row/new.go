package row

func New(
	index int,
	isHeading bool,
	isVisible bool,
) (r RowInterface, err error) {

	r = &Row{}
	r.SetIndex(index)
	r.SetVisible(isVisible)
	r.SetHeader(isHeading)

	return r, nil
}
