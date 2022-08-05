package strx

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	BYTE = 1.0 << (10 * iota)
	KILOBYTE
	MEGABYTE
	GIGABYTE
	TERABYTE
)

// FormatBytes format bytes unit
func FormatBytes(bytes int64) string {
	unit := ""
	value := float32(bytes)

	switch {

	case bytes >= TERABYTE:
		unit = "TB"
		value = value / TERABYTE
	case bytes >= GIGABYTE:
		unit = "GB"
		value = value / GIGABYTE
	case bytes >= MEGABYTE:
		unit = "MB"
		value = value / MEGABYTE
	case bytes >= KILOBYTE:
		unit = "KB"
		value = value / KILOBYTE
	case bytes == 0:
		return "0"

	}

	result := fmt.Sprintf("%.2f", value)
	result = strings.TrimSuffix(result, ".00")
	return fmt.Sprintf("%s%s", result, unit)
}

// Substring source[start:end)
func Substring(source string, start int, end int) string {
	var r = []rune(source)
	length := len(r)

	if start < 0 || end > length || start > end {
		return ""
	}

	if start == 0 && end == length {
		return source
	}

	if start == end {
		return string(r[start])
	}

	var substring = ""
	for i := 0; i < length; i++ {
		if i < start {
			continue
		}
		if i >= end {
			break
		}
		substring += string(r[i])
	}

	return substring
}

// Truncate truncate string
func Truncate(s string, size int) string {
	if len(s) < size {
		return s
	}

	return Substring(s, 0, size)
}

// ToInt cast string to int,default 0
func ToInt(val string) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		return 0
	}
	return i
}
