package create

import "strings"

const (
	PADDING_SIZE         = 2
	NEW_LINE             = "\n"
	TABLE_JOINT_SYMBOL   = "+"
	TABLE_V_SPLIT_SYMBOL = "|"
	TABLE_H_SPLIT_SYMBOL = "-"
)

func generateTable(headersList []string, rowsList [][]string,
	overRiddenHeaderHeight ...int) string {
	stringBuiler := strings.Builder{}
	var rowHeight int
	if length := len(overRiddenHeaderHeight); length > 0 {
		rowHeight = overRiddenHeaderHeight[0]
	} else {
		rowHeight = 1
	}
	columnMaxWidthMapping := getMaximumWidhtofTable(headersList, rowsList)
	createRowline(&stringBuiler, len(headersList), columnMaxWidthMapping)
	stringBuiler.Write([]byte(NEW_LINE))
	for headerIndex := 0; headerIndex < len(headersList); headerIndex++ {
		fillCell(&stringBuiler, headersList[headerIndex], headerIndex, columnMaxWidthMapping)
	}
	stringBuiler.Write([]byte(NEW_LINE))
	createRowline(&stringBuiler, len(headersList), columnMaxWidthMapping)
	for _, row := range rowsList {
		for i := 0; i < rowHeight; i++ {
			stringBuiler.Write([]byte(NEW_LINE))
		}
		for cellIndex := 0; cellIndex < len(row); cellIndex++ {
			fillCell(&stringBuiler, row[cellIndex], cellIndex, columnMaxWidthMapping)
		}
	}
	stringBuiler.Write([]byte(NEW_LINE))
	createRowline(&stringBuiler, len(headersList), columnMaxWidthMapping)
	return stringBuiler.String()
}
func fillSpace(stringbuilder *strings.Builder, length int) {
	for i := 0; i < length-1; i++ {
		stringbuilder.Write([]byte(" "))
	}
}
func createRowline(stringbuilder *strings.Builder, headersListSize int, columnMaxWidthMapping map[int]int) {
	for i := 0; i < headersListSize; i++ {
		if i == 0 {
			stringbuilder.Write([]byte(TABLE_JOINT_SYMBOL))
		}
		for j := 0; j < columnMaxWidthMapping[i]+PADDING_SIZE*2; j++ {
			stringbuilder.Write([]byte(TABLE_H_SPLIT_SYMBOL))
		}
		stringbuilder.Write([]byte(TABLE_JOINT_SYMBOL))
	}

}
func getMaximumWidhtofTable(headersList []string, rowsList [][]string) map[int]int {
	columnMaxWidthMapping := map[int]int{}
	for columnIndex := 0; columnIndex < len(headersList); columnIndex++ {
		columnMaxWidthMapping[columnIndex] = 0
	}
	for columnIndex := 0; columnIndex < len(headersList); columnIndex++ {
		if len(headersList[columnIndex]) > columnMaxWidthMapping[columnIndex] {
			columnMaxWidthMapping[columnIndex] = len(headersList[columnIndex])
		}
	}
	for _, row := range rowsList {
		for columnIndex := 0; columnIndex < len(row); columnIndex++ {
			if len(row[columnIndex]) > columnMaxWidthMapping[columnIndex] {
				columnMaxWidthMapping[columnIndex] = len(row[columnIndex])
			}
		}
	}
	for columnIndex := 0; columnIndex < len(headersList); columnIndex++ {
		if columnMaxWidthMapping[columnIndex]%2 != 0 {
			columnMaxWidthMapping[columnIndex] = columnMaxWidthMapping[columnIndex] + 1
		}
	}
	return columnMaxWidthMapping
}
func getOptimumCellPadding(cellIndex int, dataLength int, columnMaxWidthMapping map[int]int, cellPaddingSize int) int {
	if dataLength%2 != 0 {
		dataLength++
	}
	if dataLength < columnMaxWidthMapping[cellIndex] {
		cellPaddingSize = cellPaddingSize + (columnMaxWidthMapping[cellIndex]-dataLength)/2
	}
	return cellPaddingSize
}
func fillCell(stringBuilder *strings.Builder, cell string, cellIndex int, columnMaxWidthMapping map[int]int) {
	cellPaddingSize := getOptimumCellPadding(cellIndex, len(cell), columnMaxWidthMapping, PADDING_SIZE)
	if cellIndex == 0 {
		stringBuilder.Write([]byte(TABLE_V_SPLIT_SYMBOL))
	}
	fillSpace(stringBuilder, 2)
	stringBuilder.Write([]byte(cell))
	if len(cell)%2 != 0 {
		stringBuilder.Write([]byte(" "))
	}
	fillSpace(stringBuilder, cellPaddingSize)
	fillSpace(stringBuilder, cellPaddingSize)
	fillSpace(stringBuilder, 2)
	stringBuilder.Write([]byte(TABLE_V_SPLIT_SYMBOL))
}
