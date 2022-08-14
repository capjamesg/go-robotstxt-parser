/*
robots parses robots.txt files for specified URLs. It returns an object that contains the parsed Allowed and Disallow directives that apply to each user agent.

robots also creates a list of all of the sitemaps listed using the Sitemap directive in a robots.txt file.
*/

package main

import (
	"fmt"
	"log"
	"strings"
	"io/ioutil"
	"net/http"
	"bufio"
)

// Handles errors that occur in the code
func handleError (err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// A RobotsFile contains information about a robots.txt file
type RobotsFile struct {
	Rules []Agent
	Sitemaps []string
}

// An Agent lists the Allowed and Disallowed directives that apply to a user agent
type Agent struct {
	Name string
	Allowed []string
	Disallowed []string
}

// checkIfAllowed accepts a RobotsFile, user agent, and path as an argument.
// The function returns true if the specified user agent is allowed to crawl the specified path, according to the provided RobotsFile object.
// If the specified user agent is not allowed to crawl the path, checkIfAllowed returns false.
func checkIfAllowed (robot RobotsFile, userAgent string, path string) bool {
	for _, rule := range robot.Rules {
		if rule.Name == userAgent || rule.Name == "*" {
			for _, directive := range rule.Disallowed {
				if directive == path || strings.HasPrefix(path, directive) {
					return false
				}
			}
		}
	}

	return true
}

// Prints the names of all of the paths a specified user agent is explicitly allowed to crawl.
// This function returns the results of Allow directives applied to a user agent only.
// This function does not tell you if a specified user agent is not allowed to crawl a resource.
// Useful for debugging.
func displayAllowed (robot RobotsFile, userAgent string) {
	for _, agent := range robot.Rules {
		if agent.Name == userAgent || agent.Name == "*" {
			for _, directive := range agent.Allowed {
				fmt.Println(directive)
			}
		}
	}
}

// Prints the names of all of the paths a specified user agent is not allowed to crawl.
// Useful for debugging.
func displayDisallowed (robot RobotsFile) {
	for _, agent := range robot.Rules {
		if agent.Name == userAgent || agent.Name == "*" {
			for _, directive := range agent.Disallowed {
				fmt.Println(directive)
			}
		}
	}
}

// Prints the URL of each sitemap listed in a site's robots.txt file to the console.
func displaySitemaps (robot RobotsFile) {
	for _, sitemap := range robot.Sitemaps {
		fmt.Println(sitemap)
	}
}

// Creates a RobotsFile object that contains a list of robots rules and sitemaps found in the site's robots.txt file.
func makeRobot (fileName string) RobotsFile {
	response, err := http.Get(fileName)
	handleError(err)

	robotsFile, err := ioutil.ReadAll(response.Body)
	handleError(err)

	scanner := bufio.NewScanner(strings.NewReader(string(robotsFile)))

	currentRobotsDirective := ""

	robot := RobotsFile{}

	var disallowRulesBuffer []string
	var allowRulesBuffer []string
	var sitemaps []string

	for scanner.Scan() {
		if len(strings.Split(scanner.Text(), " ")) == 1 {
			continue
		}

		key := strings.Split(scanner.Text(), ": ")[0]
		value := strings.Split(scanner.Text(), ": ")[1]
		
		if key == "User-agent" && currentRobotsDirective != "" {
			robot.Rules = append(robot.Rules, Agent{currentRobotsDirective, allowRulesBuffer, disallowRulesBuffer})
			currentRobotsDirective = value
			allowRulesBuffer = []string{}
			disallowRulesBuffer = []string{}
		} else if key == "User-agent" {
			currentRobotsDirective = value
		} else if key == "Disallow" {
			disallowRulesBuffer = append(disallowRulesBuffer, value)
		} else if key == "Allow" {
			allowRulesBuffer = append(allowRulesBuffer, value)
		} else if key == "Sitemap" {
			robot.Sitemaps = append(sitemaps, value)
		}
	}

	robot.Rules = append(robot.Rules, Agent{currentRobotsDirective, allowRulesBuffer, disallowRulesBuffer})

	return robot
}