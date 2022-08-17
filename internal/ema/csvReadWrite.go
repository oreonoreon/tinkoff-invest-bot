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

	Ema10_1min float64 `csv:",omitempty"`
	//Ema10_5min  float64 `csv:"-"`
	//Ema10_15min float64 `csv:"-"`
	//Ema10_30min float64 `csv:"-"`
	//Ema10_1h    float64 `csv:"-"`
	//
	//Ema20_1min  float64 `csv:"-"`
	//Ema20_5min  float64 `csv:"-"`
	//Ema20_15min float64 `csv:"-"`
	//Ema20_30min float64 `csv:"-"`
	//Ema20_1h    float64 `csv:"-"`
	//
	//Ema50_1min  float64 `csv:"-"`
	//Ema50_5min  float64 `csv:"-"`
	//Ema50_15min float64 `csv:"-"`
	//Ema50_30min float64 `csv:"-"`
	//Ema50_1h    float64 `csv:"-"`
}

//`csv:",omitempty"`

var Candles []Candle

func ReadWriteCSV(figi string) {
	dirEntry, err := os.ReadDir(figi)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	for _, v := range dirEntry {
		Candles = ReadCSV(v.Name(), figi)
	}

	WriteCSV(Candles)

}

func ReadCSV(filename string, figi string) []Candle {
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
		if len(Candles) != 0 {
			candle.interpolate()
		} else {
			Candles = append(Candles, candle)
		}

	}

	for {
		oldCandle := Candles[len(Candles)-1]
		if (oldCandle.Time.Minute()+1)%30 == 0 {
			break
		} else {
			////oldCandle.Time = oldCandle.Time.Add(time.Minute)
			//oldCandle = Candle{
			//	Name:       oldCandle.Name,
			//	Time:       oldCandle.Time.Add(time.Minute),
			//	OpenPrice:  oldCandle.ClosePrice,
			//	ClosePrice: oldCandle.ClosePrice,
			//	MaxPrice:   oldCandle.ClosePrice,
			//	LowPrice:   oldCandle.ClosePrice,
			//	VolumeLot:  0,
			//}
			////oldCandle.VolumeLot = 0
			//Candles = append(Candles, oldCandle)
			newInterpolateCandle(oldCandle)
		}

	}

	return Candles
}
func newInterpolateCandle(oldCandle Candle) {

	//oldCandle.Time = oldCandle.Time.Add(time.Minute)
	oldCandle = Candle{
		Name:       oldCandle.Name,
		Time:       oldCandle.Time.Add(time.Minute),
		OpenPrice:  oldCandle.ClosePrice,
		ClosePrice: oldCandle.ClosePrice,
		MaxPrice:   oldCandle.ClosePrice,
		LowPrice:   oldCandle.ClosePrice,
		VolumeLot:  0,
	}
	//oldCandle.VolumeLot = 0
	Candles = append(Candles, oldCandle)
}
func (candle Candle) interpolate() {
	oldCandle := Candles[len(Candles)-1]
	dur := candle.Time.Sub(oldCandle.Time)
	if !(dur == time.Minute || dur >= time.Hour) {
		var i float64
		for i = 1; i < dur.Minutes(); i++ {

			//oldCandle = Candle{
			//	Name:       oldCandle.Name,
			//	Time:       oldCandle.Time.Add(time.Minute),
			//	OpenPrice:  oldCandle.ClosePrice,
			//	ClosePrice: oldCandle.ClosePrice,
			//	MaxPrice:   oldCandle.ClosePrice,
			//	LowPrice:   oldCandle.ClosePrice,
			//	VolumeLot:  0,
			//}
			////oldCandle.VolumeLot = 0
			//Candles = append(Candles, oldCandle)
			newInterpolateCandle(oldCandle)
			oldCandle.Time = oldCandle.Time.Add(time.Minute)
		}
	}
	Candles = append(Candles, candle)
}

func CountEma(Candles []Candle, averageCoefficient float64) {
	a := 2 / (averageCoefficient + 1)
	for k, v := range Candles {
		if k != 0 {
			v.Ema10_1min = a*v.ClosePrice + (1-a)*Candles[k-1].Ema10_1min
		} else {
			v.Ema10_1min = v.ClosePrice
		}
		Candles[k] = v
	}
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

func NewTimeFrame(n int) {

	file, err := os.Create(fmt.Sprint("1_", n, ".csv"))
	if err != nil {
		log.Fatalf("Error when create file: %s", err)
	}
	defer file.Close()
	//
	w := csv.NewWriter(file)
	w.Comma = ';'
	for k, v := range Candles {
		if v.Time.Minute()%n == 0 {
			//smax,smin,vol:=MAXLOWpriceANDVolume(Candles[k : k+n])
			newCandle := Candle{
				Name:       v.Name,
				Time:       v.Time,
				OpenPrice:  v.OpenPrice,
				ClosePrice: Candles[k+n-1].ClosePrice,
				MaxPrice:   MaxPriceCount(Candles[k : k+n]),
				LowPrice:   LowPriceCount(Candles[k : k+n]),
				VolumeLot:  VolumeLotCount(Candles[k : k+n]),
			}
			enc := csvutil.NewEncoder(w)
			enc.AutoHeader = false
			if err = enc.Encode(newCandle); err != nil {
				fmt.Println("error:", err)
			}
		}
	}
	//w := csv.NewWriter(file)
	//w.Comma = ';'
	//if err = csvutil.NewEncoder(w).Encode(Candles); err != nil {
	//	fmt.Println("error:", err)
	//}

	w.Flush()
	if err = w.Error(); err != nil {
		fmt.Println("error:", err)
	}
}

func VolumeLotCount(candles []Candle) int64 {
	var s int64
	for _, v := range candles {
		s += v.VolumeLot
	}
	return s
}
func LowPriceCount(candles []Candle) float64 {
	var s float64 = candles[0].LowPrice
	for _, v := range candles {
		if v.LowPrice < s {
			s = v.LowPrice
		}
	}
	return s
}
func MaxPriceCount(candles []Candle) float64 {
	var s float64
	for _, v := range candles {
		if v.MaxPrice > s {
			s = v.MaxPrice
		}
	}

	return s
}

func MAXLOWpriceANDVolume(candles []Candle) (float64, float64, int64) {
	var smax, smin float64
	var vol int64
	smin = candles[0].LowPrice
	for _, v := range candles {
		if v.MaxPrice > smax {
			smax = v.MaxPrice
		}
		if v.LowPrice < smin {
			smin = v.LowPrice
		}
		vol += v.VolumeLot
	}
	return smax, smin, vol
}
