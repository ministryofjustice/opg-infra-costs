package convert

func Convert(
	dataset []map[string]string,
	groupColumns []string,
	transposeColumns map[string]string,
	otherColumns []string,
) (converted map[string]map[string][]string) {

	// get the skeleton data set with keys and columns set for each
	converted = Skeleton(dataset, groupColumns, transposeColumns, otherColumns)
	for _, data := range dataset {
		key := Key(data, groupColumns)
		// now we look at the values and add them in
		for _, f := range groupColumns {
			if v, ok := data[f]; ok {
				converted[key][f] = append(converted[key][f], v)
			}
		}
		// now look at expanded columns that need to be fetched
		for labelCol, valueCol := range transposeColumns {
			field := data[labelCol]
			if val, ok := data[valueCol]; ok {
				converted[key][field] = append(converted[key][field], val)
			} else {
				converted[key][field] = append(converted[key][field], "0.00")
			}
		}
		// now we look at the values and add them in
		for _, f := range otherColumns {
			if v, ok := data[f]; ok {
				converted[key][f] = append(converted[key][f], v)
			}
		}
	}
	return
}
