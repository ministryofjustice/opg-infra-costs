package row

func (r *Row) GetHeader() bool {
	return r.header
}

func (r *Row) SetHeader(h bool) {
	r.header = h
}
