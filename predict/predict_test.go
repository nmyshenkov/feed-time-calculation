package predict

import (
	"feed-time-calculation/common"
	"reflect"
	"sync"
	"testing"
)

func TestGetPredict(t *testing.T) {

	httpPool := common.InitHttpPool(5)

	type args struct {
		clientPool *sync.Pool
		target     Coordinate
		source     []Coordinate
	}
	tests := []struct {
		name    string
		args    args
		want    Response
		wantErr bool
	}{
		{name: "Test 1", args: args{
			clientPool: httpPool,
			target:     Coordinate{Lat: 55.752992, Lng: 37.618333},
			source: []Coordinate{
				{Lat: 55.74837156167371, Lng: 37.61180107665421},
				{Lat: 55.7532706, Lng: 37.6076902},
			},
		}, want: Response{1, 1}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetPredict(tt.args.clientPool, tt.args.target, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPredict() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPredict() got = %v, want %v", got, tt.want)
			}
		})
	}
}
