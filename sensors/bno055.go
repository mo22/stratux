package sensors

import (
	"fmt"
	"github.com/kpeu3i/bno055"
)

type BNO055 struct {
	mpu *bno055.BNO055
}

func NewBNO055(i2cbus *embd.I2CBus) (*BNO055, error) {
	var (
		m   BNO055
		// mpu *bno055.BNO055
		err error
	)

	fmt.Printf("NewBNO055\n")

	sensor, err := bno055.NewSensor(0x28, 1)
	if err != nil {
		return nil, err
	}

	err = sensor.UseExternalCrystal(true)
	if err != nil {
		return nil, err
	}

	status, err := sensor.Status()
	if err != nil {
		return nil, err
	}

	revision, err := sensor.Revision()
	if err != nil {
		return nil, err
	}
	fmt.Printf(
		"*** Revision: software=%v, bootloader=%v, accelerometer=%v, gyroscope=%v, magnetometer=%v\n",
		revision.Software,
		revision.Bootloader,
		revision.Accelerometer,
		revision.Gyroscope,
		revision.Magnetometer,
	)

	axisConfig, err := sensor.AxisConfig()
	if err != nil {
		return nil, err
	}
	fmt.Printf(
		"*** Axis: x=%v, y=%v, z=%v, sign_x=%v, sign_y=%v, sign_z=%v\n",
		axisConfig.X,
		axisConfig.Y,
		axisConfig.Z,
		axisConfig.SignX,
		axisConfig.SignY,
		axisConfig.SignZ,
	)

	m.mpu = sensor

	return &m, nil
}

func (m *BNO055) Read() (T int64, G1, G2, G3, A1, A2, A3, M1, M2, M3 float64, GAError, MAGError error) {
/*
	var (
		data *icm20948.MPUData
		i    int8
	)
	data = new(icm20948.MPUData)

	for data.N == 0 && i < 5 {
		data = <-m.mpu.CAvg
		T = data.T.UnixNano()
		G1 = data.G1
		G2 = data.G2
		G3 = data.G3
		A1 = data.A1
		A2 = data.A2
		A3 = data.A3
		M1 = data.M1
		M2 = data.M2
		M3 = data.M3
		GAError = data.GAError
		MAGError = data.MagError
		i++
	}
*/
	return
}

func (m *BNO055) ReadOne() (T int64, G1, G2, G3, A1, A2, A3, M1, M2, M3 float64, GAError, MAGError error) {
/*
	var (
		data *icm20948.MPUData
	)
	data = new(icm20948.MPUData)

	data = <-m.mpu.C
	T = data.T.UnixNano()
	G1 = data.G1
	G2 = data.G2
	G3 = data.G3
	A1 = data.A1
	A2 = data.A2
	A3 = data.A3
	M1 = data.M1
	M2 = data.M2
	M3 = data.M3
	GAError = data.GAError
	MAGError = data.MagError
*/
	return
}

func (m *BNO055) Close() {
//	m.mpu.CloseMPU()
}
