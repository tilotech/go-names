# go-names

go-names provides name lists for working with frequent names and nicknames.

## Common (Frequent Names)

To get a list of the 10 most frequent first names, use it like this:

```
import (
  "fmt"
  "github.com/tilotech/go-names"
)

func main() {
  common, err := names.NewCommonPreset("US_FIRST_NAME")
  if err != nil {
    panic(err)
  }

  top10 := common.Top(10)
  fmt.Println(top10)
}
```

The common names comes with ready-to-use presets, but you can also provide your
own using the `NewCommon` constructor.

### US_FIRST_NAME and US_LAST_NAME Presets

Those two presets are based on the publicly available
[voters registry of North Carolina](https://www.ncsbe.gov/results-data/voter-registration-data)
and represent the 5,000 most common first/middle and last names, as well as
their relative frequency. E.g. the name `Michael` is the most popular name in
that source, with 1.33% of the entries having this as either their first or
middle name.

## Canonical (Base Name)

Using the canonical name it is possible to get a single representation of a name,
independent from different spellings or other effects such as aliases. Such a
canonical name can then be used e.g. in data matching. Depending on the used
data such a canonical name may not seem very intuitive.

```
import (
  "fmt"
  "github.com/tilotech/go-names"
)

func main() {
  canonical, err := names.NewCanonicalPreset("NICKNAMES")
  if err != nil {
    panic(err)
  }

  name := canonical.Of("mickey")
  fmt.Println(name) // prints: "michael"
}
```

### NICKNAMES Preset

The nicknames preset is based on various open source nickname lists, hand curated
to represent the idea of a canonical name as good as possible. However, since
nicknames can be assigned to various names, the canonical name may not make much
sense in a few cases. E.g. the four entries `clement`, `clem`, `clementine` and
`clemmie` are in the same canonical name group and all resolve to `clement`.
Other examples are even more extreme, e.g. `lucy` resolves to `louisa`. Hence,
the preset follows a best guess approach and is far from perfect.

Nickname sources:
* https://github.com/carltonnorthern/nicknames (Apache-2.0)
* https://github.com/onyxrev/common_nickname_csv (data under Public Domain)
* https://github.com/HaJongler/diminutives.db (data under Public Domain)