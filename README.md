# csv_QL
1. First argument is a csv filepath
2. Program must be able to decide if a parameter is a filename or a condition
3. Returns the lines to stdout where the condition is true

example:
  $ csv_QL  1:"hello" ./asdas.csv

Packages:
- os
 * Open
 * Args
 * Stdout
 * (bonus) Stdin
- fmt
 * Fprintf
- strings
 * Split
 * (Trim or Trimhassuffix)
- strconv
 * Atoi
- encoding/csv
 * NewReader


Bonus:
1. Add multiple csvs as argument
2. Add multiple conditions as argument
3. If no csv arguments given, use the Stdin
4. Able to process bigger file than the current computer's memory

-----
Why is it good/bad when data dominating
Why is null object pattern useful
Can I optimise this code even more
net/http package explore.
create new program to find matching lines on specific url, return the number of appearances
