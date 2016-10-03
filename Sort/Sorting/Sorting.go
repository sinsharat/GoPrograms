package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

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

type byArtist []*Track

func (x byArtist) Len() int {
	return len(x)
}

func (x byArtist) Less(i, j int) bool {
	return x[i].Artist < x[j].Artist
}

func (x byArtist) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type byAlbum []*Track

func (x byAlbum) Len() int {
	return len(x)
}

func (x byAlbum) Less(i, j int) bool {
	return x[i].Album < x[j].Album
}

func (x byAlbum) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type byTitle []*Track

func (x byTitle) Len() int {
	return len(x)
}

func (x byTitle) Less(i, j int) bool {
	return x[i].Title < x[j].Title
}

func (x byTitle) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type byDuration []*Track

func (x byDuration) Len() int {
	return len(x)
}

func (x byDuration) Less(i, j int) bool {
	return x[i].Length < x[j].Length
}

func (x byDuration) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

type byYear []*Track

func (x byYear) Len() int {
	return len(x)
}

func (x byYear) Less(i, j int) bool {
	return x[i].Year < x[j].Year
}

func (x byYear) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func main() {
	fmt.Println("Sort by Artist")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Reverse sort by Artist")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Sort by Title")
	sort.Sort(byTitle(tracks))
	printTracks(tracks)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Reverse sort by Title")
	sort.Sort(sort.Reverse(byTitle(tracks)))
	printTracks(tracks)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Sort by Album")
	sort.Sort(byAlbum(tracks))
	printTracks(tracks)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Reverse sort by Album")
	sort.Sort(sort.Reverse(byAlbum(tracks)))
	printTracks(tracks)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Sort by Duration")
	sort.Sort(byDuration(tracks))
	printTracks(tracks)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Reverse sort by Duration")
	sort.Sort(sort.Reverse(byDuration(tracks)))
	printTracks(tracks)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Sort by Year")
	sort.Sort(byYear(tracks))
	printTracks(tracks)
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("Reverse sort by Year")
	sort.Sort(sort.Reverse(byYear(tracks)))
	printTracks(tracks)
	fmt.Println("----------------------------------------------------------------")
}
