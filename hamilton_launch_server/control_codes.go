package main

import (
	"fmt"
)

type ControlCodes struct {
	SoftwareArm         string `yaml:"software_arm"`
	LaunchSystemsArm    string `yaml:"launch_systems_arm"`
	VPRocketsArm        string `yaml:"vp_rockets_arm"`
	SoftwareLaunch      string `yaml:"software_launch"`
	LaunchSystemsLaunch string `yaml:"launch_systems_launch"`
	VPRocketsLaunch     string `yaml:"vp_rockets_launch"`
	Abort               string `yaml:"abort"`
	FillControl         string `yaml:"fill_control"`
}

var controlCodes ControlCodes

func setControlCodes(codes ControlCodes) error {
	if err := validateControlCode(codes.SoftwareArm); err != nil {
		return err
	}
	if err := validateControlCode(codes.LaunchSystemsArm); err != nil {
		return err
	}
	if err := validateControlCode(codes.VPRocketsArm); err != nil {
		return err
	}
	if err := validateControlCode(codes.SoftwareLaunch); err != nil {
		return err
	}
	if err := validateControlCode(codes.LaunchSystemsLaunch); err != nil {
		return err
	}
	if err := validateControlCode(codes.VPRocketsLaunch); err != nil {
		return err
	}
	if err := validateControlCode(codes.Abort); err != nil {
		return err
	}

	if err := validateControlCode(codes.FillControl); err != nil {
		return err
	}

	controlCodes = codes
	return nil
}

func validateControlCode(code string) error {
	if len(code) < 3 {
		return fmt.Errorf("Control code '%s' too short, must be at least 3 characters", code)
	} else if len(code) > 25 {
		return fmt.Errorf("Control code '%s' too long, must be at most 25 characters", code)
	}
	return nil
}
