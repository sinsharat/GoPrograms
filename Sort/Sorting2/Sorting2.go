package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

const (
	Title  = "Title"
	Artist = "Artist"
	Album  = "Album"
	Year   = "Year"
	Length = "Length"
	Custom = "Custom"
)

type tracksorting interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
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

func length(duration string) time.Duration {
	d, err := time.ParseDuration(duration)

	if err != nil {
		fmt.Printf("Failed to parse duration for : %v and the err is %v\n", duration, err)
		panic(duration)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 32, 8, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type sortTrack struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x sortTrack) Len() int {
	return len(x.t)
}

func (x sortTrack) Less(i, j int) bool {
	return x.less(x.t[i], x.t[j])
}

func (x sortTrack) Swap(i, j int) {
	x.t[i], x.t[j] = x.t[j], x.t[i]
}

func (x *sortTrack) SortTracks(sortType string, asc bool) {
	switch sortType {
	case Title:
		x.less = func(x, y *Track) bool { return x.Title < y.Title }
	case Album:
		x.less = func(x, y *Track) bool { return x.Album < y.Album }
	case Artist:
		x.less = func(x, y *Track) bool { return x.Artist < y.Artist }
	case Length:
		x.less = func(x, y *Track) bool { return x.Length < y.Length }
	case Year:
		x.less = func(x, y *Track) bool { return x.Year < y.Year }
	default:
		x.less = func(x, y *Track) bool {

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
		}
	}

	x.Sort(asc)
}

func (x *sortTrack) Sort(asc bool) {
	if asc {
		sort.Sort(x)
	} else {
		sort.Sort(sort.Reverse(x))
	}
}

func main() {
	st := sortTrack{t: tracks}

	fmt.Println("Sort by Artist")
	st.SortTracks(Artist, true)
	printTracks(st.t)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Reverse sort by Artist")
	st.SortTracks(Artist, false)
	printTracks(st.t)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Sort by Album")
	st.SortTracks(Album, true)
	printTracks(st.t)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Reverse sort by Album")
	st.SortTracks(Album, false)
	printTracks(st.t)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Sort by Length")
	st.SortTracks(Length, true)
	printTracks(st.t)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Reverse sort by Length")
	st.SortTracks(Length, false)
	printTracks(st.t)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Sort by Year")
	st.SortTracks(Year, true)
	printTracks(st.t)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Reverse sort by Year")
	st.SortTracks(Year, false)
	printTracks(st.t)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Sort by Title")
	st.SortTracks(Title, true)
	printTracks(st.t)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Reverse sort by Title")
	st.SortTracks(Title, false)
	printTracks(st.t)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Custom Sort")
	st.SortTracks(Custom, true)
	printTracks(st.t)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Reverse Custom Sort")
	st.SortTracks(Custom, false)
	printTracks(st.t)
	fmt.Println("----------------------------------------------------------------")
}
