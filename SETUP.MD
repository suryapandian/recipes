### Setup
- Pull the branch locally
- Build the docker image using the command, `docker build -t "surya:recipe-count-test" .`
- Run the docker image using the command,  `docker run -i -t "surya:recipe-count-test"`

A sample JSON file is present in the folder `test`.
To run with different sample, please copy the sample JSON into test folder and provide the file name as environment variable `FILE_NAME` while running the docker.
Example:
`docker run -e FILE_NAME=testData.json -i -t "surya:recipe-count-test"`

Similarly, other command line arguments also could be passed as environment variables.
Example:
`docker run -e POSTCODE=10208  -e TIME_RANGE=7AM-12AM -e SEARCH_STR=Steak,Chicken,Honey -i -t "surya:recipe-count-test"`

The code has also handles default values for the arguments. When no arguments are passed the code runs with the default values.