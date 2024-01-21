package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"eko", "kurniawan", "Khannedy"})
	_ = writer.Write([]string{"budi", "pratama", "nugrara"})
	_ = writer.Write([]string{"john", "doe", "nugrara"})

	writer.Flush()
}
