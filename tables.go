package tables

import (
	"fmt"
	"math"
	"strings"
)

//type Row interface{}
type Table []interface{}

type MaxSizes [][]int

type RowSizes struct {
	maxIndex int
	numCols  int
	rowWidth int
	colWidth []int
}
type TableSizes struct {
	numRows     int
	maxWidth    int
	rows        []RowSizes
	maxSizes    MaxSizes
	maxRowWidth []int
}

func (table Table) Size() int {
	return len(table)
}

func (rows Table) Sizes() (sizes TableSizes) {
	sizes.numRows = len(rows)
	sizes.rows = make([]RowSizes, sizes.numRows)
	sizes.maxSizes = make([][]int, sizes.numRows)
	sizes.maxRowWidth = make([]int, sizes.numRows)
	maxIndex := -1
	prevCols := -1
	cols := 0
	for i, row := range rows {
		width := 0
		switch row := row.(type) {
		case Table:
			sizes.rows[i].numCols = len(row)
			sizes.rows[i].colWidth = make([]int, sizes.rows[i].numCols)
			cols = len(row)
			if cols != prevCols {
				maxIndex++
				sizes.maxSizes[maxIndex] = make([]int, len(row))
				sizes.maxRowWidth[maxIndex] = 0
			}
			sizes.rows[i].maxIndex = maxIndex
			for j, col := range row {
				sizes.rows[i].colWidth[j] = len(fmt.Sprint(col)) + 2
				width += sizes.rows[i].colWidth[j]
				if cols != prevCols || sizes.rows[i].colWidth[j] > sizes.maxSizes[maxIndex][j] {
					sizes.maxSizes[maxIndex][j] = sizes.rows[i].colWidth[j]
				}
			}
			width += sizes.rows[i].numCols - 1
		default:
			cols = 1
			if cols != prevCols {
				maxIndex++
				sizes.maxRowWidth[maxIndex] = 0
			}
			sizes.rows[i].maxIndex = maxIndex
			sizes.rows[i].numCols = 1
			sizes.rows[i].colWidth[0] = len(fmt.Sprint(row)) + 2
			width = sizes.rows[i].colWidth[0]
			if cols != prevCols || sizes.rows[i].colWidth[0] > sizes.maxSizes[maxIndex][0] {
				sizes.maxSizes[maxIndex][0] = sizes.rows[i].colWidth[0]
			}
		}
		if width > sizes.maxRowWidth[maxIndex] {
			sizes.maxRowWidth[maxIndex] = width
		}
		if width > sizes.maxWidth {
			sizes.maxWidth = width
		}
		prevCols = cols
	}
	return sizes
}

func (row Table) Sprint(sizes RowSizes, maxSizes []int, maxRowWidth int, maxWidth int) (output string) {
	output = ""
	output2 := ""
	total := float64(maxWidth - maxRowWidth)
	count := math.Floor(float64(total) / float64(len(row)))
	//fmt.Println(sizes, maxSizes, maxRowWidth, maxWidth, (float64(maxWidth) / float64(maxRowWidth)), count, total)
	for j, col := range row {
		if j == 0 {
			output += "│"
		}
		spaces := math.Round((float64(maxSizes[j] - sizes.colWidth[j] + 2)))
		if j == len(row)-1 {
			count = total
		}
		spaces += count
		total -= count
		beforeSpaces := int(math.Floor(float64(spaces) / 2))
		afterSpaces := int(spaces) - beforeSpaces
		output += strings.Repeat(" ", beforeSpaces)
		switch col := col.(type) {
		case Table:
			output += fmt.Sprint(col)
		default:
			output += fmt.Sprint(col)
		}
		output += strings.Repeat(" ", afterSpaces)
		output += "│"
		output2 += fmt.Sprint(" (", spaces, "=", beforeSpaces, "+", afterSpaces, ")")
	}
	//output += output2
	return output
}

func iif(cond bool, value interface{}, value2 interface{}) interface{} {
	if cond {
		return value
	} else {
		return value2
	}
}

func (table Table) Print() {
	sizes := table.Sizes()
	//fmt.Println(table)
	//fmt.Println("=>", sizes)
	var output string
	var prevOutput string
	for i, row := range table {
		switch row := row.(type) {
		case Table:
			output = row.Sprint(sizes.rows[i], sizes.maxSizes[sizes.rows[i].maxIndex], sizes.maxRowWidth[sizes.rows[i].maxIndex], sizes.maxWidth)
		default:
			output = fmt.Sprint("│" + fmt.Sprint(row) + "│")
		}
		if i == 0 || (i > 0 && sizes.rows[i-1].numCols != sizes.rows[i].numCols) {
			// separation row
			sep := ""
			runes := []rune(output)
			for j, value := range runes {
				if j == 0 {
					sep += iif(i == 0, "╭", "├").(string)
				} else if j == len(runes)-1 {
					sep += iif(i == 0, "╮", "┤").(string)
				} else if value == '│' {
					sep += "┬"
				} else {
					sep += "─"
				}
			}
			sepRunes := []rune(sep)
			sep2 := ""
			if i > 0 {
				prevRunes := []rune(prevOutput)
				for j, value := range prevRunes {
					if j == 0 {
						sep2 += "├"
					} else if j == len(prevRunes)-1 {
						sep2 += "┤"
					} else if value == '│' && runes[j] == '│' {
						sep2 += "┼"
					} else if value == '│' {
						sep2 += "┴"
					} else {
						sep2 += string(sepRunes[j])
					}
				}
			} else {
				sep2=sep
			}
			fmt.Println(sep2)
		}
		fmt.Println(output)
		prevOutput = output
	}
	prevRunes := []rune(prevOutput)
	sep2:=""
	for j, value := range prevRunes {
		if j == 0 {
			sep2 += "╰"
		} else if j == len(prevRunes)-1 {
			sep2 += "╯"
		} else if value == '│' {
			sep2 += "┴"
		} else {
			sep2 += "─"
		}
	}
	fmt.Println(sep2)
}
