package stats

import "testing"

var allTestData = []struct {
	actual   []interface{}
	expected Float64Data
}{
	{
		[]interface{}{1.0, "2", 3.0, 4, "4.0", 5},
		Float64Data{1.0, 2.0, 3.0, 4.0, 4.0, 5.0},
	},
	{
		[]interface{}{"345", "223", "654.4", "194"},
		Float64Data{345.0, 223.0, 654.4, 194.0},
	},
	{
		[]interface{}{7862, 4234, 9872.1, 8794},
		Float64Data{7862.0, 4234.0, 9872.1, 8794.0},
	},
	{
		[]interface{}{true, false, true, false, false},
		Float64Data{1.0, 0.0, 1.0, 0.0, 0.0},
	},
	{
		[]interface{}{14.3, 26, 17.7, "shoe"},
		Float64Data{14.3, 26.0, 17.7},
	},
}

func equal(actual, expected Float64Data) bool {
	if len(actual) != len(expected) {
		return false
	}
	for k, actualVal := range actual {
		if actualVal != expected[k] {
			return false
		}
	}
	return true
}

func TestLoadRawData(t *testing.T) {
	for _, data := range allTestData {
		actual := LoadRawData(data.actual)
		if !equal(actual, data.expected) {
			t.Fatalf("Transform(%v). Expected [%v], Actual [%v]", data.actual, data.expected, actual)
		}
	}
}
