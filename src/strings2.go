package main

import (
	"sort"
	"strings"
)

type Strings2 struct {
}

func (s *Strings2) Split(ss string, sep string) []string {
	str := strings.Split(ss, sep)
	return str
}

func (s *Strings2) Trim(s2 string) string {
	return strings.TrimSpace(s2)
}

func (s *Strings2) Substring(ss string, start int, length int) string {
	return ss[start:length]
}

func (s *Strings2) SubstringStart(ss string, start int) string {
	if len(ss) == 0 {
		return ""
	}
	return ss[start : len(ss)-start+1]
}

func (s *Strings2) ToUpperCase(ss string) string {
	return strings.ToUpper(ss)
}

func (s *Strings2) ToLowerCase(ss string) string {
	return strings.ToLower(ss)
}

func (s *Strings2) StartsWith(ss string, start string) bool {
	if len(ss) < len(start) {
		return false
	}

	ext := ss[0:len(start)]

	if s.ToUpperCase(ext) == s.ToUpperCase(start) {
		return true
	}
	return false
}

func (s *Strings2) EndsWith(ss string, end string) bool {
	if len(ss) < len(end) {
		return false
	}

	ext := ss[len(ss)-len(end) : len(ss)]
	if s.ToUpperCase(ext) == s.ToUpperCase(end) {
		return true
	}
	return false
}

func (s *Strings2) Contains(ss string, sub string) bool {
	return strings.Contains(ss, sub)
}

func (sx *Strings2) SliceContains(s []string, searchterm string) bool {
	i := sort.SearchStrings(s, searchterm)
	return i < len(s) && s[i] == searchterm
}

func (s *Strings2) InsertAt(array []string, element string, i int) []string {
	return append(array[:i], append([]string{element}, array[i:]...)...)
}
