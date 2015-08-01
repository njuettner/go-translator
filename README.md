[![Build Status](https://travis-ci.org/njuettner/go-translator.svg?branch=master)](https://travis-ci.org/njuettner/go-translator)
# Dead Simple Translator

## Installation

* clone the repository

```
git clone ...
```

* get dependencies
```
go get github.com/PuerkitoBio/goquery
```

* install it
```
go install .
```

## How to use it

```
go-translator news
news {sg} - Nachrichten {pl}
news {sg} - Neuigkeiten {pl}
news - Nachricht {f}
news - Neuigkeit {f}
news [treated as sg.] - Kunde {f} [geh.] [veraltend] [Nachricht]
news - Zeitung {f} [veraltet für: Neuigkeit]
news {sg} - News {pl} [ugs.]
the news [newscast] {sg} - die Nachrichten [Nachrichtensendung] {pl}
news-crammed {adj} - mit Neuigkeiten vollgestopft
...
```
