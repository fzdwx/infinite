package strx

import (
	"testing"
)

func TestFormatBytes(t *testing.T) {
	type args struct {
		bytes int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test zero", args{bytes: 0}, "0"},
		{"test kb", args{bytes: 1024}, "1KB"},
		{"test mb", args{bytes: 1024 * 1024}, "1MB"},
		{"test gb", args{bytes: 1024 * 1024 * 1024}, "1GB"},
		{"test tb", args{bytes: 1024 * 1024 * 1024 * 1024}, "1TB"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatBytes(tt.args.bytes); got != tt.want {
				t.Errorf("FormatBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubstring(t *testing.T) {
	type args struct {
		source string
		start  int
		end    int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"sub 1", args{"01234567", 1, 3}, "12"},
		{"sub 2", args{"hello world", 0, 100}, ""},
		{"sub 3", args{"hello world", 0, 10}, "hello worl"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Substring(tt.args.source, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("Substring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTruncate(t *testing.T) {
	type args struct {
		s    string
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{"123", 1}, "1"},
		{"t2", args{"123", 10}, "123"},
		{"t3", args{"123", 2}, "12"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Truncate(tt.args.s, tt.args.size); got != tt.want {
				t.Errorf("Truncate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToInt(t *testing.T) {
	type args struct {
		val string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"cast string err", args{val: "asdsad"}, 0},
		{"cast negative number", args{val: "-10"}, -10},
		{"cast number", args{val: "10"}, 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToInt(tt.args.val); got != tt.want {
				t.Errorf("ToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
