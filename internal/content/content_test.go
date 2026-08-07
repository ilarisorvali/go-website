package content

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func BenchmarkLoadContentFiles(b *testing.B) {
	for count := 10; count <= 1000; count += 10 {
		b.Run(strconv.Itoa(count), func(b *testing.B) {
			dir := createBenchmarkContentDir(b, count)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_, err := LoadContentFiles(dir)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func BenchmarkLoadContentFilesParallel(b *testing.B) {
	for count := 10; count <= 1000; count += 10 {
		b.Run(strconv.Itoa(count), func(b *testing.B) {
			dir := createBenchmarkContentDir(b, count)

			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_, err := LoadContentFilesParallel(dir)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

func createBenchmarkContentDir(b *testing.B, count int) string {
	b.Helper()

	dir := b.TempDir()

	content := []byte(`---
title: Test Page
slug: test-page
---

# Test Page

This is some test markdown content.
`)

	for i := 0; i < count; i++ {
		path := filepath.Join(
			dir,
			"content-"+strconv.Itoa(i)+".md",
		)

		if err := os.WriteFile(path, content, 0644); err != nil {
			b.Fatal(err)
		}
	}

	return dir
}
