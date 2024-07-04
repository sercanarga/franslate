# franslate
AI based translation tool for translating text from one language to another. It uses the Gemini AI API to perform the translation. The tool is built using Golang and the [Fyne](https://github.com/fyne-io/fyne) library.

![last commit](https://badgen.net/github/last-commit/sercanarga/franslate) ![license](https://badgen.net/github/license/sercanarga/franslate)

## Why?
Because I got fed up with the ridiculous pricing policies of AI translate tools (you know who I'm talking about), so I decided to develop my own solution.

![Screenshot](https://github.com/sercanarga/franslate/blob/main/screenshot.png?raw=true)

## How to use
1. Get an API key from [aistudio.google.com/app/apikey](https://aistudio.google.com/app/apikey)
2. Download the latest release from the [releases page](https://github.com/sercanarga/franslate/releases)
3. Run the tool and enter your API key
4. Enter the text you want to translate

## How to build
1. Clone the repository
2. Run `go install github.com/fyne-io/fyne-cross@latest`
3. Run `fyne-cross {OS} --app-id franslate.app` (replace {OS} with your [target OS](https://github.com/fyne-io/fyne-cross?tab=readme-ov-file#usage))
4. The binary will be in the `dist` folder
5. Run the binary and follow the instructions
6. Enjoy!

## Contributing
I realize that not everything is perfect, so I welcome any contribution. Before making major changes, please open a topic to discuss what you want to change.

## License
This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details.