# haystacks-go

## Build

    ./build.sh

## Info

Check if the needles exist in the haystack. Example:

    cat test/haystack | haystacks --needles test/needles

Reads the haystack from `stdin` and outputs matches to `stdout`.


## Generating performance data

Try out `needlehay` on a bit of data. This test uses the `ossp-uuid` tool (1) to generate uuids
and `shuf`/`gshuf` (2) to make a random permutation:

    uuid -v4 -n100000000 > hundred_million_uuids.txt
    gshuf -n100000 hundred_million_uuids.txt > hundred_million_uuids_sample100k.txt

then

    cat hundred_million_uuids.txt | needlehay --needles hundred_million_uuids_sample100k.txt


* (1) `brew install ossp-uuid`
* (2) `brew install coreutils`
