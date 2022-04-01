// Copyright (c) 2015-2021 The btcsuite developers
// Copyright (c) 2015-2021 The Decred developers

package btcec

import (
	secp "github.com/decred/dcrd/dcrec/secp256k1/v4"
)

// JacobianPoint is an element of the group formed by the secp256k1 curve in
// Jacobian projective coordinates and thus represents a point on the curve.
type JacobianPoint = secp.JacobianPoint

// MakeJacobianPoint returns a Jacobian point with the provided X, Y, and Z
// coordinates.
func MakeJacobianPoint(x, y, z *FieldVal) JacobianPoint {
	return secp.MakeJacobianPoint(x, y, z)
}

// AddNonConst adds the passed Jacobian points together and stores the result
// in the provided result param in *non-constant* time.
func AddNonConst(p1, p2, result *JacobianPoint) {
	secp.AddNonConst(p1, p2, result)
}

// DecompressY attempts to calculate the Y coordinate for the given X
// coordinate such that the result pair is a point on the secp256k1 curve. It
// adjusts Y based on the desired oddness and returns whether or not it was
// successful since not all X coordinates are valid.
//
// The magnitude of the provided X coordinate field val must be a max of 8 for
// a correct result. The resulting Y field val will have a max magnitude of 2.
func DecompressY(x *FieldVal, odd bool, resultY *FieldVal) bool {
	return secp.DecompressY(x, odd, resultY)
}

// DoubleNonConst doubles the passed Jacobian point and stores the result in
// the provided result parameter in *non-constant* time.
//
// NOTE: The point must be normalized for this function to return the correct
// result. The resulting point will be normalized.
func DoubleNonConst(p, result *JacobianPoint) {
	secp.DoubleNonConst(p, result)
}

// ScalarBaseMultNonConst multiplies k*G where G is the base point of the group
// and k is a big endian integer. The result is stored in Jacobian coordinates
// (x1, y1, z1).
//
// NOTE: The resulting point will be normalized.
func ScalarBaseMultNonConst(k *ModNScalar, result *JacobianPoint) {
	secp.ScalarBaseMultNonConst(k, result)
}

// ScalarMultNonConst multiplies k*P where k is a big endian integer modulo the
// curve order and P is a point in Jacobian projective coordinates and stores
// the result in the provided Jacobian point.
//
// NOTE: The point must be normalized for this function to return the correct
// result. The resulting point will be normalized.
func ScalarMultNonConst(k *ModNScalar, point, result *JacobianPoint) {
	secp.ScalarMultNonConst(k, point, result)
}
