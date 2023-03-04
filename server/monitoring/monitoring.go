package monitoring

import (
	"bufio"
	"os"
	"strconv"

	pb "github.com/Xacor/go-sysmon/proto"
)

func LoadAvg(src string) (*pb.LoadAverage, error) {
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
