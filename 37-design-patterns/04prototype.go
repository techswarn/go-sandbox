package main
import "fmt"

type iCloneable interface {
    print()
    clone() iCloneable
}

type file struct {
	name  string
    // Imagine a file has a whole load of other complex fields
    // This sample is only demonstrating how to implement prototype pattern
    ext   string
    owner string
    // ... and so on and so forth
}

func newFile() iCloneable {
    return &file{
        name: "selfie",
        ext: "jpg",
        owner: "root",
    }
}

func (f *file) print() {
	fmt.Println(f.name + "." + f.ext + " - " + f.owner)
}

func (f *file) clone() iCloneable {
	return &file{
        name: f.name + "_clone",
        ext: f.ext,
        owner: f.owner,
    }
}

func main() {
    file1 := newFile()
    file1.print()
    file2 := file1.clone()
    file2.print()
}