package main

var AlgoSeparator = "   "

type DataSet struct {
}

func (ds *DataSet) GetDataSourcePath() string {
	d := &Directory{}
	p := &Path{}

	path := p.Combine(d.GetCurrentDirectory(), "4DSource")
	if d.Exists(path) == false {
		d.CreateDirectory(path)
	}
	return path
}

func (ds *DataSet) GetDataSetPath(file string) string {
	p := &Path{}
	d := &Directory{}
	path := p.Combine(d.GetCurrentDirectory(), "DataSet")
	if file == "" {
		return path
	}
	if d.Exists(path) == false {
		d.CreateDirectory(path)
	}
	return p.Combine(path, file)
}

func (ds *DataSet) GetDrawFile(i int) string {
	p := &Path{}
	c := &Convert{}
	file := p.Combine(ds.GetDataSourcePath(), "4D-Result-"+c.IntToString(i)+".txt")
	return file
}

func (ds *DataSet) GetLastDrawNo() int {
	var i int = 1

	f := &File2{}
	for {
		file := ds.GetDrawFile(i)
		if f.Exists(file) == false {
			break
		}
		i++
	}
	return i - 1
}

func (ds *DataSet) Trim(ss []string) []string {
	s := &Strings2{}
	var result []string
	for _, str := range ss {
		if str == "" {
			continue
		}
		result = append(result, s.Trim(str))
	}
	return result
}

func (ds *DataSet) GetNoByDrawNo(i int) []string {
	f := &File2{}
	file := ds.GetDrawFile(i)
	if f.Exists(file) == false {
		return nil
	}
	str := f.ReadAllText(file)

	s := &Strings2{}
	nos := s.Split(str, "\n")
	return ds.Trim(nos)
}

func (ds *DataSet) GetAllNos() []string {
	var result []string
	lastno := ds.GetLastDrawNo()
	for i := 1; i <= lastno; i++ {
		temp := ds.GetNoByDrawNo(i)

		result = append(result, temp...)
	}
	return result
}

func (ds *DataSet) WriteDataSet(lastdraw string, header string, result []string, name string) {
	f := &File2{}
	fname := ds.GetDataSetPath(lastdraw + "_" + name + ".txt")
	if lastdraw == "" {
		fname = ds.GetDataSetPath(name + ".txt")
	}
	if f.Exists(fname) {
		f.Delete(fname)
	}
	result = append([]string{header}, result...)
	f.WriteAllLines(fname, result)
}

func (ds *DataSet) GetDataSetName() []string {
	d := &Directory{}
	s := &Strings2{}
	files := d.GetFiles(ds.GetDataSetPath(""))
	var result []string
	for _, file := range files {
		if s.EndsWith(file, ".txt") {
			result = append(result, d.GetFilenameWithoutExtension(d.GetFilenameOnly(file)))
		}
	}
	return result
}

func (ds *DataSet) GetDataSetNo(dsname string) []string {
	f := &File2{}
	s := &Strings2{}

	file := ds.GetDataSetPath(dsname + ".txt")
	if f.Exists(file) == false {
		return nil
	}
	str := f.ReadAllText(file)
	lines := ds.Trim(s.Split(str, "\n"))
	return lines[1:len(lines)]
}

func (ds *DataSet) GetDataSetHeader(dsname string) string {
	strs := ds.GetDataSetNo(dsname)
	if len(strs) > 0 {
		return strs[0]
	}
	return ""
}

func (ds *DataSet) CreateDataSet(top3 bool, starter bool, consolation bool) []string {
	d := &Directory{}
	if d.Exists(ds.GetDataSourcePath()) == false {
		return nil
	}
	var result []string
	lastno := ds.GetLastDrawNo()
	for i := 1; i < lastno; i++ {
		temp := ds.GetNoByDrawNo(i)
		if temp == nil {
			continue
		}
		if top3 && len(temp) >= 3 {
			result = append(result, temp[0:3]...)
		} else if top3 {
			result = append(result, temp[0:len(temp)]...)
		}
		if starter && len(temp) >= 13 {
			result = append(result, temp[3:13]...)
		} else if starter {
			result = append(result, temp[3:len(temp)]...)
		}
		if consolation && len(temp) >= 23 {
			result = append(result, temp[13:len(temp)]...)
		}
	}
	return result
}

func (ds *DataSet) GetNoByRaw1(str string) []string {
	s := &Strings2{}
	return ds.Trim(s.Split(str, "\n"))
}

func (ds *DataSet) GetNoByRaw2(str string) []string {
	s := &Strings2{}
	ss := ds.Trim(s.Split(str, "\n"))
	var result []string
	for _, s2 := range ss {
		sss := s.Split(s2, AlgoSeparator)
		result = append(result, sss[0])
	}
	return result
}

func (ds *DataSet) GetNoByRaw2Filter(str string, filter string) []string {
	s := &Strings2{}
	ss := ds.Trim(s.Split(str, "\n"))
	var result []string
	for _, s2 := range ss {
		sss := s.Split(s2, AlgoSeparator)
		if sss[0] == "" {
			continue
		}
		if filter == "" {
			result = append(result, sss[0])
		} else if len(sss) >= 2 && s.ToUpperCase(sss[1]) == s.ToUpperCase(filter) {
			result = append(result, sss[0])
		}
	}
	return result
}

func (ds *DataSet) GetNoByRaw2NoFilter(str string, filter string) []string {
	s := &Strings2{}
	ss := ds.Trim(s.Split(str, "\n"))
	var result []string
	for _, s2 := range ss {
		sss := s.Split(s2, AlgoSeparator)
		if sss[0] == "" {
			continue
		}
		if filter == "" || s.ToUpperCase(sss[0]) == s.ToUpperCase(filter) {
			result = append(result, sss[0])
		} else if len(filter) == 4 && len(sss[0]) == 4 {
			match := true
			for i := 0; i < len(sss[0]); i++ {
				if s.ToUpperCase(string(sss[0][i])) != s.ToLowerCase(string(filter[i])) && filter[i] != '?' {
					match = false
					break
				}
			}
			if match {
				result = append(result, sss[0])
			}
		}
	}
	return result
}

func (ds *DataSet) GetNoByRaw2NoContainsFilter(str string, filter string) []string {
	s := &Strings2{}
	ss := ds.Trim(s.Split(str, "\n"))
	var result []string
	for _, s2 := range ss {
		sss := s.Split(s2, AlgoSeparator)
		if sss[0] == "" {
			continue
		}
		if len(sss) < 2 {
			continue
		}

		if filter == "" || s.ToUpperCase(sss[0]) == s.ToUpperCase(filter) {
			result = append(result, sss[0])
		} else {
			if s.Contains(s.ToUpperCase(sss[0]), s.ToUpperCase(filter)) {
				result = append(result, sss[0])
			}
		}
	}
	return result
}

func (ds *DataSet) GetNoByRaw2AlgoFilter(str string, filter string) []string {
	s := &Strings2{}
	ss := ds.Trim(s.Split(str, "\n"))
	var result []string
	for _, s2 := range ss {
		sss := s.Split(s2, AlgoSeparator)
		if sss[0] == "" {
			continue
		}
		if filter == "" || s.ToUpperCase(sss[1]) == s.ToUpperCase(filter) {
			result = append(result, sss[0])
		} else if len(filter) == 4 && len(sss[1]) == 4 {
			match := true
			for i := 0; i < len(sss[1]); i++ {
				if s.ToUpperCase(string(sss[1][i])) != s.ToUpperCase(string(filter[i])) && filter[i] != '?' {
					match = false
					break
				}
			}
			if match {
				result = append(result, sss[0])
			}
		}
	}
	return result
}

func (ds *DataSet) GetRaw3ByAlgoCount(m map[string]int) string {
	c := &Convert{}
	result := ""
	for key := range m {
		result = result + key + AlgoSeparator + c.IntToString(m[key]) + "\r\n"
	}
	return result
}

func (ds *DataSet) GetAlgoCountByRaw2(raw string) []kvs {
	result := make(map[string]int)
	s := &Strings2{}
	ss := s.Split(raw, "\n")
	for _, sx := range ss {
		s2 := s.Split(sx, AlgoSeparator)
		if len(s2) < 2 {
			continue
		}
		value := s.Trim(s2[1])
		val, ok := result[value]
		if ok {
			result[value] = val + 1
		} else if !ok {
			result[value] = 1
		}
	}

	kvs := SortStringByValueAsc(result)
	return kvs
}

func (ds *DataSet) GetNoByRaw2StatisticFilter(r string, filter string) []string {
	stats := ds.GetAlgoCountByRaw2(r)

	c := &Convert{}
	str := ds.GetNoByRaw2(r)
	var result []string
	for i := 0; i < len(str); i++ {
		key := str[i]
		ok := GetKvsKey(stats, key)
		if ok == nil {
			continue
		}
		appear := GetKvsKey(stats, key).Value
		if c.IntToString(appear) == filter {
			result = append(result, key)
		}
	}
	return result
}

func (ds *DataSet) GetNoByRaw2AlgoContainsFilter(str string, filter string) []string {
	s := &Strings2{}
	ss := ds.Trim(s.Split(str, "\n"))
	var result []string
	for _, s2 := range ss {
		sss := s.Split(s2, AlgoSeparator)
		if sss[0] == "" {
			continue
		}
		if filter == "" || sss[1] == filter {
			result = append(result, sss[0])
		} else {
			if s.Contains(s.ToUpperCase(sss[1]), s.ToUpperCase(filter)) {
				result = append(result, sss[0])
			}
		}
	}
	return result
}

func (ds *DataSet) GetBallStats(lines []string) []kv {
	c := &Convert{}
	m := make(map[int]int)
	for _, s := range lines {
		for i := 0; i < len(s); i++ {
			key := c.ToInt32(string(s[i]))
			val, ok := m[key]
			if !ok {
				m[key] = 1
			} else {
				m[key] = val + 1
			}
		}
	}

	kv := SortIntByValueAsc(m)
	return kv
}

func (ds *DataSet) RemoveRepeat(lines []string) []string {
	s := &Strings2{}
	var result []string
	for _, sx := range lines {
		if s.SliceContains(result, sx) == false {
			result = append(result, sx)
		}
	}
	return result
}
