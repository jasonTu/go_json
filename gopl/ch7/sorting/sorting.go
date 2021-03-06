package sorting

import (
	"os"
	"fmt"
	"time"
	"text/tabwriter"
)

type StringSlice []string

func (p StringSlice) Len() int {
	return len(p)
}

func (p StringSlice) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type ByArtist []*Track

func (x ByArtist) Len() int {
	return len(x)
}

func (x ByArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x ByArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func PrintTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()    // calculate column widths and print table
}

func Length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

/*
type CustomSort struct {
	t []*Track
	less func(x, y *Track) bool
}

func (x CustomSort) Len() int {
	return len(x.t)
}

func (x CustomSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}

func (x CustomSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}
*/
