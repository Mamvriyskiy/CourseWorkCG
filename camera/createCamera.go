package camera

import (
	"../inter"
	"../mathfunc"
)

// func CreateCamera(height, width int) [4][4]float64 {

// }

func InitCamera() inter.Camera {
	var camera inter.Camera
    camera.VCamera = mathfunc.MakeVec3(0, 2, -5)
	camera.VTarget = mathfunc.MakeVec3(0, 0, 1)

	return camera
}

func CreateCamera(camera inter.Camera, proj [4][4]float64) [4][4]float64 {
	vUp := mathfunc.MakeVec3(0, 1, 0)
	t := mathfunc.MakeVec4(0, 0, 1, 0)

	camera.VLookDir = mathfunc.MulVecMat(t, mathfunc.MakeRotationYM(camera.FYaw))
	camera.VTarget = mathfunc.Vec3Sum(camera.VCamera, camera.VLookDir.Vec3)
	camera.VForward = mathfunc.Vec3Mul(camera.VLookDir.Vec3, 1.0)

	mCamera := mathfunc.MakePointAtM(camera.VCamera, camera.VTarget, vUp)
	view := mathfunc.InverseTranslationM(mCamera)

	newM := mathfunc.MulMatrices(view, proj)

	return newM
}


// type Camera struct {
// 	Pos Vec3

// 	VCamera  Vec3
// 	VLookDir Vec4
// 	FYaw     float64
// 	FTheta   float64
// 	VTarget  Vec3

// 	VForward Vec3
// }

