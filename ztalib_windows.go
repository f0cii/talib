package talib

/* TA_Initialize() initialize the ressources used by TA-Lib. This
 * function must be called once prior to any other functions declared in
 * this file.
 */
func Initialize() error {
	_, err := ta_Initialize()
	return err
}

/* TA_Shutdown() allows to free all ressources used by TA-Lib. Following
 * a shutdown, TA_Initialize() must be called again for re-using TA-Lib.
 *
 * TA_Shutdown() should be called prior to exiting the application code.
 */
func Shutdown() error {
	_, err := ta_Shutdown()
	return err
}

/*
 * TA_ACOS - Vector Trigonometric ACos
 *
 * Input  = double
 * Output = double
 *
 */
func Acos(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Acos(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_AD - Chaikin A/D Line
 *
 * Input  = High, Low, Close, Volume
 * Output = double
 *
 */
func Ad(inHigh []float64, inLow []float64, inClose []float64, inVolume []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_Ad(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), (*float64)(&inVolume[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_ADD - Vector Arithmetic Add
 *
 * Input  = double, double
 * Output = double
 *
 */
func Add(inReal0 []float64, inReal1 []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal0)
	outReal = make([]float64, n)
	ta_Add(0, n-1, (*float64)(&inReal0[0]), (*float64)(&inReal1[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_ADOSC - Chaikin A/D Oscillator
 *
 * Input  = High, Low, Close, Volume
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInFastPeriod:(From 2 to 100000)
 *    Number of period for the fast MA
 *
 * optInSlowPeriod:(From 2 to 100000)
 *    Number of period for the slow MA
 *
 *
 */
func AdOsc(inHigh []float64, inLow []float64, inClose []float64, inVolume []float64, optInFastPeriod int, optInSlowPeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_AdOsc(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), (*float64)(&inVolume[0]), optInFastPeriod, optInSlowPeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_ADX - Average Directional Movement Index
 *
 * Input  = High, Low, Close
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Adx(inHigh []float64, inLow []float64, inClose []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_Adx(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_ADXR - Average Directional Movement Index Rating
 *
 * Input  = High, Low, Close
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Adxr(inHigh []float64, inLow []float64, inClose []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_Adxr(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_APO - Absolute Price Oscillator
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInFastPeriod:(From 2 to 100000)
 *    Number of period for the fast MA
 *
 * optInSlowPeriod:(From 2 to 100000)
 *    Number of period for the slow MA
 *
 * optInMAType:
 *    Type of Moving Average
 *
 *
 */
func Apo(inReal []float64, optInFastPeriod int, optInSlowPeriod int, optInMAType int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Apo(0, n-1, (*float64)(&inReal[0]), optInFastPeriod, optInSlowPeriod, optInMAType, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_AROON - Aroon
 *
 * Input  = High, Low
 * Output = double, double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func AroOn(inHigh []float64, inLow []float64, optInTimePeriod int) (outAroonDown []float64, outAroonUp []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outAroonDown = make([]float64, n)
	outAroonUp = make([]float64, n)
	ta_AroOn(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outAroonDown[0]), (*float64)(&outAroonUp[0]))
	outAroonDown = append(make([]float64, outBegIdx), outAroonDown[:outNBElement]...)
	outAroonUp = append(make([]float64, outBegIdx), outAroonUp[:outNBElement]...)
	return
}

/*
 * TA_AROONOSC - Aroon Oscillator
 *
 * Input  = High, Low
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func AroOnOsc(inHigh []float64, inLow []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_AroOnOsc(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_ASIN - Vector Trigonometric ASin
 *
 * Input  = double
 * Output = double
 *
 */
func ASin(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_ASin(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_ATAN - Vector Trigonometric ATan
 *
 * Input  = double
 * Output = double
 *
 */
func Atan(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Atan(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_ATR - Average True Range
 *
 * Input  = High, Low, Close
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func Atr(inHigh []float64, inLow []float64, inClose []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_Atr(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_AVGPRICE - Average Price
 *
 * Input  = Open, High, Low, Close
 * Output = double
 *
 */
func AvgPrice(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outReal = make([]float64, n)
	ta_AvgPrice(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_BBANDS - Bollinger Bands
 *
 * Input  = double
 * Output = double, double, double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 * optInNbDevUp:(From TA_REAL_MIN to TA_REAL_MAX)
 *    Deviation multiplier for upper band
 *
 * optInNbDevDn:(From TA_REAL_MIN to TA_REAL_MAX)
 *    Deviation multiplier for lower band
 *
 * optInMAType:
 *    Type of Moving Average
 *
 *
 */
func BBands(inReal []float64, optInTimePeriod int, optInNbDevUp float64, optInNbDevDn float64, optInMAType int) (outRealUpperBand []float64, outRealMiddleBand []float64, outRealLowerBand []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outRealUpperBand = make([]float64, n)
	outRealMiddleBand = make([]float64, n)
	outRealLowerBand = make([]float64, n)
	ta_BBands(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, optInNbDevUp, optInNbDevDn, optInMAType, &outBegIdx, &outNBElement, (*float64)(&outRealUpperBand[0]), (*float64)(&outRealMiddleBand[0]), (*float64)(&outRealLowerBand[0]))
	outRealUpperBand = append(make([]float64, outBegIdx), outRealUpperBand[:outNBElement]...)
	outRealMiddleBand = append(make([]float64, outBegIdx), outRealMiddleBand[:outNBElement]...)
	outRealLowerBand = append(make([]float64, outBegIdx), outRealLowerBand[:outNBElement]...)
	return
}

/*
 * TA_BETA - Beta
 *
 * Input  = double, double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func Beta(inReal0 []float64, inReal1 []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal0)
	outReal = make([]float64, n)
	ta_Beta(0, n-1, (*float64)(&inReal0[0]), (*float64)(&inReal1[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_BOP - Balance Of Power
 *
 * Input  = Open, High, Low, Close
 * Output = double
 *
 */
func Bop(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outReal = make([]float64, n)
	ta_Bop(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_CCI - Commodity Channel Index
 *
 * Input  = High, Low, Close
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Cci(inHigh []float64, inLow []float64, inClose []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_Cci(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_CDL2CROWS - Two Crows
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func Cdl2Crows(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_Cdl2Crows(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDL3BLACKCROWS - Three Black Crows
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func Cdl3BlackCrows(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_Cdl3BlackCrows(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDL3INSIDE - Three Inside Up/Down
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func Cdl3InSide(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_Cdl3InSide(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDL3LINESTRIKE - Three-Line Strike
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func Cdl3LineStrike(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_Cdl3LineStrike(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDL3OUTSIDE - Three Outside Up/Down
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func Cdl3OutSide(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_Cdl3OutSide(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDL3STARSINSOUTH - Three Stars In The South
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func Cdl3StarsInSouth(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_Cdl3StarsInSouth(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDL3WHITESOLDIERS - Three Advancing White Soldiers
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func Cdl3WhiteSoldiers(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_Cdl3WhiteSoldiers(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLABANDONEDBABY - Abandoned Baby
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 * Optional Parameters
 * -------------------
 * optInPenetration:(From 0 to TA_REAL_MAX)
 *    Percentage of penetration of a candle within another candle
 *
 *
 */
func CdlAbandOnedBaBy(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64, optInPenetration float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlAbandOnedBaBy(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInPenetration, &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLADVANCEBLOCK - Advance Block
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlAdvanceBlock(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlAdvanceBlock(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLBELTHOLD - Belt-hold
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlBeltHold(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlBeltHold(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLBREAKAWAY - Breakaway
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlBreakaway(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlBreakaway(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLCLOSINGMARUBOZU - Closing Marubozu
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlCloSingMarubozu(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlCloSingMarubozu(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLCONCEALBABYSWALL - Concealing Baby Swallow
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlCOncealBaBySwall(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlCOncealBaBySwall(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLCOUNTERATTACK - Counterattack
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlCounterattack(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlCounterattack(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLDARKCLOUDCOVER - Dark Cloud Cover
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 * Optional Parameters
 * -------------------
 * optInPenetration:(From 0 to TA_REAL_MAX)
 *    Percentage of penetration of a candle within another candle
 *
 *
 */
func CdlDarkCloudCover(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64, optInPenetration float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlDarkCloudCover(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInPenetration, &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLDOJI - Doji
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlDoji(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlDoji(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLDOJISTAR - Doji Star
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlDojiStar(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlDojiStar(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLDRAGONFLYDOJI - Dragonfly Doji
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlDragOnflyDoji(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlDragOnflyDoji(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLENGULFING - Engulfing Pattern
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlEngulfing(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlEngulfing(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLEVENINGDOJISTAR - Evening Doji Star
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 * Optional Parameters
 * -------------------
 * optInPenetration:(From 0 to TA_REAL_MAX)
 *    Percentage of penetration of a candle within another candle
 *
 *
 */
func CdlEveningDojiStar(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64, optInPenetration float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlEveningDojiStar(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInPenetration, &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLEVENINGSTAR - Evening Star
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 * Optional Parameters
 * -------------------
 * optInPenetration:(From 0 to TA_REAL_MAX)
 *    Percentage of penetration of a candle within another candle
 *
 *
 */
func CdlEveningStar(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64, optInPenetration float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlEveningStar(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInPenetration, &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLGAPSIDESIDEWHITE - Up/Down-gap side-by-side white lines
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlGapSidesideWhite(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlGapSidesideWhite(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLGRAVESTONEDOJI - Gravestone Doji
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlGravestOneDoji(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlGravestOneDoji(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLHAMMER - Hammer
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlHammer(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlHammer(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLHANGINGMAN - Hanging Man
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlHangingMan(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlHangingMan(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLHARAMI - Harami Pattern
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlHarami(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlHarami(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLHARAMICROSS - Harami Cross Pattern
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlHaramiCross(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlHaramiCross(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLHIGHWAVE - High-Wave Candle
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlHighWave(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlHighWave(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLHIKKAKE - Hikkake Pattern
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlHikkake(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlHikkake(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLHIKKAKEMOD - Modified Hikkake Pattern
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlHikkakeMod(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlHikkakeMod(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLHOMINGPIGEON - Homing Pigeon
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlHoMingPigeOn(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlHoMingPigeOn(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLIDENTICAL3CROWS - Identical Three Crows
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlIdentical3Crows(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlIdentical3Crows(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLINNECK - In-Neck Pattern
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlinNeck(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlinNeck(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLINVERTEDHAMMER - Inverted Hammer
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlinvertedHammer(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlinvertedHammer(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLKICKING - Kicking
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlKicking(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlKicking(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLKICKINGBYLENGTH - Kicking - bull/bear determined by the longer marubozu
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlKickingByLength(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlKickingByLength(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLLADDERBOTTOM - Ladder Bottom
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlLadderBottom(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlLadderBottom(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLLONGLEGGEDDOJI - Long Legged Doji
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlLOngLeggedDoji(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlLOngLeggedDoji(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLLONGLINE - Long Line Candle
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlLOngLine(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlLOngLine(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLMARUBOZU - Marubozu
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlMarubozu(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlMarubozu(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLMATCHINGLOW - Matching Low
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlMatchingLow(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlMatchingLow(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLMATHOLD - Mat Hold
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 * Optional Parameters
 * -------------------
 * optInPenetration:(From 0 to TA_REAL_MAX)
 *    Percentage of penetration of a candle within another candle
 *
 *
 */
func CdlMatHold(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64, optInPenetration float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlMatHold(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInPenetration, &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLMORNINGDOJISTAR - Morning Doji Star
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 * Optional Parameters
 * -------------------
 * optInPenetration:(From 0 to TA_REAL_MAX)
 *    Percentage of penetration of a candle within another candle
 *
 *
 */
func CdlMorningDojiStar(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64, optInPenetration float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlMorningDojiStar(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInPenetration, &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLMORNINGSTAR - Morning Star
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 * Optional Parameters
 * -------------------
 * optInPenetration:(From 0 to TA_REAL_MAX)
 *    Percentage of penetration of a candle within another candle
 *
 *
 */
func CdlMorningStar(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64, optInPenetration float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlMorningStar(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInPenetration, &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLONNECK - On-Neck Pattern
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlOnNeck(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlOnNeck(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLPIERCING - Piercing Pattern
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlPiercing(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlPiercing(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLRICKSHAWMAN - Rickshaw Man
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlRickshawMan(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlRickshawMan(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLRISEFALL3METHODS - Rising/Falling Three Methods
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlRiseFall3Methods(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlRiseFall3Methods(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLSEPARATINGLINES - Separating Lines
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlSeparatingLines(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlSeparatingLines(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLSHOOTINGSTAR - Shooting Star
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlShootingStar(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlShootingStar(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLSHORTLINE - Short Line Candle
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlShortLine(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlShortLine(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLSPINNINGTOP - Spinning Top
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlSpinningTop(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlSpinningTop(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLSTALLEDPATTERN - Stalled Pattern
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlStalledPattern(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlStalledPattern(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLSTICKSANDWICH - Stick Sandwich
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlStickSandwich(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlStickSandwich(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLTAKURI - Takuri (Dragonfly Doji with very long lower shadow)
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlTakuri(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlTakuri(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLTASUKIGAP - Tasuki Gap
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlTasukiGap(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlTasukiGap(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLTHRUSTING - Thrusting Pattern
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlThrusting(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlThrusting(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLTRISTAR - Tristar Pattern
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdltriStar(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdltriStar(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLUNIQUE3RIVER - Unique 3 River
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlUnique3River(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlUnique3River(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLUPSIDEGAP2CROWS - Upside Gap Two Crows
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlupSideGap2Crows(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlupSideGap2Crows(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CDLXSIDEGAP3METHODS - Upside/Downside Gap Three Methods
 *
 * Input  = Open, High, Low, Close
 * Output = int
 *
 */
func CdlxSideGap3Methods(inOpen []float64, inHigh []float64, inLow []float64, inClose []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inOpen)
	outInteger = make([]int, n)
	ta_CdlxSideGap3Methods(0, n-1, (*float64)(&inOpen[0]), (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_CEIL - Vector Ceil
 *
 * Input  = double
 * Output = double
 *
 */
func Ceil(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Ceil(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_CMO - Chande Momentum Oscillator
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Cmo(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Cmo(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_CORREL - Pearson's Correlation Coefficient (r)
 *
 * Input  = double, double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func Correl(inReal0 []float64, inReal1 []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal0)
	outReal = make([]float64, n)
	ta_Correl(0, n-1, (*float64)(&inReal0[0]), (*float64)(&inReal1[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_COS - Vector Trigonometric Cos
 *
 * Input  = double
 * Output = double
 *
 */
func Cos(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Cos(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_COSH - Vector Trigonometric Cosh
 *
 * Input  = double
 * Output = double
 *
 */
func Cosh(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Cosh(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_DEMA - Double Exponential Moving Average
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Dema(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Dema(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_DIV - Vector Arithmetic Div
 *
 * Input  = double, double
 * Output = double
 *
 */
func Div(inReal0 []float64, inReal1 []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal0)
	outReal = make([]float64, n)
	ta_Div(0, n-1, (*float64)(&inReal0[0]), (*float64)(&inReal1[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_DX - Directional Movement Index
 *
 * Input  = High, Low, Close
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Dx(inHigh []float64, inLow []float64, inClose []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_Dx(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_EMA - Exponential Moving Average
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Ema(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Ema(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_EXP - Vector Arithmetic Exp
 *
 * Input  = double
 * Output = double
 *
 */
func Exp(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Exp(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_FLOOR - Vector Floor
 *
 * Input  = double
 * Output = double
 *
 */
func Floor(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Floor(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_HT_DCPERIOD - Hilbert Transform - Dominant Cycle Period
 *
 * Input  = double
 * Output = double
 *
 */
func HtDcPeriod(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_HtDcPeriod(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_HT_DCPHASE - Hilbert Transform - Dominant Cycle Phase
 *
 * Input  = double
 * Output = double
 *
 */
func HtDcPhase(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_HtDcPhase(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_HT_PHASOR - Hilbert Transform - Phasor Components
 *
 * Input  = double
 * Output = double, double
 *
 */
func HtPhasor(inReal []float64) (outInPhase []float64, outQuadrature []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outInPhase = make([]float64, n)
	outQuadrature = make([]float64, n)
	ta_HtPhasor(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outInPhase[0]), (*float64)(&outQuadrature[0]))
	outInPhase = append(make([]float64, outBegIdx), outInPhase[:outNBElement]...)
	outQuadrature = append(make([]float64, outBegIdx), outQuadrature[:outNBElement]...)
	return
}

/*
 * TA_HT_SINE - Hilbert Transform - SineWave
 *
 * Input  = double
 * Output = double, double
 *
 */
func HtSine(inReal []float64) (outSine []float64, outLeadSine []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outSine = make([]float64, n)
	outLeadSine = make([]float64, n)
	ta_HtSine(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outSine[0]), (*float64)(&outLeadSine[0]))
	outSine = append(make([]float64, outBegIdx), outSine[:outNBElement]...)
	outLeadSine = append(make([]float64, outBegIdx), outLeadSine[:outNBElement]...)
	return
}

/*
 * TA_HT_TRENDLINE - Hilbert Transform - Instantaneous Trendline
 *
 * Input  = double
 * Output = double
 *
 */
func HtTrendLine(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_HtTrendLine(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_HT_TRENDMODE - Hilbert Transform - Trend vs Cycle Mode
 *
 * Input  = double
 * Output = int
 *
 */
func HtTrendMode(inReal []float64) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outInteger = make([]int, n)
	ta_HtTrendMode(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_KAMA - Kaufman Adaptive Moving Average
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Kama(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Kama(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_LINEARREG - Linear Regression
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func LinearReg(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_LinearReg(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_LINEARREG_ANGLE - Linear Regression Angle
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func LinearRegAngle(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_LinearRegAngle(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_LINEARREG_INTERCEPT - Linear Regression Intercept
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func LinearRegIntercept(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_LinearRegIntercept(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_LINEARREG_SLOPE - Linear Regression Slope
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func LinearRegSlope(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_LinearRegSlope(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_LN - Vector Log Natural
 *
 * Input  = double
 * Output = double
 *
 */
func Ln(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Ln(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_LOG10 - Vector Log10
 *
 * Input  = double
 * Output = double
 *
 */
func Log10(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Log10(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_MA - Moving average
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 * optInMAType:
 *    Type of Moving Average
 *
 *
 */
func Ma(inReal []float64, optInTimePeriod int, optInMAType int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Ma(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, optInMAType, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_MACD - Moving Average Convergence/Divergence
 *
 * Input  = double
 * Output = double, double, double
 *
 * Optional Parameters
 * -------------------
 * optInFastPeriod:(From 2 to 100000)
 *    Number of period for the fast MA
 *
 * optInSlowPeriod:(From 2 to 100000)
 *    Number of period for the slow MA
 *
 * optInSignalPeriod:(From 1 to 100000)
 *    Smoothing for the signal line (nb of period)
 *
 *
 */
func Macd(inReal []float64, optInFastPeriod int, optInSlowPeriod int, optInSignalPeriod int) (outMACD []float64, outMACDSignal []float64, outMACDHist []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outMACD = make([]float64, n)
	outMACDSignal = make([]float64, n)
	outMACDHist = make([]float64, n)
	ta_Macd(0, n-1, (*float64)(&inReal[0]), optInFastPeriod, optInSlowPeriod, optInSignalPeriod, &outBegIdx, &outNBElement, (*float64)(&outMACD[0]), (*float64)(&outMACDSignal[0]), (*float64)(&outMACDHist[0]))
	outMACD = append(make([]float64, outBegIdx), outMACD[:outNBElement]...)
	outMACDSignal = append(make([]float64, outBegIdx), outMACDSignal[:outNBElement]...)
	outMACDHist = append(make([]float64, outBegIdx), outMACDHist[:outNBElement]...)
	return
}

/*
 * TA_MACDEXT - MACD with controllable MA type
 *
 * Input  = double
 * Output = double, double, double
 *
 * Optional Parameters
 * -------------------
 * optInFastPeriod:(From 2 to 100000)
 *    Number of period for the fast MA
 *
 * optInFastMAType:
 *    Type of Moving Average for fast MA
 *
 * optInSlowPeriod:(From 2 to 100000)
 *    Number of period for the slow MA
 *
 * optInSlowMAType:
 *    Type of Moving Average for slow MA
 *
 * optInSignalPeriod:(From 1 to 100000)
 *    Smoothing for the signal line (nb of period)
 *
 * optInSignalMAType:
 *    Type of Moving Average for signal line
 *
 *
 */
func MacdExt(inReal []float64, optInFastPeriod int, optInFastMAType int, optInSlowPeriod int, optInSlowMAType int, optInSignalPeriod int, optInSignalMAType int) (outMACD []float64, outMACDSignal []float64, outMACDHist []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outMACD = make([]float64, n)
	outMACDSignal = make([]float64, n)
	outMACDHist = make([]float64, n)
	ta_MacdExt(0, n-1, (*float64)(&inReal[0]), optInFastPeriod, optInFastMAType, optInSlowPeriod, optInSlowMAType, optInSignalPeriod, optInSignalMAType, &outBegIdx, &outNBElement, (*float64)(&outMACD[0]), (*float64)(&outMACDSignal[0]), (*float64)(&outMACDHist[0]))
	outMACD = append(make([]float64, outBegIdx), outMACD[:outNBElement]...)
	outMACDSignal = append(make([]float64, outBegIdx), outMACDSignal[:outNBElement]...)
	outMACDHist = append(make([]float64, outBegIdx), outMACDHist[:outNBElement]...)
	return
}

/*
 * TA_MACDFIX - Moving Average Convergence/Divergence Fix 12/26
 *
 * Input  = double
 * Output = double, double, double
 *
 * Optional Parameters
 * -------------------
 * optInSignalPeriod:(From 1 to 100000)
 *    Smoothing for the signal line (nb of period)
 *
 *
 */
func MacdFix(inReal []float64, optInSignalPeriod int) (outMACD []float64, outMACDSignal []float64, outMACDHist []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outMACD = make([]float64, n)
	outMACDSignal = make([]float64, n)
	outMACDHist = make([]float64, n)
	ta_MacdFix(0, n-1, (*float64)(&inReal[0]), optInSignalPeriod, &outBegIdx, &outNBElement, (*float64)(&outMACD[0]), (*float64)(&outMACDSignal[0]), (*float64)(&outMACDHist[0]))
	outMACD = append(make([]float64, outBegIdx), outMACD[:outNBElement]...)
	outMACDSignal = append(make([]float64, outBegIdx), outMACDSignal[:outNBElement]...)
	outMACDHist = append(make([]float64, outBegIdx), outMACDHist[:outNBElement]...)
	return
}

/*
 * TA_MAMA - MESA Adaptive Moving Average
 *
 * Input  = double
 * Output = double, double
 *
 * Optional Parameters
 * -------------------
 * optInFastLimit:(From 0.01 to 0.99)
 *    Upper limit use in the adaptive algorithm
 *
 * optInSlowLimit:(From 0.01 to 0.99)
 *    Lower limit use in the adaptive algorithm
 *
 *
 */
func Mama(inReal []float64, optInFastLimit float64, optInSlowLimit float64) (outMAMA []float64, outFAMA []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outMAMA = make([]float64, n)
	outFAMA = make([]float64, n)
	ta_Mama(0, n-1, (*float64)(&inReal[0]), optInFastLimit, optInSlowLimit, &outBegIdx, &outNBElement, (*float64)(&outMAMA[0]), (*float64)(&outFAMA[0]))
	outMAMA = append(make([]float64, outBegIdx), outMAMA[:outNBElement]...)
	outFAMA = append(make([]float64, outBegIdx), outFAMA[:outNBElement]...)
	return
}

/*
 * TA_MAVP - Moving average with variable period
 *
 * Input  = double, double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInMinPeriod:(From 2 to 100000)
 *    Value less than minimum will be changed to Minimum period
 *
 * optInMaxPeriod:(From 2 to 100000)
 *    Value higher than maximum will be changed to Maximum period
 *
 * optInMAType:
 *    Type of Moving Average
 *
 *
 */
func Mavp(inReal []float64, inPeriods []float64, optInMinPeriod int, optInMaxPeriod int, optInMAType int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Mavp(0, n-1, (*float64)(&inReal[0]), (*float64)(&inPeriods[0]), optInMinPeriod, optInMaxPeriod, optInMAType, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_MAX - Highest value over a specified period
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Max(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Max(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_MAXINDEX - Index of highest value over a specified period
 *
 * Input  = double
 * Output = int
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func MaxIndex(inReal []float64, optInTimePeriod int) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outInteger = make([]int, n)
	ta_MaxIndex(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_MEDPRICE - Median Price
 *
 * Input  = High, Low
 * Output = double
 *
 */
func MedPrice(inHigh []float64, inLow []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_MedPrice(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_MFI - Money Flow Index
 *
 * Input  = High, Low, Close, Volume
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Mfi(inHigh []float64, inLow []float64, inClose []float64, inVolume []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_Mfi(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), (*float64)(&inVolume[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_MIDPOINT - MidPoint over period
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func MidPoint(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_MidPoint(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_MIDPRICE - Midpoint Price over period
 *
 * Input  = High, Low
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func MidPrice(inHigh []float64, inLow []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_MidPrice(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_MIN - Lowest value over a specified period
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Min(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Min(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_MININDEX - Index of lowest value over a specified period
 *
 * Input  = double
 * Output = int
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func MinIndex(inReal []float64, optInTimePeriod int) (outInteger []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outInteger = make([]int, n)
	ta_MinIndex(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*int)(&outInteger[0]))
	outInteger = append(make([]int, outBegIdx), outInteger[:outNBElement]...)
	return
}

/*
 * TA_MINMAX - Lowest and highest values over a specified period
 *
 * Input  = double
 * Output = double, double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func MinMax(inReal []float64, optInTimePeriod int) (outMin []float64, outMax []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outMin = make([]float64, n)
	outMax = make([]float64, n)
	ta_MinMax(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outMin[0]), (*float64)(&outMax[0]))
	outMin = append(make([]float64, outBegIdx), outMin[:outNBElement]...)
	outMax = append(make([]float64, outBegIdx), outMax[:outNBElement]...)
	return
}

/*
 * TA_MINMAXINDEX - Indexes of lowest and highest values over a specified period
 *
 * Input  = double
 * Output = int, int
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func MinMaxIndex(inReal []float64, optInTimePeriod int) (outMinIdx []int, outMaxIdx []int) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outMinIdx = make([]int, n)
	outMaxIdx = make([]int, n)
	ta_MinMaxIndex(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*int)(&outMinIdx[0]), (*int)(&outMaxIdx[0]))
	outMinIdx = append(make([]int, outBegIdx), outMinIdx[:outNBElement]...)
	outMaxIdx = append(make([]int, outBegIdx), outMaxIdx[:outNBElement]...)
	return
}

/*
 * TA_MINUS_DI - Minus Directional Indicator
 *
 * Input  = High, Low, Close
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func MinusDi(inHigh []float64, inLow []float64, inClose []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_MinusDi(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_MINUS_DM - Minus Directional Movement
 *
 * Input  = High, Low
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func MinusDm(inHigh []float64, inLow []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_MinusDm(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_MOM - Momentum
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func Mom(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Mom(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_MULT - Vector Arithmetic Mult
 *
 * Input  = double, double
 * Output = double
 *
 */
func Mult(inReal0 []float64, inReal1 []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal0)
	outReal = make([]float64, n)
	ta_Mult(0, n-1, (*float64)(&inReal0[0]), (*float64)(&inReal1[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_NATR - Normalized Average True Range
 *
 * Input  = High, Low, Close
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func Natr(inHigh []float64, inLow []float64, inClose []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_Natr(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_OBV - On Balance Volume
 *
 * Input  = double, Volume
 * Output = double
 *
 */
func Obv(inReal []float64, inVolume []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Obv(0, n-1, (*float64)(&inReal[0]), (*float64)(&inVolume[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_PLUS_DI - Plus Directional Indicator
 *
 * Input  = High, Low, Close
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func PlusDi(inHigh []float64, inLow []float64, inClose []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_PlusDi(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_PLUS_DM - Plus Directional Movement
 *
 * Input  = High, Low
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func PlusDm(inHigh []float64, inLow []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_PlusDm(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_PPO - Percentage Price Oscillator
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInFastPeriod:(From 2 to 100000)
 *    Number of period for the fast MA
 *
 * optInSlowPeriod:(From 2 to 100000)
 *    Number of period for the slow MA
 *
 * optInMAType:
 *    Type of Moving Average
 *
 *
 */
func Ppo(inReal []float64, optInFastPeriod int, optInSlowPeriod int, optInMAType int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Ppo(0, n-1, (*float64)(&inReal[0]), optInFastPeriod, optInSlowPeriod, optInMAType, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_ROC - Rate of change : ((price/prevPrice)-1)*100
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func Roc(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Roc(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_ROCP - Rate of change Percentage: (price-prevPrice)/prevPrice
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func Rocp(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Rocp(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_ROCR - Rate of change ratio: (price/prevPrice)
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func Rocr(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Rocr(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_ROCR100 - Rate of change ratio 100 scale: (price/prevPrice)*100
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func Rocr100(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Rocr100(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_RSI - Relative Strength Index
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Rsi(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Rsi(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_SAR - Parabolic SAR
 *
 * Input  = High, Low
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInAcceleration:(From 0 to TA_REAL_MAX)
 *    Acceleration Factor used up to the Maximum value
 *
 * optInMaximum:(From 0 to TA_REAL_MAX)
 *    Acceleration Factor Maximum value
 *
 *
 */
func Sar(inHigh []float64, inLow []float64, optInAcceleration float64, optInMaximum float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_Sar(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), optInAcceleration, optInMaximum, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_SAREXT - Parabolic SAR - Extended
 *
 * Input  = High, Low
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInStartValue:(From TA_REAL_MIN to TA_REAL_MAX)
 *    Start value and direction. 0 for Auto, >0 for Long, <0 for Short
 *
 * optInOffsetOnReverse:(From 0 to TA_REAL_MAX)
 *    Percent offset added/removed to initial stop on short/long reversal
 *
 * optInAccelerationInitLong:(From 0 to TA_REAL_MAX)
 *    Acceleration Factor initial value for the Long direction
 *
 * optInAccelerationLong:(From 0 to TA_REAL_MAX)
 *    Acceleration Factor for the Long direction
 *
 * optInAccelerationMaxLong:(From 0 to TA_REAL_MAX)
 *    Acceleration Factor maximum value for the Long direction
 *
 * optInAccelerationInitShort:(From 0 to TA_REAL_MAX)
 *    Acceleration Factor initial value for the Short direction
 *
 * optInAccelerationShort:(From 0 to TA_REAL_MAX)
 *    Acceleration Factor for the Short direction
 *
 * optInAccelerationMaxShort:(From 0 to TA_REAL_MAX)
 *    Acceleration Factor maximum value for the Short direction
 *
 *
 */
func SarExt(inHigh []float64, inLow []float64, optInStartValue float64, optInOffsetOnReverse float64, optInAccelerationInitLong float64, optInAccelerationLong float64, optInAccelerationMaxLong float64, optInAccelerationInitShort float64, optInAccelerationShort float64, optInAccelerationMaxShort float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_SarExt(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), optInStartValue, optInOffsetOnReverse, optInAccelerationInitLong, optInAccelerationLong, optInAccelerationMaxLong, optInAccelerationInitShort, optInAccelerationShort, optInAccelerationMaxShort, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_SIN - Vector Trigonometric Sin
 *
 * Input  = double
 * Output = double
 *
 */
func Sin(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Sin(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_SINH - Vector Trigonometric Sinh
 *
 * Input  = double
 * Output = double
 *
 */
func Sinh(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Sinh(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_SMA - Simple Moving Average
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Sma(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Sma(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_SQRT - Vector Square Root
 *
 * Input  = double
 * Output = double
 *
 */
func Sqrt(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Sqrt(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_STDDEV - Standard Deviation
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 * optInNbDev:(From TA_REAL_MIN to TA_REAL_MAX)
 *    Nb of deviations
 *
 *
 */
func StdDev(inReal []float64, optInTimePeriod int, optInNbDev float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_StdDev(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, optInNbDev, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_STOCH - Stochastic
 *
 * Input  = High, Low, Close
 * Output = double, double
 *
 * Optional Parameters
 * -------------------
 * optInFastK_Period:(From 1 to 100000)
 *    Time period for building the Fast-K line
 *
 * optInSlowK_Period:(From 1 to 100000)
 *    Smoothing for making the Slow-K line. Usually set to 3
 *
 * optInSlowK_MAType:
 *    Type of Moving Average for Slow-K
 *
 * optInSlowD_Period:(From 1 to 100000)
 *    Smoothing for making the Slow-D line
 *
 * optInSlowD_MAType:
 *    Type of Moving Average for Slow-D
 *
 *
 */
func Stoch(inHigh []float64, inLow []float64, inClose []float64, optInFastK_Period int, optInSlowK_Period int, optInSlowK_MAType int, optInSlowD_Period int, optInSlowD_MAType int) (outSlowK []float64, outSlowD []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outSlowK = make([]float64, n)
	outSlowD = make([]float64, n)
	ta_Stoch(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInFastK_Period, optInSlowK_Period, optInSlowK_MAType, optInSlowD_Period, optInSlowD_MAType, &outBegIdx, &outNBElement, (*float64)(&outSlowK[0]), (*float64)(&outSlowD[0]))
	outSlowK = append(make([]float64, outBegIdx), outSlowK[:outNBElement]...)
	outSlowD = append(make([]float64, outBegIdx), outSlowD[:outNBElement]...)
	return
}

/*
 * TA_STOCHF - Stochastic Fast
 *
 * Input  = High, Low, Close
 * Output = double, double
 *
 * Optional Parameters
 * -------------------
 * optInFastK_Period:(From 1 to 100000)
 *    Time period for building the Fast-K line
 *
 * optInFastD_Period:(From 1 to 100000)
 *    Smoothing for making the Fast-D line. Usually set to 3
 *
 * optInFastD_MAType:
 *    Type of Moving Average for Fast-D
 *
 *
 */
func Stochf(inHigh []float64, inLow []float64, inClose []float64, optInFastK_Period int, optInFastD_Period int, optInFastD_MAType int) (outFastK []float64, outFastD []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outFastK = make([]float64, n)
	outFastD = make([]float64, n)
	ta_Stochf(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInFastK_Period, optInFastD_Period, optInFastD_MAType, &outBegIdx, &outNBElement, (*float64)(&outFastK[0]), (*float64)(&outFastD[0]))
	outFastK = append(make([]float64, outBegIdx), outFastK[:outNBElement]...)
	outFastD = append(make([]float64, outBegIdx), outFastD[:outNBElement]...)
	return
}

/*
 * TA_STOCHRSI - Stochastic Relative Strength Index
 *
 * Input  = double
 * Output = double, double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 * optInFastK_Period:(From 1 to 100000)
 *    Time period for building the Fast-K line
 *
 * optInFastD_Period:(From 1 to 100000)
 *    Smoothing for making the Fast-D line. Usually set to 3
 *
 * optInFastD_MAType:
 *    Type of Moving Average for Fast-D
 *
 *
 */
func StochRsi(inReal []float64, optInTimePeriod int, optInFastK_Period int, optInFastD_Period int, optInFastD_MAType int) (outFastK []float64, outFastD []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outFastK = make([]float64, n)
	outFastD = make([]float64, n)
	ta_StochRsi(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, optInFastK_Period, optInFastD_Period, optInFastD_MAType, &outBegIdx, &outNBElement, (*float64)(&outFastK[0]), (*float64)(&outFastD[0]))
	outFastK = append(make([]float64, outBegIdx), outFastK[:outNBElement]...)
	outFastD = append(make([]float64, outBegIdx), outFastD[:outNBElement]...)
	return
}

/*
 * TA_SUB - Vector Arithmetic Substraction
 *
 * Input  = double, double
 * Output = double
 *
 */
func Sub(inReal0 []float64, inReal1 []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal0)
	outReal = make([]float64, n)
	ta_Sub(0, n-1, (*float64)(&inReal0[0]), (*float64)(&inReal1[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_SUM - Summation
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Sum(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Sum(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_T3 - Triple Exponential Moving Average (T3)
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 * optInVFactor:(From 0 to 1)
 *    Volume Factor
 *
 *
 */
func T3(inReal []float64, optInTimePeriod int, optInVFactor float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_T3(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, optInVFactor, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_TAN - Vector Trigonometric Tan
 *
 * Input  = double
 * Output = double
 *
 */
func Tan(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Tan(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_TANH - Vector Trigonometric Tanh
 *
 * Input  = double
 * Output = double
 *
 */
func Tanh(inReal []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Tanh(0, n-1, (*float64)(&inReal[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_TEMA - Triple Exponential Moving Average
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Tema(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Tema(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_TRANGE - True Range
 *
 * Input  = High, Low, Close
 * Output = double
 *
 */
func Trange(inHigh []float64, inLow []float64, inClose []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_Trange(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_TRIMA - Triangular Moving Average
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Trima(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Trima(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_TRIX - 1-day Rate-Of-Change (ROC) of a Triple Smooth EMA
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 *
 */
func Trix(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Trix(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_TSF - Time Series Forecast
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Tsf(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Tsf(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_TYPPRICE - Typical Price
 *
 * Input  = High, Low, Close
 * Output = double
 *
 */
func TypPrice(inHigh []float64, inLow []float64, inClose []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_TypPrice(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_ULTOSC - Ultimate Oscillator
 *
 * Input  = High, Low, Close
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod1:(From 1 to 100000)
 *    Number of bars for 1st period.
 *
 * optInTimePeriod2:(From 1 to 100000)
 *    Number of bars fro 2nd period
 *
 * optInTimePeriod3:(From 1 to 100000)
 *    Number of bars for 3rd period
 *
 *
 */
func UltOsc(inHigh []float64, inLow []float64, inClose []float64, optInTimePeriod1 int, optInTimePeriod2 int, optInTimePeriod3 int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_UltOsc(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInTimePeriod1, optInTimePeriod2, optInTimePeriod3, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_VAR - Variance
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 1 to 100000)
 *    Number of period
 *
 * optInNbDev:(From TA_REAL_MIN to TA_REAL_MAX)
 *    Nb of deviations
 *
 *
 */
func Var(inReal []float64, optInTimePeriod int, optInNbDev float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Var(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, optInNbDev, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_WCLPRICE - Weighted Close Price
 *
 * Input  = High, Low, Close
 * Output = double
 *
 */
func WclPrice(inHigh []float64, inLow []float64, inClose []float64) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_WclPrice(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_WILLR - Williams' %R
 *
 * Input  = High, Low, Close
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Willr(inHigh []float64, inLow []float64, inClose []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inHigh)
	outReal = make([]float64, n)
	ta_Willr(0, n-1, (*float64)(&inHigh[0]), (*float64)(&inLow[0]), (*float64)(&inClose[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}

/*
 * TA_WMA - Weighted Moving Average
 *
 * Input  = double
 * Output = double
 *
 * Optional Parameters
 * -------------------
 * optInTimePeriod:(From 2 to 100000)
 *    Number of period
 *
 *
 */
func Wma(inReal []float64, optInTimePeriod int) (outReal []float64) {
	var outBegIdx int
	var outNBElement int
	n := len(inReal)
	outReal = make([]float64, n)
	ta_Wma(0, n-1, (*float64)(&inReal[0]), optInTimePeriod, &outBegIdx, &outNBElement, (*float64)(&outReal[0]))
	outReal = append(make([]float64, outBegIdx), outReal[:outNBElement]...)
	return
}
