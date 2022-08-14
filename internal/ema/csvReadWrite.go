package ema

import (
	"encoding/csv"
	"fmt"
	"github.com/jszwec/csvutil"
	"io"
	"log"
	"os"
	"time"
)

type Candle struct {
	Name       string
	Time       time.Time
	OpenPrice  float64
	ClosePrice float64
	MaxPrice   float64
	LowPrice   float64
	VolumeLot  int64

	Ema10_1min  float64 `csv:",omitempty"`
	Ema10_5min  float64 `csv:"-"`
	Ema10_15min float64 `csv:"-"`
	Ema10_30min float64 `csv:"-"`
	Ema10_1h    float64 `csv:"-"`

	Ema20_1min  float64 `csv:"-"`
	Ema20_5min  float64 `csv:"-"`
	Ema20_15min float64 `csv:"-"`
	Ema20_30min float64 `csv:"-"`
	Ema20_1h    float64 `csv:"-"`

	Ema50_1min  float64 `csv:"-"`
	Ema50_5min  float64 `csv:"-"`
	Ema50_15min float64 `csv:"-"`
	Ema50_30min float64 `csv:"-"`
	Ema50_1h    float64 `csv:"-"`
}

//`csv:",omitempty"`

var Candles []Candle

func ReadCSV(figi string) [][]string {
	//file, err := os.Open(fmt.Sprint(figi, "/", figi, "_", "20220101", ".csv"))
	dirEntry, err := os.ReadDir(figi)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	filename := dirEntry[0].Name()
	file, err := os.Open(figi + "/" + filename)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comma = ';'
	rows, err := reader.ReadAll()
	if err != nil {
		log.Println("Cannot read CSV file:", err)
	}
	return rows
	/*
		fileScanner := bufio.NewScanner(file)
		// read line by line
		for fileScanner.Scan() {
			fmt.Println(fileScanner.Text())
		}
		// handle first encountered error while reading
		if err := fileScanner.Err(); err != nil {
			log.Fatalf("Error while reading file: %s", err)
		}
	*/
}

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

func ReadWriteCSV(figi string) {
	dirEntry, err := os.ReadDir(figi)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	filename := dirEntry[0].Name()
	file, err := os.Open(figi + "/" + filename)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	csvReader.Comma = ';'

	// in real application this should be done once in init function.
	userHeader, err := csvutil.Header(Candle{}, "csv")
	if err != nil {
		log.Fatal("error while csvutil.Header ", err)
	}

	dec, err := csvutil.NewDecoder(csvReader, userHeader...)
	if err != nil {
		log.Fatalln("error while csvutil.NewDecoder ", err)
	}

	for {
		var candle Candle
		if err = dec.Decode(&candle); err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln("error while dec.Decode ", err)
		}
		candle.countEma(10)
		fmt.Println(candle)
		Candles = append(Candles, candle)
	}
	WriteCSV(Candles)
	//fmt.Printf("%+v", Candles)
}

func (candle *Candle) countEma(averageCoefficient float64) {
	a := 2 / (averageCoefficient + 1)
	price := candle.ClosePrice //цена закрытия

	candle.Ema10_1min = a*price + (1-a)*candle.Ema10_1min
}

func WriteCSV(Candles []Candle) {

	file, err := os.Create("1.csv")
	if err != nil {
		log.Fatalf("Error when create file: %s", err)
	}
	defer file.Close()

	w := csv.NewWriter(file)
	w.Comma = ';'
	if err = csvutil.NewEncoder(w).Encode(Candles); err != nil {
		fmt.Println("error:", err)
	}

	w.Flush()
	if err = w.Error(); err != nil {
		fmt.Println("error:", err)
	}

}
