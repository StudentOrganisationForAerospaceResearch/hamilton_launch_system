package main

import (
	"errors"
	"testing"
)

var buildAccelGyroMagnetismMsgTests = []struct {
	testName    string
	input       []byte
	expectedMsg AccelGyroMagnetismMsg
	expectedErr error
}{
	{
		"normal buffer",
		[]byte{
			0x31,
			0x00, 0x00, 0x29, 0xab,
			0xff, 0xf4, 0x50, 0x55,
			0x00, 0x00, 0x00, 0x05,
			0x00, 0x00, 0x29, 0xab,
			0xff, 0xf4, 0x50, 0x55,
			0x00, 0x00, 0x00, 0x05,
			0x00, 0x00, 0x29, 0xab,
			0xff, 0xf4, 0x50, 0x55,
			0x00, 0x00, 0x00, 0x05,
			0x00},
		AccelGyroMagnetismMsg{
			"accelGyroMagnetism",
			Vec3{
				10667,
				-765867,
				5,
			},
			Vec3{
				10667,
				-765867,
				5,
			},
			Vec3{
				10667,
				-765867,
				5,
			},
		},
		nil,
	},
	{
		"invalid length",
		[]byte{
			0x31,
			0x00, 0x00, 0x29, 0xab,
			0xff, 0xf4, 0x50,
			0x00},
		AccelGyroMagnetismMsg{},
		errors.New("accelGyroMagnetism length invalid, found 9, expected 38"),
	},
}

func TestBuildAccelGyroMagnetismMsg(t *testing.T) {
	for _, val := range buildAccelGyroMagnetismMsgTests {
		msg, err := buildAccelGyroMagnetismMsg(val.input)
		compareResults(val.testName, msg, err, val.expectedMsg, val.expectedErr, t)
	}
}

var buildBarometerMsgTests = []struct {
	testName    string
	input       []byte
	expectedMsg BarometerMsg
	expectedErr error
}{
	{
		"normal buffer",
		[]byte{
			0x31,
			0x00, 0x00, 0x29, 0xab,
			0xff, 0xf4, 0x50, 0x55,
			0x00},
		BarometerMsg{
			"barometer",
			10667,
			-765867,
		},
		nil,
	},
	{
		"invalid length",
		[]byte{
			0x31,
			0x00, 0x00, 0x29, 0xab,
			0xff, 0xf4, 0x50,
			0x00},
		BarometerMsg{},
		errors.New("barometer length invalid, found 9, expected 10"),
	},
}

func TestBuildBarometerMsg(t *testing.T) {
	for _, val := range buildBarometerMsgTests {
		msg, err := buildBarometerMsg(val.input)
		compareResults(val.testName, msg, err, val.expectedMsg, val.expectedErr, t)
	}
}

var buildGpsMsgTests = []struct {
	testName    string
	input       []byte
	expectedMsg GpsMsg
	expectedErr error
}{
	{
		"normal buffer",
		[]byte{
			0x31,
			0x00, 0x00, 0x29, 0xab,
			0xff, 0xf4, 0x50, 0x55,
			0x00, 0x00, 0x29, 0xab,
			0xff, 0xf4, 0x50, 0x55,
			0x00},
		GpsMsg{
			"gps",
			10667,
			-765867,
			10667,
			-765867,
		},
		nil,
	},
	{
		"invalid length",
		[]byte{
			0x31,
			0x00, 0x00, 0x29, 0xab,
			0xff, 0xf4, 0x50,
			0x00},
		GpsMsg{},
		errors.New("gps length invalid, found 9, expected 18"),
	},
}

func TestBuildGpsMsg(t *testing.T) {
	for _, val := range buildGpsMsgTests {
		msg, err := buildGpsMsg(val.input)
		compareResults(val.testName, msg, err, val.expectedMsg, val.expectedErr, t)
	}
}

var buildOxidizerTankConditionsMsgTests = []struct {
	testName    string
	input       []byte
	expectedMsg OxidizerTankConditionsMsg
	expectedErr error
}{
	{
		"normal buffer",
		[]byte{
			0x31,
			0x00, 0x00, 0x29, 0xab,
			0xff, 0xf4, 0x50, 0x55,
			0x00},
		OxidizerTankConditionsMsg{
			"oxidizerTankConditions",
			10667,
			-765867,
		},
		nil,
	},
	{
		"invalid length",
		[]byte{
			0x31,
			0x00, 0x00, 0x29, 0xab,
			0xff, 0xf4, 0x50,
			0x00},
		OxidizerTankConditionsMsg{},
		errors.New("oxidizerTankConditions length invalid, found 9, expected 10"),
	},
}

func TestBuildOxidizerTankConditionsMsg(t *testing.T) {
	for _, val := range buildOxidizerTankConditionsMsgTests {
		msg, err := buildOxidizerTankConditionsMsg(val.input)
		compareResults(val.testName, msg, err, val.expectedMsg, val.expectedErr, t)
	}
}

func compareResults(
	testName string,
	actualMsg interface{},
	actualErr error,
	expectedMsg interface{},
	expectedErr error,
	t *testing.T) {

	if actualMsg != expectedMsg {
		t.Errorf("Test <%s> Expected msg: \n|%v|\n But got \n|%v|\n instead\n", testName, expectedMsg, actualMsg)
	} else if actualErr == nil && expectedErr != nil {
		t.Errorf("Test <%s> expected error: \n|%v|\n but got no error\n", testName, expectedErr, actualErr)
	} else if actualErr != nil && expectedErr == nil {
		t.Errorf("Test <%s> expected no error, but got error:\n|%v|\n", testName, expectedErr, actualErr)
	} else if actualErr == expectedErr {
		// pass
	} else if actualErr.Error() != expectedErr.Error() {
		t.Errorf("Test <%s> expected error: \n|%v|\n but got\n|%v| instead\n", testName, expectedErr, actualErr)
	}
}
