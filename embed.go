package main

import "embed"

// these are not really comments
// this is a build time trick to bake the static assets and html templates into the binary
// TODO this makes the binary way too big with many static assets, implement better

//go:embed markdown/*.md
var MarkdownFS embed.FS

//go:embed ui/html/**
var TemplatesFS embed.FS

//go:embed static/*
var StaticFS embed.FS
