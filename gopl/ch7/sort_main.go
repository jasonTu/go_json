package main

import (
	"fmt"
	"sort"
	. "go_json/gopl/ch7/sorting"
)

type customSort struct {
	t []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int {
	return len(x.t)
}

func (x customSort) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}

func (x customSort) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func main() {
	names := []string{"Greate Job", "Good Job", "Well Done"}
	fmt.Println(names)    // [Greate Job Good Job Well Done]
	sort.Sort(StringSlice(names))
	fmt.Println(names)    // [Good Job Greate Job Well Done]

	var tracks = []*Track{
		{"Go", "Delilah", "From the Roots Up", 2012, Length("3m38s")},
		{"Go", "Moby", "Moby", 1992, Length("3m37s")},
		{"Ready 2 Go", "Martin Solveig", "Smash", 2011, Length("4m24s")},
		{"Go Ahead", "Alicia Keys", "As I Am", 2007, Length("4m35s")},
	}
	sort.Sort(ByArtist(tracks))
	PrintTracks(tracks)
	sort.Sort(sort.Reverse(ByArtist(tracks)))
	PrintTracks(tracks)

	// Customized sort function
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	PrintTracks(tracks)
}
