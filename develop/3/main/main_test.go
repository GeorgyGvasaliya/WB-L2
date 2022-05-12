package main

import (
	"L2/develop/3/config"
	"bufio"
	"os"
	"testing"
)

func TestTask(t *testing.T) {
	testTable := []struct {
		input    config.Config
		expected []string
	}{
		{
			input: config.Config{
				K:        0,
				Filename: "",
				N:        false,
				R:        false,
				U:        false,
			},
			expected: []string{"44477ddeeeeffffhhjjjkkkksuwxyyyz", "44477ddeeeeffffhhjjjkkkksuwxyyyz",
				"444777eefiijklrttttuxyyyyz", "ads", "ads", "ads", "aaaaddddssss", "3addddffffgghss"},
		},
		{
			input: config.Config{
				K:        1,
				Filename: "",
				N:        false,
				R:        false,
				U:        false,
			},
			expected: []string{"44477ddeeeeffffhhjjjkkkksuwxyyyz", "kufdehjwyskf4ehde7fyejzk474fxjyk",
				"ui4jxytrfklezi7yt47tye47ty", "asd", "asd", "dsa", "asdasdasdasd", "asfdgdfgfdhsdf3"},
		},
		{
			input: config.Config{
				K:        4,
				Filename: "",
				N:        false,
				R:        false,
				U:        false,
			},
			expected: []string{"kufdehjwyskf4ehde7fyejzk474fxjyk", "kufdehjwyskf4ehde7fyejzk474fxjyk",
				"ui4jxytrfklezi7yt47tye47ty", "asd", "asd", "dsa", "asdasdasdasd", "asfdgdfgfdhsdf3"},
		},
		{
			input: config.Config{
				K:        0,
				Filename: "",
				N:        true,
				R:        false,
				U:        false,
			},
			expected: []string{"44477", "44477", "444777", "3"},
		},
		{
			input: config.Config{
				K:        0,
				Filename: "",
				N:        false,
				R:        true,
				U:        false,
			},
			expected: []string{"zyyyxwuskkkkjjjhhffffeeeedd77444", "zyyyxwuskkkkjjjhhffffeeeedd77444",
				"zyyyyxuttttrlkjiifee777444", "sda", "sda", "sda", "ssssddddaaaa", "sshggffffdddda3"},
		},
		{
			input: config.Config{
				K:        0,
				Filename: "",
				N:        false,
				R:        false,
				U:        true,
			},
			expected: []string{"44477ddeeeeffffhhjjjkkkksuwxyyyz", "444777eefiijklrttttuxyyyyz", "ads", "aaaaddddssss",
				"3addddffffgghss"},
		},
	}

	f, err := os.Open("file.txt")
	if err != nil {
		t.Errorf("error with open file")
	}
	text := make([]string, 0)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		t.Errorf("error with scanner")
	}
	for _, tt := range testTable {
		got := task(text, tt.input)
		if Equal(got, tt.expected) {
			t.Errorf("Expected %s, got %s", tt.expected, got)
		}
	}
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
