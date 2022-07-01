package main

type Sector struct {
	ID      int
	Name    string
	Owner   int
	Planet  int
	Station int
	Links   [8]int
}

type Planet struct {
	ID             int
	Name           string
	Owner          int
	Equipment      int
	EquipmentRatio float64
	Ore            int
	OreRatio       float64
	Organics       int
	OrganicsRatio  float64
	Goods          int
	GoodsRatio     float64
	Energy         int
	EnergyRatio    float64
}

type Station struct {
	ID             int
	Name           string
	Owner          int
	Equipment      int
	EquipmentRatio float64
	Ore            int
	OreRatio       float64
	Organics       int
	OrganicsRatio  float64
	Goods          int
	GoodsRatio     float64
	Energy         int
	EnergyRatio    float64
}
