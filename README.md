# Biblika API

Biblika API or Bible API is third party backend to serve Bible content in JSON format. The content direct based on https://sabda.org API.

## Development

Biblika API is write in Go/Golang with protobuf or gRPC and hosted in Heroku.

## Getting Started

### Book List

Method | URL
-- | --
GET | https://biblika.herokuapp.com/v1/book/list

``` bash
{
    "book": [
        {
            "number": 1,
            "abbr": {
                "id": "Kej",
                "en": "Gen"
            },
            "name": {
                "id": "Kejadian",
                "en": "Genesis"
            },
            "chapter": 50
        },
        ...
    ]
}
```

### Passage Content

Method | URL
-- | --
GET | https://biblika.herokuapp.com/v1/passage/Gen/1/net

``` bash
{
    "verse": [
        {
            "verse": 0,
            "type": "title",
            "content": "The Creation of the World"
        },
        {
            "verse": 1,
            "type": "text",
            "content": "In the beginning God created the heavens and the earth."
        },
        ...
    ]
}
```

#### Parameter

https://biblika.herokuapp.com/v1/passage/[book]/[chapter]/[version]

Query | Description | Required
-- | -- | --
`book` | Book name of the passage | `true`
`chapter` | The chapter of passage | `true`
`version` | The code version of Bible | `true`

#### Available Version

Language | Code | Description
-- | -- | --
Bahasa Indonesia | `tl` | Terjemahan Lama
&nbsp; | `tb` | Terjemahan Baru
&nbsp; | `ayt` | Alkitab Yang Terbuka
&nbsp; | `vmd` | Versi Mudah Dibaca
&nbsp; | `tsi` | Terjemahan Sederhana Indonesia
&nbsp; | `bis` | Bahasa Indonesia Sehari-hari
&nbsp; | `milt` | Modified Indonesian Literal Translation
&nbsp; | `kskk` | Kitab Suci Komunitas Kristiani
&nbsp; | `fayh` | Firman Allah Yang Hidup
Jawa | `jawa` | Javanese Language
Bali | `bali` | Balinese Language
Toba | `toba` | Tobanese Language
karo | `karo` | Karonese Language
Sunda | `sunda` | Sundanese Language
Bugis | `bugis` | Buginese Language
Madura | `madura` | Maduranese Language
Toraja | `toraja` | Torajan Language
Makassar | `makasar` | Makassarese Language
Simalungun | `simalungun` | Simalungunese Language
Malay | `tmv` | Today Malay Version
English | `net` | New English Translation
&nbsp; | `kjv` | King James Version
&nbsp; | `leb` | Lexham English Bible
&nbsp; | `niv` | New International Version
&nbsp; | `esv` | English Standard Version
&nbsp; | `nlt` | New Living Translation
&nbsp; | `nasb` | New American Standard Bible
&nbsp; | `hcsb` | Holman Christian Standard Bible
&nbsp; | `nrsv` | New Revised Standard Version
&nbsp; | `nkjv` | New King James Version

## Status Code

Code | Description
-- | --
200 | OK
400 | Bad Request
404 | Not Found
500 | Internal Server Error

## Contact

Due to this things is still on development, I'm very open for feedback. If you have some concern or question, please let me know by shot an email to erwindosianipar@gmail.com.

## License

``` bash
Copyright © 2022 Erwindo Sianipar

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the “Software”), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

> Soli Deo Gloria
