package math

func SatoshisToXBT(sats int) (xbts float64) {

	fSats := float64(sats)
	return fSats / 100000000
}
