// Copyright ©2016 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package amos

import (
	"math"
	"math/cmplx"
)

// These routines are the versions directly modified from the Fortran code.
// They are used to ensure that code style improvements do not change the
// code output.

func iabs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func min0(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max0(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func zairyOrig(ZR, ZI float64, ID, KODE int) (AIR, AII float64, NZ int) {
	// zairy is adapted from the original Netlib code by Donald Amos.
	// http://www.netlib.no/netlib/amos/zairy.f

	// Original comment:
	/*
		C***BEGIN PROLOGUE  ZAIRY
		C***DATE WRITTEN   830501   (YYMMDD)
		C***REVISION DATE  890801   (YYMMDD)
		C***CATEGORY NO.  B5K
		C***KEYWORDS  AIRY FUNCTION,BESSEL FUNCTIONS OF ORDER ONE THIRD
		C***AUTHOR  AMOS, DONALD E., SANDIA NATIONAL LABORATORIES
		C***PURPOSE  TO COMPUTE AIRY FUNCTIONS AI(Z) AND DAI(Z) FOR COMPLEX Z
		C***DESCRIPTION
		C
		C                      ***A DOUBLE PRECISION ROUTINE***
		C         ON KODE=1, ZAIRY COMPUTES THE COMPLEX AIRY FUNCTION AI(Z) OR
		C         ITS DERIVATIVE DAI(Z)/DZ ON ID=0 OR ID=1 RESPECTIVELY. ON
		C         KODE=2, A SCALING OPTION CEXP(ZTA)*AI(Z) OR CEXP(ZTA)*
		C         DAI(Z)/DZ IS PROVIDED TO REMOVE THE EXPONENTIAL DECAY IN
		C         -PI/3<ARG(Z)<PI/3 AND THE EXPONENTIAL GROWTH IN
		C         PI/3<ABS(ARG(Z))<PI WHERE ZTA=(2/3)*Z*CSQRT(Z).
		C
		C         WHILE THE AIRY FUNCTIONS AI(Z) AND DAI(Z)/DZ ARE ANALYTIC IN
		C         THE WHOLE Z PLANE, THE CORRESPONDING SCALED FUNCTIONS DEFINED
		C         FOR KODE=2 HAVE A CUT ALONG THE NEGATIVE REAL AXIS.
		C         DEFINTIONS AND NOTATION ARE FOUND IN THE NBS HANDBOOK OF
		C         MATHEMATICAL FUNCTIONS (REF. 1).
		C
		C         INPUT      ZR,ZI ARE DOUBLE PRECISION
		C           ZR,ZI  - Z=CMPLX(ZR,ZI)
		C           ID     - ORDER OF DERIVATIVE, ID=0 OR ID=1
		C           KODE   - A PARAMETER TO INDICATE THE SCALING OPTION
		C                    KODE= 1  returnS
		C                             AI=AI(Z)                ON ID=0 OR
		C                             AI=DAI(Z)/DZ            ON ID=1
		C                        = 2  returnS
		C                             AI=CEXP(ZTA)*AI(Z)       ON ID=0 OR
		C                             AI=CEXP(ZTA)*DAI(Z)/DZ   ON ID=1 WHERE
		C                             ZTA=(2/3)*Z*CSQRT(Z)
		C
		C         OUTPUT     AIR,AII ARE DOUBLE PRECISION
		C           AIR,AII- COMPLEX ANSWER DEPENDING ON THE CHOICES FOR ID AND
		C                    KODE
		C           NZ     - UNDERFLOW INDICATOR
		C                    NZ= 0   , NORMAL return
		C                    NZ= 1   , AI=CMPLX(0.0E0,0.0E0) DUE TO UNDERFLOW IN
		C                              -PI/3<ARG(Z)<PI/3 ON KODE=1
		C           IERR   - ERROR FLAG
		C                    IERR=0, NORMAL return - COMPUTATION COMPLETED
		C                    IERR=1, INPUT ERROR   - NO COMPUTATION
		C                    IERR=2, OVERFLOW      - NO COMPUTATION, REAL(ZTA)
		C                            TOO LARGE ON KODE=1
		C                    IERR=3, CABS(Z) LARGE      - COMPUTATION COMPLETED
		C                            LOSSES OF SIGNIFCANCE BY ARGUMENT REDUCTION
		C                            PRODUCE LESS THAN HALF OF MACHINE ACCURACY
		C                    IERR=4, CABS(Z) TOO LARGE  - NO COMPUTATION
		C                            COMPLETE LOSS OF ACCURACY BY ARGUMENT
		C                            REDUCTION
		C                    IERR=5, ERROR              - NO COMPUTATION,
		C                            ALGORITHM TERMINATION CONDITION NOT MET
		C
		C***LONG DESCRIPTION
		C
		C         AI AND DAI ARE COMPUTED FOR CABS(Z)>1.0 FROM THE K BESSEL
		C         FUNCTIONS BY
		C
		C            AI(Z)=C*SQRT(Z)*K(1/3,ZTA) , DAI(Z)=-C*Z*K(2/3,ZTA)
		C                           C=1.0/(PI*SQRT(3.0))
		C                            ZTA=(2/3)*Z**(3/2)
		C
		C         WITH THE POWER SERIES FOR CABS(Z)<=1.0.
		C
		C         IN MOST COMPLEX VARIABLE COMPUTATION, ONE MUST EVALUATE ELE-
		C         MENTARY FUNCTIONS. WHEN THE MAGNITUDE OF Z IS LARGE, LOSSES
		C         OF SIGNIFICANCE BY ARGUMENT REDUCTION OCCUR. CONSEQUENTLY, IF
		C         THE MAGNITUDE OF ZETA=(2/3)*Z**1.5 EXCEEDS U1=SQRT(0.5/UR),
		C         THEN LOSSES EXCEEDING HALF PRECISION ARE LIKELY AND AN ERROR
		C         FLAG IERR=3 IS TRIGGERED WHERE UR=dmax(dmach[4),1.0D-18) IS
		C         DOUBLE PRECISION UNIT ROUNDOFF LIMITED TO 18 DIGITS PRECISION.
		C         ALSO, if THE MAGNITUDE OF ZETA IS LARGER THAN U2=0.5/UR, THEN
		C         ALL SIGNIFICANCE IS LOST AND IERR=4. IN ORDER TO USE THE INT
		C         FUNCTION, ZETA MUST BE FURTHER RESTRICTED NOT TO EXCEED THE
		C         LARGEST INTEGER, U3=I1MACH(9). THUS, THE MAGNITUDE OF ZETA
		C         MUST BE RESTRICTED BY MIN(U2,U3). ON 32 BIT MACHINES, U1,U2,
		C         AND U3 ARE APPROXIMATELY 2.0E+3, 4.2E+6, 2.1E+9 IN SINGLE
		C         PRECISION ARITHMETIC AND 1.3E+8, 1.8E+16, 2.1E+9 IN DOUBLE
		C         PRECISION ARITHMETIC RESPECTIVELY. THIS MAKES U2 AND U3 LIMIT-
		C         ING IN THEIR RESPECTIVE ARITHMETICS. THIS MEANS THAT THE MAG-
		C         NITUDE OF Z CANNOT EXCEED 3.1E+4 IN SINGLE AND 2.1E+6 IN
		C         DOUBLE PRECISION ARITHMETIC. THIS ALSO MEANS THAT ONE CAN
		C         EXPECT TO RETAIN, IN THE WORST CASES ON 32 BIT MACHINES,
		C         NO DIGITS IN SINGLE PRECISION AND ONLY 7 DIGITS IN DOUBLE
		C         PRECISION ARITHMETIC. SIMILAR CONSIDERATIONS HOLD FOR OTHER
		C         MACHINES.
		C
		C         THE APPROXIMATE RELATIVE ERROR IN THE MAGNITUDE OF A COMPLEX
		C         BESSEL FUNCTION CAN BE EXPRESSED BY P*10**S WHERE P=MAX(UNIT
		C         ROUNDOFF,1.0E-18) IS THE NOMINAL PRECISION AND 10**S REPRE-
		C         SENTS THE INCREASE IN ERROR DUE TO ARGUMENT REDUCTION IN THE
		C         ELEMENTARY FUNCTIONS. HERE, S=MAX(1,ABS(LOG10(CABS(Z))),
		C         ABS(LOG10(FNU))) APPROXIMATELY (I.E. S=MAX(1,ABS(EXPONENT OF
		C         CABS(Z),ABS(EXPONENT OF FNU)) ). HOWEVER, THE PHASE ANGLE MAY
		C         HAVE ONLY ABSOLUTE ACCURACY. THIS IS MOST LIKELY TO OCCUR WHEN
		C         ONE COMPONENT (IN ABSOLUTE VALUE) IS LARGER THAN THE OTHER BY
		C         SEVERAL ORDERS OF MAGNITUDE. if ONE COMPONENT IS 10**K LARGER
		C         THAN THE OTHER, THEN ONE CAN EXPECT ONLY MAX(ABS(LOG10(P))-K,
		C         0) SIGNIFICANT DIGITS; OR, STATED ANOTHER WAY, WHEN K EXCEEDS
		C         THE EXPONENT OF P, NO SIGNIFICANT DIGITS REMAIN IN THE SMALLER
		C         COMPONENT. HOWEVER, THE PHASE ANGLE RETAINS ABSOLUTE ACCURACY
		C         BECAUSE, IN COMPLEX ARITHMETIC WITH PRECISION P, THE SMALLER
		C         COMPONENT WILL NOT (AS A RULE) DECREASE BELOW P TIMES THE
		C         MAGNITUDE OF THE LARGER COMPONENT. IN THESE EXTREME CASES,
		C         THE PRINCIPAL PHASE ANGLE IS ON THE ORDER OF +P, -P, PI/2-P,
		C         OR -PI/2+P.
		C
		C***REFERENCES  HANDBOOK OF MATHEMATICAL FUNCTIONS BY M. ABRAMOWITZ
		C                 AND I. A. STEGUN, NBS AMS SERIES 55, U.S. DEPT. OF
		C                 COMMERCE, 1955.
		C
		C               COMPUTATION OF BESSEL FUNCTIONS OF COMPLEX ARGUMENT
		C                 AND LARGE ORDER BY D. E. AMOS, SAND83-0643, MAY, 1983
		C
		C               A SUBROUTINE PACKAGE FOR BESSEL FUNCTIONS OF A COMPLEX
		C                 ARGUMENT AND NONNEGATIVE ORDER BY D. E. AMOS, SAND85-
		C                 1018, MAY, 1985
		C
		C               A PORTABLE PACKAGE FOR BESSEL FUNCTIONS OF A COMPLEX
		C                 ARGUMENT AND NONNEGATIVE ORDER BY D. E. AMOS, TRANS.
		C                 MATH. SOFTWARE, 1986
	*/
	var AI, CONE, CSQ, CY, S1, S2, TRM1, TRM2, Z, ZTA, Z3 complex128
	var AA, AD, AK, ALIM, ATRM, AZ, AZ3, BK,
		CC, CK, COEF, CONEI, CONER, CSQI, CSQR, C1, C2, DIG,
		DK, D1, D2, ELIM, FID, FNU, PTR, RL, R1M5, SFAC, STI, STR,
		S1I, S1R, S2I, S2R, TOL, TRM1I, TRM1R, TRM2I, TRM2R, TTH, ZEROI,
		ZEROR, ZTAI, ZTAR, Z3I, Z3R, ALAZ, BB float64
	var IERR, IFLAG, K, K1, K2, MR, NN int
	var tmp complex128

	// Extra element for padding.
	CYR := []float64{math.NaN(), 0}
	CYI := []float64{math.NaN(), 0}

	_ = AI
	_ = CONE
	_ = CSQ
	_ = CY
	_ = S1
	_ = S2
	_ = TRM1
	_ = TRM2
	_ = Z
	_ = ZTA
	_ = Z3

	TTH = 6.66666666666666667E-01
	C1 = 3.55028053887817240E-01
	C2 = 2.58819403792806799E-01
	COEF = 1.83776298473930683E-01
	ZEROR = 0
	ZEROI = 0
	CONER = 1
	CONEI = 0

	NZ = 0
	if ID < 0 || ID > 1 {
		IERR = 1
	}
	if KODE < 1 || KODE > 2 {
		IERR = 1
	}
	if IERR != 0 {
		return
	}
	AZ = zabs(complex(ZR, ZI))
	TOL = dmax(dmach[4], 1.0E-18)
	FID = float64(ID)
	if AZ > 1.0E0 {
		goto Seventy
	}

	// POWER SERIES FOR CABS(Z)<=1.
	S1R = CONER
	S1I = CONEI
	S2R = CONER
	S2I = CONEI
	if AZ < TOL {
		goto OneSeventy
	}
	AA = AZ * AZ
	if AA < TOL/AZ {
		goto Forty
	}
	TRM1R = CONER
	TRM1I = CONEI
	TRM2R = CONER
	TRM2I = CONEI
	ATRM = 1.0E0
	STR = ZR*ZR - ZI*ZI
	STI = ZR*ZI + ZI*ZR
	Z3R = STR*ZR - STI*ZI
	Z3I = STR*ZI + STI*ZR
	AZ3 = AZ * AA
	AK = 2.0E0 + FID
	BK = 3.0E0 - FID - FID
	CK = 4.0E0 - FID
	DK = 3.0E0 + FID + FID
	D1 = AK * DK
	D2 = BK * CK
	AD = dmin(D1, D2)
	AK = 24.0E0 + 9.0E0*FID
	BK = 30.0E0 - 9.0E0*FID
	for K = 1; K <= 25; K++ {
		STR = (TRM1R*Z3R - TRM1I*Z3I) / D1
		TRM1I = (TRM1R*Z3I + TRM1I*Z3R) / D1
		TRM1R = STR
		S1R = S1R + TRM1R
		S1I = S1I + TRM1I
		STR = (TRM2R*Z3R - TRM2I*Z3I) / D2
		TRM2I = (TRM2R*Z3I + TRM2I*Z3R) / D2
		TRM2R = STR
		S2R = S2R + TRM2R
		S2I = S2I + TRM2I
		ATRM = ATRM * AZ3 / AD
		D1 = D1 + AK
		D2 = D2 + BK
		AD = dmin(D1, D2)
		if ATRM < TOL*AD {
			goto Forty
		}
		AK = AK + 18.0E0
		BK = BK + 18.0E0
	}
Forty:
	if ID == 1 {
		goto Fifty
	}
	AIR = S1R*C1 - C2*(ZR*S2R-ZI*S2I)
	AII = S1I*C1 - C2*(ZR*S2I+ZI*S2R)
	if KODE == 1 {
		return
	}
	tmp = zsqrt(complex(ZR, ZI))
	STR = real(tmp)
	STI = imag(tmp)
	ZTAR = TTH * (ZR*STR - ZI*STI)
	ZTAI = TTH * (ZR*STI + ZI*STR)
	tmp = zexp(complex(ZTAR, ZTAI))
	STR = real(tmp)
	STI = imag(tmp)
	PTR = AIR*STR - AII*STI
	AII = AIR*STI + AII*STR
	AIR = PTR
	return

Fifty:
	AIR = -S2R * C2
	AII = -S2I * C2
	if AZ <= TOL {
		goto Sixty
	}
	STR = ZR*S1R - ZI*S1I
	STI = ZR*S1I + ZI*S1R
	CC = C1 / (1.0E0 + FID)
	AIR = AIR + CC*(STR*ZR-STI*ZI)
	AII = AII + CC*(STR*ZI+STI*ZR)

Sixty:
	if KODE == 1 {
		return
	}
	tmp = zsqrt(complex(ZR, ZI))
	STR = real(tmp)
	STI = imag(tmp)
	ZTAR = TTH * (ZR*STR - ZI*STI)
	ZTAI = TTH * (ZR*STI + ZI*STR)
	tmp = zexp(complex(ZTAR, ZTAI))
	STR = real(tmp)
	STI = imag(tmp)
	PTR = STR*AIR - STI*AII
	AII = STR*AII + STI*AIR
	AIR = PTR
	return

	// CASE FOR CABS(Z)>1.0.
Seventy:
	FNU = (1.0E0 + FID) / 3.0E0

	/*
	   SET PARAMETERS RELATED TO MACHINE CONSTANTS.
	   TOL IS THE APPROXIMATE UNIT ROUNDOFF LIMITED TO 1.0D-18.
	   ELIM IS THE APPROXIMATE EXPONENTIAL OVER-&&UNDERFLOW LIMIT.
	   EXP(-ELIM)<EXP(-ALIM)=EXP(-ELIM)/TOL    AND
	   EXP(ELIM)>EXP(ALIM)=EXP(ELIM)*TOL       ARE INTERVALS NEAR
	   UNDERFLOW&&OVERFLOW LIMITS WHERE SCALED ARITHMETIC IS DONE.
	   RL IS THE LOWER BOUNDARY OF THE ASYMPTOTIC EXPANSION FOR LA>=Z.
	   DIG = NUMBER OF BASE 10 DIGITS IN TOL = 10**(-DIG).
	*/
	K1 = imach[15]
	K2 = imach[16]
	R1M5 = dmach[5]

	K = min0(iabs(K1), iabs(K2))
	ELIM = 2.303E0 * (float64(K)*R1M5 - 3.0E0)
	K1 = imach[14] - 1
	AA = R1M5 * float64(K1)
	DIG = dmin(AA, 18.0E0)
	AA = AA * 2.303E0
	ALIM = ELIM + dmax(-AA, -41.45E0)
	RL = 1.2E0*DIG + 3.0E0
	ALAZ = dlog(AZ)

	// TEST FOR PROPER RANGE.
	AA = 0.5E0 / TOL
	BB = float64(float32(imach[9])) * 0.5E0
	AA = dmin(AA, BB)
	AA = math.Pow(AA, TTH)
	if AZ > AA {
		goto TwoSixty
	}
	AA = dsqrt(AA)
	if AZ > AA {
		IERR = 3
	}
	tmp = zsqrt(complex(ZR, ZI))
	CSQR = real(tmp)
	CSQI = imag(tmp)
	ZTAR = TTH * (ZR*CSQR - ZI*CSQI)
	ZTAI = TTH * (ZR*CSQI + ZI*CSQR)

	//  RE(ZTA)<=0 WHEN RE(Z)<0, ESPECIALLY WHEN IM(Z) IS SMALL.
	IFLAG = 0
	SFAC = 1.0E0
	AK = ZTAI
	if ZR >= 0.0E0 {
		goto Eighty
	}
	BK = ZTAR
	CK = -dabs(BK)
	ZTAR = CK
	ZTAI = AK

Eighty:
	if ZI != 0.0E0 {
		goto Ninety
	}
	if ZR > 0.0E0 {
		goto Ninety
	}
	ZTAR = 0.0E0
	ZTAI = AK
Ninety:
	AA = ZTAR
	if AA >= 0.0E0 && ZR > 0.0E0 {
		goto OneTen
	}
	if KODE == 2 {
		goto OneHundred
	}

	// OVERFLOW TEST.
	if AA > (-ALIM) {
		goto OneHundred
	}
	AA = -AA + 0.25E0*ALAZ
	IFLAG = 1
	SFAC = TOL
	if AA > ELIM {
		goto TwoSeventy
	}

OneHundred:
	// CBKNU AND CACON return EXP(ZTA)*K(FNU,ZTA) ON KODE=2.
	MR = 1
	if ZI < 0.0E0 {
		MR = -1
	}
	ZTAR, ZTAI, FNU, KODE, MR, _, CYR, CYI, NN, RL, TOL, ELIM, ALIM = zacaiOrig(ZTAR, ZTAI, FNU, KODE, MR, 1, CYR, CYI, NN, RL, TOL, ELIM, ALIM)
	if NN < 0 {
		goto TwoEighty
	}
	NZ = NZ + NN
	goto OneThirty

OneTen:
	if KODE == 2 {
		goto OneTwenty
	}

	// UNDERFLOW TEST.
	if AA < ALIM {
		goto OneTwenty
	}
	AA = -AA - 0.25E0*ALAZ
	IFLAG = 2
	SFAC = 1.0E0 / TOL
	if AA < (-ELIM) {
		goto TwoTen
	}
OneTwenty:
	ZTAR, ZTAI, FNU, KODE, _, CYR, CYI, NZ, TOL, ELIM, ALIM = zbknuOrig(ZTAR, ZTAI, FNU, KODE, 1, CYR, CYI, NZ, TOL, ELIM, ALIM)

OneThirty:
	S1R = CYR[1] * COEF
	S1I = CYI[1] * COEF
	if IFLAG != 0 {
		goto OneFifty
	}
	if ID == 1 {
		goto OneFourty
	}
	AIR = CSQR*S1R - CSQI*S1I
	AII = CSQR*S1I + CSQI*S1R
	return
OneFourty:
	AIR = -(ZR*S1R - ZI*S1I)
	AII = -(ZR*S1I + ZI*S1R)
	return
OneFifty:
	S1R = S1R * SFAC
	S1I = S1I * SFAC
	if ID == 1 {
		goto OneSixty
	}
	STR = S1R*CSQR - S1I*CSQI
	S1I = S1R*CSQI + S1I*CSQR
	S1R = STR
	AIR = S1R / SFAC
	AII = S1I / SFAC
	return
OneSixty:
	STR = -(S1R*ZR - S1I*ZI)
	S1I = -(S1R*ZI + S1I*ZR)
	S1R = STR
	AIR = S1R / SFAC
	AII = S1I / SFAC
	return
OneSeventy:
	AA = 1.0E+3 * dmach[1]
	S1R = ZEROR
	S1I = ZEROI
	if ID == 1 {
		goto OneNinety
	}
	if AZ <= AA {
		goto OneEighty
	}
	S1R = C2 * ZR
	S1I = C2 * ZI
OneEighty:
	AIR = C1 - S1R
	AII = -S1I
	return
OneNinety:
	AIR = -C2
	AII = 0.0E0
	AA = dsqrt(AA)
	if AZ <= AA {
		goto TwoHundred
	}
	S1R = 0.5E0 * (ZR*ZR - ZI*ZI)
	S1I = ZR * ZI
TwoHundred:
	AIR = AIR + C1*S1R
	AII = AII + C1*S1I
	return
TwoTen:
	NZ = 1
	AIR = ZEROR
	AII = ZEROI
	return
TwoSeventy:
	NZ = 0
	IERR = 2
	return
TwoEighty:
	if NN == (-1) {
		goto TwoSeventy
	}
	NZ = 0
	IERR = 5
	return
TwoSixty:
	IERR = 4
	NZ = 0
	return
}

// sbknu computes the k bessel function in the right half z plane.
func zbknuOrig(ZR, ZI, FNU float64, KODE, N int, YR, YI []float64, NZ int, TOL, ELIM, ALIM float64) (ZRout, ZIout, FNUout float64, KODEout, Nout int, YRout, YIout []float64, NZout int, TOLout, ELIMout, ALIMout float64) {
	/* Old dimension comment.
		DIMENSION YR(N), YI(N), CC(8), CSSR(3), CSRR(3), BRY(3), CYR(2),
	     * CYI(2)
	*/

	// TODO(btracey): Find which of these are inputs/outputs/both and clean up
	// the function call.
	// YR and YI have length n (but n+1 with better indexing)
	var AA, AK, ASCLE, A1, A2, BB, BK, CAZ,
		CBI, CBR, CCHI, CCHR, CKI, CKR, COEFI, COEFR, CONEI, CONER,
		CRSCR, CSCLR, CSHI, CSHR, CSI, CSR, CTWOR,
		CZEROI, CZEROR, CZI, CZR, DNU, DNU2, DPI, ETEST, FC, FHS,
		FI, FK, FKS, FMUI, FMUR, FPI, FR, G1, G2, HPI, PI, PR, PTI,
		PTR, P1I, P1R, P2I, P2M, P2R, QI, QR, RAK, RCAZ, RTHPI, RZI,
		RZR, R1, S, SMUI, SMUR, SPI, STI, STR, S1I, S1R, S2I, S2R, TM,
		TTH, T1, T2, ELM, CELMR, ZDR, ZDI, AS, ALAS, HELIM float64

	var I, IFLAG, INU, K, KFLAG, KK, KMAX, KODED, IDUM, J, IC, INUB, NW int

	var tmp complex128
	var CSSR, CSRR, BRY [4]float64
	var CYR, CYI [3]float64

	KMAX = 30
	CZEROR = 0
	CZEROI = 0
	CONER = 1
	CONEI = 0
	CTWOR = 2
	R1 = 2

	DPI = 3.14159265358979324E0
	RTHPI = 1.25331413731550025E0
	SPI = 1.90985931710274403E0
	HPI = 1.57079632679489662E0
	FPI = 1.89769999331517738E0
	TTH = 6.66666666666666666E-01

	CC := [9]float64{math.NaN(), 5.77215664901532861E-01, -4.20026350340952355E-02,
		-4.21977345555443367E-02, 7.21894324666309954E-03,
		-2.15241674114950973E-04, -2.01348547807882387E-05,
		1.13302723198169588E-06, 6.11609510448141582E-09}

	CAZ = zabs(complex(ZR, ZI))
	CSCLR = 1.0E0 / TOL
	CRSCR = TOL
	CSSR[1] = CSCLR
	CSSR[2] = 1.0E0
	CSSR[3] = CRSCR
	CSRR[1] = CRSCR
	CSRR[2] = 1.0E0
	CSRR[3] = CSCLR
	BRY[1] = 1.0E+3 * dmach[1] / TOL
	BRY[2] = 1.0E0 / BRY[1]
	BRY[3] = dmach[2]
	NZ = 0
	IFLAG = 0
	KODED = KODE
	RCAZ = 1.0E0 / CAZ
	STR = ZR * RCAZ
	STI = -ZI * RCAZ
	RZR = (STR + STR) * RCAZ
	RZI = (STI + STI) * RCAZ
	INU = int(float32(FNU + 0.5))
	DNU = FNU - float64(INU)
	if dabs(DNU) == 0.5E0 {
		goto OneTen
	}
	DNU2 = 0.0E0
	if dabs(DNU) > TOL {
		DNU2 = DNU * DNU
	}
	if CAZ > R1 {
		goto OneTen
	}

	// SERIES FOR CABS(Z)<=R1.
	FC = 1.0E0
	tmp = zlog(complex(RZR, RZI))
	SMUR = real(tmp)
	SMUI = imag(tmp)
	FMUR = SMUR * DNU
	FMUI = SMUI * DNU
	FMUR, FMUI, CSHR, CSHI, CCHR, CCHI = zshchOrig(FMUR, FMUI, CSHR, CSHI, CCHR, CCHI)
	if DNU == 0.0E0 {
		goto Ten
	}
	FC = DNU * DPI
	FC = FC / dsin(FC)
	SMUR = CSHR / DNU
	SMUI = CSHI / DNU
Ten:
	A2 = 1.0E0 + DNU

	// GAM(1-Z)*GAM(1+Z)=PI*Z/SIN(PI*Z), T1=1/GAM(1-DNU), T2=1/GAM(1+DNU).
	T2 = dexp(-dgamln(A2, IDUM))
	T1 = 1.0E0 / (T2 * FC)
	if dabs(DNU) > 0.1E0 {
		goto Forty
	}

	// SERIES FOR F0 TO RESOLVE INDETERMINACY FOR SMALL ABS(DNU).
	AK = 1.0E0
	S = CC[1]
	for K = 2; K <= 8; K++ {
		AK = AK * DNU2
		TM = CC[K] * AK
		S = S + TM
		if dabs(TM) < TOL {
			goto Thirty
		}
	}
Thirty:
	G1 = -S
	goto Fifty
Forty:
	G1 = (T1 - T2) / (DNU + DNU)
Fifty:
	G2 = (T1 + T2) * 0.5E0
	FR = FC * (CCHR*G1 + SMUR*G2)
	FI = FC * (CCHI*G1 + SMUI*G2)
	tmp = zexp(complex(FMUR, FMUI))
	STR = real(tmp)
	STI = imag(tmp)
	PR = 0.5E0 * STR / T2
	PI = 0.5E0 * STI / T2
	tmp = zdiv(complex(0.5, 0), complex(STR, STI))
	PTR = real(tmp)
	PTI = imag(tmp)
	QR = PTR / T1
	QI = PTI / T1
	S1R = FR
	S1I = FI
	S2R = PR
	S2I = PI
	AK = 1.0E0
	A1 = 1.0E0
	CKR = CONER
	CKI = CONEI
	BK = 1.0E0 - DNU2
	if INU > 0 || N > 1 {
		goto Eighty
	}

	// GENERATE K(FNU,Z), 0.0E0 <= FNU < 0.5E0 AND N=1.
	if CAZ < TOL {
		goto Seventy
	}
	tmp = zmlt(complex(ZR, ZI), complex(ZR, ZI))
	CZR = real(tmp)
	CZI = imag(tmp)
	CZR = 0.25E0 * CZR
	CZI = 0.25E0 * CZI
	T1 = 0.25E0 * CAZ * CAZ
Sixty:
	FR = (FR*AK + PR + QR) / BK
	FI = (FI*AK + PI + QI) / BK
	STR = 1.0E0 / (AK - DNU)
	PR = PR * STR
	PI = PI * STR
	STR = 1.0E0 / (AK + DNU)
	QR = QR * STR
	QI = QI * STR
	STR = CKR*CZR - CKI*CZI
	RAK = 1.0E0 / AK
	CKI = (CKR*CZI + CKI*CZR) * RAK
	CKR = STR * RAK
	S1R = CKR*FR - CKI*FI + S1R
	S1I = CKR*FI + CKI*FR + S1I
	A1 = A1 * T1 * RAK
	BK = BK + AK + AK + 1.0E0
	AK = AK + 1.0E0
	if A1 > TOL {
		goto Sixty
	}
Seventy:
	YR[1] = S1R
	YI[1] = S1I
	if KODED == 1 {
		return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
	}
	tmp = zexp(complex(ZR, ZI))
	STR = real(tmp)
	STI = imag(tmp)
	tmp = zmlt(complex(S1R, S1I), complex(STR, STI))
	YR[1] = real(tmp)
	YI[1] = imag(tmp)
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM

	// GENERATE K(DNU,Z) AND K(DNU+1,Z) FOR FORWARD RECURRENCE.
Eighty:
	if CAZ < TOL {
		goto OneHundred
	}
	tmp = zmlt(complex(ZR, ZI), complex(ZR, ZI))
	CZR = real(tmp)
	CZI = imag(tmp)
	CZR = 0.25E0 * CZR
	CZI = 0.25E0 * CZI
	T1 = 0.25E0 * CAZ * CAZ
Ninety:
	FR = (FR*AK + PR + QR) / BK
	FI = (FI*AK + PI + QI) / BK
	STR = 1.0E0 / (AK - DNU)
	PR = PR * STR
	PI = PI * STR
	STR = 1.0E0 / (AK + DNU)
	QR = QR * STR
	QI = QI * STR
	STR = CKR*CZR - CKI*CZI
	RAK = 1.0E0 / AK
	CKI = (CKR*CZI + CKI*CZR) * RAK
	CKR = STR * RAK
	S1R = CKR*FR - CKI*FI + S1R
	S1I = CKR*FI + CKI*FR + S1I
	STR = PR - FR*AK
	STI = PI - FI*AK
	S2R = CKR*STR - CKI*STI + S2R
	S2I = CKR*STI + CKI*STR + S2I
	A1 = A1 * T1 * RAK
	BK = BK + AK + AK + 1.0E0
	AK = AK + 1.0E0
	if A1 > TOL {
		goto Ninety
	}
OneHundred:
	KFLAG = 2
	A1 = FNU + 1.0E0
	AK = A1 * dabs(SMUR)
	if AK > ALIM {
		KFLAG = 3
	}
	STR = CSSR[KFLAG]
	P2R = S2R * STR
	P2I = S2I * STR
	tmp = zmlt(complex(P2R, P2I), complex(RZR, RZI))
	S2R = real(tmp)
	S2I = imag(tmp)
	S1R = S1R * STR
	S1I = S1I * STR
	if KODED == 1 {
		goto TwoTen
	}
	tmp = zexp(complex(ZR, ZI))
	FR = real(tmp)
	FI = imag(tmp)
	tmp = zmlt(complex(S1R, S1I), complex(FR, FI))
	S1R = real(tmp)
	S1I = imag(tmp)
	tmp = zmlt(complex(S2R, S2I), complex(FR, FI))
	S2R = real(tmp)
	S2I = imag(tmp)
	goto TwoTen

	// IFLAG=0 MEANS NO UNDERFLOW OCCURRED
	// IFLAG=1 MEANS AN UNDERFLOW OCCURRED- COMPUTATION PROCEEDS WITH
	// KODED=2 AND A TEST FOR ON SCALE VALUES IS MADE DURING FORWARD RECURSION
OneTen:
	tmp = zsqrt(complex(ZR, ZI))
	STR = real(tmp)
	STI = imag(tmp)
	tmp = zdiv(complex(RTHPI, CZEROI), complex(STR, STI))
	COEFR = real(tmp)
	COEFI = imag(tmp)
	KFLAG = 2
	if KODED == 2 {
		goto OneTwenty
	}
	if ZR > ALIM {
		goto TwoNinety
	}

	STR = dexp(-ZR) * CSSR[KFLAG]
	STI = -STR * dsin(ZI)
	STR = STR * dcos(ZI)
	tmp = zmlt(complex(COEFR, COEFI), complex(STR, STI))
	COEFR = real(tmp)
	COEFI = imag(tmp)
OneTwenty:
	if dabs(DNU) == 0.5E0 {
		goto ThreeHundred
	}
	// MILLER ALGORITHM FOR CABS(Z)>R1.
	AK = dcos(DPI * DNU)
	AK = dabs(AK)
	if AK == CZEROR {
		goto ThreeHundred
	}
	FHS = dabs(0.25E0 - DNU2)
	if FHS == CZEROR {
		goto ThreeHundred
	}

	// COMPUTE R2=F(E). if CABS(Z)>=R2, USE FORWARD RECURRENCE TO
	// DETERMINE THE BACKWARD INDEX K. R2=F(E) IS A STRAIGHT LINE ON
	// 12<=E<=60. E IS COMPUTED FROM 2**(-E)=B**(1-I1MACH(14))=
	// TOL WHERE B IS THE BASE OF THE ARITHMETIC.
	T1 = float64(imach[14] - 1)
	T1 = T1 * dmach[5] * 3.321928094E0
	T1 = dmax(T1, 12.0E0)
	T1 = dmin(T1, 60.0E0)
	T2 = TTH*T1 - 6.0E0
	if ZR != 0.0E0 {
		goto OneThirty
	}
	T1 = HPI
	goto OneFourty
OneThirty:
	T1 = datan(ZI / ZR)
	T1 = dabs(T1)
OneFourty:
	if T2 > CAZ {
		goto OneSeventy
	}
	// FORWARD RECURRENCE LOOP WHEN CABS(Z)>=R2.
	ETEST = AK / (DPI * CAZ * TOL)
	FK = CONER
	if ETEST < CONER {
		goto OneEighty
	}
	FKS = CTWOR
	CKR = CAZ + CAZ + CTWOR
	P1R = CZEROR
	P2R = CONER
	for I = 1; I <= KMAX; I++ {
		AK = FHS / FKS
		CBR = CKR / (FK + CONER)
		PTR = P2R
		P2R = CBR*P2R - P1R*AK
		P1R = PTR
		CKR = CKR + CTWOR
		FKS = FKS + FK + FK + CTWOR
		FHS = FHS + FK + FK
		FK = FK + CONER
		STR = dabs(P2R) * FK
		if ETEST < STR {
			goto OneSixty
		}
	}
	goto ThreeTen
OneSixty:
	FK = FK + SPI*T1*dsqrt(T2/CAZ)
	FHS = dabs(0.25 - DNU2)
	goto OneEighty
OneSeventy:
	// COMPUTE BACKWARD INDEX K FOR CABS(Z)<R2.
	A2 = dsqrt(CAZ)
	AK = FPI * AK / (TOL * dsqrt(A2))
	AA = 3.0E0 * T1 / (1.0E0 + CAZ)
	BB = 14.7E0 * T1 / (28.0E0 + CAZ)
	AK = (dlog(AK) + CAZ*dcos(AA)/(1.0E0+0.008E0*CAZ)) / dcos(BB)
	FK = 0.12125E0*AK*AK/CAZ + 1.5E0
OneEighty:
	// BACKWARD RECURRENCE LOOP FOR MILLER ALGORITHM.
	K = int(float32(FK))
	FK = float64(K)
	FKS = FK * FK
	P1R = CZEROR
	P1I = CZEROI
	P2R = TOL
	P2I = CZEROI
	CSR = P2R
	CSI = P2I
	for I = 1; I <= K; I++ {
		A1 = FKS - FK
		AK = (FKS + FK) / (A1 + FHS)
		RAK = 2.0E0 / (FK + CONER)
		CBR = (FK + ZR) * RAK
		CBI = ZI * RAK
		PTR = P2R
		PTI = P2I
		P2R = (PTR*CBR - PTI*CBI - P1R) * AK
		P2I = (PTI*CBR + PTR*CBI - P1I) * AK
		P1R = PTR
		P1I = PTI
		CSR = CSR + P2R
		CSI = CSI + P2I
		FKS = A1 - FK + CONER
		FK = FK - CONER
	}
	// COMPUTE (P2/CS)=(P2/CABS(CS))*(CONJG(CS)/CABS(CS)) FOR BETTER SCALING.
	TM = zabs(complex(CSR, CSI))
	PTR = 1.0E0 / TM
	S1R = P2R * PTR
	S1I = P2I * PTR
	CSR = CSR * PTR
	CSI = -CSI * PTR
	tmp = zmlt(complex(COEFR, COEFI), complex(S1R, S1I))
	STR = real(tmp)
	STI = imag(tmp)
	tmp = zmlt(complex(STR, STI), complex(CSR, CSI))
	S1R = real(tmp)
	S1I = imag(tmp)
	if INU > 0 || N > 1 {
		goto TwoHundred
	}
	ZDR = ZR
	ZDI = ZI
	if IFLAG == 1 {
		goto TwoSeventy
	}
	goto TwoFourty
TwoHundred:
	// COMPUTE P1/P2=(P1/CABS(P2)*CONJG(P2)/CABS(P2) FOR SCALING.
	TM = zabs(complex(P2R, P2I))
	PTR = 1.0E0 / TM
	P1R = P1R * PTR
	P1I = P1I * PTR
	P2R = P2R * PTR
	P2I = -P2I * PTR
	tmp = zmlt(complex(P1R, P1I), complex(P2R, P2I))
	PTR = real(tmp)
	PTI = imag(tmp)
	STR = DNU + 0.5E0 - PTR
	STI = -PTI
	tmp = zdiv(complex(STR, STI), complex(ZR, ZI))
	STR = real(tmp)
	STI = imag(tmp)
	STR = STR + 1.0E0
	tmp = zmlt(complex(STR, STI), complex(S1R, S1I))
	S2R = real(tmp)
	S2I = imag(tmp)

	// FORWARD RECURSION ON THE THREE TERM RECURSION WITH RELATION WITH
	// SCALING NEAR EXPONENT EXTREMES ON KFLAG=1 OR KFLAG=3
TwoTen:
	STR = DNU + 1.0E0
	CKR = STR * RZR
	CKI = STR * RZI
	if N == 1 {
		INU = INU - 1
	}
	if INU > 0 {
		goto TwoTwenty
	}
	if N > 1 {
		goto TwoFifteen
	}
	S1R = S2R
	S1I = S2I
TwoFifteen:
	ZDR = ZR
	ZDI = ZI
	if IFLAG == 1 {
		goto TwoSeventy
	}
	goto TwoFourty
TwoTwenty:
	INUB = 1
	if IFLAG == 1 {
		goto TwoSixtyOne
	}
TwoTwentyFive:
	P1R = CSRR[KFLAG]
	ASCLE = BRY[KFLAG]
	for I = INUB; I <= INU; I++ {
		STR = S2R
		STI = S2I
		S2R = CKR*STR - CKI*STI + S1R
		S2I = CKR*STI + CKI*STR + S1I
		S1R = STR
		S1I = STI
		CKR = CKR + RZR
		CKI = CKI + RZI
		if KFLAG >= 3 {
			continue
		}
		P2R = S2R * P1R
		P2I = S2I * P1R
		STR = dabs(P2R)
		STI = dabs(P2I)
		P2M = dmax(STR, STI)
		if P2M <= ASCLE {
			continue
		}
		KFLAG = KFLAG + 1
		ASCLE = BRY[KFLAG]
		S1R = S1R * P1R
		S1I = S1I * P1R
		S2R = P2R
		S2I = P2I
		STR = CSSR[KFLAG]
		S1R = S1R * STR
		S1I = S1I * STR
		S2R = S2R * STR
		S2I = S2I * STR
		P1R = CSRR[KFLAG]
	}
	if N != 1 {
		goto TwoFourty
	}
	S1R = S2R
	S1I = S2I
TwoFourty:
	STR = CSRR[KFLAG]
	YR[1] = S1R * STR
	YI[1] = S1I * STR
	if N == 1 {
		return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
	}
	YR[2] = S2R * STR
	YI[2] = S2I * STR
	if N == 2 {
		return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
	}
	KK = 2
TwoFifty:
	KK = KK + 1
	if KK > N {
		return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
	}
	P1R = CSRR[KFLAG]
	ASCLE = BRY[KFLAG]
	for I = KK; I <= N; I++ {
		P2R = S2R
		P2I = S2I
		S2R = CKR*P2R - CKI*P2I + S1R
		S2I = CKI*P2R + CKR*P2I + S1I
		S1R = P2R
		S1I = P2I
		CKR = CKR + RZR
		CKI = CKI + RZI
		P2R = S2R * P1R
		P2I = S2I * P1R
		YR[I] = P2R
		YI[I] = P2I
		if KFLAG >= 3 {
			continue
		}
		STR = dabs(P2R)
		STI = dabs(P2I)
		P2M = dmax(STR, STI)
		if P2M <= ASCLE {
			continue
		}
		KFLAG = KFLAG + 1
		ASCLE = BRY[KFLAG]
		S1R = S1R * P1R
		S1I = S1I * P1R
		S2R = P2R
		S2I = P2I
		STR = CSSR[KFLAG]
		S1R = S1R * STR
		S1I = S1I * STR
		S2R = S2R * STR
		S2I = S2I * STR
		P1R = CSRR[KFLAG]
	}
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM

	// IFLAG=1 CASES, FORWARD RECURRENCE ON SCALED VALUES ON UNDERFLOW.
TwoSixtyOne:
	HELIM = 0.5E0 * ELIM
	ELM = dexp(-ELIM)
	CELMR = ELM
	ASCLE = BRY[1]
	ZDR = ZR
	ZDI = ZI
	IC = -1
	J = 2
	for I = 1; I <= INU; I++ {
		STR = S2R
		STI = S2I
		S2R = STR*CKR - STI*CKI + S1R
		S2I = STI*CKR + STR*CKI + S1I
		S1R = STR
		S1I = STI
		CKR = CKR + RZR
		CKI = CKI + RZI
		AS = zabs(complex(S2R, S2I))
		ALAS = dlog(AS)
		P2R = -ZDR + ALAS
		if P2R < (-ELIM) {
			goto TwoSixtyThree
		}
		tmp = zlog(complex(S2R, S2I))
		STR = real(tmp)
		STI = imag(tmp)
		P2R = -ZDR + STR
		P2I = -ZDI + STI
		P2M = dexp(P2R) / TOL
		P1R = P2M * dcos(P2I)
		P1I = P2M * dsin(P2I)
		P1R, P1I, NW, ASCLE, TOL = zuchkOrig(P1R, P1I, NW, ASCLE, TOL)
		if NW != 0 {
			goto TwoSixtyThree
		}
		J = 3 - J
		CYR[J] = P1R
		CYI[J] = P1I
		if IC == (I - 1) {
			goto TwoSixtyFour
		}
		IC = I
		continue
	TwoSixtyThree:
		if ALAS < HELIM {
			continue
		}
		ZDR = ZDR - ELIM
		S1R = S1R * CELMR
		S1I = S1I * CELMR
		S2R = S2R * CELMR
		S2I = S2I * CELMR
	}
	if N != 1 {
		goto TwoSeventy
	}
	S1R = S2R
	S1I = S2I
	goto TwoSeventy
TwoSixtyFour:
	KFLAG = 1
	INUB = I + 1
	S2R = CYR[J]
	S2I = CYI[J]
	J = 3 - J
	S1R = CYR[J]
	S1I = CYI[J]
	if INUB <= INU {
		goto TwoTwentyFive
	}
	if N != 1 {
		goto TwoFourty
	}
	S1R = S2R
	S1I = S2I
	goto TwoFourty
TwoSeventy:
	YR[1] = S1R
	YI[1] = S1I
	if N == 1 {
		goto TwoEighty
	}
	YR[2] = S2R
	YI[2] = S2I
TwoEighty:
	ASCLE = BRY[1]
	ZDR, ZDI, FNU, N, YR, YI, NZ, RZR, RZI, ASCLE, TOL, ELIM = zksclOrig(ZDR, ZDI, FNU, N, YR, YI, NZ, RZR, RZI, ASCLE, TOL, ELIM)
	INU = N - NZ
	if INU <= 0 {
		return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
	}
	KK = NZ + 1
	S1R = YR[KK]
	S1I = YI[KK]
	YR[KK] = S1R * CSRR[1]
	YI[KK] = S1I * CSRR[1]
	if INU == 1 {
		return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
	}
	KK = NZ + 2
	S2R = YR[KK]
	S2I = YI[KK]
	YR[KK] = S2R * CSRR[1]
	YI[KK] = S2I * CSRR[1]
	if INU == 2 {
		return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
	}
	T2 = FNU + float64(float32(KK-1))
	CKR = T2 * RZR
	CKI = T2 * RZI
	KFLAG = 1
	goto TwoFifty
TwoNinety:

	// SCALE BY dexp(Z), IFLAG = 1 CASES.

	KODED = 2
	IFLAG = 1
	KFLAG = 2
	goto OneTwenty

	// FNU=HALF ODD INTEGER CASE, DNU=-0.5
ThreeHundred:
	S1R = COEFR
	S1I = COEFI
	S2R = COEFR
	S2I = COEFI
	goto TwoTen

ThreeTen:
	NZ = -2
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
}

// SET K FUNCTIONS TO ZERO ON UNDERFLOW, CONTINUE RECURRENCE
// ON SCALED FUNCTIONS UNTIL TWO MEMBERS COME ON SCALE, THEN
// return WITH MIN(NZ+2,N) VALUES SCALED BY 1/TOL.
func zksclOrig(ZRR, ZRI, FNU float64, N int, YR, YI []float64, NZ int, RZR, RZI, ASCLE, TOL, ELIM float64) (
	ZRRout, ZRIout, FNUout float64, Nout int, YRout, YIout []float64, NZout int, RZRout, RZIout, ASCLEout, TOLout, ELIMout float64) {
	var ACS, AS, CKI, CKR, CSI, CSR, FN, STR, S1I, S1R, S2I,
		S2R, ZEROI, ZEROR, ZDR, ZDI, CELMR, ELM, HELIM, ALAS float64

	var I, IC, KK, NN, NW int
	var tmp complex128
	var CYR, CYI [3]float64
	// DIMENSION YR(N), YI(N), CYR(2), CYI(2)
	ZEROR = 0
	ZEROI = 0
	NZ = 0
	IC = 0
	NN = min0(2, N)
	for I = 1; I <= NN; I++ {
		S1R = YR[I]
		S1I = YI[I]
		CYR[I] = S1R
		CYI[I] = S1I
		AS = zabs(complex(S1R, S1I))
		ACS = -ZRR + dlog(AS)
		NZ = NZ + 1
		YR[I] = ZEROR
		YI[I] = ZEROI
		if ACS < (-ELIM) {
			continue
		}

		tmp = zlog(complex(S1R, S1I))
		CSR = real(tmp)
		CSI = imag(tmp)
		CSR = CSR - ZRR
		CSI = CSI - ZRI
		STR = dexp(CSR) / TOL
		CSR = STR * dcos(CSI)
		CSI = STR * dsin(CSI)
		CSR, CSI, NW, ASCLE, TOL = zuchkOrig(CSR, CSI, NW, ASCLE, TOL)
		if NW != 0 {
			continue
		}
		YR[I] = CSR
		YI[I] = CSI
		IC = I
		NZ = NZ - 1
	}
	if N == 1 {
		return ZRR, ZRI, FNU, N, YR, YI, NZ, RZR, RZI, ASCLE, TOL, ELIM
	}
	if IC > 1 {
		goto Twenty
	}
	YR[1] = ZEROR
	YI[1] = ZEROI
	NZ = 2
Twenty:
	if N == 2 {
		return ZRR, ZRI, FNU, N, YR, YI, NZ, RZR, RZI, ASCLE, TOL, ELIM
	}
	if NZ == 0 {
		return ZRR, ZRI, FNU, N, YR, YI, NZ, RZR, RZI, ASCLE, TOL, ELIM
	}
	FN = FNU + 1.0E0
	CKR = FN * RZR
	CKI = FN * RZI
	S1R = CYR[1]
	S1I = CYI[1]
	S2R = CYR[2]
	S2I = CYI[2]
	HELIM = 0.5E0 * ELIM
	ELM = dexp(-ELIM)
	CELMR = ELM
	ZDR = ZRR
	ZDI = ZRI

	// FIND TWO CONSECUTIVE Y VALUES ON SCALE. SCALE RECURRENCE IF
	// S2 GETS LARGER THAN EXP(ELIM/2)
	for I = 3; I <= N; I++ {
		KK = I
		CSR = S2R
		CSI = S2I
		S2R = CKR*CSR - CKI*CSI + S1R
		S2I = CKI*CSR + CKR*CSI + S1I
		S1R = CSR
		S1I = CSI
		CKR = CKR + RZR
		CKI = CKI + RZI
		AS = zabs(complex(S2R, S2I))
		ALAS = dlog(AS)
		ACS = -ZDR + ALAS
		NZ = NZ + 1
		YR[I] = ZEROR
		YI[I] = ZEROI
		if ACS < (-ELIM) {
			goto TwentyFive
		}
		tmp = zlog(complex(S2R, S2I))
		CSR = real(tmp)
		CSI = imag(tmp)
		CSR = CSR - ZDR
		CSI = CSI - ZDI
		STR = dexp(CSR) / TOL
		CSR = STR * dcos(CSI)
		CSI = STR * dsin(CSI)
		CSR, CSI, NW, ASCLE, TOL = zuchkOrig(CSR, CSI, NW, ASCLE, TOL)
		if NW != 0 {
			goto TwentyFive
		}
		YR[I] = CSR
		YI[I] = CSI
		NZ = NZ - 1
		if IC == KK-1 {
			goto Forty
		}
		IC = KK
		continue
	TwentyFive:
		if ALAS < HELIM {
			continue
		}
		ZDR = ZDR - ELIM
		S1R = S1R * CELMR
		S1I = S1I * CELMR
		S2R = S2R * CELMR
		S2I = S2I * CELMR
	}
	NZ = N
	if IC == N {
		NZ = N - 1
	}
	goto FourtyFive
Forty:
	NZ = KK - 2
FourtyFive:
	for I = 1; I <= NZ; I++ {
		YR[I] = ZEROR
		YI[I] = ZEROI
	}
	return ZRR, ZRI, FNU, N, YR, YI, NZ, RZR, RZI, ASCLE, TOL, ELIM
}

// Y ENTERS AS A SCALED QUANTITY WHOSE MAGNITUDE IS GREATER THAN
// EXP(-ALIM)=ASCLE=1.0E+3*dmach[1)/TOL. THE TEST IS MADE TO SEE
// if THE MAGNITUDE OF THE REAL OR IMAGINARY PART WOULD UNDERFLOW
// WHEN Y IS SCALED (BY TOL) TO ITS PROPER VALUE. Y IS ACCEPTED
// if THE UNDERFLOW IS AT LEAST ONE PRECISION BELOW THE MAGNITUDE
// OF THE LARGEST COMPONENT; OTHERWISE THE PHASE ANGLE DOES NOT HAVE
// ABSOLUTE ACCURACY AND AN UNDERFLOW IS ASSUMED.
func zuchkOrig(YR, YI float64, NZ int, ASCLE, TOL float64) (YRout, YIout float64, NZout int, ASCLEout, TOLout float64) {
	var SS, ST, WR, WI float64
	NZ = 0
	WR = dabs(YR)
	WI = dabs(YI)
	ST = dmin(WR, WI)
	if ST > ASCLE {
		return YR, YI, NZ, ASCLE, TOL
	}
	SS = dmax(WR, WI)
	ST = ST / TOL
	if SS < ST {
		NZ = 1
	}
	return YR, YI, NZ, ASCLE, TOL
}

// ZACAI APPLIES THE ANALYTIC CONTINUATION FORMULA
//
//  K(FNU,ZN*EXP(MP))=K(FNU,ZN)*EXP(-MP*FNU) - MP*I(FNU,ZN)
//        MP=PI*MR*CMPLX(0.0,1.0)
//
// TO CONTINUE THE K FUNCTION FROM THE RIGHT HALF TO THE LEFT
// HALF Z PLANE FOR USE WITH ZAIRY WHERE FNU=1/3 OR 2/3 AND N=1.
// ZACAI IS THE SAME AS ZACON WITH THE PARTS FOR LARGER ORDERS AND
// RECURRENCE REMOVED. A RECURSIVE CALL TO ZACON CAN RESULT if ZACON
// IS CALLED FROM ZAIRY.
func zacaiOrig(ZR, ZI, FNU float64, KODE, MR, N int, YR, YI []float64, NZ int, RL, TOL, ELIM, ALIM float64) (
	ZRout, ZIout, FNUout float64, KODEout, MRout, Nout int, YRout, YIout []float64, NZout int, RLout, TOLout, ELIMout, ALIMout float64) {
	var ARG, ASCLE, AZ, CSGNR, CSGNI, CSPNR,
		CSPNI, C1R, C1I, C2R, C2I, DFNU, FMR, PI,
		SGN, YY, ZNR, ZNI float64
	var INU, IUF, NN, NW int
	CYR := []float64{math.NaN(), 0, 0}
	CYI := []float64{math.NaN(), 0, 0}

	PI = math.Pi
	NZ = 0
	ZNR = -ZR
	ZNI = -ZI
	AZ = zabs(complex(ZR, ZI))
	NN = N
	DFNU = FNU + float64(float32(N-1))
	if AZ <= 2.0E0 {
		goto Ten
	}
	if AZ*AZ*0.25 > DFNU+1.0E0 {
		goto Twenty
	}
Ten:
	// POWER SERIES FOR THE I FUNCTION.
	ZNR, ZNI, FNU, KODE, NN, YR, YI, NW, TOL, ELIM, ALIM = zseriOrig(ZNR, ZNI, FNU, KODE, NN, YR, YI, NW, TOL, ELIM, ALIM)
	goto Forty
Twenty:
	if AZ < RL {
		goto Thirty
	}
	// ASYMPTOTIC EXPANSION FOR LARGE Z FOR THE I FUNCTION.
	ZNR, ZNI, FNU, KODE, NN, YR, YI, NW, RL, TOL, ELIM, ALIM = zasyiOrig(ZNR, ZNI, FNU, KODE, NN, YR, YI, NW, RL, TOL, ELIM, ALIM)
	if NW < 0 {
		goto Eighty
	}
	goto Forty
Thirty:
	// MILLER ALGORITHM NORMALIZED BY THE SERIES FOR THE I FUNCTION
	ZNR, ZNI, FNU, KODE, NN, YR, YI, NW, TOL = zmlriOrig(ZNR, ZNI, FNU, KODE, NN, YR, YI, NW, TOL)
	if NW < 0 {
		goto Eighty
	}
Forty:
	// ANALYTIC CONTINUATION TO THE LEFT HALF PLANE FOR THE K FUNCTION.
	ZNR, ZNI, FNU, KODE, _, CYR, CYI, NW, TOL, ELIM, ALIM = zbknuOrig(ZNR, ZNI, FNU, KODE, 1, CYR, CYI, NW, TOL, ELIM, ALIM)
	if NW != 0 {
		goto Eighty
	}
	FMR = float64(float32(MR))
	SGN = -math.Copysign(PI, FMR)
	CSGNR = 0.0E0
	CSGNI = SGN
	if KODE == 1 {
		goto Fifty
	}
	YY = -ZNI
	CSGNR = -CSGNI * dsin(YY)
	CSGNI = CSGNI * dcos(YY)
Fifty:
	// CALCULATE CSPN=EXP(FNU*PI*I) TO MINIMIZE LOSSES OF SIGNIFICANCE
	// WHEN FNU IS LARGE
	INU = int(float32(FNU))
	ARG = (FNU - float64(float32(INU))) * SGN
	CSPNR = dcos(ARG)
	CSPNI = dsin(ARG)
	if INU%2 == 0 {
		goto Sixty
	}
	CSPNR = -CSPNR
	CSPNI = -CSPNI
Sixty:
	C1R = CYR[1]
	C1I = CYI[1]
	C2R = YR[1]
	C2I = YI[1]
	if KODE == 1 {
		goto Seventy
	}
	IUF = 0
	ASCLE = 1.0E+3 * dmach[1] / TOL
	ZNR, ZNI, C1R, C1I, C2R, C2I, NW, ASCLE, ALIM, IUF = zs1s2Orig(ZNR, ZNI, C1R, C1I, C2R, C2I, NW, ASCLE, ALIM, IUF)
	NZ = NZ + NW
Seventy:
	YR[1] = CSPNR*C1R - CSPNI*C1I + CSGNR*C2R - CSGNI*C2I
	YI[1] = CSPNR*C1I + CSPNI*C1R + CSGNR*C2I + CSGNI*C2R
	return ZR, ZI, FNU, KODE, MR, N, YR, YI, NZ, RL, TOL, ELIM, ALIM
Eighty:
	NZ = -1
	if NW == -2 {
		NZ = -2
	}
	return ZR, ZI, FNU, KODE, MR, N, YR, YI, NZ, RL, TOL, ELIM, ALIM
}

// ZASYI COMPUTES THE I BESSEL FUNCTION FOR REAL(Z)>=0.0 BY
// MEANS OF THE ASYMPTOTIC EXPANSION FOR LARGE CABS(Z) IN THE
// REGION CABS(Z)>MAX(RL,FNU*FNU/2). NZ=0 IS A NORMAL return.
// NZ<0 INDICATES AN OVERFLOW ON KODE=1.
func zasyiOrig(ZR, ZI, FNU float64, KODE, N int, YR, YI []float64, NZ int, RL, TOL, ELIM, ALIM float64) (
	ZRout, ZIout, FNUout float64, KODEout, Nout int, YRout, YIout []float64, NZout int, RLout, TOLout, ELIMout, ALIMout float64) {
	var AA, AEZ, AK, AK1I, AK1R, ARG, ARM, ATOL,
		AZ, BB, BK, CKI, CKR, CONEI, CONER, CS1I, CS1R, CS2I, CS2R, CZI,
		CZR, DFNU, DKI, DKR, DNU2, EZI, EZR, FDN, PI, P1I,
		P1R, RAZ, RTPI, RTR1, RZI, RZR, S, SGN, SQK, STI, STR, S2I,
		S2R, TZI, TZR, ZEROI, ZEROR float64

	var I, IB, IL, INU, J, JL, K, KODED, M, NN int
	var tmp complex128

	PI = math.Pi
	RTPI = 0.159154943091895336E0
	ZEROR = 0
	ZEROI = 0
	CONER = 1
	CONEI = 0

	NZ = 0
	AZ = zabs(complex(ZR, ZI))
	ARM = 1.0E3 * dmach[1]
	RTR1 = dsqrt(ARM)
	IL = min0(2, N)
	DFNU = FNU + float64(float32(N-IL))

	// OVERFLOW TEST
	RAZ = 1.0E0 / AZ
	STR = ZR * RAZ
	STI = -ZI * RAZ
	AK1R = RTPI * STR * RAZ
	AK1I = RTPI * STI * RAZ
	tmp = zsqrt(complex(AK1R, AK1I))
	AK1R = real(tmp)
	AK1I = imag(tmp)
	CZR = ZR
	CZI = ZI
	if KODE != 2 {
		goto Ten
	}
	CZR = ZEROR
	CZI = ZI
Ten:
	if dabs(CZR) > ELIM {
		goto OneHundred
	}
	DNU2 = DFNU + DFNU
	KODED = 1
	if (dabs(CZR) > ALIM) && (N > 2) {
		goto Twenty
	}
	KODED = 0
	tmp = zexp(complex(CZR, CZI))
	STR = real(tmp)
	STI = imag(tmp)
	tmp = zmlt(complex(AK1R, AK1I), complex(STR, STI))
	AK1R = real(tmp)
	AK1I = imag(tmp)
Twenty:
	FDN = 0.0E0
	if DNU2 > RTR1 {
		FDN = DNU2 * DNU2
	}
	EZR = ZR * 8.0E0
	EZI = ZI * 8.0E0

	// WHEN Z IS IMAGINARY, THE ERROR TEST MUST BE MADE RELATIVE TO THE
	// FIRST RECIPROCAL POWER SINCE THIS IS THE LEADING TERM OF THE
	// EXPANSION FOR THE IMAGINARY PART.
	AEZ = 8.0E0 * AZ
	S = TOL / AEZ
	JL = int(float32(RL+RL)) + 2
	P1R = ZEROR
	P1I = ZEROI
	if ZI == 0.0E0 {
		goto Thirty
	}

	// CALCULATE EXP(PI*(0.5+FNU+N-IL)*I) TO MINIMIZE LOSSES OF
	// SIGNIFICANCE WHEN FNU OR N IS LARGE
	INU = int(float32(FNU))
	ARG = (FNU - float64(float32(INU))) * PI
	INU = INU + N - IL
	AK = -dsin(ARG)
	BK = dcos(ARG)
	if ZI < 0.0E0 {
		BK = -BK
	}
	P1R = AK
	P1I = BK
	if INU%2 == 0 {
		goto Thirty
	}
	P1R = -P1R
	P1I = -P1I
Thirty:
	for K = 1; K <= IL; K++ {
		SQK = FDN - 1.0E0
		ATOL = S * dabs(SQK)
		SGN = 1.0E0
		CS1R = CONER
		CS1I = CONEI
		CS2R = CONER
		CS2I = CONEI
		CKR = CONER
		CKI = CONEI
		AK = 0.0E0
		AA = 1.0E0
		BB = AEZ
		DKR = EZR
		DKI = EZI
		// TODO(btracey): This loop is executed tens of thousands of times. Why?
		// is that really necessary?
		for J = 1; J <= JL; J++ {
			tmp = zdiv(complex(CKR, CKI), complex(DKR, DKI))
			STR = real(tmp)
			STI = imag(tmp)
			CKR = STR * SQK
			CKI = STI * SQK
			CS2R = CS2R + CKR
			CS2I = CS2I + CKI
			SGN = -SGN
			CS1R = CS1R + CKR*SGN
			CS1I = CS1I + CKI*SGN
			DKR = DKR + EZR
			DKI = DKI + EZI
			AA = AA * dabs(SQK) / BB
			BB = BB + AEZ
			AK = AK + 8.0E0
			SQK = SQK - AK
			if AA <= ATOL {
				goto Fifty
			}
		}
		goto OneTen
	Fifty:
		S2R = CS1R
		S2I = CS1I
		if ZR+ZR >= ELIM {
			goto Sixty
		}
		TZR = ZR + ZR
		TZI = ZI + ZI
		tmp = zexp(complex(-TZR, -TZI))
		STR = real(tmp)
		STI = imag(tmp)
		tmp = zmlt(complex(STR, STI), complex(P1R, P1I))
		STR = real(tmp)
		STI = imag(tmp)
		tmp = zmlt(complex(STR, STI), complex(CS2R, CS2I))
		STR = real(tmp)
		STI = imag(tmp)
		S2R = S2R + STR
		S2I = S2I + STI
	Sixty:
		FDN = FDN + 8.0E0*DFNU + 4.0E0
		P1R = -P1R
		P1I = -P1I
		M = N - IL + K
		YR[M] = S2R*AK1R - S2I*AK1I
		YI[M] = S2R*AK1I + S2I*AK1R
	}
	if N <= 2 {
		return ZR, ZI, FNU, KODE, N, YR, YI, NZ, RL, TOL, ELIM, ALIM
	}
	NN = N
	K = NN - 2
	AK = float64(float32(K))
	STR = ZR * RAZ
	STI = -ZI * RAZ
	RZR = (STR + STR) * RAZ
	RZI = (STI + STI) * RAZ
	IB = 3
	for I = IB; I <= NN; I++ {
		YR[K] = (AK+FNU)*(RZR*YR[K+1]-RZI*YI[K+1]) + YR[K+2]
		YI[K] = (AK+FNU)*(RZR*YI[K+1]+RZI*YR[K+1]) + YI[K+2]
		AK = AK - 1.0E0
		K = K - 1
	}
	if KODED == 0 {
		return ZR, ZI, FNU, KODE, N, YR, YI, NZ, RL, TOL, ELIM, ALIM
	}
	tmp = zexp(complex(CZR, CZI))
	CKR = real(tmp)
	CKI = imag(tmp)
	for I = 1; I <= NN; I++ {
		STR = YR[I]*CKR - YI[I]*CKI
		YI[I] = YR[I]*CKI + YI[I]*CKR
		YR[I] = STR
	}
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, RL, TOL, ELIM, ALIM
OneHundred:
	NZ = -1
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, RL, TOL, ELIM, ALIM
OneTen:
	NZ = -2
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, RL, TOL, ELIM, ALIM
}

// ZMLRI COMPUTES THE I BESSEL FUNCTION FOR RE(Z)>=0.0 BY THE
// MILLER ALGORITHM NORMALIZED BY A NEUMANN SERIES.
func zmlriOrig(ZR, ZI, FNU float64, KODE, N int, YR, YI []float64, NZ int, TOL float64) (
	ZRout, ZIout, FNUout float64, KODEout, Nout int, YRout, YIout []float64, NZout int, TOLout float64) {
	var ACK, AK, AP, AT, AZ, BK, CKI, CKR, CNORMI,
		CNORMR, CONEI, CONER, FKAP, FKK, FLAM, FNF, PTI, PTR, P1I,
		P1R, P2I, P2R, RAZ, RHO, RHO2, RZI, RZR, SCLE, STI, STR, SUMI,
		SUMR, TFNF, TST, ZEROI, ZEROR float64
	var I, IAZ, IDUM, IFNU, INU, ITIME, K, KK, KM, M int
	var tmp complex128
	ZEROR = 0
	ZEROI = 0
	CONER = 1
	CONEI = 0

	SCLE = dmach[1] / TOL
	NZ = 0
	AZ = zabs(complex(ZR, ZI))
	IAZ = int(float32(AZ))
	IFNU = int(float32(FNU))
	INU = IFNU + N - 1
	AT = float64(float32(IAZ)) + 1.0E0
	RAZ = 1.0E0 / AZ
	STR = ZR * RAZ
	STI = -ZI * RAZ
	CKR = STR * AT * RAZ
	CKI = STI * AT * RAZ
	RZR = (STR + STR) * RAZ
	RZI = (STI + STI) * RAZ
	P1R = ZEROR
	P1I = ZEROI
	P2R = CONER
	P2I = CONEI
	ACK = (AT + 1.0E0) * RAZ
	RHO = ACK + dsqrt(ACK*ACK-1.0E0)
	RHO2 = RHO * RHO
	TST = (RHO2 + RHO2) / ((RHO2 - 1.0E0) * (RHO - 1.0E0))
	TST = TST / TOL

	// COMPUTE RELATIVE TRUNCATION ERROR INDEX FOR SERIES.
	//fmt.Println("before loop", P2R, P2I, CKR, CKI, RZR, RZI, TST, AK)
	AK = AT
	for I = 1; I <= 80; I++ {
		PTR = P2R
		PTI = P2I
		P2R = P1R - (CKR*PTR - CKI*PTI)
		P2I = P1I - (CKI*PTR + CKR*PTI)
		P1R = PTR
		P1I = PTI
		CKR = CKR + RZR
		CKI = CKI + RZI
		AP = zabs(complex(P2R, P2I))
		if AP > TST*AK*AK {
			goto Twenty
		}
		AK = AK + 1.0E0
	}
	goto OneTen
Twenty:
	I = I + 1
	K = 0
	if INU < IAZ {
		goto Forty
	}
	// COMPUTE RELATIVE TRUNCATION ERROR FOR RATIOS.
	P1R = ZEROR
	P1I = ZEROI
	P2R = CONER
	P2I = CONEI
	AT = float64(float32(INU)) + 1.0E0
	STR = ZR * RAZ
	STI = -ZI * RAZ
	CKR = STR * AT * RAZ
	CKI = STI * AT * RAZ
	ACK = AT * RAZ
	TST = dsqrt(ACK / TOL)
	ITIME = 1
	for K = 1; K <= 80; K++ {
		PTR = P2R
		PTI = P2I
		P2R = P1R - (CKR*PTR - CKI*PTI)
		P2I = P1I - (CKR*PTI + CKI*PTR)
		P1R = PTR
		P1I = PTI
		CKR = CKR + RZR
		CKI = CKI + RZI
		AP = zabs(complex(P2R, P2I))
		if AP < TST {
			continue
		}
		if ITIME == 2 {
			goto Forty
		}
		ACK = zabs(complex(CKR, CKI))
		FLAM = ACK + dsqrt(ACK*ACK-1.0E0)
		FKAP = AP / zabs(complex(P1R, P1I))
		RHO = dmin(FLAM, FKAP)
		TST = TST * dsqrt(RHO/(RHO*RHO-1.0E0))
		ITIME = 2
	}
	goto OneTen
Forty:
	// BACKWARD RECURRENCE AND SUM NORMALIZING RELATION.
	K = K + 1
	KK = max0(I+IAZ, K+INU)
	FKK = float64(float32(KK))
	P1R = ZEROR
	P1I = ZEROI

	// SCALE P2 AND SUM BY SCLE.
	P2R = SCLE
	P2I = ZEROI
	FNF = FNU - float64(float32(IFNU))
	TFNF = FNF + FNF
	BK = dgamln(FKK+TFNF+1.0E0, IDUM) - dgamln(FKK+1.0E0, IDUM) - dgamln(TFNF+1.0E0, IDUM)
	BK = dexp(BK)
	SUMR = ZEROR
	SUMI = ZEROI
	KM = KK - INU
	for I = 1; I <= KM; I++ {
		PTR = P2R
		PTI = P2I
		P2R = P1R + (FKK+FNF)*(RZR*PTR-RZI*PTI)
		P2I = P1I + (FKK+FNF)*(RZI*PTR+RZR*PTI)
		P1R = PTR
		P1I = PTI
		AK = 1.0E0 - TFNF/(FKK+TFNF)
		ACK = BK * AK
		SUMR = SUMR + (ACK+BK)*P1R
		SUMI = SUMI + (ACK+BK)*P1I
		BK = ACK
		FKK = FKK - 1.0E0
	}
	YR[N] = P2R
	YI[N] = P2I
	if N == 1 {
		goto Seventy
	}
	for I = 2; I <= N; I++ {
		PTR = P2R
		PTI = P2I
		P2R = P1R + (FKK+FNF)*(RZR*PTR-RZI*PTI)
		P2I = P1I + (FKK+FNF)*(RZI*PTR+RZR*PTI)
		P1R = PTR
		P1I = PTI
		AK = 1.0E0 - TFNF/(FKK+TFNF)
		ACK = BK * AK
		SUMR = SUMR + (ACK+BK)*P1R
		SUMI = SUMI + (ACK+BK)*P1I
		BK = ACK
		FKK = FKK - 1.0E0
		M = N - I + 1
		YR[M] = P2R
		YI[M] = P2I
	}
Seventy:
	if IFNU <= 0 {
		goto Ninety
	}
	for I = 1; I <= IFNU; I++ {
		PTR = P2R
		PTI = P2I
		P2R = P1R + (FKK+FNF)*(RZR*PTR-RZI*PTI)
		P2I = P1I + (FKK+FNF)*(RZR*PTI+RZI*PTR)
		P1R = PTR
		P1I = PTI
		AK = 1.0E0 - TFNF/(FKK+TFNF)
		ACK = BK * AK
		SUMR = SUMR + (ACK+BK)*P1R
		SUMI = SUMI + (ACK+BK)*P1I
		BK = ACK
		FKK = FKK - 1.0E0
	}
Ninety:
	PTR = ZR
	PTI = ZI
	if KODE == 2 {
		PTR = ZEROR
	}
	tmp = zlog(complex(RZR, RZI))
	STR = real(tmp)
	STI = imag(tmp)
	P1R = -FNF*STR + PTR
	P1I = -FNF*STI + PTI
	AP = dgamln(1.0E0+FNF, IDUM)
	PTR = P1R - AP
	PTI = P1I

	// THE DIVISION CEXP(PT)/(SUM+P2) IS ALTERED TO AVOID OVERFLOW
	// IN THE DENOMINATOR BY SQUARING LARGE QUANTITIES.
	P2R = P2R + SUMR
	P2I = P2I + SUMI
	AP = zabs(complex(P2R, P2I))
	P1R = 1.0E0 / AP
	tmp = zexp(complex(PTR, PTI))
	STR = real(tmp)
	STI = imag(tmp)
	CKR = STR * P1R
	CKI = STI * P1R
	PTR = P2R * P1R
	PTI = -P2I * P1R
	tmp = zmlt(complex(CKR, CKI), complex(PTR, PTI))
	CNORMR = real(tmp)
	CNORMI = imag(tmp)
	for I = 1; I <= N; I++ {
		STR = YR[I]*CNORMR - YI[I]*CNORMI
		YI[I] = YR[I]*CNORMI + YI[I]*CNORMR
		YR[I] = STR
	}
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL
OneTen:
	NZ = -2
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL
}

// ZSERI COMPUTES THE I BESSEL FUNCTION FOR REAL(Z)>=0.0 BY
// MEANS OF THE POWER SERIES FOR LARGE CABS(Z) IN THE
// REGION CABS(Z)<=2*SQRT(FNU+1). NZ=0 IS A NORMAL return.
// NZ>0 MEANS THAT THE LAST NZ COMPONENTS WERE SET TO ZERO
// DUE TO UNDERFLOW. NZ<0 MEANS UNDERFLOW OCCURRED, BUT THE
// CONDITION CABS(Z)<=2*SQRT(FNU+1) WAS VIOLATED AND THE
// COMPUTATION MUST BE COMPLETED IN ANOTHER ROUTINE WITH N=N-ABS(NZ).
func zseriOrig(ZR, ZI, FNU float64, KODE, N int, YR, YI []float64, NZ int, TOL, ELIM, ALIM float64) (
	ZRout, ZIout, FNUout float64, KODEout, Nout int, YRout, YIout []float64, NZout int, TOLout, ELIMout, ALIMout float64) {
	var AA, ACZ, AK, AK1I, AK1R, ARM, ASCLE, ATOL,
		AZ, CKI, CKR, COEFI, COEFR, CONEI, CONER, CRSCR, CZI, CZR, DFNU,
		FNUP, HZI, HZR, RAZ, RS, RTR1, RZI, RZR, S, SS, STI,
		STR, S1I, S1R, S2I, S2R, ZEROI, ZEROR float64
	var I, IB, IDUM, IFLAG, IL, K, L, M, NN, NW int
	var WR, WI [3]float64
	var tmp complex128

	CONER = 1.0
	NZ = 0
	AZ = zabs(complex(ZR, ZI))
	if AZ == 0.0E0 {
		goto OneSixty
	}
	// TODO(btracey)
	// The original fortran line is "ARM = 1.0D+3*D1MACH(1)". Evidently, in Fortran
	// this is interpreted as one to the power of +3*D1MACH(1). While it is possible
	// this was intentional, it seems unlikely.
	//ARM = 1.0E0 + 3*dmach[1]
	//math.Pow(1, 3*dmach[1])
	ARM = 1000 * dmach[1]
	RTR1 = dsqrt(ARM)
	CRSCR = 1.0E0
	IFLAG = 0
	if AZ < ARM {
		goto OneFifty
	}
	HZR = 0.5E0 * ZR
	HZI = 0.5E0 * ZI
	CZR = ZEROR
	CZI = ZEROI
	if AZ <= RTR1 {
		goto Ten
	}
	tmp = zmlt(complex(HZR, HZI), complex(HZR, HZI))
	CZR = real(tmp)
	CZI = imag(tmp)
Ten:
	ACZ = zabs(complex(CZR, CZI))
	NN = N
	tmp = zlog(complex(HZR, HZI))
	CKR = real(tmp)
	CKI = imag(tmp)
Twenty:
	DFNU = FNU + float64(float32(NN-1))
	FNUP = DFNU + 1.0E0

	// UNDERFLOW TEST.
	AK1R = CKR * DFNU
	AK1I = CKI * DFNU
	AK = dgamln(FNUP, IDUM)
	AK1R = AK1R - AK
	if KODE == 2 {
		AK1R = AK1R - ZR
	}
	if AK1R > (-ELIM) {
		goto Forty
	}
Thirty:
	NZ = NZ + 1
	YR[NN] = ZEROR
	YI[NN] = ZEROI
	if ACZ > DFNU {
		goto OneNinety
	}
	NN = NN - 1
	if NN == 0 {
		return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
	}
	goto Twenty
Forty:
	if AK1R > (-ALIM) {
		goto Fifty
	}
	IFLAG = 1
	SS = 1.0E0 / TOL
	CRSCR = TOL
	ASCLE = ARM * SS
Fifty:
	AA = dexp(AK1R)
	if IFLAG == 1 {
		AA = AA * SS
	}
	COEFR = AA * dcos(AK1I)
	COEFI = AA * dsin(AK1I)
	ATOL = TOL * ACZ / FNUP
	IL = min0(2, NN)
	for I = 1; I <= IL; I++ {
		DFNU = FNU + float64(float32(NN-I))
		FNUP = DFNU + 1.0E0
		S1R = CONER
		S1I = CONEI
		if ACZ < TOL*FNUP {
			goto Seventy
		}
		AK1R = CONER
		AK1I = CONEI
		AK = FNUP + 2.0E0
		S = FNUP
		AA = 2.0E0
	Sixty:
		RS = 1.0E0 / S
		STR = AK1R*CZR - AK1I*CZI
		STI = AK1R*CZI + AK1I*CZR
		AK1R = STR * RS
		AK1I = STI * RS
		S1R = S1R + AK1R
		S1I = S1I + AK1I
		S = S + AK
		AK = AK + 2.0E0
		AA = AA * ACZ * RS
		if AA > ATOL {
			goto Sixty
		}
	Seventy:
		S2R = S1R*COEFR - S1I*COEFI
		S2I = S1R*COEFI + S1I*COEFR
		WR[I] = S2R
		WI[I] = S2I
		if IFLAG == 0 {
			goto Eighty
		}
		S2R, S2I, NW, ASCLE, TOL = zuchkOrig(S2R, S2I, NW, ASCLE, TOL)
		if NW != 0 {
			goto Thirty
		}
	Eighty:
		M = NN - I + 1
		YR[M] = S2R * CRSCR
		YI[M] = S2I * CRSCR
		if I == IL {
			continue
		}
		tmp = zdiv(complex(COEFR, COEFI), complex(HZR, HZI))
		STR = real(tmp)
		STI = imag(tmp)
		COEFR = STR * DFNU
		COEFI = STI * DFNU
	}
	if NN <= 2 {
		return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
	}
	K = NN - 2
	AK = float64(float32(K))
	RAZ = 1.0E0 / AZ
	STR = ZR * RAZ
	STI = -ZI * RAZ
	RZR = (STR + STR) * RAZ
	RZI = (STI + STI) * RAZ
	if IFLAG == 1 {
		goto OneTwenty
	}
	IB = 3
OneHundred:
	for I = IB; I <= NN; I++ {
		YR[K] = (AK+FNU)*(RZR*YR[K+1]-RZI*YI[K+1]) + YR[K+2]
		YI[K] = (AK+FNU)*(RZR*YI[K+1]+RZI*YR[K+1]) + YI[K+2]
		AK = AK - 1.0E0
		K = K - 1
	}
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM

	// RECUR BACKWARD WITH SCALED VALUES.
OneTwenty:
	// EXP(-ALIM)=EXP(-ELIM)/TOL=APPROX. ONE PRECISION ABOVE THE
	// UNDERFLOW LIMIT = ASCLE = dmach[1)*SS*1.0D+3.
	S1R = WR[1]
	S1I = WI[1]
	S2R = WR[2]
	S2I = WI[2]
	for L = 3; L <= NN; L++ {
		CKR = S2R
		CKI = S2I
		S2R = S1R + (AK+FNU)*(RZR*CKR-RZI*CKI)
		S2I = S1I + (AK+FNU)*(RZR*CKI+RZI*CKR)
		S1R = CKR
		S1I = CKI
		CKR = S2R * CRSCR
		CKI = S2I * CRSCR
		YR[K] = CKR
		YI[K] = CKI
		AK = AK - 1.0E0
		K = K - 1
		if zabs(complex(CKR, CKI)) > ASCLE {
			goto OneFourty
		}
	}
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
OneFourty:
	IB = L + 1
	if IB > NN {
		return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
	}
	goto OneHundred
OneFifty:
	NZ = N
	if FNU == 0.0E0 {
		NZ = NZ - 1
	}
OneSixty:
	YR[1] = ZEROR
	YI[1] = ZEROI
	if FNU != 0.0E0 {
		goto OneSeventy
	}
	YR[1] = CONER
	YI[1] = CONEI
OneSeventy:
	if N == 1 {
		return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
	}
	for I = 2; I <= N; I++ {
		YR[I] = ZEROR
		YI[I] = ZEROI
	}
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM

	// return WITH NZ<0 if CABS(Z*Z/4)>FNU+N-NZ-1 COMPLETE
	// THE CALCULATION IN CBINU WITH N=N-IABS(NZ)

OneNinety:
	NZ = -NZ
	return ZR, ZI, FNU, KODE, N, YR, YI, NZ, TOL, ELIM, ALIM
}

// ZS1S2 TESTS FOR A POSSIBLE UNDERFLOW RESULTING FROM THE
// ADDITION OF THE I AND K FUNCTIONS IN THE ANALYTIC CON-
// TINUATION FORMULA WHERE S1=K FUNCTION AND S2=I FUNCTION.
// ON KODE=1 THE I AND K FUNCTIONS ARE DIFFERENT ORDERS OF
// MAGNITUDE, BUT FOR KODE=2 THEY CAN BE OF THE SAME ORDER
// OF MAGNITUDE AND THE MAXIMUM MUST BE AT LEAST ONE
// PRECISION ABOVE THE UNDERFLOW LIMIT.
func zs1s2Orig(ZRR, ZRI, S1R, S1I, S2R, S2I float64, NZ int, ASCLE, ALIM float64, IUF int) (
	ZRRout, ZRIout, S1Rout, S1Iout, S2Rout, S2Iout float64, NZout int, ASCLEout, ALIMout float64, IUFout int) {
	var AA, ALN, AS1, AS2, C1I, C1R, S1DI, S1DR, ZEROI, ZEROR float64
	var tmp complex128

	ZEROR = 0
	ZEROI = 0
	NZ = 0
	AS1 = zabs(complex(S1R, S1I))
	AS2 = zabs(complex(S2R, S2I))
	if S1R == 0.0E0 && S1I == 0.0E0 {
		goto Ten
	}
	if AS1 == 0.0E0 {
		goto Ten
	}
	ALN = -ZRR - ZRR + dlog(AS1)
	S1DR = S1R
	S1DI = S1I
	S1R = ZEROR
	S1I = ZEROI
	AS1 = ZEROR
	if ALN < (-ALIM) {
		goto Ten
	}
	tmp = zlog(complex(S1DR, S1DI))
	C1R = real(tmp)
	C1I = imag(tmp)

	C1R = C1R - ZRR - ZRR
	C1I = C1I - ZRI - ZRI
	tmp = zexp(complex(C1R, C1I))
	S1R = real(tmp)
	S1I = imag(tmp)
	AS1 = zabs(complex(S1R, S1I))
	IUF = IUF + 1
Ten:
	AA = dmax(AS1, AS2)
	if AA > ASCLE {
		return ZRR, ZRI, S1R, S1I, S2R, S2I, NZ, ASCLE, ALIM, IUF
	}
	S1R = ZEROR
	S1I = ZEROI
	S2R = ZEROR
	S2I = ZEROI
	NZ = 1
	IUF = 0
	return ZRR, ZRI, S1R, S1I, S2R, S2I, NZ, ASCLE, ALIM, IUF
}

// ZSHCH COMPUTES THE COMPLEX HYPERBOLIC FUNCTIONS CSH=SINH(X+iY) AND
// CCH=COSH(X+I*Y), WHERE I**2=-1.
// TODO(btracey): use cmplx.Sinh and cmplx.Cosh.
func zshchOrig(ZR, ZI, CSHR, CSHI, CCHR, CCHI float64) (ZRout, ZIout, CSHRout, CSHIout, CCHRout, CCHIout float64) {
	var CH, CN, SH, SN float64
	SH = math.Sinh(ZR)
	CH = math.Cosh(ZR)
	SN = dsin(ZI)
	CN = dcos(ZI)
	CSHR = SH * CN
	CSHI = CH * SN
	CCHR = CH * CN
	CCHI = SH * SN
	return ZR, ZI, CSHR, CSHI, CCHR, CCHI
}

func dmax(a, b float64) float64 {
	return math.Max(a, b)
}

func dmin(a, b float64) float64 {
	return math.Min(a, b)
}

func dabs(a float64) float64 {
	return math.Abs(a)
}

func datan(a float64) float64 {
	return math.Atan(a)
}

func dtan(a float64) float64 {
	return math.Tan(a)
}

func dlog(a float64) float64 {
	return math.Log(a)
}

func dsin(a float64) float64 {
	return math.Sin(a)
}

func dcos(a float64) float64 {
	return math.Cos(a)
}

func dexp(a float64) float64 {
	return math.Exp(a)
}

func dsqrt(a float64) float64 {
	return math.Sqrt(a)
}

func zmlt(a, b complex128) complex128 {
	return a * b
}

func zdiv(a, b complex128) complex128 {
	return a / b
}

func zabs(a complex128) float64 {
	return cmplx.Abs(a)
}

func zsqrt(a complex128) complex128 {
	return cmplx.Sqrt(a)
}

func zexp(a complex128) complex128 {
	return cmplx.Exp(a)
}

func zlog(a complex128) complex128 {
	return cmplx.Log(a)
}

// Zshch computes the hyperbolic sin and cosine of the input z.
func Zshch(z complex128) (sinh, cosh complex128) {
	return cmplx.Sinh(z), cmplx.Cosh(z)
}
