# Anonymize Mysqldump Tool

[![Test Main Branch](https://github.com/DekodeInteraktiv/anonymize-mysqldump/actions/workflows/test-main.yml/badge.svg?branch=main)](https://github.com/DekodeInteraktiv/anonymize-mysqldump/actions/workflows/test-main.yml)

Allows you to pipe data from `mysqldump` and anonymize it:

```sh
mysqldump -u yada -pbadpass -h db | anonymize-mysqldump --config config.json > anonymized.sql
```

You may also pipe in the content of an existing SQL dump using `cat`:

```sh
cat database.sql | anonymize-mysqldump --config config.json > anonymized.sql
```

You can also define the locale for the fake data generation (defaults to `en`):

```sh
cat database.sql | anonymize-mysqldump --config config.json --locale nb_no > anonymized.sql
```

Here is the help command output as a guide.

```sh
Usage:
  anonymize-mysqldump [flags]

Flags:
  --help -h      Outputs help text and exits.
  --config       The path to a custom config file.
  --locale       The faker locale.

Config:
  The anonymizer will use a default config suitable for WordPress, but you can override this by providing your own.
```

## Installation

### Built binaries
You can download the binary for your system from the [Releases](https://github.com/DekodeInteraktiv/anonymize-mysqldump/releases/) page. Once downloaded and unarchived, move it to a location in your path such as `/usr/local/bin` and make it executable. For instance, to download the MacOS binary for 64 bit platforms (this is most common):

```sh
curl -OL https://github.com/DekodeInteraktiv/anonymize-mysqldump/releases/download/latest/anonymize-mysqldump_darwin_amd64.gz
gunzip anonymize-mysqldump_darwin_amd64.gz
mv anonymize-mysqldump_darwin_amd64 /usr/local/bin/anonymize-mysqldump
chmod +x /usr/local/bin/anonymize-mysqldump
```

### Direct from source
You can clone this Git repository and build the project locally your self, you will then need a local [Golang](https://go.dev/) environment.

```sh
git clone https://github.com/DekodeInteraktiv/anonymize-mysqldump anonymize-mysqldump
cd anonymize-mysqldump
go install ./...
```

The above will install a copy to your `PATH`. You can also just build the executable and move it where you prefer using `go build ./...`, do take note of the **three** dots after the forward slash, which will make sure the build includes subdirectories.

#### Running tests

Golang comes with a test suite built in, you can run the tests on every file in a project using `go test ./...`

## Usage

This tool is designed to read a file stream over STDIN and produce an output over STDOUT. A config file is required and can be provided via the `-c` or `--config` flag. An example config for anonymizing a WordPress database is provided at [`config.example.json`](./config.example.json):

```sh
curl -LO https://raw.githubusercontent.com/DekodeInteraktiv/anonymize-mysqldump/main/config.example.json
```

Whenever the tool experiences an error, it will output a log to STDERR. If you wish to not see that output while the command is running, redirect it to some other file (or `/dev/null` if you don't care):

```sh
mysqldump -u yada -pbadpass -h db | anonymize-mysqldump --config config.json 2> path/to/errors.log > anonymized.sql
```

## Caveats

Important things to be aware of!

- Currently this only modifies `INSERT` statements. Should you wish to modify other fields, feel free to submit a PR.
- **Verify the output file has been modified.** This is a friendly reminder this tool is still in its early days and you should verify the output sql file before distributing it to ensure the desired modifications have been applied.

## Config File

An example config for anonymizing a WordPress database is provided at [`config.example.json`](./config.example.json).

```json
{
  "patterns": [
    {
      "tableName": "wp_users",
      "purge": false,
      "fields": []
    }
  ]
}
```

The config is composed of many objects in the `patterns` array:

- `patterns`: an array of objects defining what modifications should be made.
  - `tableName`: the name of the table the data will be stored in (used to parse `INSERT` statements to determine if the query should be modified.). You can also use regex to identify the relevant tables, required for multisite compatibility e.g. `.*_comments`.
  - `purge`: Optional `boolean` field, if set to `true` then any `INSERT` query matching the table name will be stripped out, avoiding accidentally including data that can't be anonymized in your final result.
  - `fields`: an array of objects defining modifications to individual values' fields
    - `field`: a string representing the name of the field. Not currently used, but still required to work and useful for debugging.
    - `position`: the 1-based index of what number column this field represents. For instance, assuming a table with 3 columns `foo`, `bar`, and `baz`, and you wished to modify the `bar` column, this value would be `2`.
    - `type`: a string representing the type of data stored in this field. Read more about field types [here](#field-types).
    - `constraints`: an array of objects defining comparison rules used to determine if a value should be modified or not. Currently these are limited to a simple string equality comparison.
      - `field`: a string representing the name of the field.
      - `position`: the 1-based index of what number column this field represents. For instance, assuming a table with 3 columns `foo`, `bar`, and `baz`, and you wished to modify the `bar` column, this value would be `2`.
      - `value`: string value to match against.

### Constraints

Supposing you have a WordPress database and you need to modify certain meta, be it user meta, post meta, or comment meta. You can use `constraints` to update data only whenever a certain condition is matched. For instance, let's say you have a user meta key `last_ip_address`. If you wanted to change that value, you can use the following config in the `fields` array:

```json
{
  "field": "meta_value",
  "position": 4,
  "type": "ipv4",
  "constraints": [
    {
      "field": "meta_key",
      "position": 3,
      "value": "last_ip_address"
    }
  ]
}

```

### Field Types

Each column stores a certain type of data, be it a name, username, email, etc. The `type` property in the config is used to define the type of data stored, and ultimately the type of random data to be inserted into the field. [https://github.com/dmgk/faker](https://github.com/dmgk/faker) is used for generating the fake data. These are the types currently supported:

| Type                   | Example output                                                                                                                                                                                                                  |
|------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| `username`             | `micah_pfeffer`                                                                                                                                                                                                                 |
| `password`             | `5ccf672d5c73748146be6b37568efa57`                                                                                                                                                                                              |
| `email`                | `maritza@farrell.org`                                                                                                                                                                                                           |
| `url`                  | `http://sporerhamill.net/kyla.schmitt`                                                                                                                                                                                          |
| `name`                 | `Natasha Hartmann`                                                                                                                                                                                                              |
| `firstName`            | `Carolina`                                                                                                                                                                                                                      |
| `lastName`             | `Kohler`                                                                                                                                                                                                                        |
| `personPrefix`         | `Dr.`                                                                                                                                                                                                                           |
| `personTitle`          | `Chief Functionality Orchestrator`                                                                                                                                                                                              |
| `phoneNumber`          | `+49-131-0003060`                                                                                                                                                                                                               |
| `billingAddressFull`   | `Carolina Kohler 6071 Heaney Island Suite 553, Ebbaville Texas 37307 JP +49-131-0003060 maritza@farrell.org`                                                                                                                    |
| `addressFull`          | `6071 Heaney Island Suite 553, Ebbaville Texas 37307`                                                                                                                                                                           |
| `addressStreet`        | `"586 Sylvester Turnpike"`                                                                                                                                                                                                      |
| `addressSecondary`     | `Apt. 411`                                                                                                                                                                                                                      |
| `addressCity`          | `North Dessie`                                                                                                                                                                                                                  |
| `addressState`         | `Maryland`                                                                                                                                                                                                                      |
| `addressPostCode`      | `31340`                                                                                                                                                                                                                         |
| `addressCountry`       | `Uruguay`                                                                                                                                                                                                                       |
| `addressCountryCode`   | `JP`                                                                                                                                                                                                                            |
| `paragraph`            | `Lorem ipsum dolor sit amet, fabulas nostrum recteque vel ea, sit ut nemore similique. Ad per dicam molestie, nostro constituto duo ad. Ex scripta impedit cum, vidisse feugiat vivendum cum ad, liber senserit mediocrem pro.` |
| `shortString`          | `wqFyJIrXYfVP7cLwqFyJIrXYfVP7cL`                                                                                                                                                                                                |
| `ipv4`                 | `121.204.82.227`                                                                                                                                                                                                                |
| `companyName`          | `Aufderhar LLC`                                                                                                                                                                                                                 |
| `companySuffix`        | `Inc`                                                                                                                                                                                                                           |
| `companyNumber`        | `123456789`                                                                                                                                                                                                                     |
| `creditCardNumber`     | `1234-2121-1221-1211`                                                                                                                                                                                                           |
| `creditCardExpiryDate` | `2015-11-11`                                                                                                                                                                                                                    |
| `creditCardType`       | `mastercard`                                                                                                                                                                                                                    |
| `norwegianSSN`         | `07026765743`                                                                                                                                                                                                                   |
| `purge`                |                                                                                                                                                                                                                                 |

If you need another type, please feel free to add support and file a PR!

## Releases

The releases are built by [GoReleaser](https://goreleaser.com) which automatically generates binaries for various operating systems and creates a Github release with them.

## Credit

Many thanks to [`Automattic/go-search-replace`](https://github.com/Automattic/go-search-replace) for serving as the starting point for this tool, and [`HumanMade/go-anonymize-mysqldump`](https://github.com/HumanMade/go-anonymize-mysqldump) for their improvements. Also many thanks to [`xwb1989/sqlparser`](https://github.com/xwb1989/sqlparser) for the SQL parsing library. We wouldn't have been able to do this without them!
