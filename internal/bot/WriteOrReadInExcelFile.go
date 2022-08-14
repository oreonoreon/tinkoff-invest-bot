package bot

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
)

func (U *Cache) writeCacheInFile() {
	log.Println("Запись кэша в файл ...")
	i := 0
	categories := map[string]string{
		"A1": "userID", "B1": "TokenSanBox",
		"C1": "TokenRealTrade", "D1": "Amount",
		"E1": "sum1", "F1": "sum2",
		"G1": "sum3", "H1": "FirstName",
		"I1": "LastName", "J1": "UserName",
	}

	f := excelize.NewFile()
	if err := f.SetColWidth("Sheet1", "A", "J", 15); err != nil {
		log.Println("Немогу установить ширину коллон в excel 'f.SetColWidth' ", err)
	}
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for key, value := range U.Users {
		i++
		k, _ := excelize.CoordinatesToCellName(1, i+1)
		f.SetCellValue("Sheet1", k, key)
		k, _ = excelize.CoordinatesToCellName(2, i+1)
		f.SetCellValue("Sheet1", k, value.TokenSanBox)
		k, _ = excelize.CoordinatesToCellName(3, i+1)
		f.SetCellValue("Sheet1", k, value.TokenRealTrade)
		k, _ = excelize.CoordinatesToCellName(8, i+1)
		f.SetCellValue("Sheet1", k, value.FirstName)
		k, _ = excelize.CoordinatesToCellName(9, i+1)
		f.SetCellValue("Sheet1", k, value.LastName)
		k, _ = excelize.CoordinatesToCellName(10, i+1)
		f.SetCellValue("Sheet1", k, value.UserName)
	}
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
	log.Println("Запись кэша в файл закончена...")
}

/*
func (U *Cache) writeInFile() {
	i := 0
	categories := map[string]string{
		"A1": "OperationId", "B1": "UserID",
		"C1": "Amount", "D1": "Datetime",
		"E1": "Status", "F1": "FirstName",
		"G1": "LastName", "H1": "UserName",
	}

	f := excelize.NewFile()
	if err := f.SetColWidth("Sheet1", "A", "H", 15); err != nil {
		log.Println("Немогу установить ширину коллон в excel 'f.SetColWidth' ", err)
	}
	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for _, value := range reqYoomoneyGetAllOperation() {
		i++
		k, _ := excelize.CoordinatesToCellName(1, i+1)
		f.SetCellValue("Sheet1", k, value.OperationId)
		k, _ = excelize.CoordinatesToCellName(2, i+1)
		f.SetCellValue("Sheet1", k, value.Label)
		k, _ = excelize.CoordinatesToCellName(3, i+1)
		f.SetCellValue("Sheet1", k, value.Amount)
		k, _ = excelize.CoordinatesToCellName(4, i+1)
		f.SetCellValue("Sheet1", k, value.Datetime)
		k, _ = excelize.CoordinatesToCellName(5, i+1)
		f.SetCellValue("Sheet1", k, value.Status)

		Id, _ := strconv.ParseInt(value.Label, 10, 64)
		k, _ = excelize.CoordinatesToCellName(6, i+1)
		f.SetCellValue("Sheet1", k, U.Users[Id].FirstName)
		k, _ = excelize.CoordinatesToCellName(7, i+1)
		f.SetCellValue("Sheet1", k, U.Users[Id].LastName)
		k, _ = excelize.CoordinatesToCellName(8, i+1)
		f.SetCellValue("Sheet1", k, U.Users[Id].UserName)

	}
	if err := f.SaveAs("Book2.xlsx"); err != nil {
		fmt.Println(err)
	}
	log.Println("Запись в файл игроков оплативших билеты ...")
}
*/
func loadFromExcel(c *Cache) {
	f, erro := excelize.OpenFile("Book1.xlsx")
	if erro != nil {
		fmt.Println(erro)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Get value from cell by given worksheet name and axis.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	var Id int64

	for i := 1; i < len(rows); i++ {

		for j := 1; j <= 10; j++ {
			k, _ := excelize.CoordinatesToCellName(j, i+1)
			cell, err := f.GetCellValue("Sheet1", k)
			if err != nil {
				fmt.Println(err)
				return
			}
			switch j {
			case 1:
				Id, _ = strconv.ParseInt(cell, 10, 64)
				c.FullFillCache(Id)
			case 2:
				if entry, ok := c.Users[Id]; ok {
					entry.TokenSanBox = cell
					c.Users[Id] = entry
				}
			case 3:
				if entry, ok := c.Users[Id]; ok {
					entry.TokenRealTrade = cell
					c.Users[Id] = entry
				}
			case 8:
				if entry, ok := c.Users[Id]; ok {
					entry.FirstName = cell
					c.Users[Id] = entry
				}
			case 9:
				if entry, ok := c.Users[Id]; ok {
					entry.LastName = cell
					c.Users[Id] = entry
				}
			case 10:
				if entry, ok := c.Users[Id]; ok {
					entry.UserName = cell
					c.Users[Id] = entry
				}
			}
		}
	}

}
