package ema

//func readCSV(figi string) [][]string {
//	//file, err := os.Open(fmt.Sprint(figi, "/", figi, "_", "20220101", ".csv"))
//	dirEntry, err := os.ReadDir(figi)
//	if err != nil {
//		log.Fatalf("Error when opening file: %s", err)
//	}
//	filename := dirEntry[0].Name()
//	file, err := os.Open(figi + "/" + filename)
//	if err != nil {
//		log.Fatalf("Error when opening file: %s", err)
//	}
//	defer file.Close()
//	reader := csv.NewReader(file)
//	reader.Comma = ';'
//	rows, err := reader.ReadAll()
//	if err != nil {
//		log.Println("Cannot read CSV file:", err)
//	}
//	return rows
//	/*
//		fileScanner := bufio.NewScanner(file)
//		// read line by line
//		for fileScanner.Scan() {
//			fmt.Println(fileScanner.Text())
//		}
//		// handle first encountered error while reading
//		if err := fileScanner.Err(); err != nil {
//			log.Fatalf("Error while reading file: %s", err)
//		}
//	*/
//}

/*
func WriteInCandle(rows [][]string) {
	for _, row := range rows {
		//fmt.Println(row[0])
		candle := Candle{
			Name:       row[0],
			Time:       row[1],
			OpenPrice:  row[2],
			ClosePrice: row[3],
			MaxPrice:   row[4],
			LowPrice:   row[5],
			VolumeLot:  row[6],
		}
		Candles = append(Candles, candle)
	}
}
*/
