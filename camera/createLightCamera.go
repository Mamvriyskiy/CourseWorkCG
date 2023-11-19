package camera

import (
	"github.com/Mamvriyskiy/CourseWorkCG/inter"
	"github.com/Mamvriyskiy/CourseWorkCG/mathfunc"
)

func InitLightCamera() inter.Camera {
	var lightSource inter.Camera

	lightSource.FYaw = 0.7
	lightSource.VTarget = mathfunc.MakeVec3(2.4, 2, -0.28)
 	lightSource.VCamera = mathfunc.MakeVec3(3, 2, -1)

	return lightSource
}

