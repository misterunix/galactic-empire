package main

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/abhinav-TB/dantdb"
)

// overly complicated function to create print out the progress of the creation
func printComplete(s string, v1, v2 int) {

	doneV := float64(v1) / float64(v2) * 100.0
	doneS := strconv.FormatFloat(doneV, 'f', 3, 64)
	fmt.Printf("%s: %s%%\r", s, doneS)

}

// Create blank sectors
func createSectors(db *dantdb.Driver) {

	fmt.Println("Creating blank sectors...")
	for i := 0; i < sectorCount; i++ {

		s := Sector{
			ID:      i,
			Name:    fmt.Sprintf("Sector %d", i),
			Owner:   -1,
			Planet:  -1,
			Station: -1,
			Links:   [8]int{-1, -1, -1, -1, -1, -1, -1, -1},
		}

		db.Write("Sector", strconv.Itoa(i), s)

		printComplete("Sectors", i, sectorCount)

	}
	// I cant stand for it to be any thing less than 100%
	printComplete("Sectors", 1, 1)
	fmt.Println()
}

// Create semi-blank planets
func createPlanets(db *dantdb.Driver, planetCount int) {
	fmt.Println("Creating semi-blank planets...")
	for i := 0; i < planetCount; i++ {

		eqr := rand.Float64() + rand.Float64() - rand.Float64()
		or := rand.Float64() + rand.Float64() - rand.Float64()
		org := rand.Float64() + rand.Float64() - rand.Float64()
		gr := rand.Float64() + rand.Float64() - rand.Float64()
		er := rand.Float64() + rand.Float64() - rand.Float64()

		p := Planet{
			ID:             i,
			Sector:         -1,
			Name:           fmt.Sprintf("Planet %d", i),
			Owner:          -1,
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

		printComplete("Planets", i, planetCount)

	}
	// I cant stand for it to be any thing less than 100%
	printComplete("Planets", 1, 1)
	fmt.Println()
}

// Create semi-blank stations
func createStations(db *dantdb.Driver, stationCount int) {
	fmt.Println("Creating semi-blank stations...")
	for i := 0; i < stationCount; i++ {

		eqr := rand.Float64() + rand.Float64() - rand.Float64()
		or := rand.Float64() + rand.Float64() - rand.Float64()
		org := rand.Float64() + rand.Float64() - rand.Float64()
		gr := rand.Float64() + rand.Float64() - rand.Float64()
		er := rand.Float64() + rand.Float64() - rand.Float64()

		s := Station{
			ID:             i,
			Sector:         -1,
			Name:           fmt.Sprintf("Station %d", i),
			Owner:          -1,
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

		printComplete("Stations", i, stationCount)

	}
	// I cant stand for it to be any thing less than 100%
	printComplete("Stations", 1, 1)
	fmt.Println()
}

// Create the universe specified by the parameters.
// sectorCount: number of sectors in the universe
// planetCount: number of planets in each sector
// stationCount: number of stations in each sector
// minDistance: minimum distance between players
func buildUniverse(planetCount int, stationCount int, minDistance int) error {

	fmt.Println("Building universe...")

	dir := "./data/"
	db, err := dantdb.New(dir)
	if err != nil {
		return err
	}

	createSectors(db)

	// Create blank sectors
	/*
		fmt.Println("Creating blank sectors...")
		for i := 0; i < sectorCount; i++ {

			s := Sector{
				ID:      i,
				Name:    fmt.Sprintf("Sector %d", i),
				Owner:   -1,
				Planet:  -1,
				Station: -1,
				Links:   [8]int{-1, -1, -1, -1, -1, -1, -1, -1},
			}

			db.Write("Sector", strconv.Itoa(i), s)

			printComplete("Sectors", i, sectorCount)

		}


		// I cant stand for it to be any thing less than 100%
		printComplete("Sectors", 1, 1)
		fmt.Println()
	*/

	createPlanets(db, planetCount)
	/*
		// Create semi-blank planets
		fmt.Println("Creating semi-blank planets...")
		for i := 0; i < planetCount; i++ {

			eqr := rand.Float64() + rand.Float64() - rand.Float64()
			or := rand.Float64() + rand.Float64() - rand.Float64()
			org := rand.Float64() + rand.Float64() - rand.Float64()
			gr := rand.Float64() + rand.Float64() - rand.Float64()
			er := rand.Float64() + rand.Float64() - rand.Float64()

			p := Planet{
				ID:             i,
				Name:           fmt.Sprintf("Planet %d", i),
				Owner:          -1,
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

			printComplete("Planets", i, planetCount)

		}
		// I cant stand for it to be any thing less than 100%
		printComplete("Planets", 1, 1)
		fmt.Println()
	*/

	createStations(db, stationCount)
	/*
		// Create semi-blank stations
		fmt.Println("Creating semi-blank stations...")
		for i := 0; i < stationCount; i++ {

			eqr := rand.Float64() + rand.Float64() - rand.Float64()
			or := rand.Float64() + rand.Float64() - rand.Float64()
			org := rand.Float64() + rand.Float64() - rand.Float64()
			gr := rand.Float64() + rand.Float64() - rand.Float64()
			er := rand.Float64() + rand.Float64() - rand.Float64()

			s := Station{
				ID:             i,
				Name:           fmt.Sprintf("Station %d", i),
				Owner:          -1,
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

			printComplete("Stations", i, stationCount)

		}
		// I cant stand for it to be any thing less than 100%
		printComplete("Stations", 1, 1)
		fmt.Println()
	*/

	// Create links between sectors
	fmt.Println("Creating links...")
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

		printComplete("Sector links", i, sectorCount)

	}
	// I cant stand for it to be any thing less than 100%
	printComplete("Sector links", 1, 1)
	fmt.Println()

	// Assign planets to sectors
	fmt.Println("Assigning planets to Sectors...")
	for i := 0; i < planetCount; i++ {

		for {

			s := rand.Intn(sectorCount) // random sector

			p := Planet{} // planet data
			err := db.Read("Planet", strconv.Itoa(i), &p)
			if err != nil {
				return err
			}

			s1 := Sector{} // sector data
			err = db.Read("Sector", strconv.Itoa(s), &s1)
			if err != nil {
				return err
			}

			if s1.Planet == -1 {
				// add the planet to the sector
				s1.Planet = i
				err = db.Write("Sector", strconv.Itoa(s), s1)
				if err != nil {
					return err
				}

				// add the sector to the planet
				p.Sector = s
				err = db.Write("Planet", strconv.Itoa(i), p)
				if err != nil {
					return err
				}
				break
			}
		}

		printComplete("Planets", i, planetCount)

	}
	// I cant stand for it to be any thing less than 100%
	printComplete("Planets", 1, 1)
	fmt.Println()

	// Assign stations to sectors
	fmt.Println("Assigning stations to Sectors...")
	for i := 0; i < stationCount; i++ {

		for {

			s := rand.Intn(sectorCount) // random sector

			s1 := Sector{} // sector data
			err := db.Read("Sector", strconv.Itoa(s), &s1)
			if err != nil {
				return err
			}

			s2 := Station{} // station data
			err = db.Read("Station", strconv.Itoa(i), &s2)
			if err != nil {
				return err
			}

			if s1.Station == -1 && s1.Planet == -1 {

				// add the station to the sector
				s1.Station = i
				err = db.Write("Sector", strconv.Itoa(s), s1)
				if err != nil {
					return err
				}

				// add the sector to the station
				s2.Sector = s
				err = db.Write("Station", strconv.Itoa(i), s2)
				if err != nil {
					return err
				}

				break
			}
		}

		printComplete("Station", i, stationCount)
	}
	// I cant stand for it to be any thing less than 100%
	printComplete("Station", 1, 1)
	fmt.Println()

	// Assign players to sectors
	var pl [6]int // player list for random planet

	fmt.Println("Assigning players to Planets...")
	for {
		bad := false

		for i := 0; i < 6; i++ {
			pl[i] = rand.Intn(planetCount)
		}

		for j := 0; j < 5; j++ {
			for k := j; k < 6; k++ {
				if j == k {
					continue
				}
				dis := distance(pl[j], pl[k])
				if dis < minDistance {
					fmt.Println("Bad distance:", pl[j], pl[k], dis)
					bad = true
				}
			}
		}

		if !bad {
			break
		}

	}

	for i := 0; i < 6; i++ {

		p := Planet{} // planet data
		err := db.Read("Planet", strconv.Itoa(pl[i]), &p)
		if err != nil {
			return err
		}
		p.Owner = i
		err = db.Write("Planet", strconv.Itoa(pl[i]), p)
		if err != nil {
			return err
		}
		fmt.Printf("Player %d is on Planet %d\n", i, pl[i])

	}

	linkdistance(db, 1, 44)

	return nil
}

// link count between two sectors
func linkdistance(db *dantdb.Driver, bs int, es int) int {

	largemap := make(map[int][]int)

	for i := 0; i < sectorCount; i++ {

		s1 := Sector{} // sector data
		err := db.Read("Sector", strconv.Itoa(bs), &s1)
		if err != nil {
			return -1
		}

		for j := 0; j < 8; j++ {
			if s1.Links[j] != -1 {
				largemap[i] = append(largemap[i], s1.Links[j])
			}
		}
	}

	for i := 0; i < sectorCount; i++ {
		fmt.Println(largemap[i])
	}

	return 1
}
