package main

import lib "github.com/tamaco489/go_sandbox/lo/lib/reduce"

func main() {
	// lo.Map
	// lib.NewUserSlicerByLoMap()

	// lo.UniqMap
	// lib.NewUserSlicerByLoUniqMap()

	// lo.SliceToMap
	// lib.NewPlayerSlicerByLoSliceToMap()

	// lo.FilterMap
	// lib.NewPlayerSlicerByLoFilterMap()

	// lo.Reduce
	lib.NewPokemonSliceByLoReducer()
}
