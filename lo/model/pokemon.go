package model

// ポケモンの能力の最終値を計算する構造体
type PokemonStats struct {
	HP         uint32
	Attack     uint32
	Defense    uint32
	Speed      uint32
	SpecialAt  uint32
	SpecialDe  uint32
	TotalStats uint32
}

// ポケモンの種族値（Base Stats）構造体
type BaseStats struct {
	HP        uint32
	Attack    uint32
	Defense   uint32
	Speed     uint32
	SpecialAt uint32
	SpecialDe uint32
}

// ポケモンの個体値（Individual Values, IVs）構造体
type IndividualValues struct {
	HP        uint32
	Attack    uint32
	Defense   uint32
	Speed     uint32
	SpecialAt uint32
	SpecialDe uint32
}

// ポケモンの努力値（Effort Values, EVs）構造体
type EffortValues struct {
	HP        uint32
	Attack    uint32
	Defense   uint32
	Speed     uint32
	SpecialAt uint32
	SpecialDe uint32
}

// ポケモンの構造体
type Pokemon struct {
	ID        int
	Name      string
	BaseStats BaseStats
	IVs       IndividualValues
	EVs       EffortValues
	Level     uint32
}

// 新しいポケモンを作成する関数（レベルと能力値の計算）
func NewPokemon(id int, name string, baseStats BaseStats, ivs IndividualValues, evs EffortValues, level uint32) *Pokemon {
	return &Pokemon{
		ID:        id,
		Name:      name,
		BaseStats: baseStats,
		IVs:       ivs,
		EVs:       evs,
		Level:     level,
	}
}

// ポケモンの能力値をレベルを基に計算する関数
func (p *Pokemon) CalculateStats() PokemonStats {

	// HPの計算
	HP := (2*p.BaseStats.HP+p.IVs.HP+p.EVs.HP/4)*p.Level/100 + p.Level + 10

	// 攻撃、守備、特攻、特防、速さの計算（全て同じ計算式）
	Attack := ((2*p.BaseStats.Attack+p.IVs.Attack+p.EVs.Attack/4)*p.Level)/100 + 5
	Defense := ((2*p.BaseStats.Defense+p.IVs.Defense+p.EVs.Defense/4)*p.Level)/100 + 5
	Speed := ((2*p.BaseStats.Speed+p.IVs.Speed+p.EVs.Speed/4)*p.Level)/100 + 5
	SpecialAt := ((2*p.BaseStats.SpecialAt+p.IVs.SpecialAt+p.EVs.SpecialAt/4)*p.Level)/100 + 5
	SpecialDe := ((2*p.BaseStats.SpecialDe+p.IVs.SpecialDe+p.EVs.SpecialDe/4)*p.Level)/100 + 5

	// 合計ステータスの計算
	totalStats := HP + Attack + Defense + Speed + SpecialAt + SpecialDe

	// 計算結果をPokemonStats構造体で返す
	return PokemonStats{
		HP:         HP,
		Attack:     Attack,
		Defense:    Defense,
		Speed:      Speed,
		SpecialAt:  SpecialAt,
		SpecialDe:  SpecialDe,
		TotalStats: totalStats,
	}
}
