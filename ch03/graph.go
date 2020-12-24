package graph

import (
	"fmt"
	"io"
)

// WriteTo dump graph to w
func WriteTo(w io.Writer, adjList [][]int) error {
	size := len(adjList)
	if _, err := fmt.Fprintf(w, "%d", size); err != nil {
		return err
	}
	for i := 0; i < size; i++ {
		adjLen := len(adjList[i])
		if _, err := fmt.Fprintf(w, "\n%d", adjLen); err != nil {
			return err
		}
		for j := 0; j < adjLen; j++ {
			if _, err := fmt.Fprintf(w, " %d", adjList[i][j]); err != nil {
				return err
			}
		}
	}
	if _, err := fmt.Fprintf(w, "\n"); err != nil {
		return err
	}
	return nil
}

// ReadFrom create graph from r
func ReadFrom(r io.Reader, adjList *[][]int) error {
	var size int
	if _, err := fmt.Fscanf(r, "%d", &size); err != nil {
		return err
	}
	*adjList = make([][]int, size)

	for i := 0; i < size; i++ {
		var adjLen int
		if _, err := fmt.Fscanf(r, "\n%d", &adjLen); err != nil {
			return err
		}
		(*adjList)[i] = make([]int, adjLen)

		for j := 0; j < adjLen; j++ {
			if _, err := fmt.Fscanf(r, " %d", &(*adjList)[i][j]); err != nil {
				return err
			}
		}
	}
	return nil
}
