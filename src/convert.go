package main

import (
	"fmt"
	"strconv"
)

type Convert struct {
}

func (c *Convert) ToInt64(str string) int64 {
	n, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		return n
	}
	return 0
}

func (c *Convert) ToInt32(str string) int {
	n, err := strconv.ParseInt(str, 10, 32)
	if err == nil {
		return int(n)
	}
	return 0
}

func (c *Convert) ToFloat(str string) float32 {
	n, err := strconv.ParseFloat(str, 32)
	if err == nil {
		return float32(n)
	}
	return 0
}

func (c *Convert) ToDouble(str string) float64 {
	n, err := strconv.ParseFloat(str, 64)
	if err == nil {
		return n
	}
	return 0
}

func (c *Convert) IntToString(val int) string {
	str := strconv.FormatInt(int64(val), 10)

	return str
}

func (c *Convert) LongToString(val int64) string {
	str := strconv.FormatInt(val, 10)

	return str
}

func (c *Convert) FloatToString(val float32) string {
	s := fmt.Sprintf("%f", val)
	return s
}

func (c *Convert) DoubleToString(val float64) string {
	s := fmt.Sprintf("%f", val)
	return s
}
