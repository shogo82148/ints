package ints

import (
	"cmp"
	"fmt"
	"math/bits"
)

// Uint1024 is a type that represents an 1024-bit unsigned integer.
type Uint1024 [16]uint64

// IsZero returns true if a is zero.
func (a Uint1024) IsZero() bool {
	var zero Uint1024
	return a == zero
}

// Add returns the sum a+b.
//
// This function's execution time does not depend on the inputs.
func (a Uint1024) Add(b Uint1024) Uint1024 {
	u15, carry := bits.Add64(a[15], b[15], 0)
	u14, carry := bits.Add64(a[14], b[14], carry)
	u13, carry := bits.Add64(a[13], b[13], carry)
	u12, carry := bits.Add64(a[12], b[12], carry)
	u11, carry := bits.Add64(a[11], b[11], carry)
	u10, carry := bits.Add64(a[10], b[10], carry)
	u9, carry := bits.Add64(a[9], b[9], carry)
	u8, carry := bits.Add64(a[8], b[8], carry)
	u7, carry := bits.Add64(a[7], b[7], carry)
	u6, carry := bits.Add64(a[6], b[6], carry)
	u5, carry := bits.Add64(a[5], b[5], carry)
	u4, carry := bits.Add64(a[4], b[4], carry)
	u3, carry := bits.Add64(a[3], b[3], carry)
	u2, carry := bits.Add64(a[2], b[2], carry)
	u1, carry := bits.Add64(a[1], b[1], carry)
	u0, _ := bits.Add64(a[0], b[0], carry)
	return Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
}

// Sub returns the difference a-b.
//
// This function's execution time does not depend on the inputs.
func (a Uint1024) Sub(b Uint1024) Uint1024 {
	u15, borrow := bits.Sub64(a[15], b[15], 0)
	u14, borrow := bits.Sub64(a[14], b[14], borrow)
	u13, borrow := bits.Sub64(a[13], b[13], borrow)
	u12, borrow := bits.Sub64(a[12], b[12], borrow)
	u11, borrow := bits.Sub64(a[11], b[11], borrow)
	u10, borrow := bits.Sub64(a[10], b[10], borrow)
	u9, borrow := bits.Sub64(a[9], b[9], borrow)
	u8, borrow := bits.Sub64(a[8], b[8], borrow)
	u7, borrow := bits.Sub64(a[7], b[7], borrow)
	u6, borrow := bits.Sub64(a[6], b[6], borrow)
	u5, borrow := bits.Sub64(a[5], b[5], borrow)
	u4, borrow := bits.Sub64(a[4], b[4], borrow)
	u3, borrow := bits.Sub64(a[3], b[3], borrow)
	u2, borrow := bits.Sub64(a[2], b[2], borrow)
	u1, borrow := bits.Sub64(a[1], b[1], borrow)
	u0, _ := bits.Sub64(a[0], b[0], borrow)
	return Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
}

// Mul returns the product a*b.
//
// This function's execution time does not depend on the inputs.
func (a Uint1024) Mul(b Uint1024) Uint1024 {
	//                  a0    a1    a2    a3    a4    a5    a6    a7    a8    a9   a10   a11   a12   a13   a14   a15
	//                x b0    b1    b2    b3    b4    b5    b6    b7    b8    b9   b10   b11   b12   b13   b14   b15
	//                ----------------------------------------------------------------------------------------------
	//                                                                                                   h1515 l1515 - 1.
	//                                                                                             h1415 l1415
	//                                                                                       h1315 l1315
	//                                                                                 h1215 l1215
	//                                                                           h1115 l1115
	//                                                                     h1015 l1015
	//                                                               h0915 l0915
	//                                                         h0815 l0815
	//                                                   h0715 l0715
	//                                             h0615 l0615
	//                                       h0515 l0515
	//                                 h0415 l0415
	//                           h0315 l0315
	//                     h0215 l0215
	//               h0115 l0115
	//         h0015 l0015
	//                                                                                             h1514 l1514       - 2.
	//                                                                                       h1414 l1414
	//                                                                                 h1314 l1314
	//                                                                           h1214 l1214
	//                                                                     h1114 l1114
	//                                                               h1014 l1014
	//                                                         h0914 l0914
	//                                                   h0814 l0814
	//                                             h0714 l0714
	//                                       h0614 l0614
	//                                 h0514 l0514
	//                           h0414 l0414
	//                     h0314 l0314
	//               h0214 l0214
	//         h0114 l0114
	//   h0014 l0014
	//                                                                                       h1513 l1513            - 3.
	//                                                                                 h1413 l1413
	//                                                                           h1313 l1313
	//                                                                     h1213 l1213
	//                                                               h1113 l1113
	//                                                         h1013 l1013
	//                                                   h0913 l0913
	//                                             h0813 l0813
	//                                       h0713 l0713
	//                                 h0613 l0613
	//                           h0513 l0513
	//                     h0413 l0413
	//               h0313 l0313
	//         h0213 l0213
	//   h0113 l0113
	// (snip)
	//                                                                                 h1512 l1512                   - 4.
	//                                                                           h1412 l1412
	//                                                                     h1312 l1312
	//                                                               h1212 l1212
	//                                                         h1112 l1112
	//                                                   h1012 l1012
	//                                             h0912 l0912
	//                                       h0812 l0812
	//                                 h0712 l0712
	//                           h0612 l0612
	//                     h0512 l0512
	//               h0412 l0412
	//         h0312 l0312
	//   h0212 l0212
	// (snip)
	//                                                                           h1511 l1511                          - 5.
	//                                                                     h1411 l1411
	//                                                               h1311 l1311
	//                                                         h1211 l1211
	//                                                   h1111 l1111
	//                                             h1011 l1011
	//                                       h0911 l0911
	//                                 h0811 l0811
	//                           h0711 l0711
	//                     h0611 l0611
	//               h0511 l0511
	//         h0411 l0411
	//   h0311 l0311
	// (snip)
	//                                                                     h1510 l1510                               - 6.
	//                                                               h1410 l1410
	//                                                         h1310 l1310
	//                                                   h1210 l1210
	//                                             h1110 l1110
	//                                       h1010 l1010
	//                                 h0910 l0910
	//                           h0810 l0810
	//                     h0710 l0710
	//               h0610 l0610
	//         h0510 l0510
	//   h0410 l0410
	// (snip)
	//                                                               h1509 l1509                                     - 7.
	//                                                         h1409 l1409
	//                                                   h1309 l1309
	//                                             h1209 l1209
	//                                       h1109 l1109
	//                                 h1009 l1009
	//                           h0909 l0909
	//                     h0809 l0809
	//               h0709 l0709
	//         h0609 l0609
	//   h0509 l0509
	// (snip)
	//                                                         h1508 l1508                                           - 8.
	//                                                   h1408 l1408
	//                                             h1308 l1308
	//                                       h1208 l1208
	//                                 h1108 l1108
	//                           h1008 l1008
	//                     h0908 l0908
	//               h0808 l0808
	//         h0708 l0708
	//   h0608 l0608
	// (snip)
	//                                                   h1507 l1507                                                 - 9.
	//                                             h1407 l1407
	//                                       h1307 l1307
	//                                 h1207 l1207
	//                           h1107 l1107
	//                     h1007 l1007
	//               h0907 l0907
	//         h0807 l0807
	//   h0707 l0707
	// (snip)
	//                                             h1506 l1506                                                      - 10.
	//                                       h1406 l1406
	//                                 h1306 l1306
	//                           h1206 l1206
	//                     h1106 l1106
	//               h1006 l1006
	//         h0906 l0906
	//   h0806 l0806
	// (snip)
	//                                       h1505 l1505                                                            - 11.
	//                                 h1405 l1405
	//                           h1305 l1305
	//                     h1205 l1205
	//               h1105 l1105
	//         h1005 l1005
	//   h0905 l0905
	// (snip)
	//                                 h1504 l1504                                                                  - 12.
	//                           h1404 l1404
	//                     h1304 l1304
	//               h1204 l1204
	//         h1104 l1104
	//   h1004 l1004
	// (snip)
	//                           h1503 l1503                                                                        - 13.
	//                     h1403 l1403
	//               h1303 l1303
	//         h1203 l1203
	//   h1103 l1103
	// (snip)
	//                     h1502 l1502                                                                              - 14.
	//               h1402 l1402
	//         h1302 l1302
	//   h1202 l1202
	// (snip)
	//               h1501 l1501                                                                                    - 15.
	//         h1401 l1401
	//   h1301 l1301
	// (snip)
	//         h1500 l1500                                                                                          - 16.
	//   h1400 l1400
	// (snip)
	// -------------------------------------------------------------------------------------------------------------
	//                  u0    u1    u2    u3    u4    u5    u6    u7    u8    u9   u10   u11   u12   u13   u14   u15

	h1515, l1515 := bits.Mul64(a[15], b[15])
	h1415, l1415 := bits.Mul64(a[14], b[15])
	h1315, l1315 := bits.Mul64(a[13], b[15])
	h1215, l1215 := bits.Mul64(a[12], b[15])
	h1115, l1115 := bits.Mul64(a[11], b[15])
	h1015, l1015 := bits.Mul64(a[10], b[15])
	h0915, l0915 := bits.Mul64(a[9], b[15])
	h0815, l0815 := bits.Mul64(a[8], b[15])
	h0715, l0715 := bits.Mul64(a[7], b[15])
	h0615, l0615 := bits.Mul64(a[6], b[15])
	h0515, l0515 := bits.Mul64(a[5], b[15])
	h0415, l0415 := bits.Mul64(a[4], b[15])
	h0315, l0315 := bits.Mul64(a[3], b[15])
	h0215, l0215 := bits.Mul64(a[2], b[15])
	h0115, l0115 := bits.Mul64(a[1], b[15])
	_, l0015 := bits.Mul64(a[0], b[15])

	h1514, l1514 := bits.Mul64(a[15], b[14])
	h1414, l1414 := bits.Mul64(a[14], b[14])
	h1314, l1314 := bits.Mul64(a[13], b[14])
	h1214, l1214 := bits.Mul64(a[12], b[14])
	h1114, l1114 := bits.Mul64(a[11], b[14])
	h1014, l1014 := bits.Mul64(a[10], b[14])
	h0914, l0914 := bits.Mul64(a[9], b[14])
	h0814, l0814 := bits.Mul64(a[8], b[14])
	h0714, l0714 := bits.Mul64(a[7], b[14])
	h0614, l0614 := bits.Mul64(a[6], b[14])
	h0514, l0514 := bits.Mul64(a[5], b[14])
	h0414, l0414 := bits.Mul64(a[4], b[14])
	h0314, l0314 := bits.Mul64(a[3], b[14])
	h0214, l0214 := bits.Mul64(a[2], b[14])
	_, l0114 := bits.Mul64(a[1], b[14])
	// h0014, l0014 := bits.Mul64(a[0], b[14])

	h1513, l1513 := bits.Mul64(a[15], b[13])
	h1413, l1413 := bits.Mul64(a[14], b[13])
	h1313, l1313 := bits.Mul64(a[13], b[13])
	h1213, l1213 := bits.Mul64(a[12], b[13])
	h1113, l1113 := bits.Mul64(a[11], b[13])
	h1013, l1013 := bits.Mul64(a[10], b[13])
	h0913, l0913 := bits.Mul64(a[9], b[13])
	h0813, l0813 := bits.Mul64(a[8], b[13])
	h0713, l0713 := bits.Mul64(a[7], b[13])
	h0613, l0613 := bits.Mul64(a[6], b[13])
	h0513, l0513 := bits.Mul64(a[5], b[13])
	h0413, l0413 := bits.Mul64(a[4], b[13])
	h0313, l0313 := bits.Mul64(a[3], b[13])
	_, l0213 := bits.Mul64(a[2], b[13])
	// h0113, l0113 := bits.Mul64(a[1], b[13])
	// h0013, l0013 := bits.Mul64(a[0], b[13])

	h1512, l1512 := bits.Mul64(a[15], b[12])
	h1412, l1412 := bits.Mul64(a[14], b[12])
	h1312, l1312 := bits.Mul64(a[13], b[12])
	h1212, l1212 := bits.Mul64(a[12], b[12])
	h1112, l1112 := bits.Mul64(a[11], b[12])
	h1012, l1012 := bits.Mul64(a[10], b[12])
	h0912, l0912 := bits.Mul64(a[9], b[12])
	h0812, l0812 := bits.Mul64(a[8], b[12])
	h0712, l0712 := bits.Mul64(a[7], b[12])
	h0612, l0612 := bits.Mul64(a[6], b[12])
	h0512, l0512 := bits.Mul64(a[5], b[12])
	h0412, l0412 := bits.Mul64(a[4], b[12])
	_, l0312 := bits.Mul64(a[3], b[12])
	// h0212, l0212 := bits.Mul64(a[2], b[12])
	// h0112, l0112 := bits.Mul64(a[1], b[12])
	// h0012, l0012 := bits.Mul64(a[0], b[12])

	h1511, l1511 := bits.Mul64(a[15], b[11])
	h1411, l1411 := bits.Mul64(a[14], b[11])
	h1311, l1311 := bits.Mul64(a[13], b[11])
	h1211, l1211 := bits.Mul64(a[12], b[11])
	h1111, l1111 := bits.Mul64(a[11], b[11])
	h1011, l1011 := bits.Mul64(a[10], b[11])
	h0911, l0911 := bits.Mul64(a[9], b[11])
	h0811, l0811 := bits.Mul64(a[8], b[11])
	h0711, l0711 := bits.Mul64(a[7], b[11])
	h0611, l0611 := bits.Mul64(a[6], b[11])
	h0511, l0511 := bits.Mul64(a[5], b[11])
	_, l0411 := bits.Mul64(a[4], b[11])
	// h0311, l0311 := bits.Mul64(a[3], b[11])
	// h0211, l0211 := bits.Mul64(a[2], b[11])
	// h0111, l0111 := bits.Mul64(a[1], b[11])
	// h0011, l0011 := bits.Mul64(a[0], b[11])

	h1510, l1510 := bits.Mul64(a[15], b[10])
	h1410, l1410 := bits.Mul64(a[14], b[10])
	h1310, l1310 := bits.Mul64(a[13], b[10])
	h1210, l1210 := bits.Mul64(a[12], b[10])
	h1110, l1110 := bits.Mul64(a[11], b[10])
	h1010, l1010 := bits.Mul64(a[10], b[10])
	h0910, l0910 := bits.Mul64(a[9], b[10])
	h0810, l0810 := bits.Mul64(a[8], b[10])
	h0710, l0710 := bits.Mul64(a[7], b[10])
	h0610, l0610 := bits.Mul64(a[6], b[10])
	_, l0510 := bits.Mul64(a[5], b[10])
	// h0410, l0410 := bits.Mul64(a[4], b[10])
	// h0310, l0310 := bits.Mul64(a[3], b[10])
	// h0210, l0210 := bits.Mul64(a[2], b[10])
	// h0110, l0110 := bits.Mul64(a[1], b[10])
	// h0010, l0010 := bits.Mul64(a[0], b[10])

	h1509, l1509 := bits.Mul64(a[15], b[9])
	h1409, l1409 := bits.Mul64(a[14], b[9])
	h1309, l1309 := bits.Mul64(a[13], b[9])
	h1209, l1209 := bits.Mul64(a[12], b[9])
	h1109, l1109 := bits.Mul64(a[11], b[9])
	h1009, l1009 := bits.Mul64(a[10], b[9])
	h0909, l0909 := bits.Mul64(a[9], b[9])
	h0809, l0809 := bits.Mul64(a[8], b[9])
	h0709, l0709 := bits.Mul64(a[7], b[9])
	_, l0609 := bits.Mul64(a[6], b[9])
	// h0509, l0509 := bits.Mul64(a[5], b[9])
	// h0409, l0409 := bits.Mul64(a[4], b[9])
	// h0309, l0309 := bits.Mul64(a[3], b[9])
	// h0209, l0209 := bits.Mul64(a[2], b[9])
	// h0109, l0109 := bits.Mul64(a[1], b[9])
	// h0009, l0009 := bits.Mul64(a[0], b[9])

	h1508, l1508 := bits.Mul64(a[15], b[8])
	h1408, l1408 := bits.Mul64(a[14], b[8])
	h1308, l1308 := bits.Mul64(a[13], b[8])
	h1208, l1208 := bits.Mul64(a[12], b[8])
	h1108, l1108 := bits.Mul64(a[11], b[8])
	h1008, l1008 := bits.Mul64(a[10], b[8])
	h0908, l0908 := bits.Mul64(a[9], b[8])
	h0808, l0808 := bits.Mul64(a[8], b[8])
	_, l0708 := bits.Mul64(a[7], b[8])
	// h0608, l0608 := bits.Mul64(a[6], b[8])
	// h0508, l0508 := bits.Mul64(a[5], b[8])
	// h0408, l0408 := bits.Mul64(a[4], b[8])
	// h0308, l0308 := bits.Mul64(a[3], b[8])
	// h0208, l0208 := bits.Mul64(a[2], b[8])
	// h0108, l0108 := bits.Mul64(a[1], b[8])
	// h0008, l0008 := bits.Mul64(a[0], b[8])

	h1507, l1507 := bits.Mul64(a[15], b[7])
	h1407, l1407 := bits.Mul64(a[14], b[7])
	h1307, l1307 := bits.Mul64(a[13], b[7])
	h1207, l1207 := bits.Mul64(a[12], b[7])
	h1107, l1107 := bits.Mul64(a[11], b[7])
	h1007, l1007 := bits.Mul64(a[10], b[7])
	h0907, l0907 := bits.Mul64(a[9], b[7])
	_, l0807 := bits.Mul64(a[8], b[7])
	// h0707, l0707 := bits.Mul64(a[7], b[7])
	// h0607, l0607 := bits.Mul64(a[6], b[7])
	// h0507, l0507 := bits.Mul64(a[5], b[7])
	// h0407, l0407 := bits.Mul64(a[4], b[7])
	// h0307, l0307 := bits.Mul64(a[3], b[7])
	// h0207, l0207 := bits.Mul64(a[2], b[7])
	// h0107, l0107 := bits.Mul64(a[1], b[7])
	// h0007, l0007 := bits.Mul64(a[0], b[7])

	h1506, l1506 := bits.Mul64(a[15], b[6])
	h1406, l1406 := bits.Mul64(a[14], b[6])
	h1306, l1306 := bits.Mul64(a[13], b[6])
	h1206, l1206 := bits.Mul64(a[12], b[6])
	h1106, l1106 := bits.Mul64(a[11], b[6])
	h1006, l1006 := bits.Mul64(a[10], b[6])
	_, l0906 := bits.Mul64(a[9], b[6])
	// h0806, l0806 := bits.Mul64(a[8], b[6])
	// h0706, l0706 := bits.Mul64(a[7], b[6])
	// h0606, l0606 := bits.Mul64(a[6], b[6])
	// h0506, l0506 := bits.Mul64(a[5], b[6])
	// h0406, l0406 := bits.Mul64(a[4], b[6])
	// h0306, l0306 := bits.Mul64(a[3], b[6])
	// h0206, l0206 := bits.Mul64(a[2], b[6])
	// h0106, l0106 := bits.Mul64(a[1], b[6])
	// h0006, l0006 := bits.Mul64(a[0], b[6])

	h1505, l1505 := bits.Mul64(a[15], b[5])
	h1405, l1405 := bits.Mul64(a[14], b[5])
	h1305, l1305 := bits.Mul64(a[13], b[5])
	h1205, l1205 := bits.Mul64(a[12], b[5])
	h1105, l1105 := bits.Mul64(a[11], b[5])
	_, l1005 := bits.Mul64(a[10], b[5])
	// h0905, l0905 := bits.Mul64(a[9], b[5])
	// h0805, l0805 := bits.Mul64(a[8], b[5])
	// h0705, l0705 := bits.Mul64(a[7], b[5])
	// h0605, l0605 := bits.Mul64(a[6], b[5])
	// h0505, l0505 := bits.Mul64(a[5], b[5])
	// h0405, l0405 := bits.Mul64(a[4], b[5])
	// h0305, l0305 := bits.Mul64(a[3], b[5])
	// h0205, l0205 := bits.Mul64(a[2], b[5])
	// h0105, l0105 := bits.Mul64(a[1], b[5])
	// h0005, l0005 := bits.Mul64(a[0], b[5])

	h1504, l1504 := bits.Mul64(a[15], b[4])
	h1404, l1404 := bits.Mul64(a[14], b[4])
	h1304, l1304 := bits.Mul64(a[13], b[4])
	h1204, l1204 := bits.Mul64(a[12], b[4])
	_, l1104 := bits.Mul64(a[11], b[4])
	// h1004, l1004 := bits.Mul64(a[10], b[4])
	// h0904, l0904 := bits.Mul64(a[9], b[4])
	// h0804, l0804 := bits.Mul64(a[8], b[4])
	// h0704, l0704 := bits.Mul64(a[7], b[4])
	// h0604, l0604 := bits.Mul64(a[6], b[4])
	// h0504, l0504 := bits.Mul64(a[5], b[4])
	// h0404, l0404 := bits.Mul64(a[4], b[4])
	// h0304, l0304 := bits.Mul64(a[3], b[4])
	// h0204, l0204 := bits.Mul64(a[2], b[4])
	// h0104, l0104 := bits.Mul64(a[1], b[4])
	// h0004, l0004 := bits.Mul64(a[0], b[4])

	h1503, l1503 := bits.Mul64(a[15], b[3])
	h1403, l1403 := bits.Mul64(a[14], b[3])
	h1303, l1303 := bits.Mul64(a[13], b[3])
	_, l1203 := bits.Mul64(a[12], b[3])
	// h1103, l1103 := bits.Mul64(a[11], b[3])
	// h1003, l1003 := bits.Mul64(a[10], b[3])
	// h0903, l0903 := bits.Mul64(a[9], b[3])
	// h0803, l0803 := bits.Mul64(a[8], b[3])
	// h0703, l0703 := bits.Mul64(a[7], b[3])
	// h0603, l0603 := bits.Mul64(a[6], b[3])
	// h0503, l0503 := bits.Mul64(a[5], b[3])
	// h0403, l0403 := bits.Mul64(a[4], b[3])
	// h0303, l0303 := bits.Mul64(a[3], b[3])
	// h0203, l0203 := bits.Mul64(a[2], b[3])
	// h0103, l0103 := bits.Mul64(a[1], b[3])
	// h0003, l0003 := bits.Mul64(a[0], b[3])

	h1502, l1502 := bits.Mul64(a[15], b[2])
	h1402, l1402 := bits.Mul64(a[14], b[2])
	_, l1302 := bits.Mul64(a[13], b[2])
	// h1202, l1202 := bits.Mul64(a[12], b[2])
	// h1102, l1102 := bits.Mul64(a[11], b[2])
	// h1002, l1002 := bits.Mul64(a[10], b[2])
	// h0902, l0902 := bits.Mul64(a[9], b[2])
	// h0802, l0802 := bits.Mul64(a[8], b[2])
	// h0702, l0702 := bits.Mul64(a[7], b[2])
	// h0602, l0602 := bits.Mul64(a[6], b[2])
	// h0502, l0502 := bits.Mul64(a[5], b[2])
	// h0402, l0402 := bits.Mul64(a[4], b[2])
	// h0302, l0302 := bits.Mul64(a[3], b[2])
	// h0202, l0202 := bits.Mul64(a[2], b[2])
	// h0102, l0102 := bits.Mul64(a[1], b[2])
	// h0002, l0002 := bits.Mul64(a[0], b[2])

	h1501, l1501 := bits.Mul64(a[15], b[1])
	_, l1401 := bits.Mul64(a[14], b[1])
	// h1301, l1301 := bits.Mul64(a[13], b[1])
	// h1201, l1201 := bits.Mul64(a[12], b[1])
	// h1101, l1101 := bits.Mul64(a[11], b[1])
	// h1001, l1001 := bits.Mul64(a[10], b[1])
	// h0901, l0901 := bits.Mul64(a[9], b[1])
	// h0801, l0801 := bits.Mul64(a[8], b[1])
	// h0701, l0701 := bits.Mul64(a[7], b[1])
	// h0601, l0601 := bits.Mul64(a[6], b[1])
	// h0501, l0501 := bits.Mul64(a[5], b[1])
	// h0401, l0401 := bits.Mul64(a[4], b[1])
	// h0301, l0301 := bits.Mul64(a[3], b[1])
	// h0201, l0201 := bits.Mul64(a[2], b[1])
	// h0101, l0101 := bits.Mul64(a[1], b[1])
	// h0001, l0001 := bits.Mul64(a[0], b[1])

	_, l1500 := bits.Mul64(a[15], b[0])
	// h1400, l1400 := bits.Mul64(a[14], b[0])
	// h1300, l1300 := bits.Mul64(a[13], b[0])
	// h1200, l1200 := bits.Mul64(a[12], b[0])
	// h1100, l1100 := bits.Mul64(a[11], b[0])
	// h1000, l1000 := bits.Mul64(a[10], b[0])
	// h0900, l0900 := bits.Mul64(a[9], b[0])
	// h0800, l0800 := bits.Mul64(a[8], b[0])
	// h0700, l0700 := bits.Mul64(a[7], b[0])
	// h0600, l0600 := bits.Mul64(a[6], b[0])
	// h0500, l0500 := bits.Mul64(a[5], b[0])
	// h0400, l0400 := bits.Mul64(a[4], b[0])
	// h0300, l0300 := bits.Mul64(a[3], b[0])
	// h0200, l0200 := bits.Mul64(a[2], b[0])
	// h0100, l0100 := bits.Mul64(a[1], b[0])
	// h0000, l0000 := bits.Mul64(a[0], b[0])

	var u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15, carry uint64
	// 1.
	u15 = l1515
	u14 = l1415
	u13 = l1315
	u12 = l1215
	u11 = l1115
	u10 = l1015
	u9 = l0915
	u8 = l0815
	u7 = l0715
	u6 = l0615
	u5 = l0515
	u4 = l0415
	u3 = l0315
	u2 = l0215
	u1 = l0115
	u0 = l0015
	u14, carry = bits.Add64(u14, h1515, 0)
	u13, carry = bits.Add64(u13, h1415, carry)
	u12, carry = bits.Add64(u12, h1315, carry)
	u11, carry = bits.Add64(u11, h1215, carry)
	u10, carry = bits.Add64(u10, h1115, carry)
	u9, carry = bits.Add64(u9, h1015, carry)
	u8, carry = bits.Add64(u8, h0915, carry)
	u7, carry = bits.Add64(u7, h0815, carry)
	u6, carry = bits.Add64(u6, h0715, carry)
	u5, carry = bits.Add64(u5, h0615, carry)
	u4, carry = bits.Add64(u4, h0515, carry)
	u3, carry = bits.Add64(u3, h0415, carry)
	u2, carry = bits.Add64(u2, h0315, carry)
	u1, carry = bits.Add64(u1, h0215, carry)
	u0, _ = bits.Add64(u0, h0115, carry)

	// 2.
	u14, carry = bits.Add64(u14, l1514, 0)
	u13, carry = bits.Add64(u13, l1414, carry)
	u12, carry = bits.Add64(u12, l1314, carry)
	u11, carry = bits.Add64(u11, l1214, carry)
	u10, carry = bits.Add64(u10, l1114, carry)
	u9, carry = bits.Add64(u9, l1014, carry)
	u8, carry = bits.Add64(u8, l0914, carry)
	u7, carry = bits.Add64(u7, l0814, carry)
	u6, carry = bits.Add64(u6, l0714, carry)
	u5, carry = bits.Add64(u5, l0614, carry)
	u4, carry = bits.Add64(u4, l0514, carry)
	u3, carry = bits.Add64(u3, l0414, carry)
	u2, carry = bits.Add64(u2, l0314, carry)
	u1, carry = bits.Add64(u1, l0214, carry)
	u0, _ = bits.Add64(u0, l0114, carry)
	u13, carry = bits.Add64(u13, h1514, 0)
	u12, carry = bits.Add64(u12, h1414, carry)
	u11, carry = bits.Add64(u11, h1314, carry)
	u10, carry = bits.Add64(u10, h1214, carry)
	u9, carry = bits.Add64(u9, h1114, carry)
	u8, carry = bits.Add64(u8, h1014, carry)
	u7, carry = bits.Add64(u7, h0914, carry)
	u6, carry = bits.Add64(u6, h0814, carry)
	u5, carry = bits.Add64(u5, h0714, carry)
	u4, carry = bits.Add64(u4, h0614, carry)
	u3, carry = bits.Add64(u3, h0514, carry)
	u2, carry = bits.Add64(u2, h0414, carry)
	u1, carry = bits.Add64(u1, h0314, carry)
	u0, _ = bits.Add64(u0, h0214, carry)

	// 3.
	u13, carry = bits.Add64(u13, l1513, 0)
	u12, carry = bits.Add64(u12, l1413, carry)
	u11, carry = bits.Add64(u11, l1313, carry)
	u10, carry = bits.Add64(u10, l1213, carry)
	u9, carry = bits.Add64(u9, l1113, carry)
	u8, carry = bits.Add64(u8, l1013, carry)
	u7, carry = bits.Add64(u7, l0913, carry)
	u6, carry = bits.Add64(u6, l0813, carry)
	u5, carry = bits.Add64(u5, l0713, carry)
	u4, carry = bits.Add64(u4, l0613, carry)
	u3, carry = bits.Add64(u3, l0513, carry)
	u2, carry = bits.Add64(u2, l0413, carry)
	u1, carry = bits.Add64(u1, l0313, carry)
	u0, _ = bits.Add64(u0, l0213, carry)
	u12, carry = bits.Add64(u12, h1513, 0)
	u11, carry = bits.Add64(u11, h1413, carry)
	u10, carry = bits.Add64(u10, h1313, carry)
	u9, carry = bits.Add64(u9, h1213, carry)
	u8, carry = bits.Add64(u8, h1113, carry)
	u7, carry = bits.Add64(u7, h1013, carry)
	u6, carry = bits.Add64(u6, h0913, carry)
	u5, carry = bits.Add64(u5, h0813, carry)
	u4, carry = bits.Add64(u4, h0713, carry)
	u3, carry = bits.Add64(u3, h0613, carry)
	u2, carry = bits.Add64(u2, h0513, carry)
	u1, carry = bits.Add64(u1, h0413, carry)
	u0, _ = bits.Add64(u0, h0313, carry)

	// 4.
	u12, carry = bits.Add64(u12, l1512, 0)
	u11, carry = bits.Add64(u11, l1412, carry)
	u10, carry = bits.Add64(u10, l1312, carry)
	u9, carry = bits.Add64(u9, l1212, carry)
	u8, carry = bits.Add64(u8, l1112, carry)
	u7, carry = bits.Add64(u7, l1012, carry)
	u6, carry = bits.Add64(u6, l0912, carry)
	u5, carry = bits.Add64(u5, l0812, carry)
	u4, carry = bits.Add64(u4, l0712, carry)
	u3, carry = bits.Add64(u3, l0612, carry)
	u2, carry = bits.Add64(u2, l0512, carry)
	u1, carry = bits.Add64(u1, l0412, carry)
	u0, _ = bits.Add64(u0, l0312, carry)
	u11, carry = bits.Add64(u11, h1512, 0)
	u10, carry = bits.Add64(u10, h1412, carry)
	u9, carry = bits.Add64(u9, h1312, carry)
	u8, carry = bits.Add64(u8, h1212, carry)
	u7, carry = bits.Add64(u7, h1112, carry)
	u6, carry = bits.Add64(u6, h1012, carry)
	u5, carry = bits.Add64(u5, h0912, carry)
	u4, carry = bits.Add64(u4, h0812, carry)
	u3, carry = bits.Add64(u3, h0712, carry)
	u2, carry = bits.Add64(u2, h0612, carry)
	u1, carry = bits.Add64(u1, h0512, carry)
	u0, _ = bits.Add64(u0, h0412, carry)

	// 5.
	u11, carry = bits.Add64(u11, l1511, 0)
	u10, carry = bits.Add64(u10, l1411, carry)
	u9, carry = bits.Add64(u9, l1311, carry)
	u8, carry = bits.Add64(u8, l1211, carry)
	u7, carry = bits.Add64(u7, l1111, carry)
	u6, carry = bits.Add64(u6, l1011, carry)
	u5, carry = bits.Add64(u5, l0911, carry)
	u4, carry = bits.Add64(u4, l0811, carry)
	u3, carry = bits.Add64(u3, l0711, carry)
	u2, carry = bits.Add64(u2, l0611, carry)
	u1, carry = bits.Add64(u1, l0511, carry)
	u0, _ = bits.Add64(u0, l0411, carry)
	u10, carry = bits.Add64(u10, h1511, 0)
	u9, carry = bits.Add64(u9, h1411, carry)
	u8, carry = bits.Add64(u8, h1311, carry)
	u7, carry = bits.Add64(u7, h1211, carry)
	u6, carry = bits.Add64(u6, h1111, carry)
	u5, carry = bits.Add64(u5, h1011, carry)
	u4, carry = bits.Add64(u4, h0911, carry)
	u3, carry = bits.Add64(u3, h0811, carry)
	u2, carry = bits.Add64(u2, h0711, carry)
	u1, carry = bits.Add64(u1, h0611, carry)
	u0, _ = bits.Add64(u0, h0511, carry)

	// 6.
	u10, carry = bits.Add64(u10, l1510, 0)
	u9, carry = bits.Add64(u9, l1410, carry)
	u8, carry = bits.Add64(u8, l1310, carry)
	u7, carry = bits.Add64(u7, l1210, carry)
	u6, carry = bits.Add64(u6, l1110, carry)
	u5, carry = bits.Add64(u5, l1010, carry)
	u4, carry = bits.Add64(u4, l0910, carry)
	u3, carry = bits.Add64(u3, l0810, carry)
	u2, carry = bits.Add64(u2, l0710, carry)
	u1, carry = bits.Add64(u1, l0610, carry)
	u0, _ = bits.Add64(u0, l0510, carry)
	u9, carry = bits.Add64(u9, h1510, 0)
	u8, carry = bits.Add64(u8, h1410, carry)
	u7, carry = bits.Add64(u7, h1310, carry)
	u6, carry = bits.Add64(u6, h1210, carry)
	u5, carry = bits.Add64(u5, h1110, carry)
	u4, carry = bits.Add64(u4, h1010, carry)
	u3, carry = bits.Add64(u3, h0910, carry)
	u2, carry = bits.Add64(u2, h0810, carry)
	u1, carry = bits.Add64(u1, h0710, carry)
	u0, _ = bits.Add64(u0, h0610, carry)

	// 7.
	u9, carry = bits.Add64(u9, l1509, 0)
	u8, carry = bits.Add64(u8, l1409, carry)
	u7, carry = bits.Add64(u7, l1309, carry)
	u6, carry = bits.Add64(u6, l1209, carry)
	u5, carry = bits.Add64(u5, l1109, carry)
	u4, carry = bits.Add64(u4, l1009, carry)
	u3, carry = bits.Add64(u3, l0909, carry)
	u2, carry = bits.Add64(u2, l0809, carry)
	u1, carry = bits.Add64(u1, l0709, carry)
	u0, _ = bits.Add64(u0, l0609, carry)
	u8, carry = bits.Add64(u8, h1509, 0)
	u7, carry = bits.Add64(u7, h1409, carry)
	u6, carry = bits.Add64(u6, h1309, carry)
	u5, carry = bits.Add64(u5, h1209, carry)
	u4, carry = bits.Add64(u4, h1109, carry)
	u3, carry = bits.Add64(u3, h1009, carry)
	u2, carry = bits.Add64(u2, h0909, carry)
	u1, carry = bits.Add64(u1, h0809, carry)
	u0, _ = bits.Add64(u0, h0709, carry)

	// 8.
	u8, carry = bits.Add64(u8, l1508, 0)
	u7, carry = bits.Add64(u7, l1408, carry)
	u6, carry = bits.Add64(u6, l1308, carry)
	u5, carry = bits.Add64(u5, l1208, carry)
	u4, carry = bits.Add64(u4, l1108, carry)
	u3, carry = bits.Add64(u3, l1008, carry)
	u2, carry = bits.Add64(u2, l0908, carry)
	u1, carry = bits.Add64(u1, l0808, carry)
	u0, _ = bits.Add64(u0, l0708, carry)
	u7, carry = bits.Add64(u7, h1508, 0)
	u6, carry = bits.Add64(u6, h1408, carry)
	u5, carry = bits.Add64(u5, h1308, carry)
	u4, carry = bits.Add64(u4, h1208, carry)
	u3, carry = bits.Add64(u3, h1108, carry)
	u2, carry = bits.Add64(u2, h1008, carry)
	u1, carry = bits.Add64(u1, h0908, carry)
	u0, _ = bits.Add64(u0, h0808, carry)

	// 9.
	u7, carry = bits.Add64(u7, l1507, 0)
	u6, carry = bits.Add64(u6, l1407, carry)
	u5, carry = bits.Add64(u5, l1307, carry)
	u4, carry = bits.Add64(u4, l1207, carry)
	u3, carry = bits.Add64(u3, l1107, carry)
	u2, carry = bits.Add64(u2, l1007, carry)
	u1, carry = bits.Add64(u1, l0907, carry)
	u0, _ = bits.Add64(u0, l0807, carry)
	u6, carry = bits.Add64(u6, h1507, 0)
	u5, carry = bits.Add64(u5, h1407, carry)
	u4, carry = bits.Add64(u4, h1307, carry)
	u3, carry = bits.Add64(u3, h1207, carry)
	u2, carry = bits.Add64(u2, h1107, carry)
	u1, carry = bits.Add64(u1, h1007, carry)
	u0, _ = bits.Add64(u0, h0907, carry)

	// 10.
	u6, carry = bits.Add64(u6, l1506, 0)
	u5, carry = bits.Add64(u5, l1406, carry)
	u4, carry = bits.Add64(u4, l1306, carry)
	u3, carry = bits.Add64(u3, l1206, carry)
	u2, carry = bits.Add64(u2, l1106, carry)
	u1, carry = bits.Add64(u1, l1006, carry)
	u0, _ = bits.Add64(u0, l0906, carry)
	u5, carry = bits.Add64(u5, h1506, 0)
	u4, carry = bits.Add64(u4, h1406, carry)
	u3, carry = bits.Add64(u3, h1306, carry)
	u2, carry = bits.Add64(u2, h1206, carry)
	u1, carry = bits.Add64(u1, h1106, carry)
	u0, _ = bits.Add64(u0, h1006, carry)

	// 11.
	u5, carry = bits.Add64(u5, l1505, 0)
	u4, carry = bits.Add64(u4, l1405, carry)
	u3, carry = bits.Add64(u3, l1305, carry)
	u2, carry = bits.Add64(u2, l1205, carry)
	u1, carry = bits.Add64(u1, l1105, carry)
	u0, _ = bits.Add64(u0, l1005, carry)
	u4, carry = bits.Add64(u4, h1505, 0)
	u3, carry = bits.Add64(u3, h1405, carry)
	u2, carry = bits.Add64(u2, h1305, carry)
	u1, carry = bits.Add64(u1, h1205, carry)
	u0, _ = bits.Add64(u0, h1105, carry)

	// 12.
	u4, carry = bits.Add64(u4, l1504, 0)
	u3, carry = bits.Add64(u3, l1404, carry)
	u2, carry = bits.Add64(u2, l1304, carry)
	u1, carry = bits.Add64(u1, l1204, carry)
	u0, _ = bits.Add64(u0, l1104, carry)
	u3, carry = bits.Add64(u3, h1504, 0)
	u2, carry = bits.Add64(u2, h1404, carry)
	u1, carry = bits.Add64(u1, h1304, carry)
	u0, _ = bits.Add64(u0, h1204, carry)

	// 13.
	u3, carry = bits.Add64(u3, l1503, 0)
	u2, carry = bits.Add64(u2, l1403, carry)
	u1, carry = bits.Add64(u1, l1303, carry)
	u0, _ = bits.Add64(u0, l1203, carry)
	u2, carry = bits.Add64(u2, h1503, 0)
	u1, carry = bits.Add64(u1, h1403, carry)
	u0, _ = bits.Add64(u0, h1303, carry)

	// 14.
	u2, carry = bits.Add64(u2, l1502, 0)
	u1, carry = bits.Add64(u1, l1402, carry)
	u0, _ = bits.Add64(u0, l1302, carry)
	u1, carry = bits.Add64(u1, h1502, 0)
	u0, _ = bits.Add64(u0, h1402, carry)

	// 15.
	u1, carry = bits.Add64(u1, l1501, 0)
	u0, _ = bits.Add64(u0, l1401, carry)
	u0, _ = bits.Add64(u0, h1501, 0)

	// 16.
	u0, _ = bits.Add64(u0, l1500, 0)

	return Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
}

// Div returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Div implements Euclidean division (unlike Go); see [Uint1024.DivMod] for more details.
func (a Uint1024) Div(b Uint1024) Uint1024 {
	q, _ := a.DivMod(b)
	return q
}

// Mod returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Mod implements Euclidean division (unlike Go); see [Uint1024.DivMod] for more details.
func (a Uint1024) Mod(b Uint1024) Uint1024 {
	_, r := a.DivMod(b)
	return r
}

// DivMod returns the quotient and remainder of a/b.
// DivMod implements Euclidean division and modulus (unlike Go):
//
//	q = a div b  such that
//	m = a - b*q  with 0 <= m < |b|
//
// (See Raymond T. Boute, “The Euclidean definition of the functions
// div and mod”. ACM Transactions on Programming Languages and
// Systems (TOPLAS), 14(2):127-144, New York, NY, USA, 4/1992.
// ACM press.)
// See [Uint1024.QuoRem] for T-division and modulus (like Go).
func (a Uint1024) DivMod(b Uint1024) (Uint1024, Uint1024) {
	if b[0] == 0 && b[1] == 0 && b[2] == 0 && b[3] == 0 && b[4] == 0 && b[5] == 0 && b[6] == 0 && b[7] == 0 {
		// optimize for uint512 / uint256
		q0, r0 := Uint512{a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7]}.DivMod(Uint512{b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15]})
		q1, r1 := div512(r0, Uint512{a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15]}, Uint512{b[8], b[9], b[10], b[11], b[12], b[13], b[14], b[15]})
		return Uint1024{
				q0[0], q0[1], q0[2], q0[3], q0[4], q0[5], q0[6], q0[7],
				q1[0], q1[1], q1[2], q1[3], q1[4], q1[5], q1[6], q1[7],
			}, Uint1024{
				0, 0, 0, 0, 0, 0, 0, 0,
				r1[0], r1[1], r1[2], r1[3], r1[4], r1[5], r1[6], r1[7],
			}
	}

	n := uint(Uint512{b[0], b[1], b[2], b[3], b[4], b[5], b[6], b[7]}.LeadingZeros())
	x := a.Rsh(1)
	y := b.Lsh(n)
	q, _ := div512(Uint512{x[0], x[1], x[2], x[3], x[4], x[5], x[6], x[7]}, Uint512{x[8], x[9], x[10], x[11], x[12], x[13], x[14], x[15]}, Uint512{y[0], y[1], y[2], y[3], y[4], y[5], y[6], y[7]})
	q = q.Rsh(511 - n)
	if q.Sign() > 0 {
		q = q.Sub(Uint512{0, 0, 0, 0, 0, 0, 0, 1})
	}

	u := b.Mul(Uint1024{0, 0, 0, 0, 0, 0, 0, 0, q[0], q[1], q[2], q[3], q[4], q[5], q[6], q[7]})
	r := a.Sub(u)
	if r.Cmp(b) >= 0 {
		q = q.Add(Uint512{0, 0, 0, 0, 0, 0, 0, 1})
		r = r.Sub(b)
	}

	return Uint1024{0, 0, 0, 0, 0, 0, 0, 0, q[0], q[1], q[2], q[3], q[4], q[5], q[6], q[7]}, r
}

// 512-bit of version of bits.Div64.
// https://github.com/golang/go/blob/c893e1cf821b06aa0602f7944ce52f0eb28fd7b5/src/math/bits/bits.go#L514-L568
func div512(hi, lo, y Uint512) (quo, rem Uint512) {
	if y.IsZero() {
		panic("division by zero")
	}
	if y.Cmp(hi) <= 0 {
		panic("division overflow")
	}

	// If high part is zero, we can directly return the results.
	if hi.IsZero() {
		return lo.DivMod(y)
	}

	s := uint(y.LeadingZeros())
	y = y.Lsh(s)

	two256 := Uint512{0, 0, 0, 1, 0, 0, 0, 0}
	yn1 := Uint512{0, 0, 0, 0, y[0], y[1], y[2], y[3]}
	yn0 := Uint512{0, 0, 0, 0, y[4], y[5], y[6], y[7]}
	un256 := hi.Lsh(s).Or(lo.Rsh(512 - s))
	un10 := lo.Lsh(s)
	un1 := Uint512{0, 0, 0, 0, un10[0], un10[1], un10[2], un10[3]}
	un0 := Uint512{0, 0, 0, 0, un10[4], un10[5], un10[6], un10[7]}
	q1 := un256.Div(yn1)
	rhat := un256.Sub(q1.Mul(yn1))

	for q1.Cmp(two256) >= 0 || q1.Mul(yn0).Cmp(two256.Mul(rhat).Add(un1)) > 0 {
		q1 = q1.Sub(Uint512{0, 0, 0, 0, 0, 0, 0, 1})
		rhat = rhat.Add(yn1)
		if rhat.Cmp(two256) >= 0 {
			break
		}
	}

	un21 := un256.Mul(two256).Add(un1).Sub(q1.Mul(y))
	q0 := un21.Div(yn1)
	rhat = un21.Sub(q0.Mul(yn1))

	for q0.Cmp(two256) >= 0 || q0.Mul(yn0).Cmp(two256.Mul(rhat).Add(un0)) > 0 {
		q0 = q0.Sub(Uint512{0, 0, 0, 0, 0, 0, 0, 1})
		rhat = rhat.Add(yn1)
		if rhat.Cmp(two256) >= 0 {
			break
		}
	}

	return q1.Mul(two256).Add(q0), un21.Mul(two256).Add(un0).Sub(q0.Mul(y)).Rsh(s)
}

// Quo returns the quotient a/b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Quo implements T-division (like Go); see [Uint1024.QuoRem] for more details.
// For unsigned integers T‑division and Euclidean division are identical,
// therefore Quo simply forwards to Div.
func (a Uint1024) Quo(b Uint1024) Uint1024 {
	return a.Div(b)
}

// Rem returns the remainder a%b for b != 0.
// If b == 0, a division-by-zero run-time panic occurs.
// Rem implements T-division (like Go); see [Uint1024.QuoRem] for more details.
// For unsigned integers T‑division and Euclidean division are identical,
// therefore Rem simply forwards to Mod.
func (a Uint1024) Rem(b Uint1024) Uint1024 {
	return a.Mod(b)
}

// QuoRem returns the quotient and remainder of a/b.
// QuoRem implements T-division and modulus (like Go):
//
//	q = a/b      with the result truncated to zero
//	r = a - b*q
//
// (See Daan Leijen, “Division and Modulus for Computer Scientists”.)
// See [Uint1024.DivMod] for Euclidean division and modulus (unlike Go).
// For unsigned integers T‑division and Euclidean division are identical,
// therefore QuoRem simply forwards to DivMod.
func (a Uint1024) QuoRem(b Uint1024) (Uint1024, Uint1024) {
	return a.DivMod(b)
}

// And returns the bitwise AND of a and b.
func (a Uint1024) And(b Uint1024) Uint1024 {
	return Uint1024{
		a[0] & b[0],
		a[1] & b[1],
		a[2] & b[2],
		a[3] & b[3],
		a[4] & b[4],
		a[5] & b[5],
		a[6] & b[6],
		a[7] & b[7],
		a[8] & b[8],
		a[9] & b[9],
		a[10] & b[10],
		a[11] & b[11],
		a[12] & b[12],
		a[13] & b[13],
		a[14] & b[14],
		a[15] & b[15],
	}
}

// AndNot returns the bitwise AND NOT of a and b.
func (a Uint1024) AndNot(b Uint1024) Uint1024 {
	return Uint1024{
		a[0] &^ b[0],
		a[1] &^ b[1],
		a[2] &^ b[2],
		a[3] &^ b[3],
		a[4] &^ b[4],
		a[5] &^ b[5],
		a[6] &^ b[6],
		a[7] &^ b[7],
		a[8] &^ b[8],
		a[9] &^ b[9],
		a[10] &^ b[10],
		a[11] &^ b[11],
		a[12] &^ b[12],
		a[13] &^ b[13],
		a[14] &^ b[14],
		a[15] &^ b[15],
	}
}

// Or returns the bitwise OR of a and b.
func (a Uint1024) Or(b Uint1024) Uint1024 {
	return Uint1024{
		a[0] | b[0],
		a[1] | b[1],
		a[2] | b[2],
		a[3] | b[3],
		a[4] | b[4],
		a[5] | b[5],
		a[6] | b[6],
		a[7] | b[7],
		a[8] | b[8],
		a[9] | b[9],
		a[10] | b[10],
		a[11] | b[11],
		a[12] | b[12],
		a[13] | b[13],
		a[14] | b[14],
		a[15] | b[15],
	}
}

// Xor returns the bitwise XOR of a and b.
func (a Uint1024) Xor(b Uint1024) Uint1024 {
	return Uint1024{
		a[0] ^ b[0],
		a[1] ^ b[1],
		a[2] ^ b[2],
		a[3] ^ b[3],
		a[4] ^ b[4],
		a[5] ^ b[5],
		a[6] ^ b[6],
		a[7] ^ b[7],
		a[8] ^ b[8],
		a[9] ^ b[9],
		a[10] ^ b[10],
		a[11] ^ b[11],
		a[12] ^ b[12],
		a[13] ^ b[13],
		a[14] ^ b[14],
		a[15] ^ b[15],
	}
}

// Not returns the bitwise NOT of a.
func (a Uint1024) Not() Uint1024 {
	return Uint1024{
		^a[0],
		^a[1],
		^a[2],
		^a[3],
		^a[4],
		^a[5],
		^a[6],
		^a[7],
		^a[8],
		^a[9],
		^a[10],
		^a[11],
		^a[12],
		^a[13],
		^a[14],
		^a[15],
	}
}

// Lsh returns the logical left shift a<<i.
//
// This function's execution time does not depend on the inputs.
func (a Uint1024) Lsh(i uint) Uint1024 {
	// This operation may overflow, but it's okay because when it overflows,
	// the result is always greater than or equal to 64.
	// And shifts of 64 bits or more always result in 0, so they don't affect the final result.
	n1 := uint(i - 64)
	n2 := uint(64 - i)
	n3 := uint(i - 128)
	n4 := uint(128 - i)
	n5 := uint(i - 192)
	n6 := uint(192 - i)
	n7 := uint(i - 256)
	n8 := uint(256 - i)
	n9 := uint(i - 320)
	n10 := uint(320 - i)
	n11 := uint(i - 384)
	n12 := uint(384 - i)
	n13 := uint(i - 448)
	n14 := uint(448 - i)
	n15 := uint(i - 512)
	n16 := uint(512 - i)
	n17 := uint(i - 576)
	n18 := uint(576 - i)
	n19 := uint(i - 640)
	n20 := uint(640 - i)
	n21 := uint(i - 704)
	n22 := uint(704 - i)
	n23 := uint(i - 768)
	n24 := uint(768 - i)
	n25 := uint(i - 832)
	n26 := uint(832 - i)
	n27 := uint(i - 896)
	n28 := uint(896 - i)
	n29 := uint(i - 960)
	n30 := uint(960 - i)

	return Uint1024{
		a[0]<<i | a[1]<<n1 | a[1]>>n2 | a[2]<<n3 | a[2]>>n4 | a[3]<<n5 | a[3]>>n6 | a[4]<<n7 | a[4]>>n8 |
			a[5]<<n9 | a[5]>>n10 | a[6]<<n11 | a[6]>>n12 | a[7]<<n13 | a[7]>>n14 | a[8]<<n15 | a[8]>>n16 |
			a[9]<<n17 | a[9]>>n18 | a[10]<<n19 | a[10]>>n20 | a[11]<<n21 | a[11]>>n22 | a[12]<<n23 | a[12]>>n24 |
			a[13]<<n25 | a[13]>>n26 | a[14]<<n27 | a[14]>>n28 | a[15]<<n29 | a[15]>>n30,
		a[1]<<i | a[2]<<n1 | a[2]>>n2 | a[3]<<n3 | a[3]>>n4 | a[4]<<n5 | a[4]>>n6 | a[5]<<n7 | a[5]>>n8 |
			a[6]<<n9 | a[6]>>n10 | a[7]<<n11 | a[7]>>n12 | a[8]<<n13 | a[8]>>n14 | a[9]<<n15 | a[9]>>n16 |
			a[10]<<n17 | a[10]>>n18 | a[11]<<n19 | a[11]>>n20 | a[12]<<n21 | a[12]>>n22 | a[13]<<n23 | a[13]>>n24 |
			a[14]<<n25 | a[14]>>n26 | a[15]<<n27 | a[15]>>n28,
		a[2]<<i | a[3]<<n1 | a[3]>>n2 | a[4]<<n3 | a[4]>>n4 | a[5]<<n5 | a[5]>>n6 | a[6]<<n7 | a[6]>>n8 |
			a[7]<<n9 | a[7]>>n10 | a[8]<<n11 | a[8]>>n12 | a[9]<<n13 | a[9]>>n14 | a[10]<<n15 | a[10]>>n16 |
			a[11]<<n17 | a[11]>>n18 | a[12]<<n19 | a[12]>>n20 | a[13]<<n21 | a[13]>>n22 | a[14]<<n23 | a[14]>>n24 |
			a[15]<<n25 | a[15]>>n26,
		a[3]<<i | a[4]<<n1 | a[4]>>n2 | a[5]<<n3 | a[5]>>n4 | a[6]<<n5 | a[6]>>n6 | a[7]<<n7 | a[7]>>n8 |
			a[8]<<n9 | a[8]>>n10 | a[9]<<n11 | a[9]>>n12 | a[10]<<n13 | a[10]>>n14 | a[11]<<n15 | a[11]>>n16 |
			a[12]<<n17 | a[12]>>n18 | a[13]<<n19 | a[13]>>n20 | a[14]<<n21 | a[14]>>n22 | a[15]<<n23 | a[15]>>n24,
		a[4]<<i | a[5]<<n1 | a[5]>>n2 | a[6]<<n3 | a[6]>>n4 | a[7]<<n5 | a[7]>>n6 | a[8]<<n7 | a[8]>>n8 |
			a[9]<<n9 | a[9]>>n10 | a[10]<<n11 | a[10]>>n12 | a[11]<<n13 | a[11]>>n14 | a[12]<<n15 | a[12]>>n16 |
			a[13]<<n17 | a[13]>>n18 | a[14]<<n19 | a[14]>>n20 | a[15]<<n21 | a[15]>>n22,
		a[5]<<i | a[6]<<n1 | a[6]>>n2 | a[7]<<n3 | a[7]>>n4 | a[8]<<n5 | a[8]>>n6 | a[9]<<n7 | a[9]>>n8 |
			a[10]<<n9 | a[10]>>n10 | a[11]<<n11 | a[11]>>n12 | a[12]<<n13 | a[12]>>n14 | a[13]<<n15 | a[13]>>n16 |
			a[14]<<n17 | a[14]>>n18 | a[15]<<n19 | a[15]>>n20,
		a[6]<<i | a[7]<<n1 | a[7]>>n2 | a[8]<<n3 | a[8]>>n4 | a[9]<<n5 | a[9]>>n6 | a[10]<<n7 | a[10]>>n8 |
			a[11]<<n9 | a[11]>>n10 | a[12]<<n11 | a[12]>>n12 | a[13]<<n13 | a[13]>>n14 | a[14]<<n15 | a[14]>>n16 |
			a[15]<<n17 | a[15]>>n18,
		a[7]<<i | a[8]<<n1 | a[8]>>n2 | a[9]<<n3 | a[9]>>n4 | a[10]<<n5 | a[10]>>n6 | a[11]<<n7 | a[11]>>n8 |
			a[12]<<n9 | a[12]>>n10 | a[13]<<n11 | a[13]>>n12 | a[14]<<n13 | a[14]>>n14 | a[15]<<n15 | a[15]>>n16,
		a[8]<<i | a[9]<<n1 | a[9]>>n2 | a[10]<<n3 | a[10]>>n4 | a[11]<<n5 | a[11]>>n6 | a[12]<<n7 | a[12]>>n8 |
			a[13]<<n9 | a[13]>>n10 | a[14]<<n11 | a[14]>>n12 | a[15]<<n13 | a[15]>>n14,
		a[9]<<i | a[10]<<n1 | a[10]>>n2 | a[11]<<n3 | a[11]>>n4 | a[12]<<n5 | a[12]>>n6 | a[13]<<n7 | a[13]>>n8 |
			a[14]<<n9 | a[14]>>n10 | a[15]<<n11 | a[15]>>n12,
		a[10]<<i | a[11]<<n1 | a[11]>>n2 | a[12]<<n3 | a[12]>>n4 | a[13]<<n5 | a[13]>>n6 | a[14]<<n7 | a[14]>>n8 |
			a[15]<<n9 | a[15]>>n10,
		a[11]<<i | a[12]<<n1 | a[12]>>n2 | a[13]<<n3 | a[13]>>n4 | a[14]<<n5 | a[14]>>n6 | a[15]<<n7 | a[15]>>n8,
		a[12]<<i | a[13]<<n1 | a[13]>>n2 | a[14]<<n3 | a[14]>>n4 | a[15]<<n5 | a[15]>>n6,
		a[13]<<i | a[14]<<n1 | a[14]>>n2 | a[15]<<n3 | a[15]>>n4,
		a[14]<<i | a[15]<<n1 | a[15]>>n2,
		a[15] << i,
	}
}

// Rsh returns the logical right shift a>>i.
//
// This function's execution time does not depend on the inputs.
func (a Uint1024) Rsh(i uint) Uint1024 {
	// This operation may overflow, but it's okay because when it overflows,
	// the result is always greater than or equal to 64.
	// And shifts of 64 bits or more always result in 0, so they don't affect the final result.
	n1 := uint(i - 64)
	n2 := uint(64 - i)
	n3 := uint(i - 128)
	n4 := uint(128 - i)
	n5 := uint(i - 192)
	n6 := uint(192 - i)
	n7 := uint(i - 256)
	n8 := uint(256 - i)
	n9 := uint(i - 320)
	n10 := uint(320 - i)
	n11 := uint(i - 384)
	n12 := uint(384 - i)
	n13 := uint(i - 448)
	n14 := uint(448 - i)
	n15 := uint(i - 512)
	n16 := uint(512 - i)
	n17 := uint(i - 576)
	n18 := uint(576 - i)
	n19 := uint(i - 640)
	n20 := uint(640 - i)
	n21 := uint(i - 704)
	n22 := uint(704 - i)
	n23 := uint(i - 768)
	n24 := uint(768 - i)
	n25 := uint(i - 832)
	n26 := uint(832 - i)
	n27 := uint(i - 896)
	n28 := uint(896 - i)
	n29 := uint(i - 960)
	n30 := uint(960 - i)

	return Uint1024{
		a[0] >> i,
		a[1]>>i | a[0]>>n1 | a[0]<<n2,
		a[2]>>i | a[1]>>n1 | a[1]<<n2 | a[0]>>n3 | a[0]<<n4,
		a[3]>>i | a[2]>>n1 | a[2]<<n2 | a[1]>>n3 | a[1]<<n4 | a[0]>>n5 | a[0]<<n6,
		a[4]>>i | a[3]>>n1 | a[3]<<n2 | a[2]>>n3 | a[2]<<n4 | a[1]>>n5 | a[1]<<n6 | a[0]>>n7 | a[0]<<n8,
		a[5]>>i | a[4]>>n1 | a[4]<<n2 | a[3]>>n3 | a[3]<<n4 | a[2]>>n5 | a[2]<<n6 | a[1]>>n7 | a[1]<<n8 |
			a[0]>>n9 | a[0]<<n10,
		a[6]>>i | a[5]>>n1 | a[5]<<n2 | a[4]>>n3 | a[4]<<n4 | a[3]>>n5 | a[3]<<n6 | a[2]>>n7 | a[2]<<n8 |
			a[1]>>n9 | a[1]<<n10 | a[0]>>n11 | a[0]<<n12,
		a[7]>>i | a[6]>>n1 | a[6]<<n2 | a[5]>>n3 | a[5]<<n4 | a[4]>>n5 | a[4]<<n6 | a[3]>>n7 | a[3]<<n8 |
			a[2]>>n9 | a[2]<<n10 | a[1]>>n11 | a[1]<<n12 | a[0]>>n13 | a[0]<<n14,
		a[8]>>i | a[7]>>n1 | a[7]<<n2 | a[6]>>n3 | a[6]<<n4 | a[5]>>n5 | a[5]<<n6 | a[4]>>n7 | a[4]<<n8 |
			a[3]>>n9 | a[3]<<n10 | a[2]>>n11 | a[2]<<n12 | a[1]>>n13 | a[1]<<n14 | a[0]>>n15 | a[0]<<n16,
		a[9]>>i | a[8]>>n1 | a[8]<<n2 | a[7]>>n3 | a[7]<<n4 | a[6]>>n5 | a[6]<<n6 | a[5]>>n7 | a[5]<<n8 |
			a[4]>>n9 | a[4]<<n10 | a[3]>>n11 | a[3]<<n12 | a[2]>>n13 | a[2]<<n14 | a[1]>>n15 | a[1]<<n16 |
			a[0]>>n17 | a[0]<<n18,
		a[10]>>i | a[9]>>n1 | a[9]<<n2 | a[8]>>n3 | a[8]<<n4 | a[7]>>n5 | a[7]<<n6 | a[6]>>n7 | a[6]<<n8 |
			a[5]>>n9 | a[5]<<n10 | a[4]>>n11 | a[4]<<n12 | a[3]>>n13 | a[3]<<n14 | a[2]>>n15 | a[2]<<n16 |
			a[1]>>n17 | a[1]<<n18 | a[0]>>n19 | a[0]<<n20,
		a[11]>>i | a[10]>>n1 | a[10]<<n2 | a[9]>>n3 | a[9]<<n4 | a[8]>>n5 | a[8]<<n6 | a[7]>>n7 | a[7]<<n8 |
			a[6]>>n9 | a[6]<<n10 | a[5]>>n11 | a[5]<<n12 | a[4]>>n13 | a[4]<<n14 | a[3]>>n15 | a[3]<<n16 |
			a[2]>>n17 | a[2]<<n18 | a[1]>>n19 | a[1]<<n20 | a[0]>>n21 | a[0]<<n22,
		a[12]>>i | a[11]>>n1 | a[11]<<n2 | a[10]>>n3 | a[10]<<n4 | a[9]>>n5 | a[9]<<n6 | a[8]>>n7 | a[8]<<n8 |
			a[7]>>n9 | a[7]<<n10 | a[6]>>n11 | a[6]<<n12 | a[5]>>n13 | a[5]<<n14 | a[4]>>n15 | a[4]<<n16 |
			a[3]>>n17 | a[3]<<n18 | a[2]>>n19 | a[2]<<n20 | a[1]>>n21 | a[1]<<n22 | a[0]>>n23 | a[0]<<n24,
		a[13]>>i | a[12]>>n1 | a[12]<<n2 | a[11]>>n3 | a[11]<<n4 | a[10]>>n5 | a[10]<<n6 | a[9]>>n7 | a[9]<<n8 |
			a[8]>>n9 | a[8]<<n10 | a[7]>>n11 | a[7]<<n12 | a[6]>>n13 | a[6]<<n14 | a[5]>>n15 | a[5]<<n16 |
			a[4]>>n17 | a[4]<<n18 | a[3]>>n19 | a[3]<<n20 | a[2]>>n21 | a[2]<<n22 | a[1]>>n23 | a[1]<<n24 |
			a[0]>>n25 | a[0]<<n26,
		a[14]>>i | a[13]>>n1 | a[13]<<n2 | a[12]>>n3 | a[12]<<n4 | a[11]>>n5 | a[11]<<n6 | a[10]>>n7 | a[10]<<n8 |
			a[9]>>n9 | a[9]<<n10 | a[8]>>n11 | a[8]<<n12 | a[7]>>n13 | a[7]<<n14 | a[6]>>n15 | a[6]<<n16 |
			a[5]>>n17 | a[5]<<n18 | a[4]>>n19 | a[4]<<n20 | a[3]>>n21 | a[3]<<n22 | a[2]>>n23 | a[2]<<n24 |
			a[1]>>n25 | a[1]<<n26 | a[0]>>n27 | a[0]<<n28,
		a[15]>>i | a[14]>>n1 | a[14]<<n2 | a[13]>>n3 | a[13]<<n4 | a[12]>>n5 | a[12]<<n6 | a[11]>>n7 | a[11]<<n8 |
			a[10]>>n9 | a[10]<<n10 | a[9]>>n11 | a[9]<<n12 | a[8]>>n13 | a[8]<<n14 | a[7]>>n15 | a[7]<<n16 |
			a[6]>>n17 | a[6]<<n18 | a[5]>>n19 | a[5]<<n20 | a[4]>>n21 | a[4]<<n22 | a[3]>>n23 | a[3]<<n24 |
			a[2]>>n25 | a[2]<<n26 | a[1]>>n27 | a[1]<<n28 | a[0]>>n29 | a[0]<<n30,
	}
}

// LeadingZeros returns the number of leading zero bits in a; the result is 1024 for a == 0.
func (a Uint1024) LeadingZeros() int {
	if a[0] != 0 {
		return bits.LeadingZeros64(a[0])
	}
	if a[1] != 0 {
		return bits.LeadingZeros64(a[1]) + 64
	}
	if a[2] != 0 {
		return bits.LeadingZeros64(a[2]) + 128
	}
	if a[3] != 0 {
		return bits.LeadingZeros64(a[3]) + 192
	}
	if a[4] != 0 {
		return bits.LeadingZeros64(a[4]) + 256
	}
	if a[5] != 0 {
		return bits.LeadingZeros64(a[5]) + 320
	}
	if a[6] != 0 {
		return bits.LeadingZeros64(a[6]) + 384
	}
	if a[7] != 0 {
		return bits.LeadingZeros64(a[7]) + 448
	}
	if a[8] != 0 {
		return bits.LeadingZeros64(a[8]) + 512
	}
	if a[9] != 0 {
		return bits.LeadingZeros64(a[9]) + 576
	}
	if a[10] != 0 {
		return bits.LeadingZeros64(a[10]) + 640
	}
	if a[11] != 0 {
		return bits.LeadingZeros64(a[11]) + 704
	}
	if a[12] != 0 {
		return bits.LeadingZeros64(a[12]) + 768
	}
	if a[13] != 0 {
		return bits.LeadingZeros64(a[13]) + 832
	}
	if a[14] != 0 {
		return bits.LeadingZeros64(a[14]) + 896
	}
	return bits.LeadingZeros64(a[15]) + 960
}

// TrailingZeros returns the number of trailing zero bits in a; the result is 1024 for a == 0.
func (a Uint1024) TrailingZeros() int {
	if a[15] != 0 {
		return bits.TrailingZeros64(a[15])
	}
	if a[14] != 0 {
		return bits.TrailingZeros64(a[14]) + 64
	}
	if a[13] != 0 {
		return bits.TrailingZeros64(a[13]) + 128
	}
	if a[12] != 0 {
		return bits.TrailingZeros64(a[12]) + 192
	}
	if a[11] != 0 {
		return bits.TrailingZeros64(a[11]) + 256
	}
	if a[10] != 0 {
		return bits.TrailingZeros64(a[10]) + 320
	}
	if a[9] != 0 {
		return bits.TrailingZeros64(a[9]) + 384
	}
	if a[8] != 0 {
		return bits.TrailingZeros64(a[8]) + 448
	}
	if a[7] != 0 {
		return bits.TrailingZeros64(a[7]) + 512
	}
	if a[6] != 0 {
		return bits.TrailingZeros64(a[6]) + 576
	}
	if a[5] != 0 {
		return bits.TrailingZeros64(a[5]) + 640
	}
	if a[4] != 0 {
		return bits.TrailingZeros64(a[4]) + 704
	}
	if a[3] != 0 {
		return bits.TrailingZeros64(a[3]) + 768
	}
	if a[2] != 0 {
		return bits.TrailingZeros64(a[2]) + 832
	}
	if a[1] != 0 {
		return bits.TrailingZeros64(a[1]) + 896
	}
	return bits.TrailingZeros64(a[0]) + 960
}

// BitLen returns the number of bits required to represent a in binary; the result is 0 for a == 0.
func (a Uint1024) BitLen() int {
	if a[0] != 0 {
		return bits.Len64(a[0]) + 960
	}
	if a[1] != 0 {
		return bits.Len64(a[1]) + 896
	}
	if a[2] != 0 {
		return bits.Len64(a[2]) + 832
	}
	if a[3] != 0 {
		return bits.Len64(a[3]) + 768
	}
	if a[4] != 0 {
		return bits.Len64(a[4]) + 704
	}
	if a[5] != 0 {
		return bits.Len64(a[5]) + 640
	}
	if a[6] != 0 {
		return bits.Len64(a[6]) + 576
	}
	if a[7] != 0 {
		return bits.Len64(a[7]) + 512
	}
	if a[8] != 0 {
		return bits.Len64(a[8]) + 448
	}
	if a[9] != 0 {
		return bits.Len64(a[9]) + 384
	}
	if a[10] != 0 {
		return bits.Len64(a[10]) + 320
	}
	if a[11] != 0 {
		return bits.Len64(a[11]) + 256
	}
	if a[12] != 0 {
		return bits.Len64(a[12]) + 192
	}
	if a[13] != 0 {
		return bits.Len64(a[13]) + 128
	}
	if a[14] != 0 {
		return bits.Len64(a[14]) + 64
	}
	return bits.Len64(a[15])
}

// Sign returns the sign of a.
// It returns 1 if a > 0, and 0 if a == 0.
// It does not return -1 because Uint128 is unsigned.
func (a Uint1024) Sign() int {
	var zero Uint1024
	if a == zero {
		return 0
	}
	return 1
}

// Neg returns the negation of a.
//
// This function's execution time does not depend on the inputs.
func (a Uint1024) Neg() Uint1024 {
	u15, borrow := bits.Sub64(0, a[15], 0)
	u14, borrow := bits.Sub64(0, a[14], borrow)
	u13, borrow := bits.Sub64(0, a[13], borrow)
	u12, borrow := bits.Sub64(0, a[12], borrow)
	u11, borrow := bits.Sub64(0, a[11], borrow)
	u10, borrow := bits.Sub64(0, a[10], borrow)
	u9, borrow := bits.Sub64(0, a[9], borrow)
	u8, borrow := bits.Sub64(0, a[8], borrow)
	u7, borrow := bits.Sub64(0, a[7], borrow)
	u6, borrow := bits.Sub64(0, a[6], borrow)
	u5, borrow := bits.Sub64(0, a[5], borrow)
	u4, borrow := bits.Sub64(0, a[4], borrow)
	u3, borrow := bits.Sub64(0, a[3], borrow)
	u2, borrow := bits.Sub64(0, a[2], borrow)
	u1, borrow := bits.Sub64(0, a[1], borrow)
	u0, _ := bits.Sub64(0, a[0], borrow)
	return Uint1024{u0, u1, u2, u3, u4, u5, u6, u7, u8, u9, u10, u11, u12, u13, u14, u15}
}

// Cmp returns the comparison result of a and b.
// It returns -1 if a < b, 0 if a == b, and 1 if a > b.
func (a Uint1024) Cmp(b Uint1024) int {
	if ret := cmp.Compare(a[0], b[0]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[1], b[1]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[2], b[2]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[3], b[3]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[4], b[4]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[5], b[5]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[6], b[6]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[7], b[7]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[8], b[8]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[9], b[9]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[10], b[10]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[11], b[11]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[12], b[12]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[13], b[13]); ret != 0 {
		return ret
	}
	if ret := cmp.Compare(a[14], b[14]); ret != 0 {
		return ret
	}
	return cmp.Compare(a[15], b[15])
}

// Text returns the string representation of a in the given base.
// Base must be between 2 and 62, inclusive.
// The result uses the lower-case letters 'a' to 'z' for digit values 10 to 35,
// and the upper-case letters 'A' to 'Z' for digit values 36 to 61. No prefix (such as "0x") is added to the string.
func (a Uint1024) Text(base int) string {
	_, s := formatBits1024(nil, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15], base, false, false)
	return s
}

// Append appends the string representation of a, as generated by a.Text(base), to buf and returns the extended buffer.
func (a Uint1024) Append(dst []byte, base int) []byte {
	d, _ := formatBits1024(dst, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15], base, false, true)
	return d
}

// AppendText implements the [encoding.TextAppender] interface.
func (a Uint1024) AppendText(dst []byte) ([]byte, error) {
	d, _ := formatBits1024(dst, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15], 10, false, true)
	return d, nil
}

// String returns the string representation of a in base 10.
func (a Uint1024) String() string {
	_, s := formatBits1024(nil, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11], a[12], a[13], a[14], a[15], 10, false, false)
	return s
}

// Format implements [fmt.Formatter].
func (a Uint1024) Format(s fmt.State, verb rune) {
	format(s, verb, a.Sign(), a)
}
