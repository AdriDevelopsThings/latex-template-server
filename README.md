# latex-template-server
Host LaTeX templates and fill them out by using an api

## Set up
### Docker
**(recommended)**
Just use the docker image (`ghcr.io/adridevelopsthings/latex-template-server:main`) and set up the volumes `/dist/configuration.yml` and `/dist/templates` or just use the [`docker-compose.yml`](docker-compose.yml) in this repository and modify it to your needs.

### From source
Build the server with go: ``go build .``. You also need `pdflatex`.

## Structure of configuration.yml

configuration.yml
```yaml
server:
    listen:
        host: 0.0.0.0
        port: 80
template_path: "templates" # directory which hosts all LaTeX templates, "templates" = default value
tmp_directory: "tmp" # "tmp" = default value
```

## Create a LaTeX template

Just create a file with the name `YOUR_TEMPLATE_NAME.tex` in the template directory configured in the `config.yml` file. Here is an example for a LaTeX file that needs two parameters: `firstname` and `lastname`:

```
\documentclass{report}

\usepackage[utf8]{inputenc}
\usepackage{datatool}

% you don't have to generate this csv file, the csv file is generated for any api request
\DTLloaddb{data}{data.csv}

\begin{document}
\DTLforeach{data}{
    \Firstname=firstname,
    \Lastname=lastname}{
        Your name is \Firstname \Lastname
    }
\end{document}
```

## Making api requests
You can generate a PDF file from a LaTeX template:
```
POST /template/YOUR_TEMPLATE_NAME HTTP/1.1
Content-Type: application/json

{
    "arguments": [{
        "firstname": "Robert",
        "lastname": "Smith"
    }]
}
```

A response does look like this:
```
HTTP/1.1 200 OK
Content-Type: application/pdf

...
```