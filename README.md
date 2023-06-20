[EDA](https://pkg.go.dev/github.com/gregoryv/eda) - Expenses data file format

A simple expenses log for families.

## Quick start

    $ go install github.com/gregoryv/eda/cmd/budget@latest
	
Create a file like the [example.eda](example.eda) and run
	
	$ budget example.eda
	3 200 000 loans left
    ---------- --------------------
         4 000 car
         2 791 car1
           666 daughter
        21 499 house
           166 life
        24 165 loan
         2 383 man
         1 500 mobile
           566 son
         2 383 wife
    ---------- --------------------
        44 896 sum
             2 people
    ---------- --------------------
        22 448 each

