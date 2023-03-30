# DALL·E 2 CLI

[![Licence](https://img.shields.io/github/license/Ileriayo/markdown-badges?style=for-the-badge)](./LICENSE)
![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

This is a simple command-line interface (CLI) application that uses the OpenAI API to interact with DALL·E 2 using the terminal.

## Requirements

To use this application, you will need:

- Go version 1.18 or higher installed on your computer
- An OpenAI API key (you can sign up for one [here](https://platform.openai.com/))
- A stable internet connection

## Installation

To install this application: 

1. Clone this repository:

```
git clone https://github.com/sesaquecruz/go-openai-dalle.git
```

2. Enter the project directory:

```
cd go-openai-dalle
```

3. Install the required packages:

```
go mod download
```

4. Set the following environment variable:

```
export OPENAI_API_KEY=<your-api-key>
```

Note: You can also set this api key in an .env file in the project directory.

5. Run the program by executing:

```
go run cmd/main.go
```

## Usage

After the program has been started, you can start creating images using DALL·E 2. Describe an image using natural language and the result generated will be saved as a PNG image in the "out" folder in the project directory.

## Example

<div style="display: flex; justify-content: space-around; flex-direction: column; margin: 2%;">
  <div style="display: flex; justify-content: center; margin: 2%;">
  	<img src="./examples/description.png" style="width: 45%;">
  </div>
  <div style="display: flex; justify-content: space-around;">
    <img src="./examples/lost-robot-1.png" style="width: 35%;">
    <img src="./examples/lost-robot-2.png" style="width: 35%;">
  </div>
</div>


## Contributing

Contributions are welcome! If you find a bug or would like to suggest an enhancement, please create a new issue. If you would like to contribute code, please fork the repository and submit a pull request.

## License

This project is licensed under the MIT License. See LICENSE file for more information.
