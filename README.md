# Data-pirates-challenge-go

This project aims to solve the challenge of building a web scraper.

The application must collect zip code and location data, as well as assign a unique ID to each record.

# Challenge details

- The application must fetch the data of two or more states directly from the web page.

- As the target system returns the information through the Cold Fusion Markup file, the solution adopted was a scraper.

![image](https://user-images.githubusercontent.com/12565936/163111260-38aea82a-be0c-49a8-8d0d-c52f357115d8.png)

- The scraper scans the page tags and searches for the indicated elements.

- Once the desired element has been found, the information is stored in lists and later saved in files with the jsonl extension.

# Running the tests locally:
- First, clone the project to your local PC respecting `GO PATH`;
- After downloading, access the project's root directory using the IDE of your choice;
- Install all dependencies by running the command: `go get -v .`
- Assuming you have the GO language installed, just run the command below;
- `go test -v ./...`

<img src="https://user-images.githubusercontent.com/12565936/163111408-e16d553c-9766-46e0-915a-a5236f5563a1.png" width="500" height="500">

# Running project

- To run the application, run the command `go run .\main.go` in the root directory.

- At the end of the execution, the files with the data will be saved in the root directory as shown below.

<img src="https://user-images.githubusercontent.com/12565936/163112307-54884511-2bc5-47ae-b36b-4348138cf361.png" width="500" height="500">

- File result example:

<img src="https://user-images.githubusercontent.com/12565936/163112474-a4f34d69-0752-40b0-97c2-4c5d63242725.png" width="500" height="300">


# PS
- Application has been developed and tested on Windows PC
