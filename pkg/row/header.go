package row

func (r *Row) GetHeader() (bool, error) {
	return r.header, nil
}

func (r *Row) SetHeader(h bool) error {
	r.header = h
	return nil
}
