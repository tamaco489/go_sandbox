package lib

import (
	"log"
	"math/rand"
	"time"

	"github.com/samber/lo"
	"github.com/tamaco489/go_sandbox/lo/model"
)

// NewPokemonSliceByLoReducer は、ランダムに選ばれた6匹のポケモンの TotalStats の合計を計算し、結果を出力します。
func NewPokemonSliceByLoReducer() {

	// ポケモンのサンプルデータ
	pokemonList := getRandomPokemonList()
	for i, v := range pokemonList {
		log.Printf("[pokemon] No: %d, name: %s, ttl_stats: %d", i+1, v.Name, v.CalculateStats().TotalStats)
	}
	/*
		2025/03/29 16:47:20 [pokemon] No: 1, name: Snorlax, ttl_stats: 747
		2025/03/29 16:47:20 [pokemon] No: 2, name: Eevee, ttl_stats: 502
		2025/03/29 16:47:20 [pokemon] No: 3, name: Bulbasaur, ttl_stats: 525
		2025/03/29 16:47:20 [pokemon] No: 4, name: Charmander, ttl_stats: 516
		2025/03/29 16:47:20 [pokemon] No: 5, name: Gyarados, ttl_stats: 747
		2025/03/29 16:47:20 [pokemon] No: 6, name: Blastoise, ttl_stats: 737
	*/

	// lo.Reduce を使用して、選ばれたポケモンのTotalStatsの合計を計算
	totalStats := lo.Reduce(pokemonList, func(acc uint32, p *model.Pokemon, index int) uint32 {
		// 各ポケモンのTotalStatsを加算する
		// acc は累積値、p は現在のポケモン、index はポケモンのインデックス
		return acc + p.CalculateStats().TotalStats
	}, 0)

	log.Printf("Total Stats of selected 6 Pokemons: %d", totalStats) // 2025/03/29 16:47:20 Total Stats of selected 6 Pokemons: 3774
}

// getRandomPokemonList は、サンプルポケモンデータをシャッフルして、最初の6匹のポケモンをランダムに選択して返します。
func getRandomPokemonList() []*model.Pokemon {

	pokemonList := genSampleData()
	randGen := rand.New(rand.NewSource(time.Now().UnixNano()))

	// スライス内の要素をシャッフル
	randGen.Shuffle(len(pokemonList), func(i, j int) {
		pokemonList[i], pokemonList[j] = pokemonList[j], pokemonList[i]
	})

	// シャッフル後、最初の6つの要素を選択
	return pokemonList[:6]
}

// サンプルのポケモンデータを生成する関数
func genSampleData() []*model.Pokemon {
	return []*model.Pokemon{
		model.NewPokemon(1, "Pikachu", model.BaseStats{HP: 35, Attack: 55, Defense: 40, Speed: 90, SpecialAt: 50, SpecialDe: 50}, model.IndividualValues{HP: 31, Attack: 31, Defense: 31, Speed: 31, SpecialAt: 31, SpecialDe: 31}, model.EffortValues{HP: 0, Attack: 0, Defense: 0, Speed: 252, SpecialAt: 0, SpecialDe: 0}, 50),
		model.NewPokemon(2, "Bulbasaur", model.BaseStats{HP: 45, Attack: 49, Defense: 49, Speed: 45, SpecialAt: 65, SpecialDe: 65}, model.IndividualValues{HP: 31, Attack: 31, Defense: 31, Speed: 31, SpecialAt: 31, SpecialDe: 31}, model.EffortValues{HP: 0, Attack: 0, Defense: 0, Speed: 252, SpecialAt: 0, SpecialDe: 0}, 50),
		model.NewPokemon(3, "Charmander", model.BaseStats{HP: 39, Attack: 52, Defense: 43, Speed: 65, SpecialAt: 60, SpecialDe: 50}, model.IndividualValues{HP: 31, Attack: 31, Defense: 31, Speed: 31, SpecialAt: 31, SpecialDe: 31}, model.EffortValues{HP: 0, Attack: 0, Defense: 0, Speed: 252, SpecialAt: 0, SpecialDe: 0}, 50),
		model.NewPokemon(4, "Squirtle", model.BaseStats{HP: 44, Attack: 48, Defense: 65, Speed: 43, SpecialAt: 50, SpecialDe: 64}, model.IndividualValues{HP: 31, Attack: 31, Defense: 31, Speed: 31, SpecialAt: 31, SpecialDe: 31}, model.EffortValues{HP: 0, Attack: 0, Defense: 0, Speed: 252, SpecialAt: 0, SpecialDe: 0}, 50),
		model.NewPokemon(5, "Eevee", model.BaseStats{HP: 55, Attack: 40, Defense: 50, Speed: 55, SpecialAt: 45, SpecialDe: 50}, model.IndividualValues{HP: 31, Attack: 31, Defense: 31, Speed: 31, SpecialAt: 31, SpecialDe: 31}, model.EffortValues{HP: 0, Attack: 0, Defense: 0, Speed: 252, SpecialAt: 0, SpecialDe: 0}, 50),
		model.NewPokemon(6, "Snorlax", model.BaseStats{HP: 160, Attack: 110, Defense: 65, Speed: 30, SpecialAt: 65, SpecialDe: 110}, model.IndividualValues{HP: 31, Attack: 31, Defense: 31, Speed: 31, SpecialAt: 31, SpecialDe: 31}, model.EffortValues{HP: 0, Attack: 0, Defense: 0, Speed: 252, SpecialAt: 0, SpecialDe: 0}, 50),
		model.NewPokemon(7, "Mewtwo", model.BaseStats{HP: 106, Attack: 110, Defense: 90, Speed: 130, SpecialAt: 154, SpecialDe: 90}, model.IndividualValues{HP: 31, Attack: 31, Defense: 31, Speed: 31, SpecialAt: 31, SpecialDe: 31}, model.EffortValues{HP: 0, Attack: 0, Defense: 0, Speed: 252, SpecialAt: 0, SpecialDe: 0}, 50),
		model.NewPokemon(8, "Charizard", model.BaseStats{HP: 78, Attack: 84, Defense: 78, Speed: 100, SpecialAt: 109, SpecialDe: 85}, model.IndividualValues{HP: 31, Attack: 31, Defense: 31, Speed: 31, SpecialAt: 31, SpecialDe: 31}, model.EffortValues{HP: 0, Attack: 0, Defense: 0, Speed: 252, SpecialAt: 0, SpecialDe: 0}, 50),
		model.NewPokemon(9, "Blastoise", model.BaseStats{HP: 79, Attack: 83, Defense: 100, Speed: 78, SpecialAt: 85, SpecialDe: 105}, model.IndividualValues{HP: 31, Attack: 31, Defense: 31, Speed: 31, SpecialAt: 31, SpecialDe: 31}, model.EffortValues{HP: 0, Attack: 0, Defense: 0, Speed: 252, SpecialAt: 0, SpecialDe: 0}, 50),
		model.NewPokemon(10, "Gyarados", model.BaseStats{HP: 95, Attack: 125, Defense: 79, Speed: 81, SpecialAt: 60, SpecialDe: 100}, model.IndividualValues{HP: 31, Attack: 31, Defense: 31, Speed: 31, SpecialAt: 31, SpecialDe: 31}, model.EffortValues{HP: 0, Attack: 0, Defense: 0, Speed: 252, SpecialAt: 0, SpecialDe: 0}, 50),
	}
}
