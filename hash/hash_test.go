package hash

import (
	"encoding/binary"
	"hash"
	"hash/fnv"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func crossCheck(h hash.Hash64, v any) uint64 {
	h.Reset()
	binary.Write(h, binary.LittleEndian, v)
	return h.Sum64()
}

func TestEnsureHasher(t *testing.T) {
	t.Run("PassNilHash", func(t *testing.T) {
		h := ensureHasher(nil)
		assert.NotNil(t, h)
	})

	t.Run("PassNotNilHash", func(t *testing.T) {
		h := ensureHasher(fnv.New64a())
		assert.NotNil(t, h)
	})
}

func TestHashFuncForBool(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val bool
	}{
		{fnv.New64(), true},
		{fnv.New64(), false},
	}

	for _, tc := range tests {
		hash := HashFuncForBool[bool](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForBoolSlice(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val []bool
	}{
		{fnv.New64(), []bool{false, false}},
		{fnv.New64(), []bool{false, true}},
		{fnv.New64(), []bool{true, false}},
		{fnv.New64(), []bool{true, true}},
	}

	for _, tc := range tests {
		hash := HashFuncForBoolSlice[[]bool](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForInt8(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val int8
	}{
		{fnv.New64(), math.MinInt8},
		{fnv.New64(), -64},
		{fnv.New64(), 0},
		{fnv.New64(), 64},
		{fnv.New64(), math.MaxInt8},
	}

	for _, tc := range tests {
		hash := HashFuncForInt8[int8](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForInt8Slice(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val []int8
	}{
		{fnv.New64(), []int8{-32, 40, -69, 50, -121, -63}},
		{fnv.New64(), []int8{-76, 51, -92, -36, -14, 55}},
		{fnv.New64(), []int8{89, 27, 62, 8, -16, 126}},
		{fnv.New64(), []int8{-30, -100, -26, 38, -38, -56}},
	}

	for _, tc := range tests {
		hash := HashFuncForInt8Slice[[]int8](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForInt16(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val int16
	}{
		{fnv.New64(), math.MinInt16},
		{fnv.New64(), -1024},
		{fnv.New64(), 0},
		{fnv.New64(), 1024},
		{fnv.New64(), math.MaxInt16},
	}

	for _, tc := range tests {
		hash := HashFuncForInt16[int16](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForInt16Slice(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val []int16
	}{
		{fnv.New64(), []int16{-28832, -32600, -25541, 26290, 28423, -6847}},
		{fnv.New64(), []int16{13364, 8883, 15908, 25948, 16754, -24649}},
		{fnv.New64(), []int16{-4135, -24933, -3138, -17016, -20112, 14078}},
		{fnv.New64(), []int16{-7070, 24604, -3226, 27558, -26278, -6328}},
	}

	for _, tc := range tests {
		hash := HashFuncForInt16Slice[[]int16](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForInt32(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val int32
	}{
		{fnv.New64(), math.MinInt32},
		{fnv.New64(), -4096},
		{fnv.New64(), 0},
		{fnv.New64(), 4096},
		{fnv.New64(), math.MaxInt32},
	}

	for _, tc := range tests {
		hash := HashFuncForInt32[int32](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForInt32Slice(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val []int32
	}{
		{fnv.New64(), []int32{-1186127860, -1930235437, -1445935215, 2017135249, 1248588448, 627997297}},
		{fnv.New64(), []int32{1450795723, -72601595, -730176221, -472406877, -355100387, -1542010938}},
		{fnv.New64(), []int32{614031270, -1714716452, 308087659, 23346265, 2067216393, -2088490033}},
		{fnv.New64(), []int32{-1357240991, 1387784369, -1410745432, -1961236436, -1378785059, 1929742318}},
	}

	for _, tc := range tests {
		hash := HashFuncForInt32Slice[[]int32](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForInt64(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val int64
	}{
		{fnv.New64(), math.MinInt64},
		{fnv.New64(), -65536},
		{fnv.New64(), 0},
		{fnv.New64(), 65536},
		{fnv.New64(), math.MaxInt64},
	}

	for _, tc := range tests {
		hash := HashFuncForInt64[int64](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForInt64Slice(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val []int64
	}{
		{fnv.New64(), []int64{9203702600390295577, 8036172665949622051, -8002370904357608127, -4557246544839785065, -8106624152729499065, -2030698037618462150}},
		{fnv.New64(), []int64{9010822915581722444, -4104213585721682217, -475583987999412205, -8563936895678788925, 5161941888241203024, 3977297171134468538}},
		{fnv.New64(), []int64{-6610219873607910607, -4208573241254926906, 4769479065153071691, 7376262298314016284, -4517096875408669349, -5501145984209057490}},
		{fnv.New64(), []int64{-1236610016984957453, 2484802110959232407, -6214509794517465184, -4798196200802047101, 419332496410755629, -2912749951390041006}},
	}

	for _, tc := range tests {
		hash := HashFuncForInt64Slice[[]int64](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForInt(t *testing.T) {
	tests := []struct {
		h            hash.Hash64
		val          int
		expectedHash uint64
	}{
		{fnv.New64(), math.MinInt, 0xa8c7f832281a3945},
		{fnv.New64(), -1048576, 0xc0e08d25d517bb2e},
		{fnv.New64(), 0, 0xa8c7f832281a39c5},
		{fnv.New64(), 1048576, 0xddd9c58a53239395},
		{fnv.New64(), math.MaxInt, 0xd65de7467f38b9ed},
	}

	for _, tc := range tests {
		hash := HashFuncForInt[int](tc.h)(tc.val)

		assert.Equal(t, tc.expectedHash, hash)
	}
}

func TestHashFuncForIntSlice(t *testing.T) {
	tests := []struct {
		h            hash.Hash64
		val          []int
		expectedHash uint64
	}{
		{
			fnv.New64(),
			[]int{-4621520736659628020, 2530432221995272849, -4576219449590091309, 5222186584675971744, -5205285703879964783, 7491004887823448689},
			0x1f472030b6d75ac1,
		},
		{
			fnv.New64(),
			[]int{2086826187330921635, -2278623272419892533, 8208023018045544733, -7726872168903200763, 1787682448534325190, -4053312076364749533},
			0x9e49a4c671f1bfa,
		},
		{
			fnv.New64(),
			[]int{-4717960579063914586, 3287695384476138585, -6288961753144655652, 8985580042855069705, -2052106792860841109, 3474165364659858383},
			0xfa93cfd2469a89fd,
		},
		{
			fnv.New64(),
			[]int{5562767588139329580, -4281968447839394463, 1988648585567234269, -4361771299686061903, 1304187130113001454 - 6642401092734174296},
			0xcbcd458d810c5111,
		},
	}

	for _, tc := range tests {
		hash := HashFuncForIntSlice[[]int](tc.h)(tc.val)

		assert.Equal(t, tc.expectedHash, hash)
	}
}

func TestHashFuncForUint8(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val uint8
	}{
		{fnv.New64(), 0},
		{fnv.New64(), 64},
		{fnv.New64(), 128},
		{fnv.New64(), math.MaxUint8},
	}

	for _, tc := range tests {
		hash := HashFuncForUint8[uint8](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForUint8Slice(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val []uint8
	}{
		{fnv.New64(), []uint8{96, 168, 59, 178, 7, 65}},
		{fnv.New64(), []uint8{52, 179, 36, 92, 114, 183}},
		{fnv.New64(), []uint8{217, 155, 190, 136, 112, 254}},
		{fnv.New64(), []uint8{98, 28, 102, 166, 90, 72}},
	}

	for _, tc := range tests {
		hash := HashFuncForUint8Slice[[]uint8](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForUint16(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val uint16
	}{
		{fnv.New64(), 0},
		{fnv.New64(), 1024},
		{fnv.New64(), 2048},
		{fnv.New64(), math.MaxUint16},
	}

	for _, tc := range tests {
		hash := HashFuncForUint16[uint16](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForUint16Slice(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val []uint16
	}{
		{fnv.New64(), []uint16{3936, 168, 7227, 59058, 61191, 25921}},
		{fnv.New64(), []uint16{46132, 41651, 48676, 58716, 49522, 8119}},
		{fnv.New64(), []uint16{28633, 7835, 29630, 15752, 12656, 46846}},
		{fnv.New64(), []uint16{25698, 57372, 29542, 60326, 6490, 26440}},
	}

	for _, tc := range tests {
		hash := HashFuncForUint16Slice[[]uint16](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForUint32(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val uint32
	}{
		{fnv.New64(), 0},
		{fnv.New64(), 4096},
		{fnv.New64(), 8192},
		{fnv.New64(), math.MaxUint32},
	}

	for _, tc := range tests {
		hash := HashFuncForUint32[uint32](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForUint32Slice(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val []uint32
	}{
		{fnv.New64(), []uint32{961355788, 217248211, 701548433, 4164618897, 3396072096, 2775480945}},
		{fnv.New64(), []uint32{3598279371, 2074882053, 1417307427, 1675076771, 1792383261, 605472710}},
		{fnv.New64(), []uint32{2761514918, 432767196, 2455571307, 2170829913, 4214700041, 58993615}},
		{fnv.New64(), []uint32{790242657, 3535268017, 736738216, 186247212, 768698589, 4077225966}},
	}

	for _, tc := range tests {
		hash := HashFuncForUint32Slice[[]uint32](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForUint64(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val uint64
	}{
		{fnv.New64(), 0},
		{fnv.New64(), 65536},
		{fnv.New64(), 262144},
		{fnv.New64(), math.MaxUint64},
	}

	for _, tc := range tests {
		hash := HashFuncForUint64[uint64](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForUint64Slice(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val []uint64
	}{
		{fnv.New64(), []uint64{4601851300195147788, 13870524624119460307, 13241458369829586833, 11753804258850048657, 5222186584675971744, 16714376924678224497}},
		{fnv.New64(), []uint64{6944748764434883275, 1496499867951575045, 14393431997344802083, 11310198224185697443, 17431395054900320541, 1787682448534325190}},
		{fnv.New64(), []uint64{13728783494645637030, 2934410283710120156, 16394637280848710507, 12511067421330914393, 18208952079709845513, 3474165364659858383}},
		{fnv.New64(), []uint64{4941403589015381345, 14084972774023489713, 2580970944120601512, 14786139624994105388, 11212020622422010077, 10527559166967777262}},
	}

	for _, tc := range tests {
		hash := HashFuncForUint64Slice[[]uint64](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForUintptr(t *testing.T) {
	tests := []struct {
		h            hash.Hash64
		val          uintptr
		expectedHash uint64
	}{
		{fnv.New64(), 17679884386518147865, 0x33f881e22147ae0e},
		{fnv.New64(), 1350474133962131410, 0x7bc0f5be41cd224},
		{fnv.New64(), 6950455016768307630, 0xeb9197d4cdf4713a},
		{fnv.New64(), 1253960824125502454, 0x792917c3e2b79e71},
	}

	for _, tc := range tests {
		hash := HashFuncForUintptr[uintptr](tc.h)(tc.val)

		assert.Equal(t, tc.expectedHash, hash)
	}
}

func TestHashFuncForUintptrSlice(t *testing.T) {
	tests := []struct {
		h            hash.Hash64
		val          []uintptr
		expectedHash uint64
	}{
		{
			fnv.New64(),
			[]uintptr{17690775450175591588, 15307102153518918142, 2784068335753094752, 14285040365087116476, 11072319963162446655, 9818173382121258855},
			0x3f0cb6cd415b56b3,
		},
		{
			fnv.New64(),
			[]uintptr{16015887064541234754, 3175921209289284192, 9249255765685461474, 7931571905377148184, 4900691200732813320, 18406028322227128020},
			0xf32f899f6c3ea2ff,
		},
		{
			fnv.New64(),
			[]uintptr{14632817385921045723, 15872437402962067597, 9307665909450800122, 4738048589706258401, 7465508200187603062, 11767191803354815518},
			0x25f4490ea2aec08,
		},
		{
			fnv.New64(),
			[]uintptr{13421605491922070479, 9057804508480927086, 4577692780544562546, 4397777384653390049, 4638171652237934993, 16672827934807222316},
			0xe51f292813fd40c2,
		},
	}

	for _, tc := range tests {
		hash := HashFuncForUintptrSlice[[]uintptr](tc.h)(tc.val)

		assert.Equal(t, tc.expectedHash, hash)
	}
}

func TestHashFuncForUint(t *testing.T) {
	tests := []struct {
		h            hash.Hash64
		val          uint
		expectedHash uint64
	}{
		{fnv.New64(), 0, 0xa8c7f832281a39c5},
		{fnv.New64(), 1048576, 0xddd9c58a53239395},
		{fnv.New64(), 4194304, 0xd480c2d17bf4d285},
		{fnv.New64(), math.MaxUint, 0xd65de7467f38b96d},
	}

	for _, tc := range tests {
		hash := HashFuncForUint[uint](tc.h)(tc.val)

		assert.Equal(t, tc.expectedHash, hash)
	}
}

func TestHashFuncForUintSlice(t *testing.T) {
	tests := []struct {
		h            hash.Hash64
		val          []uint
		expectedHash uint64
	}{
		{
			fnv.New64(),
			[]uint{18420584253445099999, 1503914795427679488, 12649918963468568785, 8404867576573022926, 6578170777616368000, 16800641873940274907},
			0xeb9438ca0ef2a192,
		},
		{
			fnv.New64(),
			[]uint{10826496560680066938, 5630992903254390578, 14737334403767407558, 9792097170767072720, 17213770151693931396, 684685762873323315},
			0xab3deeff7c57682a,
		},
		{
			fnv.New64(),
			[]uint{17540597631615965111, 9640978507436659677, 13079401454582967108, 190215522736641748, 17984956834308201458, 11223996373968164989},
			0xb5e13db36e455a77,
		},
		{
			fnv.New64(),
			[]uint{17383009236727847062, 7844017944328667086, 17014663605106000514, 15851945451968287086, 12348676365202506957, 869957720694551175},
			0x6096f266fc28f9,
		},
	}

	for _, tc := range tests {
		hash := HashFuncForUintSlice[[]uint](tc.h)(tc.val)

		assert.Equal(t, tc.expectedHash, hash)
	}
}

func TestHashFuncForFloat32(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val float32
	}{
		{fnv.New64(), math.SmallestNonzeroFloat32},
		{fnv.New64(), -225.763367},
		{fnv.New64(), 0.0},
		{fnv.New64(), 609.844360},
		{fnv.New64(), math.MaxFloat32},
	}

	for _, tc := range tests {
		hash := HashFuncForFloat32[float32](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForFloat32Slice(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val []float32
	}{
		{fnv.New64(), []float32{-2.132507, 7.690552, -128.716431, -451.299988, 132.381226, 624.352783}},
		{fnv.New64(), []float32{505.902344, -675.498291, 121.078003, -547.491699, 779.831299, -612.358154}},
		{fnv.New64(), []float32{-23.044617, -363.701233, 555.020264, -287.094727, 948.437134, -246.660461}},
		{fnv.New64(), []float32{71.496094, 54.191650, -440.341125, 206.232910, -568.780579, -717.199463}},
	}

	for _, tc := range tests {
		hash := HashFuncForFloat32Slice[[]float32](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForFloat64(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val float64
	}{
		{fnv.New64(), math.SmallestNonzeroFloat64},
		{fnv.New64(), -479653.214375},
		{fnv.New64(), 0},
		{fnv.New64(), 510255.375435},
		{fnv.New64(), math.MaxFloat64},
	}

	for _, tc := range tests {
		hash := HashFuncForFloat64[float64](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForFloat64Slice(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val []float64
	}{
		{fnv.New64(), []float64{-2132.564575, 7690.586197, -128716.413711, -451299.977517, 132381.208046, 624352.754696}},
		{fnv.New64(), []float64{505902.339553, -675498.318409, 121078.048209, -547491.702819, 779831.277595, -612358.161117}},
		{fnv.New64(), []float64{-23044.621904, -363701.199088, 555020.271402, -287094.704336, 948437.080701, -246660.472812}},
		{fnv.New64(), []float64{71496.101268, 54191.616199, -440341.139052, 206232.940819, -568780.576644, -717199.496041}},
	}

	for _, tc := range tests {
		hash := HashFuncForFloat64Slice[[]float64](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForComplex64(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val complex64
	}{
		{fnv.New64(), 21.093628 + -865.444519i},
		{fnv.New64(), 96.960571 + 113.996460i},
		{fnv.New64(), 198.922974 + -304.801758i},
		{fnv.New64(), 799.151855 + -265.779297i},
	}

	for _, tc := range tests {
		hash := HashFuncForComplex64[complex64](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForComplex64Slice(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val []complex64
	}{
		{fnv.New64(), []complex64{836.075439 + 319.198730i, -396.301392 + 97.574341i, -599.073364 + -871.023010i, 472.891968 + -311.331848i, -994.387390 + 719.885498i, 62.667969 + 991.171143i}},
		{fnv.New64(), []complex64{172.986450 + 441.786133i, -981.721680 + 27.400513i, 618.823975 + -448.397034i, -89.653259 + 964.098267i, -7.371155 + -46.384033i, 5.743164 + 615.343262i}},
		{fnv.New64(), []complex64{-402.763855 + 833.713745i, -465.684143 + 336.901855i, 507.139648 + -43.720581i, -859.695190 + -934.495911i, 836.147339 + -707.162598i, -6.975281 + -728.090576i}},
		{fnv.New64(), []complex64{-325.706421 + -973.684631i, 257.356201 + -930.497437i, -320.695740 + 712.320923i, 863.462646 + -694.631226i, -27.068787 + -684.384644i, -875.297852 + -526.355408i}},
	}

	for _, tc := range tests {
		hash := HashFuncForComplex64Slice[[]complex64](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForComplex128(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val complex128
	}{
		{fnv.New64(), 602841.153600 + -612222.876161i},
		{fnv.New64(), -373850.755769 + -684413.525167i},
		{fnv.New64(), 346502.604442 + -606823.462103i},
		{fnv.New64(), 715607.556959 + -285332.925894i},
	}

	for _, tc := range tests {
		hash := HashFuncForComplex128[complex128](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForComplex128Slice(t *testing.T) {
	tests := []struct {
		h   hash.Hash64
		val []complex128
	}{
		{fnv.New64(), []complex128{836075.435207 + 319198.681861i, -396301.412406 + 97574.359574i, -599073.328297 + -871023.017853i, 472891.909932 + -311331.865049i, -994387.360994 + 719885.498207i, 62668.009303 + 991171.178758i}},
		{fnv.New64(), []complex128{172986.479880 + 441786.222986i, -981721.680041 + 27400.514860i, 618824.041870 + -448397.016550i, -89653.233483 + 964098.265209i, -7371.108472 + -46384.041090i, 5743.156343 + 615343.253679i}},
		{fnv.New64(), []complex128{-402763.825069 + 833713.812231i, -465684.117015 + 336901.849280i, 507139.685789 + -43720.622509i, -859695.216181 + -934495.893126i, 836147.298213 + -707162.602015i, -6975.269076 + -728090.590054i}},
		{fnv.New64(), []complex128{-325706.417534 + -973684.645112i, 257356.132494 + -930497.409866i, -320695.784318 + 712320.938727i, 863462.698141 + -694631.262919i, -27068.816069 + -684384.650171i, -875297.836951 + -526355.406724i}},
	}

	for _, tc := range tests {
		hash := HashFuncForComplex128Slice[[]complex128](tc.h)(tc.val)
		expectedHash := crossCheck(tc.h, tc.val)

		assert.Equal(t, expectedHash, hash)
	}
}

func TestHashFuncForString(t *testing.T) {
	tests := []struct {
		h            hash.Hash64
		val          string
		expectedHash uint64
	}{
		{fnv.New64(), "Cardinal", 0xd29df436bca7c26d},
		{fnv.New64(), "Blue Jay", 0xfc4bc384ef670779},
		{fnv.New64(), "Peacock", 0x22f3bf65080a9989},
		{fnv.New64(), "Mandarin Duck", 0x1c8c5d4aba378ed8},
		{fnv.New64(), "Scarlet Macaw", 0x32c9e271a5062380},
		{fnv.New64(), "Kingfisher", 0x3443e01f092a67d3},
		{fnv.New64(), "Quetzal", 0xac8fa31b127fbccd},
		{fnv.New64(), "Hummingbird", 0xd8d5a5f62cc2ea3f},
		{fnv.New64(), "Flamingo", 0xd9e38805f051499a},
	}

	for _, tc := range tests {
		hash := HashFuncForString[string](tc.h)(tc.val)

		assert.Equal(t, tc.expectedHash, hash)
	}
}

func TestHashFuncForStringSlice(t *testing.T) {
	tests := []struct {
		h            hash.Hash64
		val          []string
		expectedHash uint64
	}{
		{
			fnv.New64(),
			[]string{"Hi, how're you?", "Hola, ¿cómo estás?", "Salut, comment ça va ?", "Ciao, come stai?", "Hallo, wie geht's?", "こんにちは、お元気ですか？"},
			0x8203d9a15f8eca16,
		},
		{
			fnv.New64(),
			[]string{"Let it go.", "Déjalo ir.", "Laisse-le aller.", "Lascialo andare.", "Lass es los.", "それを放っておけ。"},
			0x57c60d1ea1b2ec27,
		},
		{
			fnv.New64(),
			[]string{"I like meditation.", "Me gusta la meditación.", "J'aime la méditation.", "Mi piace la meditazione.", "Ich mag Meditation.", "私は瞑想が好きです。"},
			0x4b905cdfdc4e14c5,
		},
		{
			fnv.New64(),
			[]string{"Nothing really matters.", "Nada realmente importa.", "Rien n'a vraiment d'importance.", "Niente importa davvero.", "Nichts ist wirklich wichtig.", "何も本当に重要ではない。"},
			0x60bc0f79f72f64a2,
		},
	}

	for _, tc := range tests {
		hash := HashFuncForStringSlice[[]string](tc.h)(tc.val)

		assert.Equal(t, tc.expectedHash, hash)
	}
}
