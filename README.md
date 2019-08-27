# Set-Intersection_GO

A program that when given two CSV files of numbers, compares them and outputs the total and distinct count of numbers in each file as well as the total and distinct overlap between the values in the two files.

## Usage

Included in this repo are two executables, one for windows and one for linux. Simply running the program will prompt the user for two files to analyse, however it is also possible to pass the filepaths as command line arguments. Passing one file will perform a single file analysis and output the total and distinct count of values in that file, passing two or more will analyse and compare the first two files given.

### Windows

From within the repository directory, run:

```
.\setIntersection.exe [pathToFirstFile] [pathToSecondFile]
```

### Linux

From within the repository directory, run:

```
./setIntersection [pathToFirstFile] [pathToSecondFile]
```


## Functionality

The program works by scanning through each line in the CSV and attempting to convert it to a numeric value, if it succeeds that value is then stored as a key in a map to a counter of the number of occurences. This means when repeated numbers occur, they are not added multiple times but instead increment the counter mapped to their value, avoiding the need for huge numbers of comparisons. While scanning through the file, it also increments a simple counter for each numeric value it comes across, this is where the total count for a given file is found while the distinct count comes from the length of the map.

Once the two files have been counted and sorted into their maps a comparison is done of the two maps, iterating over one and checking the count of the value in the other map. The distinct overlap is then however many values are present in the both maps while the total overlap is the count of all the occurences of any values that are present in both.

## Notes

* Should the need arise for analysis of more than two files, the program would not be hard to adapt. The individual file analysis step can be performed on as many files as need be and the file comparison would only need to be adapted to check each map against eachother.
* Due to the way the files are individually counted first it would be relatively simple to parallelise this step, leading to slight speed improvements in the two file case but further improvements should more than two files need to be processed at once.