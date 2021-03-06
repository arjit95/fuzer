<p align="center">
  <h3 align="center">Fuzer</h3>
  <p align="center">    
    A fuzzy search library written in golang with webassembly support
    <br />
    <br />
    <a href="https://github.com/arjit95/fuzer/issues">Report Bug</a>
    ·
    <a href="https://github.com/arjit95/fuzer/issues">Request Feature</a>
    .
    <a href="https://blissful-mcclintock-54e7e3.netlify.app/">Demo</a>
  </p>
</p>

<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the project](#about-the-project)
* [Installing](#installing)
* [Usage](#usage)
    * [Using with webassembly](using-with-webassembly)
* [License](#license)

## About the project
Fuzer is a simple library to fuzzy autocomplete suggestions based on a simple ranking algorithm

## Installing

First, use go get to install the latest version of the library. This command will download fuzer with all its dependencies:

```bash
go get -u github.com/arjit95/fuzer
```

## Usage
Check out [fuzer.go](cmd/fuzer.go) for sample usage.

### Using with webassembly
First, we need to generate a wasm file which could be embedded inside a web page.

```bash
// clone the repository
git clone https://github.com/arjit95/fuzer
cd fuzer

// Build the wasm file
GOOS=js GOARCH=wasm go build -o main.wasm web/api.go

cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" <path to your website folder>
cp main.wasm <path to your website folder>
```

A simple loader file is also provided. [fuzer.js](web/fuzer.js) is present for loading the binary. After embedding the JS file in your page, it will export `Fuzer` to your global namespace.

```html
<html>
    <head>Fuzer Demo</head>
    <body>
    <script type="text/javascript" src="fuzer.js">
    <script type="text/javascript">
        async function loadLibrary() {
            await Fuzer.load(<path to wasm binary>);

            // your fuzer instance
            const instance = Fuzer.getInstance();
        }
    </script>
    </body>
```

A working sample is available at [worker.js](https://github.com/arjit95/fuzer-frontend/static/worker.js).

## License
Distributed under the MIT License. See `LICENSE` for more information.