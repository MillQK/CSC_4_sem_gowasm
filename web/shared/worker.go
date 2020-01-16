package shared

import "sync"

func RowWorker(wg *sync.WaitGroup, rowsChannel chan uint32, columnSize uint32, handler func(rowNum uint32, columnNum uint32)) {
	defer wg.Done()
	for rowNumber := range rowsChannel {
		for i := uint32(0); i < columnSize; i++ {
			handler(rowNumber, i)
		}
	}
}
