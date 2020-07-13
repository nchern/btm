package main

func (e *event) date() string {

	// 0 - date
	// example:         ["09/Jul/2020 16:34:27",
	return e.raw[0].(string)
}

func (e *event) device() string {

	// 1 - device
	// example:          "RockBLOCK 18388",
	return e.raw[1].(string)
}

func (e *event) direction() string {

	// 2 - direction. MO: from device to the float64ernet, MT: from the float64ernet to device.
	// example:          "MO",
	return e.raw[2].(string)
}

func (e *event) hex() string {

	// 3: hex data
	// example:          "372c392c31362c33342c31372c33362e39363039382c2d3132322e30303135312c34352c20342c32312c31332e34312c302e30302c31332e34302c302e31342c31342e35392c302e33302c313030302c302c33302e302c312c33362e39333536302c2d3132322e30303037392c3235322c2d3133302c323832342c322e30332c3431",
	return e.raw[3].(string)
}

func (e *event) Lat() float64 {

	// 4: Iridium lat; very imprecise.
	// example:          36.973366666666664,
	return e.raw[4].(float64)
}

func (e *event) Long() float64 {

	// 5: Iridium long; very imprecise.
	// example:          -121.98155,
	return e.raw[5].(float64)
}

func (e *event) month() float64 {

	// 6: month
	// example:          7,
	return e.raw[6].(float64)
}
func (e *event) day() float64 {

	// 7: day
	// example:          9,
	return e.raw[7].(float64)
}
func (e *event) hour() float64 {

	// 8: hour
	// example:          16,
	return e.raw[8].(float64)
}
func (e *event) minute() float64 {

	// 9: minute
	// example:          34,
	return e.raw[9].(float64)
}
func (e *event) second() float64 {

	// 10: second
	// example:          17,
	return e.raw[10].(float64)
}
func (e *event) latitude() float64 {

	// 11: latitude
	// example:          36.96098,
	return e.raw[11].(float64)
}
func (e *event) longitude() float64 {

	// 12: longitude
	// example:          -122.00151,
	return e.raw[12].(float64)
}
func (e *event) heading() float64 {

	// 13: heading
	// example:          45,
	return e.raw[13].(float64)
}
func (e *event) pitch() float64 {

	// 14: pitch
	// example:          4,
	return e.raw[14].(float64)
}
func (e *event) roll() float64 {

	// 15: roll
	// example:          21,
	return e.raw[15].(float64)
}
func (e *event) thrusterV() float64 {

	// 16: thrusterV
	// example:          13.41,
	return e.raw[16].(float64)
}
func (e *event) thrusterA() float64 {

	// 17: thrusterA
	// example:          0.00,
	return e.raw[17].(float64)
}
func (e *event) hotelV() float64 {

	// 18: hotelV
	// example:          13.40,
	return e.raw[18].(float64)
}
func (e *event) hotelA() float64 {

	// 19: hotelA
	// example:          0.14,
	return e.raw[19].(float64)
}
func (e *event) solarV() float64 {

	// 20: solarV
	// example:          14.59,
	return e.raw[20].(float64)
}
func (e *event) solarA() float64 {

	// 21: solarA
	// example:          0.30,
	return e.raw[21].(float64)
}
func (e *event) throttle() float64 {

	// 22: throttle
	// example:          1000,
	return e.raw[22].(float64)
}
func (e *event) rpm() float64 {

	// 23: rpm
	// example:          0,
	return e.raw[23].(float64)
}
func (e *event) rudderAngle() float64 {

	// 24: rudderAngle
	// example:          30.0,
	return e.raw[24].(float64)
}
func (e *event) nextWaypofloat64Number() float64 {

	// 25: nextWaypofloat64Number
	// example:          1,
	return e.raw[25].(float64)
}
func (e *event) nextWaypofloat64Lat() float64 {

	// 26: nextWaypofloat64Lat
	// example:          36.93560,
	return e.raw[26].(float64)
}
func (e *event) nextWaypofloat64Long() float64 {

	// 27: nextWaypofloat64Long
	// example:          -122.00079,
	return e.raw[27].(float64)
}
func (e *event) targetHeading() float64 {

	// 28: targetHeading
	// example:          252,
	return e.raw[28].(float64)
}
func (e *event) crossTrackError() float64 {

	// 29: crossTrackError
	// example:          -130,
	return e.raw[29].(float64)
}
func (e *event) nextWaypofloat64Distance() float64 {

	// 30: nextWaypofloat64Distance
	// example:          2824,
	return e.raw[30].(float64)
}
func (e *event) averageSpeed() float64 {

	// 31: averageSpeed
	// example:          2.03,
	return e.raw[31].(float64)
}
func (e *event) uptimeMins() float64 {

	// 32: uptimeMins
	// example:          41,
	return e.raw[32].(float64)
}
