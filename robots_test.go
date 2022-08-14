package main

import (
	"fmt"
	"testing"
	"reflect"
)

// Helper function to execute a test
func runTest (t *testing.T, robot RobotsFile, userAgent string, path string, expected bool) {
	allowed := checkIfAllowed(robot, userAgent, path)

	if allowed != expected {
		t.Errorf("Expected %v, got %v, for user agent %v and path %v", expected, allowed, userAgent, path)
	}
}

// Tests the checkIfAllowed function
func TestRobotsParsing (t *testing.T) {
	fmt.Println("Running tests...")
	robot := makeRobot("https://jamesg.blog/robots.txt")

	runTest(t, robot, "*", "/s/", false)
	runTest(t, robot, "*", "/", true)
	runTest(t, robot, "ia_archiver", "/", false)
	// uses dash instead of underscore
	runTest(t, robot, "ia-archiver", "/", true)
}

// Tests the value of the Sitemaps item in a RobotsFile
func TestSitemaps (t *testing.T) {
	robot := makeRobot("https://jamesg.blog/robots.txt")

	expectedSitemaps := []string{"https://jamesg.blog/sitemap.xml"}

	if reflect.DeepEqual(robot.Sitemaps, expectedSitemaps) {
		t.Errorf("Expected %v, got %v", expectedSitemaps, robot.Sitemaps)
	}
}