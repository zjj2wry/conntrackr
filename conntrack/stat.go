package conntrack

import (
	"bufio"
	"os"
	"strings"
)

// Gets info from /proc/net/stat/nf_conntrack
func Stat(fname string) (*StatResultList, error) {
	fd, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fd.Close()

	result := NewStatResultList()

	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		sr, err := NewStatResultFromFields(i, fields)
		if err != nil && !strings.Contains(err.Error(), "Probably a header") {
			return nil, err
		}

		if sr != nil {
			result.Append(sr)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
