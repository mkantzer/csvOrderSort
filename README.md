## csvOrderSort
project for a friend for handling a simple CSV squish, used as a learning excersize


# Generation:

In `generate.go`, we create a CSV file 10,000 lines long, composed of a Part Number and a Quanity.
Each line is represents a seperate "order" for some parts. 

An example of the generated file is in `testdata/orderList.csv`.

# Colapsing:

in `sort.go`, we take in a list of file location as an argument. We parse these files, and output the sum of all of the orders for any given part number.

An example output is in `testdata/collapsed.csv`
