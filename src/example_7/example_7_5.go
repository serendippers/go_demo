package example_7

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

const debug = true

func Test_7_5() {
	example_7_5()
}

func example_7_5() {

	var buf *bytes.Buffer

	if debug {
		buf = new(bytes.Buffer)

	}
	f(buf)

	fmt.Println(buf.String())

	printTracks(tracks)
}

func f(writer io.Writer) {
	fmt.Printf("writer type is %T value is %s is nil : %t", writer, writer, writer == nil)
	if writer != nil {
		writer.Write([]byte("do f fun"))
	}
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")


	sort.Sort(byArtist(tracks))
	//倒序
	sort.Sort(sort.Reverse(byArtist(tracks)))
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap (i, j int) {
	x[i], x[j] = x[j], x[i]
}
