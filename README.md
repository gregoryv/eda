[EDA](https://pkg.go.dev/github.com/gregoryv/eda) - Expenses data file format

A simple expenses log for families.

## Quick start

    $ go install github.com/gregoryv/eda/cmd/budget@latest
	
Create a file like the [example.eda](example.eda) and run
	
	$ budget example.eda

Read more about the file format on https://gregoryv.github.io/eda.


## todo

Add number of people in the format instead as a flag. The names can be
used as tags to have the summary reflect if one person should pay more
than the other.

    # people name1 name2
	
	

