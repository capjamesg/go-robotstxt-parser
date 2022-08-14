# Golang Robots.txt Parser

Parses a robots.txt file and provides a function to check whether a robot is or is not allowed to crawl a given URL.

## Project Documentation

package main // import "robots-parser"


### FUNCTIONS

#### func checkIfAllowed(robot RobotsFile, userAgent string, path string) bool
    checkIfAllowed accepts a RobotsFile, user agent, and path as an argument.
    The function returns true if the specified user agent is allowed to
    crawl the specified path, according to the provided RobotsFile object.
    If the specified user agent is not allowed to crawl the path, checkIfAllowed
    returns false.

#### func displayAllowed(robot RobotsFile, userAgent string)
    Prints the names of all of the paths a specified user agent is explicitly
    allowed to crawl. This function returns the results of Allow directives
    applied to a user agent only. This function does not tell you if a specified
    user agent is not allowed to crawl a resource. Useful for debugging.

#### func displayDisallowed(robot RobotsFile)
    Prints the names of all of the paths a specified user agent is not allowed
    to crawl. Useful for debugging.

#### func displaySitemaps(robot RobotsFile)
    Prints the URL of each sitemap listed in a site's robots.txt file to the
    console.

#### func handleError(err error)
    Handles errors that occur in the code

#### func makeRobot(fileName string) RobotsFile
    Creates a RobotsFile object that contains a list of robots rules and
    sitemaps found in the site's robots.txt file.


### TYPES

    type Agent struct {
            Name       string
            Allowed    []string
            Disallowed []string
    }
    An Agent lists the Allowed and Disallowed directives that apply to a user
    agent

    type RobotsFile struct {
            Rules    []Agent
            Sitemaps []string
    }
    A RobotsFile contains information about a robots.txt file

## License

This project is licensed under the [MIT license](LICENSE).

## Contributors

- capjamesg