package cmd

// TODO: implement math sort and sort by length structures
type MathSort []string

func (x MathSort) Len() int           { return len(x) }
func (x MathSort) Less(i, j int) bool { return x[i] > x[j] }
func (x MathSort) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type LenSort []string

func (x LenSort) Len() int           { return len(x) }
func (x LenSort) Less(i, j int) bool { return len(x[i]) > len(x[j]) }
func (x LenSort) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }
