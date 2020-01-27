# YouTube-DL Web Backend
___

This repository contains all the necessary code and tools to run the backend of the YouTube-DL Web Application.

## Get started

First things first, the only requirement is to have a recent version of Golang (+1.13) installed. 

All it's needed to run the backend application is a config.json file (an example can be found [here](https://github.com/youtubedl-web/backend/blob/master/cmd/backend/config_example.json)).

Note: The config.json resides in the same folder in which you run the backend application
```
go get github.com/youtubedl-web/backend

# Compile the backend application
go install github.com/youtubedl-web/backend/cmd/backend

# Run the compiled application
backend
```

## Contributing

First of all thank you so much for your desire to contribute to this project. To make things easier all the architecture, design choices and engineering stuff will be saved on the Wiki tab of this repository in order to keep this README short and clear.