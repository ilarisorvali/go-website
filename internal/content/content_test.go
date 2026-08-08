package content

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func BenchmarkLoadContentFiles(b *testing.B) {
	for count := 1; count <= 10; count += 1 {
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
	for count := 1; count <= 10; count += 1 {
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

### asdasdasdasd

Lokim ipsum dolor sit amet, consectetur adipiscing elit. Mauris ut sagittis elit, sed consequat massa. Sed quis eleifend metus, ut gravida nulla. Aliquam erat volutpat. Fusce dictum laoreet orci ac vehicula. Nunc malesuada feugiat est, et pellentesque diam faucibus efficitur. Maecenas viverra, nisi vitae mattis rhoncus, massa risus pellentesque orci, in vestibulum est sem quis lorem. Nullam facilisis tincidunt nibh, ac cursus erat interdum ut. Vestibulum malesuada, nisi mattis fermentum volutpat, mauris erat fringilla risus, in laoreet magna augue non purus. Maecenas mattis erat eget ipsum fermentum maximus. Integer quis interdum ipsum. Integer porttitor sapien ullamcorper erat mattis porttitor. Praesent accumsan ex in metus tincidunt, id pharetra orci luctus.



![image](/images/screenshot.png)
This should come under the image, won't definitely pass an a11y test lol


[^1]: See https://pkg.go.dev/net/http
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
