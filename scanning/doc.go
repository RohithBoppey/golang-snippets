/*
Package scanning holds helpers for reading user input.

Ways to read input in Go:

	scan    - keep reading values constantly and fill them in variables like: Scan(&a, &b, ....)
	scanf   - you give the type of the variable too like: Scanf("%s %d", a, b)
	scanln  - similar to scan, but only works for the first line: hello/n world only gives hello in the answer

	bufio scanner - creates a buffer pointing to the input and reads it line by line
*/
package scanning
