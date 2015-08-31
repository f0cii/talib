#ifndef TA_EXPORT_H
#define TA_EXPORT_H

#include "ta_func.h"
#include "ta_abstract.h"

#define TA_EXPORT extern "C" __declspec(dllexport)

TA_EXPORT TA_RetCode Z_Initialize()
{
	return TA_Initialize();
}

TA_EXPORT TA_RetCode Z_Shutdown()
{
	return TA_Shutdown();
}

TA_EXPORT TA_RetCode Z_Acos(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_ACOS(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Ad(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, const double *inVolume, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_AD(startIdx, endIdx, inHigh, inLow, inClose, inVolume, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Add(int startIdx, int endIdx, const double *inReal0, const double *inReal1, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_ADD(startIdx, endIdx, inReal0, inReal1, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_AdOsc(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, const double *inVolume, int optInFastPeriod, /* From 2 to 100000 */ int optInSlowPeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_ADOSC(startIdx, endIdx, inHigh, inLow, inClose, inVolume, optInFastPeriod, optInSlowPeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Adx(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_ADX(startIdx, endIdx, inHigh, inLow, inClose, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Adxr(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_ADXR(startIdx, endIdx, inHigh, inLow, inClose, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Apo(int startIdx, int endIdx, const double *inReal, int optInFastPeriod, /* From 2 to 100000 */ int optInSlowPeriod, /* From 2 to 100000 */ TA_MAType optInMAType, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_APO(startIdx, endIdx, inReal, optInFastPeriod, optInSlowPeriod, optInMAType, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_AroOn(int startIdx, int endIdx, const double *inHigh, const double *inLow, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outAroonDown, double *outAroonUp) {
    return TA_AROON(startIdx, endIdx, inHigh, inLow, optInTimePeriod, outBegIdx, outNBElement, outAroonDown, outAroonUp);
}

TA_EXPORT TA_RetCode Z_AroOnOsc(int startIdx, int endIdx, const double *inHigh, const double *inLow, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_AROONOSC(startIdx, endIdx, inHigh, inLow, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_ASin(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_ASIN(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Atan(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_ATAN(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Atr(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_ATR(startIdx, endIdx, inHigh, inLow, inClose, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_AvgPrice(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_AVGPRICE(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_BBands(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ double optInNbDevUp, /* From TA_REAL_MIN to TA_REAL_MAX */ double optInNbDevDn, /* From TA_REAL_MIN to TA_REAL_MAX */ TA_MAType optInMAType, int *outBegIdx, int *outNBElement, double *outRealUpperBand, double *outRealMiddleBand, double *outRealLowerBand) {
    return TA_BBANDS(startIdx, endIdx, inReal, optInTimePeriod, optInNbDevUp, optInNbDevDn, optInMAType, outBegIdx, outNBElement, outRealUpperBand, outRealMiddleBand, outRealLowerBand);
}

TA_EXPORT TA_RetCode Z_Beta(int startIdx, int endIdx, const double *inReal0, const double *inReal1, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_BETA(startIdx, endIdx, inReal0, inReal1, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Bop(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_BOP(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Cci(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_CCI(startIdx, endIdx, inHigh, inLow, inClose, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Cdl2Crows(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDL2CROWS(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_Cdl3BlackCrows(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDL3BLACKCROWS(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_Cdl3InSide(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDL3INSIDE(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_Cdl3LineStrike(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDL3LINESTRIKE(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_Cdl3OutSide(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDL3OUTSIDE(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_Cdl3StarsInSouth(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDL3STARSINSOUTH(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_Cdl3WhiteSoldiers(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDL3WHITESOLDIERS(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlAbandOnedBaBy(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, double optInPenetration, /* From 0 to TA_REAL_MAX */ int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLABANDONEDBABY(startIdx, endIdx, inOpen, inHigh, inLow, inClose, optInPenetration, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlAdvanceBlock(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLADVANCEBLOCK(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlBeltHold(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLBELTHOLD(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlBreakaway(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLBREAKAWAY(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlCloSingMarubozu(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLCLOSINGMARUBOZU(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlCOncealBaBySwall(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLCONCEALBABYSWALL(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlCounterattack(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLCOUNTERATTACK(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlDarkCloudCover(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, double optInPenetration, /* From 0 to TA_REAL_MAX */ int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLDARKCLOUDCOVER(startIdx, endIdx, inOpen, inHigh, inLow, inClose, optInPenetration, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlDoji(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLDOJI(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlDojiStar(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLDOJISTAR(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlDragOnflyDoji(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLDRAGONFLYDOJI(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlEngulfing(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLENGULFING(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlEveningDojiStar(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, double optInPenetration, /* From 0 to TA_REAL_MAX */ int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLEVENINGDOJISTAR(startIdx, endIdx, inOpen, inHigh, inLow, inClose, optInPenetration, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlEveningStar(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, double optInPenetration, /* From 0 to TA_REAL_MAX */ int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLEVENINGSTAR(startIdx, endIdx, inOpen, inHigh, inLow, inClose, optInPenetration, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlGapSidesideWhite(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLGAPSIDESIDEWHITE(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlGravestOneDoji(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLGRAVESTONEDOJI(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlHammer(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLHAMMER(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlHangingMan(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLHANGINGMAN(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlHarami(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLHARAMI(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlHaramiCross(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLHARAMICROSS(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlHighWave(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLHIGHWAVE(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlHikkake(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLHIKKAKE(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlHikkakeMod(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLHIKKAKEMOD(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlHoMingPigeOn(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLHOMINGPIGEON(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlIdentical3Crows(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLIDENTICAL3CROWS(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlinNeck(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLINNECK(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlinvertedHammer(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLINVERTEDHAMMER(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlKicking(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLKICKING(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlKickingByLength(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLKICKINGBYLENGTH(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlLadderBottom(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLLADDERBOTTOM(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlLOngLeggedDoji(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLLONGLEGGEDDOJI(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlLOngLine(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLLONGLINE(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlMarubozu(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLMARUBOZU(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlMatchingLow(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLMATCHINGLOW(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlMatHold(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, double optInPenetration, /* From 0 to TA_REAL_MAX */ int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLMATHOLD(startIdx, endIdx, inOpen, inHigh, inLow, inClose, optInPenetration, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlMorningDojiStar(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, double optInPenetration, /* From 0 to TA_REAL_MAX */ int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLMORNINGDOJISTAR(startIdx, endIdx, inOpen, inHigh, inLow, inClose, optInPenetration, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlMorningStar(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, double optInPenetration, /* From 0 to TA_REAL_MAX */ int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLMORNINGSTAR(startIdx, endIdx, inOpen, inHigh, inLow, inClose, optInPenetration, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlOnNeck(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLONNECK(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlPiercing(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLPIERCING(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlRickshawMan(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLRICKSHAWMAN(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlRiseFall3Methods(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLRISEFALL3METHODS(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlSeparatingLines(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLSEPARATINGLINES(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlShootingStar(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLSHOOTINGSTAR(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlShortLine(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLSHORTLINE(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlSpinningTop(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLSPINNINGTOP(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlStalledPattern(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLSTALLEDPATTERN(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlStickSandwich(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLSTICKSANDWICH(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlTakuri(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLTAKURI(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlTasukiGap(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLTASUKIGAP(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlThrusting(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLTHRUSTING(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdltriStar(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLTRISTAR(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlUnique3River(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLUNIQUE3RIVER(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlupSideGap2Crows(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLUPSIDEGAP2CROWS(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_CdlxSideGap3Methods(int startIdx, int endIdx, const double *inOpen, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_CDLXSIDEGAP3METHODS(startIdx, endIdx, inOpen, inHigh, inLow, inClose, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_Ceil(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_CEIL(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Cmo(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_CMO(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Correl(int startIdx, int endIdx, const double *inReal0, const double *inReal1, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_CORREL(startIdx, endIdx, inReal0, inReal1, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Cos(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_COS(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Cosh(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_COSH(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Dema(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_DEMA(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Div(int startIdx, int endIdx, const double *inReal0, const double *inReal1, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_DIV(startIdx, endIdx, inReal0, inReal1, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Dx(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_DX(startIdx, endIdx, inHigh, inLow, inClose, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Ema(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_EMA(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Exp(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_EXP(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Floor(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_FLOOR(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_HtDcPeriod(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_HT_DCPERIOD(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_HtDcPhase(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_HT_DCPHASE(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_HtPhasor(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outInPhase, double *outQuadrature) {
    return TA_HT_PHASOR(startIdx, endIdx, inReal, outBegIdx, outNBElement, outInPhase, outQuadrature);
}

TA_EXPORT TA_RetCode Z_HtSine(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outSine, double *outLeadSine) {
    return TA_HT_SINE(startIdx, endIdx, inReal, outBegIdx, outNBElement, outSine, outLeadSine);
}

TA_EXPORT TA_RetCode Z_HtTrendLine(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_HT_TRENDLINE(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_HtTrendMode(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_HT_TRENDMODE(startIdx, endIdx, inReal, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_Kama(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_KAMA(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_LinearReg(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_LINEARREG(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_LinearRegAngle(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_LINEARREG_ANGLE(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_LinearRegIntercept(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_LINEARREG_INTERCEPT(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_LinearRegSlope(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_LINEARREG_SLOPE(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Ln(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_LN(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Log10(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_LOG10(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Ma(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 1 to 100000 */ TA_MAType optInMAType, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_MA(startIdx, endIdx, inReal, optInTimePeriod, optInMAType, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Macd(int startIdx, int endIdx, const double *inReal, int optInFastPeriod, /* From 2 to 100000 */ int optInSlowPeriod, /* From 2 to 100000 */ int optInSignalPeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outMACD, double *outMACDSignal, double *outMACDHist) {
    return TA_MACD(startIdx, endIdx, inReal, optInFastPeriod, optInSlowPeriod, optInSignalPeriod, outBegIdx, outNBElement, outMACD, outMACDSignal, outMACDHist);
}

TA_EXPORT TA_RetCode Z_MacdExt(int startIdx, int endIdx, const double *inReal, int optInFastPeriod, /* From 2 to 100000 */ TA_MAType optInFastMAType, int optInSlowPeriod, /* From 2 to 100000 */ TA_MAType optInSlowMAType, int optInSignalPeriod, /* From 1 to 100000 */ TA_MAType optInSignalMAType, int *outBegIdx, int *outNBElement, double *outMACD, double *outMACDSignal, double *outMACDHist) {
    return TA_MACDEXT(startIdx, endIdx, inReal, optInFastPeriod, optInFastMAType, optInSlowPeriod, optInSlowMAType, optInSignalPeriod, optInSignalMAType, outBegIdx, outNBElement, outMACD, outMACDSignal, outMACDHist);
}

TA_EXPORT TA_RetCode Z_MacdFix(int startIdx, int endIdx, const double *inReal, int optInSignalPeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outMACD, double *outMACDSignal, double *outMACDHist) {
    return TA_MACDFIX(startIdx, endIdx, inReal, optInSignalPeriod, outBegIdx, outNBElement, outMACD, outMACDSignal, outMACDHist);
}

TA_EXPORT TA_RetCode Z_Mama(int startIdx, int endIdx, const double *inReal, double optInFastLimit, /* From 0.01 to 0.99 */ double optInSlowLimit, /* From 0.01 to 0.99 */ int *outBegIdx, int *outNBElement, double *outMAMA, double *outFAMA) {
    return TA_MAMA(startIdx, endIdx, inReal, optInFastLimit, optInSlowLimit, outBegIdx, outNBElement, outMAMA, outFAMA);
}

TA_EXPORT TA_RetCode Z_Mavp(int startIdx, int endIdx, const double *inReal, const double *inPeriods, int optInMinPeriod, /* From 2 to 100000 */ int optInMaxPeriod, /* From 2 to 100000 */ TA_MAType optInMAType, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_MAVP(startIdx, endIdx, inReal, inPeriods, optInMinPeriod, optInMaxPeriod, optInMAType, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Max(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_MAX(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_MaxIndex(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_MAXINDEX(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_MedPrice(int startIdx, int endIdx, const double *inHigh, const double *inLow, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_MEDPRICE(startIdx, endIdx, inHigh, inLow, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Mfi(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, const double *inVolume, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_MFI(startIdx, endIdx, inHigh, inLow, inClose, inVolume, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_MidPoint(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_MIDPOINT(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_MidPrice(int startIdx, int endIdx, const double *inHigh, const double *inLow, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_MIDPRICE(startIdx, endIdx, inHigh, inLow, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Min(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_MIN(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_MinIndex(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, int *outInteger) {
    return TA_MININDEX(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outInteger);
}

TA_EXPORT TA_RetCode Z_MinMax(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outMin, double *outMax) {
    return TA_MINMAX(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outMin, outMax);
}

TA_EXPORT TA_RetCode Z_MinMaxIndex(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, int *outMinIdx, int *outMaxIdx) {
    return TA_MINMAXINDEX(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outMinIdx, outMaxIdx);
}

TA_EXPORT TA_RetCode Z_MinusDi(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_MINUS_DI(startIdx, endIdx, inHigh, inLow, inClose, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_MinusDm(int startIdx, int endIdx, const double *inHigh, const double *inLow, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_MINUS_DM(startIdx, endIdx, inHigh, inLow, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Mom(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_MOM(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Mult(int startIdx, int endIdx, const double *inReal0, const double *inReal1, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_MULT(startIdx, endIdx, inReal0, inReal1, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Natr(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_NATR(startIdx, endIdx, inHigh, inLow, inClose, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Obv(int startIdx, int endIdx, const double *inReal, const double *inVolume, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_OBV(startIdx, endIdx, inReal, inVolume, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_PlusDi(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_PLUS_DI(startIdx, endIdx, inHigh, inLow, inClose, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_PlusDm(int startIdx, int endIdx, const double *inHigh, const double *inLow, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_PLUS_DM(startIdx, endIdx, inHigh, inLow, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Ppo(int startIdx, int endIdx, const double *inReal, int optInFastPeriod, /* From 2 to 100000 */ int optInSlowPeriod, /* From 2 to 100000 */ TA_MAType optInMAType, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_PPO(startIdx, endIdx, inReal, optInFastPeriod, optInSlowPeriod, optInMAType, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Roc(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_ROC(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Rocp(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_ROCP(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Rocr(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_ROCR(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Rocr100(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_ROCR100(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Rsi(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_RSI(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Sar(int startIdx, int endIdx, const double *inHigh, const double *inLow, double optInAcceleration, /* From 0 to TA_REAL_MAX */ double optInMaximum, /* From 0 to TA_REAL_MAX */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_SAR(startIdx, endIdx, inHigh, inLow, optInAcceleration, optInMaximum, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_SarExt(int startIdx, int endIdx, const double *inHigh, const double *inLow, double optInStartValue, /* From TA_REAL_MIN to TA_REAL_MAX */ double optInOffsetOnReverse, /* From 0 to TA_REAL_MAX */ double optInAccelerationInitLong, /* From 0 to TA_REAL_MAX */ double optInAccelerationLong, /* From 0 to TA_REAL_MAX */ double optInAccelerationMaxLong, /* From 0 to TA_REAL_MAX */ double optInAccelerationInitShort, /* From 0 to TA_REAL_MAX */ double optInAccelerationShort, /* From 0 to TA_REAL_MAX */ double optInAccelerationMaxShort, /* From 0 to TA_REAL_MAX */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_SAREXT(startIdx, endIdx, inHigh, inLow, optInStartValue, optInOffsetOnReverse, optInAccelerationInitLong, optInAccelerationLong, optInAccelerationMaxLong, optInAccelerationInitShort, optInAccelerationShort, optInAccelerationMaxShort, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Sin(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_SIN(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Sinh(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_SINH(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Sma(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_SMA(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Sqrt(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_SQRT(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_StdDev(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ double optInNbDev, /* From TA_REAL_MIN to TA_REAL_MAX */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_STDDEV(startIdx, endIdx, inReal, optInTimePeriod, optInNbDev, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Stoch(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int optInFastK_Period, /* From 1 to 100000 */ int optInSlowK_Period, /* From 1 to 100000 */ TA_MAType optInSlowK_MAType, int optInSlowD_Period, /* From 1 to 100000 */ TA_MAType optInSlowD_MAType, int *outBegIdx, int *outNBElement, double *outSlowK, double *outSlowD) {
    return TA_STOCH(startIdx, endIdx, inHigh, inLow, inClose, optInFastK_Period, optInSlowK_Period, optInSlowK_MAType, optInSlowD_Period, optInSlowD_MAType, outBegIdx, outNBElement, outSlowK, outSlowD);
}

TA_EXPORT TA_RetCode Z_Stochf(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int optInFastK_Period, /* From 1 to 100000 */ int optInFastD_Period, /* From 1 to 100000 */ TA_MAType optInFastD_MAType, int *outBegIdx, int *outNBElement, double *outFastK, double *outFastD) {
    return TA_STOCHF(startIdx, endIdx, inHigh, inLow, inClose, optInFastK_Period, optInFastD_Period, optInFastD_MAType, outBegIdx, outNBElement, outFastK, outFastD);
}

TA_EXPORT TA_RetCode Z_StochRsi(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int optInFastK_Period, /* From 1 to 100000 */ int optInFastD_Period, /* From 1 to 100000 */ TA_MAType optInFastD_MAType, int *outBegIdx, int *outNBElement, double *outFastK, double *outFastD) {
    return TA_STOCHRSI(startIdx, endIdx, inReal, optInTimePeriod, optInFastK_Period, optInFastD_Period, optInFastD_MAType, outBegIdx, outNBElement, outFastK, outFastD);
}

TA_EXPORT TA_RetCode Z_Sub(int startIdx, int endIdx, const double *inReal0, const double *inReal1, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_SUB(startIdx, endIdx, inReal0, inReal1, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Sum(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_SUM(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_T3(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ double optInVFactor, /* From 0 to 1 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_T3(startIdx, endIdx, inReal, optInTimePeriod, optInVFactor, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Tan(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_TAN(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Tanh(int startIdx, int endIdx, const double *inReal, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_TANH(startIdx, endIdx, inReal, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Tema(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_TEMA(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Trange(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_TRANGE(startIdx, endIdx, inHigh, inLow, inClose, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Trima(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_TRIMA(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Trix(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_TRIX(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Tsf(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_TSF(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_TypPrice(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_TYPPRICE(startIdx, endIdx, inHigh, inLow, inClose, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_UltOsc(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int optInTimePeriod1, /* From 1 to 100000 */ int optInTimePeriod2, /* From 1 to 100000 */ int optInTimePeriod3, /* From 1 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_ULTOSC(startIdx, endIdx, inHigh, inLow, inClose, optInTimePeriod1, optInTimePeriod2, optInTimePeriod3, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Var(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 1 to 100000 */ double optInNbDev, /* From TA_REAL_MIN to TA_REAL_MAX */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_VAR(startIdx, endIdx, inReal, optInTimePeriod, optInNbDev, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_WclPrice(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_WCLPRICE(startIdx, endIdx, inHigh, inLow, inClose, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Willr(int startIdx, int endIdx, const double *inHigh, const double *inLow, const double *inClose, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_WILLR(startIdx, endIdx, inHigh, inLow, inClose, optInTimePeriod, outBegIdx, outNBElement, outReal);
}

TA_EXPORT TA_RetCode Z_Wma(int startIdx, int endIdx, const double *inReal, int optInTimePeriod, /* From 2 to 100000 */ int *outBegIdx, int *outNBElement, double *outReal) {
    return TA_WMA(startIdx, endIdx, inReal, optInTimePeriod, outBegIdx, outNBElement, outReal);
}


#endif // TA_EXPORT_H
