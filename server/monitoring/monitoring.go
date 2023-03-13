package monitoring

import (
	"bufio"
	"math"
	"os"
	"strconv"

	pb "github.com/Xacor/go-sysmon/proto"
)

// LoadAvg parses values from src (usually /proc/loadavg) into protobuf LoadAverage and return it.
func GetLoadAvg(src string) (*pb.LoadAverage, error) {
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
	return &pb.LoadAverage{Load1: buf[0], Load5: buf[1], Load15: buf[2]}, nil
}

// ProcStat parses values from src (usually /proc/stat) and returns protobuf ProcStat,
// containing persantage for each cpu state.
// For more about cpu states see https://www.kernel.org/doc/Documentation/filesystems/proc.txt.
func GetProcStat(src string) (*pb.ProcStat, error) {
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
	return &pb.ProcStat{
		Us: float32(math.Floor(us*10) / 10),
		Sy: float32(math.Floor(sy*10) / 10),
		Id: float32(math.Floor(id*10) / 10),
	}, nil
}
