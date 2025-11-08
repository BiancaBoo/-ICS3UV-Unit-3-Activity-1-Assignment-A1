/*
 * @author Bianca Bertinato
 * @version 1.0.0
 * @date 2025-10-08
 * @fileoverview This program computes and outputs the area of a circle with a radius of 5.6cm.
 */
package main

import "fmt"

func main() {
	  // initialize numbers as constants
		const PI float64 = 3.14
		const RADIUS float64 = 5.6

		// INPUT - none 

    // PROCESS
		// calculate the area of the circle
		area := PI * RADIUS * RADIUS 

		// OUTPUT
		// display the result
		fmt.Println("THE AREA OF A CIRCLE WITH A RADIUS", RADIUS, "CM IS", area, "CM²")

		fmt.Println("\nDone.")
}
