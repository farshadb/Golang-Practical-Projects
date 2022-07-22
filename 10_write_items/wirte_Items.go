package wirtIterms

import (
	"encoding/csv"
	"os"
)

type Item struct {
	SKU,Name  string
}
// use this program to close open files to decerease memory usage

func writeItems (fileName string, items []Item) {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	row := []string{"sku", "name"}
	
	wtr := csv.NewWriter(file)
	defer wtr.Flush()

	if err := wtr.Write(row); err != nil {
		return err
	}
	for _, item := range items {
		row[0] = item.SKU
		row[1] = item.Name
		if err := wtr.Write(row); err != nil {
			return err
		}

}


}
