package row

func (r *Row) GetIndex() (int, error) {
	return r.index, nil
}

func (r *Row) SetIndex(i int) error {
	r.index = i
	return nil
}
