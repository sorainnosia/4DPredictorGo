package main

type Plugins interface {
	GetNo(string) string
}

var occurencePlugins *OccurencePlugins

func GetPluginsNo(plugin string, str string) string {
	if plugin == "Identity" {
		i := &IdentityPlugins{}
		return i.GetNo(str)
	} else if plugin == "ABCD" {
		i := &ABCDPlugins{}
		return i.GetNo(str)
	} else if plugin == "Sum" {
		i := &SumPlugins{}
		return i.GetNo(str)
	} else if plugin == "Occurence" {
		if occurencePlugins == nil {
			occurencePlugins = &OccurencePlugins{}
			occurencePlugins.Init()
		}
		return occurencePlugins.GetNo(str)
	}
	return ""
}

type SumPlugins struct {
}

func (s *SumPlugins) GetNo(str string) string {
	var result int
	c := &Convert{}
	for i := 0; i < len(str); i++ {
		temp := c.ToInt32(string(str[i]))
		result += temp
	}
	return c.IntToString(result)
}

type ABCDPlugins struct {
}

func (a *ABCDPlugins) GetNo(str string) string {
	start := 64
	m := make(map[string]string)
	var result string
	for i := 0; i < len(str); i++ {
		val, ok := m[string(str[i])]
		if !ok {
			start++
			m[string(str[i])] = string(start)
			val = string(start)
		}
		result = result + val
	}
	return result
}

type IdentityPlugins struct {
}

func (s *IdentityPlugins) GetNo(str string) string {
	return str
}

type OccurencePlugins struct {
	NotoResult map[string]int
}

func (o *OccurencePlugins) Init() {
	o.NotoResult = make(map[string]int)
	lastno := o.GetLastDrawNo()
	for i := 1; i <= lastno; i++ {
		temp := o.GetNoByDrawNo(i)
		for _, t := range temp {
			val, ok := o.NotoResult[t]
			if !ok {
				o.NotoResult[t] = 1
			} else {
				o.NotoResult[t] = val + 1
			}
		}
	}
}

func (o *OccurencePlugins) GetDrawFile(i int) string {
	p := &Path{}
	c := &Convert{}
	ds := &DataSet{}
	file := p.Combine(ds.GetDataSourcePath(), c.IntToString(i)+".txt")
	return file
}

func (o *OccurencePlugins) GetLastDrawNo() int {
	var i int = 1

	f := &File2{}
	ds := &DataSet{}
	for {
		file := ds.GetDrawFile(i)
		if f.Exists(file) == false {
			break
		}
		i++
	}
	return i
}

func (o *OccurencePlugins) GetNoByDrawNo(i int) []string {
	f := &File2{}
	ds := &DataSet{}
	file := ds.GetDrawFile(i)
	if f.Exists(file) == false {
		return nil
	}
	str := f.ReadAllText(file)

	s := &Strings2{}
	nos := s.Split(str, "\n")
	return ds.Trim(nos)
}

func (o *OccurencePlugins) GetNo(str string) string {
	if len(o.NotoResult) == 0 {
		o.Init()
	}
	c := &Convert{}
	_, ok := o.NotoResult[str]
	if !ok {
		return "0"
	}
	return c.IntToString(o.NotoResult[str])
}
