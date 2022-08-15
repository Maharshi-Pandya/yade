package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

// TODO: Make a new package for Linear Algebra utils

// Helper function to print a byte matrix
func MatPrint(matName string, mat [][]byte) {
	fmt.Printf("\n%s is:", matName)
	for i := range mat {
		fmt.Println()
		for j := range mat[0] {
			fmt.Printf("%v ", mat[i][j])
		}
	}

	fmt.Println()
}

// Helper function to generate a random Key byte matrix of size N.
func MatGenerateSquare(num int64, n int) ([][]byte, error) {
	// TODO: Random number between 0-29 and populate the matrix
	// restricting the key size to 10
	if n > 10 {
		return nil, errors.New("Provide input size less than 10")
	}

	// init
	mat := make([][]byte, n)
	for i := range mat {
		mat[i] = make([]byte, n)
	}

	// populate the matrix
	for i := range mat {
		for j := range mat[0] {
			num, err := rand.Int(rand.Reader, big.NewInt(num))
			if err != nil {
				return nil, err
			}

			// A = 65
			randByte := num.Int64()
			mat[i][j] = byte(randByte)
		}
	}

	return mat, nil
}

// Helper function to multiply a matrix and a vector
func MatProdVec(mat [][]byte, vec []byte) ([]int, error) {
	if len(mat[0]) != len(vec) {
		return nil, errors.New("Matrix and Vector not compatible for multiplication")
	}

	// resultant vector
	res := make([]int, len(vec))

	// dot product of each matrix row with components of vector
	for i := range mat {
		tempSum := 0
		for j := range mat[0] {
			tempSum += int(mat[i][j]) * int(vec[j])
		}
		res[i] = tempSum
	}

	return res, nil
}

// Helper function to calculate modulus with constant num
func MatModConst(mat [][]int, num int) ([][]byte, error) {
	r, c := len(mat), len(mat[0])

	if r == 0 || c == 0 {
		return nil, errors.New("Matrix dimensions cannot be 0")
	}

	// init new mat
	newMat := make([][]byte, r)
	for i := range newMat {
		newMat[i] = make([]byte, c)
	}

	// perform mod
	for i := range mat {
		for j := range mat[0] {
			newMat[i][j] = byte(mat[i][j] % num)
		}
	}

	return newMat, nil
}

// Helper function to calculate Vector modulus with constant num
func VecModConst(vec []int, num int) ([]byte, error) {
	r := len(vec)

	if r == 0 {
		return nil, errors.New("Vector length cannot be 0")
	}

	// init new vector
	newVec := make([]byte, r)

	// perform modulus
	for i := range vec {
		newVec[i] = byte(vec[i] % num)
	}

	return newVec, nil
}

// Helper function to calculate the determinant of a NxN byte matrix
func MatDet(mat [][]byte) (float32, error) {
	// Make a new matrix for operations
	matcp := make([][]int, len(mat))
	for i := range matcp {
		matcp[i] = make([]int, len(mat[0]))
	}
	for i := range matcp {
		for j := range matcp[0] {
			matcp[i][j] = int(mat[i][j])
		}
	}

	// Converting the matrix to Echelon form using Gaussian elimination
	// Determinant will be the product of diagonal elements.
	var det float32 = 1.0

	n := len(matcp)
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			// skip when pivot is 0
			if matcp[j][i] == 0 {
				continue
			} else {
				// when zero division might occur, swap the rows and continue
				if matcp[j-1][i] == 0 {
					for x := 0; x < n; x++ {
						tem := matcp[j][x]
						matcp[j][x] = matcp[j-1][x]
						matcp[j-1][x] = tem
					}
					continue
				}
				// calculate ratio and subtract
				ratio := float32(matcp[j][i]) / float32(matcp[j-1][i])
				for k := 0; k < n; k++ {
					matcp[j][k] = int(float32(matcp[j][k]) - (ratio * float32(matcp[j-1][k])))
				}
			}
		}
	}

	// multiply the diagonal elements to get determinant
	for i := range matcp {
		det *= float32(matcp[i][i])
	}
	return det, nil
}
