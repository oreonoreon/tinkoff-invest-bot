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

	EMAS CandleEma `csv:",omitempty"`
}
type CandleEma struct {
	Ema_10 float64
	Ema_20 float64
	Ema_50 float64
	//Time   time.Time
}

var Candles []Candle

func ReadWriteCSV(figi string) {
	dirEntry, err := os.ReadDir(figi)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	for _, v := range dirEntry {
		Candles = readCSV(v.Name(), figi)
	}

	writeCSV(Candles)

}

func readCSV(filename string, figi string) []Candle {
	file, err := os.Open(figi + "/" + filename)
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}
	defer file.Close()
	csvReader := csv.NewReader(file)
	//csvReader.FieldsPerRecord = 7
	csvReader.Comma = ';'

	// in real application this should be done once in init function.
	userHeader, err := csvutil.Header(Candle{}, "csv")
	if err != nil {
		log.Fatal("error while csvutil.Header ", err)
	}

	dec, err := csvutil.NewDecoder(csvReader, userHeader...)
	dec.DisallowMissingColumns = false
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
	oldCandle := Candles[len(Candles)-1]
	for {
		if (oldCandle.Time.Minute()+1)%30 == 0 {
			break
		} else {
			oldCandle = newInterpolateCandle(oldCandle)
		}
	}

	return Candles
}
func newInterpolateCandle(oldCandle Candle) Candle {
	oldCandle = Candle{
		Name:       oldCandle.Name,
		Time:       oldCandle.Time.Add(time.Minute),
		OpenPrice:  oldCandle.ClosePrice,
		ClosePrice: oldCandle.ClosePrice,
		MaxPrice:   oldCandle.ClosePrice,
		LowPrice:   oldCandle.ClosePrice,
		VolumeLot:  0,
	}
	Candles = append(Candles, oldCandle)
	return oldCandle
}
func (candle Candle) interpolate() {
	oldCandle := Candles[len(Candles)-1]
	dur := candle.Time.Sub(oldCandle.Time)
	if !(dur == time.Minute || dur >= time.Hour) {
		var i float64
		for i = 1; i < dur.Minutes(); i++ {
			oldCandle = newInterpolateCandle(oldCandle)
		}
	}
	Candles = append(Candles, candle)
}

func writeCSV(Candles []Candle) {

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

// NewTimeFrame тут n целое число отбражающее таймфрейм свечей, к примеру 5 минутные, 15 минутные и т.д
func NewTimeFrame(n int) {
	var Candles1 []Candle

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
			var eMas CandleEma
			if len(Candles1) != 0 {
				eMas = CountEma(Candles[k+n-1].ClosePrice, Candles1[len(Candles1)-1].EMAS)
			} else {
				eMas = CandleEma{
					Ema_10: Candles[k+n-1].ClosePrice,
					Ema_20: Candles[k+n-1].ClosePrice,
					Ema_50: Candles[k+n-1].ClosePrice,
				}
			}
			newCandle := Candle{
				Name:       v.Name,
				Time:       v.Time,
				OpenPrice:  v.OpenPrice,
				ClosePrice: Candles[k+n-1].ClosePrice,
				MaxPrice:   MaxPriceCount(Candles[k : k+n]),
				LowPrice:   LowPriceCount(Candles[k : k+n]),
				VolumeLot:  VolumeLotCount(Candles[k : k+n]),
				EMAS:       eMas,
			}
			Candles1 = append(Candles1, newCandle)
			//enc := csvutil.NewEncoder(w)
			//enc.AutoHeader = false
			//if err = enc.Encode(newCandle); err != nil {
			//	fmt.Println("error:", err)
			//}
		}
	}
	//b, err := csvutil.Marshal(Candles1)
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//os.WriteFile(fmt.Sprint("1_", n, ".csv"), b, 0755)
	enc := csvutil.NewEncoder(w)
	//enc.AutoHeader = false
	if err = enc.Encode(Candles1); err != nil {
		fmt.Println("error:", err)
	}

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

func CountEma(price float64, eMas CandleEma) CandleEma {
	aver := [3]float64{10, 20, 50}
	for k, _ := range aver {
		a := 2 / (aver[k] + 1)
		switch k {
		case 0:
			eMas.Ema_10 = a*price + (1-a)*eMas.Ema_10
		case 1:
			eMas.Ema_20 = a*price + (1-a)*eMas.Ema_20
		case 2:
			eMas.Ema_50 = a*price + (1-a)*eMas.Ema_50
		}
	}
	return eMas
}
