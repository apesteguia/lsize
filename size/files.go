package size

import (
	"fmt"
	"os"
	"sort"
)

type Files struct {
	F           []File
	Self        File
	N           int
	LongestName int
}

func (f *Files) Init(path string) error {
	err := f.Self.GetSize(path)
	f.N = 0
	f.LongestName = 0
	if err != nil {
		return err
	}

	return nil
}

func (f *Files) GetSizes() error {
	dir, err := os.Open(f.Self.Path)
	if err != nil {
		return err
	}
	defer dir.Close()

	infos, err := dir.Readdir(-1)
	if err != nil {
		return err
	}

	f.F = make([]File, len(infos))

	for i, info := range infos {
		if info.IsDir() {
			f.F[i].File = false
		} else {
			f.F[i].File = true
		}

		if len(info.Name()) > f.LongestName {
			f.LongestName = len(info.Name())
		}

		f.F[i].GetSize(f.Self.Path + info.Name())
		f.N += 1
	}

	f.Sort()

	return nil
}

func (f *Files) List(order bool) {
	fmt.Print("\033[A\033[K")
	separators := make([]byte, f.LongestName+19)
	for i := range separators {
		separators[i] = '='
	}

	fmt.Println(f.Self.Path)
	fmt.Println(string(separators))

	if order {
		for _, v := range f.F {
			fileType := "[File]  "
			if !v.File {
				fileType = "[Folder]"
			}
			// realSizeStr := fmt.Sprintf("%-*s", f.LongestName, v.RealSize)
			fmt.Printf("%s  %-*s  (%s)\n", fileType, f.LongestName, v.Name, v.RealSize)
		}
	} else {
		for i := len(f.F) - 1; i >= 0; i-- {
			v := f.F[i]
			realSizeStr := fmt.Sprintf("%-*s", f.LongestName, v.RealSize)
			fmt.Printf("%-*s %s\n", f.LongestName, v.Name, realSizeStr)
		}
	}

	fmt.Println(string(separators))
	fmt.Println("Total: ", f.Self.RealSize)
}

func (f *Files) ListRaw(order bool) {
	if order {
		for _, v := range f.F {
			fmt.Println(v)
		}
	} else {
		for i := len(f.F) - 1; i >= 0; i-- {
			fmt.Println(f.F[i])
		}
	}
}

func (f *Files) Sort() {
	sort.Slice(f.F, func(i, j int) bool {
		return f.F[i].Size > f.F[j].Size
	})
}