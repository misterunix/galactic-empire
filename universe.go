package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/abhinav-TB/dantdb"
)

func buildUniverse(sectorCount int, planetCount int, stationCount int) error {

	dir := "./"
	db, err := dantdb.New(dir)
	if err != nil {
		return err
	}

	for i := 0; i < sectorCount; i++ {
		s := Sector{
			ID:      i,
			Name:    fmt.Sprintf("Sector %d", i),
			Owner:   0,
			Planet:  0,
			Station: 0,
			Links:   [8]int{-1, -1, -1, -1, -1, -1, -1, -1},
		}
		db.Write("Sector", strconv.Itoa(i), s)
	}

	for i := 0; i < planetCount; i++ {
		p := Planet{
			ID:             i,
			Name:           fmt.Sprintf("Planet %d", i),
			Owner:          0,
			Equipment:      0,
			EquipmentRatio: 0.0,
			Ore:            0,
			OreRatio:       0.0,
			Organics:       0,
			OrganicsRatio:  0.0,
			Goods:          0,
			GoodsRatio:     0.0,
			Energy:         0,
			EnergyRatio:    0.0,
		}
		db.Write("Planet", strconv.Itoa(i), p)
	}

	for i := 0; i < stationCount; i++ {
		s := Station{
			ID:             i,
			Name:           fmt.Sprintf("Station %d", i),
			Owner:          0,
			Equipment:      0,
			EquipmentRatio: 0.0,
			Ore:            0,
			OreRatio:       0.0,
			Organics:       0,
			OrganicsRatio:  0.0,
			Goods:          0,
			GoodsRatio:     0.0,
			Energy:         0,
			EnergyRatio:    0.0,
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
	return nil
}
