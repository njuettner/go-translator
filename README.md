[![Build Status](https://travis-ci.org/njuettner/go-translator.svg?branch=master)](https://travis-ci.org/njuettner/go-translator)
# Dead Simple Translator

## Installation

```console
make install
```

## How to use it

```console
$ go-translator news
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
