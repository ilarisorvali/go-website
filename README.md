# go-website

# Personal Go Website

A fast, minimal personal website written in Go, using  [`goldmark`](https://github.com/yuin/goldmark) for Markdown parsing.

'''
podman run -p 8080:9000 -v /path/to/markdown/content/directory:/app/markdown:Z go-website:latest
'''
