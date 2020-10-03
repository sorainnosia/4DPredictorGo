package main

import "sort"

type kv struct {
	Key   int
	Value int
}

type kvs struct {
	Key   string
	Value int
}

func GetKvsKey(obj []kvs, key string) *kvs {
	var o *kvs
	for _, a := range obj {
		if a.Key == key {
			o = &a
			break
		}
	}
	return o
}

func GetKvKey(obj []kv, key int) *kv {
	var o *kv
	for _, a := range obj {
		if a.Key == key {
			o = &a
			break
		}
	}
	return o
}
func SortIntByValueDesc(mps map[int]int) []kv {
	var ss []kv
	for k, v := range mps {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}

func SortIntByValueAsc(mps map[int]int) []kv {
	var ss []kv
	for k, v := range mps {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	return ss
}

func SortStringByValueDesc(mps map[string]int) []kvs {
	var ss []kvs
	for k, v := range mps {
		ss = append(ss, kvs{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}

func SortStringByValueAsc(mps map[string]int) []kvs {
	var ss []kvs
	for k, v := range mps {
		ss = append(ss, kvs{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value < ss[j].Value
	})

	return ss
}
