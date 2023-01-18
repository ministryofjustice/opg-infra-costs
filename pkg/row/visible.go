package row

func (r *Row) GetVisible() bool {
	return r.visible
}

func (r *Row) SetVisible(v bool) {
	r.visible = v
}
