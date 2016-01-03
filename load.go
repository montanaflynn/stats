package stats

import (
	"fmt"
	"strconv"
)

// LoadRawData parses and converts a slice of mixed data types to floats
func LoadRawData(raw []interface{}) (f Float64Data) {
	for _, v := range raw {
		switch t := v.(type) {
		case int:
			a := float64(t)
			f = append(f, a)
		case uint:
			f = append(f, float64(t))
		case float64:
			f = append(f, t)
		case string:
			fl, err := strconv.ParseFloat(t, 64)
			if err != nil {
				fmt.Println(err)
			} else {
				f = append(f, fl)
			}
		case bool:
			if t == true {
				f = append(f, 1.0)
			} else {
				f = append(f, 0.0)
			}
		}
	}
	return f
}
