A transliteration of Russ Cox's [`tab.c`](https://bsky.app/profile/swtch.com/post/3lnnlg24jgs2a) program in Go. Intended for use with [Acme](https://research.swtch.com/acme) when editing text which is indented with spaces instead of tabs.

Git can automatically convert between tabs and spaces using clean and smudge [filters](https://git-scm.com/docs/gitattributes#_filter).

Define filters for different tab widths:

```
for (i in 2 4 8) {
	git config --global filter.tab$i.clean 'acme-tab -u -n '$i
	git config --global filter.tab$i.smudge 'acme-tab -n '$i
}
```

In a `.gitattributes` file, or [`.git/info/attributes` to avoid committing](https://www.kenmuse.com/blog/creating-a-gitattributes-without-committing/)

```
*.yaml filter=tab2
```