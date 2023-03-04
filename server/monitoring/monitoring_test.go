package monitoring

import (
	"reflect"
	"testing"

	pb "github.com/Xacor/go-sysmon/proto"
)

func TestLoadAvg(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.LoadAverage
		wantErr bool
	}{
		{"TestOK", args{src: "../testData/loadavg"}, &pb.LoadAverage{Load1: 1.29, Load5: 1.27, Load15: 0.94}, false},
		{"TestFileNotFound", args{src: "/qwerty"}, nil, true},
		{"TestParseFloat", args{src: "../testData/badLoadavg"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := LoadAvg(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadAvg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadAvg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcStat(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name    string
		args    args
		want    *pb.ProcStat
		wantErr bool
	}{
		{"TestOK", args{src: "../testData/stat"}, &pb.ProcStat{Us: 1.6, Sy: 0.5, Id: 97.6}, false},
		{"TestFileNotFound", args{src: "/qwerty"}, nil, true},
		{"TestParseUintError", args{src: "../testData/badStat"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ProcStat(tt.args.src)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProcStat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProcStat() = %v, want %v", got, tt.want)
			}
		})
	}
}
