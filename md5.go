package main

import (
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func parseSignatures(path string) (map[string]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	signs := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for lnum := 1; scanner.Scan(); lnum++ {
		fields := strings.Fields(scanner.Text())
		if len(fields) != 2 {
			return nil, fmt.Errorf("%s:%d bad line", path, lnum)
		}
		signs[fields[1]] = fields[0]
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return signs, nil
}

func fileMD5(path string) (string, error) {
	file, err := os.Open(path)

	if err != nil {
		return "", err
	}

	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

type md5CheckResult struct {
	path   string
	matcth bool
	err    error
}

func md5Worker(path string, sig string, out chan *md5CheckResult) {
	r := &md5CheckResult{path: path}

	s, err := fileMD5(path)
	if err != nil {
		r.err = err
		out <- r
		return
	}

	r.matcth = (s == sig)
	out <- r
}

func md5Challenge() {
	sigs, err := parseSignatures("nasa-logs/md5sum.txt")
	if err != nil {
		log.Fatalf("ERROR cannot read signature file - %s", err)
	}

	out := make(chan *md5CheckResult)
	for path, sig := range sigs {
		go md5Worker(path, sig, out)
	}

	for range sigs {
		r := <-out
		switch {
		case r.err != nil:
			fmt.Printf("%s: error - %s\n", r.path, r.err)
		case !r.matcth:
			fmt.Printf("%s: signature mismatch\n", r.path)
		default:
			fmt.Printf("%s: signare matched\n", r.path)
		}
	}

}
