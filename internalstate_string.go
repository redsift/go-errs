// Code generated by "stringer -type=InternalState"; DO NOT EDIT.

package errs

import "strconv"

const _InternalState_name = "MochasippiBrevePapiInstantTurkishIrishCremaCappuccinoUnknownLatteFlatwhiteMelangeCortadoGalaoKopisusuAffogatoMacchiatoBicerinBombónMochaCaphesuadaCarajilloEspressoEiskaffeeFrappuccinoIcedIndianfilterKopiluwakKopitubrukViennaYuanyangNoneAmericanoCubanoZorroDoppioRomanoGuillermoRistrettoAntoccinoAulaitKulaMelyaMarocchinoMielMazagranPalazzoMediciToubaPocilloMistoFrappe"

var _InternalState_index = [...]uint16{0, 10, 15, 19, 26, 33, 38, 43, 53, 60, 65, 74, 81, 88, 93, 101, 109, 118, 125, 132, 137, 147, 156, 164, 173, 184, 188, 200, 209, 219, 225, 233, 237, 246, 252, 257, 263, 269, 278, 287, 296, 302, 306, 311, 321, 325, 333, 340, 346, 351, 358, 363, 369}

func (i InternalState) String() string {
	if i < 0 || i >= InternalState(len(_InternalState_index)-1) {
		return "InternalState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _InternalState_name[_InternalState_index[i]:_InternalState_index[i+1]]
}
