package collector

const (
	PasardanaCatMix int = iota
	PasardanaCatEq
	PasardanaCatFI
	PasardanaCatMM
	PasardanaCatProt
	PasardanaCatLim
	PasardanaCatDire
	PasardanaCatGlobal
)

const (
	CatMM int = iota
	CatFI
	CatMix
	CatEq
	CatIdx
	CatProt
	CatGlobal
	CatOthers
)

func GetPasardanaMappedCategory(c int) int {
	m := map[int]int{
		PasardanaCatMix:    CatMix,
		PasardanaCatEq:     CatEq,
		PasardanaCatFI:     CatFI,
		PasardanaCatMM:     CatMM,
		PasardanaCatProt:   CatProt,
		PasardanaCatGlobal: CatGlobal,
	}

	v, ok := m[c]
	if !ok {
		return CatOthers
	}
	return v
}
