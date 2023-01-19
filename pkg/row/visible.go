package row

func (r *Row) GetVisible() (bool, error) {
	return r.visible, nil
}

func (r *Row) SetVisible(v bool) error {
	r.visible = v
	return nil
}
