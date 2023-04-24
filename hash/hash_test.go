package hash

import (
	"testing"
)

func TestMD5(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		wanted string
	}{
		{
			name:   "normal test",
			text:   "hello",
			wanted: "5d41402abc4b2a76b9719d911017c592",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.text); got != tt.wanted {
				t.Errorf("%s MD5() = %v, want %v", tt.name, got, tt.wanted)
			}
		})
	}
}

func TestSHA1(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		wanted string
	}{
		{
			name:   "normal test",
			text:   "hello",
			wanted: "aaf4c61ddcc5e8a2dabede0f3b482cd9aea9434d",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA1(tt.text); got != tt.wanted {
				t.Errorf("%s SHA1() = %v, want %v", tt.name, got, tt.wanted)
			}
		})
	}
}

func TestSHA1Short(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		wanted string
	}{
		{
			name:   "normal test",
			text:   "hello",
			wanted: "aaf4c6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA1Short(tt.text); got != tt.wanted {
				t.Errorf("%s SHA1Short() = %v, want %v", tt.name, got, tt.wanted)
			}
		})
	}
}

func TestSHA256(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		wanted string
	}{
		{
			name:   "normal test",
			text:   "hello",
			wanted: "2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SHA256(tt.text); got != tt.wanted {
				t.Errorf("%s SHA256() = %v, want %v", tt.name, got, tt.wanted)
			}
		})
	}
}

func TestFNV32(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		wanted string
	}{
		{
			name:   "normal test",
			text:   "hello",
			wanted: "4f9f2cab",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FNV32(tt.text); got != tt.wanted {
				t.Errorf("%s FNV32() = %v, want %v", tt.name, got, tt.wanted)
			}
		})
	}
}

func TestFNV64(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		wanted string
	}{
		{
			name:   "normal test",
			text:   "hello",
			wanted: "a430d84680aabd0b",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FNV64(tt.text); got != tt.wanted {
				t.Errorf("%s FNV64() = %v, want %v", tt.name, got, tt.wanted)
			}
		})
	}
}

func TestFNV128(t *testing.T) {
	tests := []struct {
		name   string
		text   string
		wanted string
	}{
		{
			name:   "normal test",
			text:   "hello",
			wanted: "e3e1efd54283d94f7081314b599d31b3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FNV128(tt.text); got != tt.wanted {
				t.Errorf("%s FNV128() = %v, want %v", tt.name, got, tt.wanted)
			}
		})
	}
}
