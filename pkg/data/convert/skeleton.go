package convert

func Skeleton(
	dataset []map[string]string,
	groupColumns []string,
	dateCostColumns map[string]string,
	otherColumns []string,
) (skel map[string]map[string][]string) {
	skel = make(map[string]map[string][]string)

	// now the dateCostColumns are expanded for their values
	expanded := make(map[string]bool)
	for _, ds := range dataset {
		for src, _ := range dateCostColumns {
			if v, ok := ds[src]; ok {
				expanded[v] = true
			}
		}
	}

	// using the groupcolumns, find all possible groupby values for
	// all the dataset
	for _, ds := range dataset {
		key := Key(ds, groupColumns)
		skel[key] = make(map[string][]string)
		// push in the groupby fields
		for _, f := range groupColumns {
			skel[key][f] = []string{}
		}
		// push the expanded date ones
		for d, _ := range expanded {
			skel[key][d] = []string{}
		}
		// push the other columns
		for _, o := range otherColumns {
			skel[key][o] = []string{}
		}
	}

	return
}
