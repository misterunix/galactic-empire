package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/abhinav-TB/dantdb"
)

func buildUniverse(sectorCount int, planetCount int, stationCount int) error {

	dir := "./data/"
	db, err := dantdb.New(dir)
	if err != nil {
		return err
	}

	for i := 0; i < sectorCount; i++ {
		s := Sector{
			ID:      i,
			Name:    fmt.Sprintf("Sector %d", i),
			Owner:   0,
			Planet:  -1,
			Station: -1,
			Links:   [8]int{-1, -1, -1, -1, -1, -1, -1, -1},
		}
		db.Write("Sector", strconv.Itoa(i), s)
	}

	for i := 0; i < planetCount; i++ {

		eqr := rand.Float64() + rand.Float64() - rand.Float64()
		or := rand.Float64() + rand.Float64() - rand.Float64()
		org := rand.Float64() + rand.Float64() - rand.Float64()
		gr := rand.Float64() + rand.Float64() - rand.Float64()
		er := rand.Float64() + rand.Float64() - rand.Float64()

		p := Planet{
			ID:             i,
			Name:           fmt.Sprintf("Planet %d", i),
			Owner:          0,
			Equipment:      0,
			EquipmentRatio: eqr,
			Ore:            0,
			OreRatio:       or,
			Organics:       0,
			OrganicsRatio:  org,
			Goods:          0,
			GoodsRatio:     gr,
			Energy:         0,
			EnergyRatio:    er,
		}
		db.Write("Planet", strconv.Itoa(i), p)
	}

	for i := 0; i < stationCount; i++ {
		eqr := rand.Float64() + rand.Float64() - rand.Float64()
		or := rand.Float64() + rand.Float64() - rand.Float64()
		org := rand.Float64() + rand.Float64() - rand.Float64()
		gr := rand.Float64() + rand.Float64() - rand.Float64()
		er := rand.Float64() + rand.Float64() - rand.Float64()

		s := Station{
			ID:             i,
			Name:           fmt.Sprintf("Station %d", i),
			Owner:          0,
			Equipment:      0,
			EquipmentRatio: eqr,
			Ore:            0,
			OreRatio:       or,
			Organics:       0,
			OrganicsRatio:  org,
			Goods:          0,
			GoodsRatio:     gr,
			Energy:         0,
			EnergyRatio:    er,
		}
		db.Write("Station", strconv.Itoa(i), s)
	}

	for i := 0; i < sectorCount; i++ {
		for jl := 0; jl < 3; jl++ {

			lp1 := rand.Intn(sectorCount)
			lp2 := rand.Intn(sectorCount)

			s1 := Sector{}
			err := db.Read("Sector", strconv.Itoa(lp1), &s1)
			if err != nil {
				return err
			}

			s2 := Sector{}
			err = db.Read("Sector", strconv.Itoa(lp2), &s2)
			if err != nil {
				return err
			}

			for j := 0; j < 8; j++ {
				if s1.Links[j] == -1 {
					s1.Links[j] = lp2
					break
				}
			}
			for j := 0; j < 8; j++ {
				if s2.Links[j] == -1 {
					s2.Links[j] = lp1
					break
				}
			}

			err = db.Write("Sector", strconv.Itoa(lp1), s1)
			if err != nil {
				return err
			}
			err = db.Write("Sector", strconv.Itoa(lp2), s2)
			if err != nil {
				return err
			}
		}
	}

	fmt.Println("Plants")
	for i := 0; i < planetCount; i++ {

		for {
			s := rand.Intn(sectorCount)
			s1 := Sector{}
			err := db.Read("Sector", strconv.Itoa(s), &s1)
			if err != nil {
				return err
			}
			if s1.Planet == -1 {
				s1.Planet = i
				err = db.Write("Sector", strconv.Itoa(s), s1)
				if err != nil {
					return err
				}
				break
			}
		}
	}

	fmt.Println("Stations")
	for i := 0; i < stationCount; i++ {

		for {
			s := rand.Intn(sectorCount)
			s1 := Sector{}
			err := db.Read("Sector", strconv.Itoa(s), &s1)
			if err != nil {
				return err
			}
			if s1.Station == -1 {
				s1.Station = i
				err = db.Write("Sector", strconv.Itoa(s), s1)
				if err != nil {
					return err
				}
				break
			}
		}
	}

	return nil
}
