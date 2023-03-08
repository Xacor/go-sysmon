package monitoring

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"

	pb "github.com/Xacor/go-sysmon/proto"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/structpb"
)

// LoadAvg parses values from src (usually /proc/loadavg) into map[string]interface{} and returns it.
func LoadAvg(src string) (map[string]any, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	fileScanner := bufio.NewScanner(file)

	fileScanner.Split(bufio.ScanWords)

	// Interested only in first 3 values stored in file
	buf := make([]float32, 0, 3)
	for i := 0; i < 3; i++ {
		fileScanner.Scan()

		val, err := strconv.ParseFloat(fileScanner.Text(), 32)
		if err != nil {
			return nil, err
		}

		buf = append(buf, float32(val))
	}

	file.Close()
	return map[string]any{
		"load1":  buf[0],
		"load5":  buf[1],
		"load15": buf[2],
	}, nil
}

func MarshallLoadAvg(avg map[string]any) (*pb.LoadAverage, error) {

	m, err := structpb.NewStruct(avg)
	if err != nil {
		return nil, fmt.Errorf("Error in marshallLoadAverage: %w", err)
	}

	bytes, err := m.MarshalJSON()
	if err != nil {
		return nil, fmt.Errorf("Error in marshallLoadAverage: %w", err)
	}

	avgpb := pb.LoadAverage{}
	err = protojson.Unmarshal(bytes, &avgpb)
	if err != nil {
		return nil, fmt.Errorf("Error in marshallLoadAverage: %w", err)
	}
	return &avgpb, nil
}

// ProcStat parses values from src (usually /proc/stat) and returns protobuf ProcStat,
// containing percentage for each cpu state.
// For more about cpu states see https://www.kernel.org/doc/Documentation/filesystems/proc.txt.
func ProcStat(src string) (map[string]any, error) {
	file, err := os.Open(src)
	if err != nil {
		return nil, err
	}

	// Setting up scanner
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanWords)

	buf := make([]uint32, 0, 7)
	var sum uint32

	// first column is "cpu", so skip
	fileScanner.Scan()

	for i := 0; i < 7; i++ {
		fileScanner.Scan()

		val, err := strconv.ParseUint(fileScanner.Text(), 10, 32)
		if err != nil {
			return nil, err
		}

		sum += uint32(val)
		buf = append(buf, uint32(val))
	}

	file.Close()
	var us float64 = 100 * (float64(buf[0]+buf[1]) / float64(sum))        // %(user+nice)
	var sy float64 = 100 * (float64(buf[2]+buf[5]+buf[6]) / float64(sum)) // %(system+hi+st)
	var id float64 = 100 * (float64(buf[3]) / float64(sum))               // %idle

	// It is ok if sum of all fields != 100%, because io-wait from /proc/stat is not reliable (see docs),
	// and Floor is used for representation
	return map[string]any{
		"us": float32(math.Floor(us*10) / 10),
		"sy": float32(math.Floor(sy*10) / 10),
		"id": float32(math.Floor(id*10) / 10),
	}, nil
}
