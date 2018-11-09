# rapidlyCode

help google translate for maplestory to chinese 

## command
`--src` or  `-s` input file path, default is `./a.html`

`--dest` or `-d` output file path , default is `./b.html`

`--rule` or `-r` rule file path, default is `./rule.md`

`--noad` don't add AD to output file content

`--adpath` don't use `--noad` than set ad file path, default is `./Tools/ad.html`

### v1.2

- add feature Single-line commentes are created simply by beginning in line with '#' or '-'

### v1.3

- add feature that in rule file can import other rule file .
```
@import "other_rule.md"
``` 

### v1.4

- upload rapidlycode source code in 'Code' folder.

- skip the content in <>.

### v1.5

- add `--noad` `--adpath`

### v1.6

- code refactoring, running faster .

- add feature that `UTC time -> BeiJin time` and  `AEDT  time -> BeiJin time`.

### v1.7

-  From a.html delete useless "<strong></strong>" tags in advance, will better work
