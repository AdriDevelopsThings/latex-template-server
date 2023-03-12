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
# it's required to set a salt value
salt: "generate a secure random salt with pwgen for example (required)"
app_url: "http://localhost:3000" # it's required to set the public url of the backend for link generation
encryption_key_size: 48 # 48 = default value
file_serve_path: "files" # directory where generated PDFs will be saved for X minutes, "files" = default value
template_path: "templates" # directory which hosts all LaTeX templates, "templates" = default value
tmp_directory: "tmp" # "tmp" = default value
delete_file_after: 600 # time after that a generated PDF file will be deleted, 600 = default value
```

## Create a LaTeX template

Just create a file with the name `YOUR_TEMPLATE_NAME.tex` in the template directory configured in the `configured.yml` file. Here is an example for a LaTeX file that needs two parameters: `firstname` and `lastname`:

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
GET /template/YOUR_TEMPLATE_NAME HTTP/1.1
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
Content-Type: application/json

{
    "link": "http://localhost:3000/file/id/encryption_key/filename
}
```

You can download this file from this link for the `delete_time_after` configured in the `configuration.yml` file. As you can see, the files are encrypted with AES and the encryption_key supplied in the link to the filename.