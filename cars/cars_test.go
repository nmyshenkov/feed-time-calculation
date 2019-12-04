package cars

import (
	"feed-time-calculation/common"
	"reflect"
	"sync"
	"testing"
)

func TestGetCars(t *testing.T) {

	httpPool := common.InitHttpPool(5)

	type args struct {
		clientPool *sync.Pool
		coord      Coordinate
	}
	tests := []struct {
		name    string
		args    args
		want    []Cars
		wantErr bool
	}{
		{name: "Test 1", args: args{
			clientPool: httpPool,
			coord: Coordinate{
				Lat:   55.752992,
				Lng:   37.618333,
				Limit: 3,
			},
		}, want: []Cars{{166, 55.7575429, 37.6135117}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCars(tt.args.clientPool, tt.args.coord)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCars() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCars() got = %v, want %v", got, tt.want)
			}
		})
	}
}
