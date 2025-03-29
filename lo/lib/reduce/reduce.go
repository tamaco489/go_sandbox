package lib

import (
	"fmt"
	"sort"

	"github.com/samber/lo"
	"github.com/tamaco489/go_sandbox/lo/model"
)

func NewPokemonSliceByLoReducer() {

	// ポケモンのサンプルデータ
	pokemonList := genSampleData()

	// Top3のポケモンを取得
	top3PokemonList := getTopPokemonListByStats(pokemonList, 10)

	// 結果を表示
	fmt.Println("Top 3 Pokemon by Stats:")
	for i, p := range top3PokemonList {
		fmt.Printf("[rank: %d] name: %s, total_stats: %d\n", i+1, p.Name, p.TotalStats)
	}

	/* [ステータスが高い上位3匹のみ出力]
	Top 3 Pokemon by Stats:
	[rank: 1] name: Pikachu, total_stats: 527
	[rank: 2] name: Bulbasaur, total_stats: 525
	[rank: 3] name: Charmander, total_stats: 516
	*/
}

// ポケモンのステータス合計を保持する構造体
type pokemonStat struct {
	Name       string
	TotalStats uint32
}

// newInitPokemonStats は、ポケモンのステータスを保持するスライスを初期化する関数。
//
// 初期値として空のスライスを返すことで、lo.Reduce の初期値として適切な型を提供する。
func newInitPokemonStats() []pokemonStat {
	return make([]pokemonStat, 0)
}

// newSetPokemonStats は、新しいポケモンのステータスエントリを作成するコンストラクタ関数。
func newSetPokemonStats(name string, totalStats uint32) *pokemonStat {
	return &pokemonStat{
		Name:       name,
		TotalStats: totalStats,
	}
}

// GetTopPokemonListByStats: ポケモンのリストを受け取り、ステータスの合計値が高い上位 topN のポケモンを取得する関数。
//
// pokemonList: 対象となるポケモンのリスト
//
// topN: 取得する上位ポケモンの数
//
// 戻り値: ステータスの高いポケモンのリスト（降順ソート済み）
func getTopPokemonListByStats(pokemonList []*model.Pokemon, topN int) []pokemonStat {

	reducer := func(acc []pokemonStat, pokemon *model.Pokemon, index int) []pokemonStat {
		// ステータス計算
		stats := pokemon.CalculateStats()

		// 合計ステータスを保存するエントリ（コンストラクタを使用）
		newEntry := newSetPokemonStats(pokemon.Name, stats.TotalStats)

		acc = append(acc, *newEntry)

		// 上位n匹のみ返す
		if len(acc) > topN {
			return acc[:topN]
		}

		return acc
	}

	// lo.Reduce を使って、ポケモンのリスト全体を畳み込み処理します。
	// 初期値として、newInitPokemonStats() を使用して空の []pokemonStat を渡しています。
	// これにより、pokemonList の各要素に対して、reducer 関数が順次適用されます。
	// 最終的に、上位 topN のポケモンのステータスエントリが topPokemonList に格納されます。
	topPokemonList := lo.Reduce(pokemonList, reducer, newInitPokemonStats())

	// 降順ソート
	sort.Slice(topPokemonList, func(i, j int) bool {
		return topPokemonList[i].TotalStats > topPokemonList[j].TotalStats
	})

	return topPokemonList
}

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
