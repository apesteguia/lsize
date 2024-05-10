package size

import (
	"fmt"
	"math"
	"os/exec"
	"strconv"
	"strings"
)

type File struct {
	Path     string
	Name     string
	Size     float64
	RealSize string
	File     bool
}

var sizes = []byte{'B', 'K', 'M', 'G', 'T'}

func (f *File) GetSize(path string) error {
	f.Path = path
	var cmd *exec.Cmd

	if path != "/" {
		splitted := strings.Split(path, "/")
		f.Name = splitted[len(splitted)-1]
		cmd = exec.Command("du", "-sh", f.Path)
	} else {
		cmd = exec.Command("sudo", "du", "-sh", f.Path)
	}

	stdout, err := cmd.Output()
	if err != nil {
		return err
	}

	ssize, size, _ := realSize(string(stdout))
	f.Size = size
	f.RealSize = ssize

	return nil
}

func realSize(str string) (string, float64, error) {
	splitter := strings.Fields(str)
	sizeStr := strings.Replace(splitter[0], ",", ".", -1)
	bsize := sizeStr[len(sizeStr)-1]

	reals := sizeStr
	sizeStr = sizeStr[:len(sizeStr)-1]

	v, err := strconv.ParseFloat(sizeStr, 64)
	if err != nil {
		return "0.0K", 0, err
	}

	for i, s := range sizes {
		if bsize == s {
			v *= math.Pow(1024, float64(i+1))
			return reals, v, nil
		}
	}

	return "No data", 0, fmt.Errorf("unidad de tamaño desconocida: %c", bsize)
}
