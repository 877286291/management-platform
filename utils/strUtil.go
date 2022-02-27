package utils

import (
	"bytes"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		_, err := b.Write(val)
		if err != nil {
			return nil
		}
	case rune:
		_, err := b.WriteRune(val)
		if err != nil {
			return nil
		}
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Out of memory!")
		}
	}()
	_, err := b.WriteString(s)
	if err != nil {
		return nil
	}
	return b
}
func Strlist2str(list []string) string {
	var str string
	for e := range list {
		str = str + list[e] + ","
	}
	strings.TrimSuffix(str, ",")
	return str
}

func Intlist2str(list []int32) string {
	var str string
	for e := range list {
		str = str + strconv.Itoa(int(list[e])) + ","
	}
	strings.TrimSuffix(str, ",")
	return str
}

func Intlist2strlist(list []string) []int32 {
	var strlist []int32
	for e := range list {
		temp, _ := strconv.Atoi(list[e])
		strlist = append(strlist, int32(temp))
	}
	return strlist
}
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}
