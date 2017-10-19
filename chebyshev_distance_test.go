package stats

import (
	"testing"
)

func TestChebyshevDistance(  t *testing.T ) { //data_points_x, data_point_y []float64 ) float64 {
	var (
		err error
		data_point_x []float64
		data_point_y []float64 
	)
	data_point_x = []float64{2, 4, 4, 4, 5, 5, 7, 9}
	data_point_y = []float64{2, 4, 4, 4, 5, 5, 7, 1}
	_, err = ComputeChebyshevDistance( data_point_x, data_point_y )
	if err != nil {
		t.Errorf( "Failed to compute disatance." )
	}
}
