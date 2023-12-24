package camera

import (
	"../inter"
	"../mathfunc"
)

func InitCamera() inter.Camera {
	var camera inter.Camera
    camera.VCamera = mathfunc.MakeVec3(0, 2, -5)
	camera.VTarget = mathfunc.MakeVec3(0, 0, 1)
	camera.VForward = mathfunc.MakeVec3(0, 0, 1)

	return camera
}

func CreateCamera(camera inter.Camera, proj [4][4]float64, flag int) ([4][4]float64, [4][4]float64) {
	vUp := mathfunc.MakeVec3(0, 1, 0)
	t := mathfunc.MakeVec4(0, 0, 1, 0)

	camera.VLookDir = mathfunc.MulVecMat(t, mathfunc.MakeRotationYM(camera.FYaw))
	camera.VLookDir = mathfunc.MulVecMat(t, mathfunc.MakeRotationYM(camera.FYaw))

	camera.VTarget = mathfunc.Vec3Sum(camera.VCamera, camera.VLookDir.Vec3)
	camera.VForward = mathfunc.Vec3Mul(camera.VLookDir.Vec3, 1.0)

	mCamera := mathfunc.MakePointAtM(camera.VCamera, camera.VTarget, vUp)
	view := mathfunc.InverseTranslationM(mCamera)

	newM := mathfunc.MulMatrices(view, proj)
	var matrix [4][4]float64
	if flag == 1 {
		matrix = Matrixminus(newM)
	}

	return newM, matrix
}

func Matrixminus(A [4][4]float64) [4][4]float64 {
	n := len(A)
    AInv := [4][4]float64{}

    // Вычисление обратной матрицы с помощью метода Гаусса-Жордана
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if i == j {
                AInv[i][j] = 1
            }
        }
    }

    for k := 0; k < n; k++ {
        // Делаем A[k][k] равным 1, деля строку k на A[k][k]
        pivot := A[k][k]
        for j := 0; j < n; j++ {
            A[k][j] /= pivot
            AInv[k][j] /= pivot
        }

        // Обнуляем все элементы в столбце k, кроме A[k][k]
        for i := 0; i < n; i++ {
            if i != k {
                factor := A[i][k]
                for j := 0; j < n; j++ {
                    A[i][j] -= factor * A[k][j]
                    AInv[i][j] -= factor * AInv[k][j]
                }
            }
        }
    }

	return AInv
}
