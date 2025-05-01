package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	nstop    int
	minspace int
	bout     *bufio.Writer
)

func untab(b *bufio.Reader) error {
	pos := 0
	for {
		r, _, err := b.ReadRune()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if r == '\t' {
			for {
				_, err := bout.WriteRune(' ')
				if err != nil {
					return err
				}
				pos++
				if pos%nstop == 0 {
					break
				}
			}
		} else {
			_, err := bout.WriteRune(r)
			if err != nil {
				return err
			}
			if r == '\n' {
				pos = 0
			} else {
				pos++
			}
		}
	}
}

func putspaces(pos, n int) error {
	if n < minspace {
		for ; n > 0; n-- {
			_, err := bout.WriteRune(' ')
			if err != nil {
				return err
			}
		}
		return nil
	}

	epos := pos + n
	for pos/nstop != epos/nstop {
		_, err := bout.WriteRune('\t')
		if err != nil {
			return err
		}
		pos += nstop - pos%nstop
	}
	for ; pos < epos; pos++ {
		_, err := bout.WriteRune(' ')
		if err != nil {
			return err
		}
	}
	return nil
}

func tab(b *bufio.Reader) error {
	pos := 0
	nspace := 0
	for {
		r, _, err := b.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		switch r {
		case ' ':
			nspace++
			pos++
			continue
		case '\t':
			nspace += nstop - pos%nstop
			pos += nstop - pos%nstop
			continue
		}

		putspaces(pos-nspace, nspace)
		nspace = 0
		_, err = bout.WriteRune(r)
		if err != nil {
			return err
		}

		if r == '\n' {
			pos = 0
		} else {
			pos++
		}
	}
	return nil
}

func main() {
	cvt := tab

	flag.IntVar(&minspace, "m", 2, "minimum spaces")
	flag.IntVar(&nstop, "n", 8, "tab stop")
	flag.BoolFunc("u", "untab", func(string) error {
		cvt = untab
		return nil
	})
	flag.Parse()

	bout = bufio.NewWriter(os.Stdout)
	defer bout.Flush()
	if flag.NArg() > 0 {
		var err error
		for i := 0; i < flag.NArg(); i++ {
			f, err := os.Open(flag.Arg(i))
			if err != nil {
				fmt.Fprintf(os.Stderr, "error opening %s: %v", os.Args[i], err)
				continue
			}
			b := bufio.NewReader(f)
			err = cvt(b)
			if err != nil {
				fmt.Fprintf(os.Stderr, "error converting %s: %v", os.Args[i], err)
			}
		}
		if err != nil {
			os.Exit(1)
		}
	} else {
		b := bufio.NewReader(os.Stdin)
		cvt(b)
	}
}
