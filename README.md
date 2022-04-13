# Data-pirates-challenge-go

This project aims to solve the challenge of building a web scraper.

The application must collect zip code and location data, as well as assign a unique ID to each record.

# Challenge details

- The application must fetch the data of two or more states directly from the web page.

- As the target system returns the information through the Cold Fusion Markup file, the solution adopted was a scraper.



- The scraper scans the page tags and searches for the indicated elements.

- Once the desired element has been found, the information is stored in lists and later saved in files with the jsonl extension.

# Running the tests locally:
- First, clone the project to your local PC respecting `GO PATH`;
- After downloading, access the project's root directory using the IDE of your choice;
- Install all dependencies by running the command: `go get ./...`
- Assuming you have the GO language installed, just run the command below;
- `go test -v ./...`



# Running project

- To run the application, run the command `go run .\main.go` in the root directory.

- At the end of the execution, the files with the data will be saved in the root directory as shown below.


- File result example:


# PS
- Application has been developed and tested on Windows PC
