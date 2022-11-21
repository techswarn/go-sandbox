#!/bin/bash

starting_dir=$(realpath .)

echo "$starting_dir"

#delete vendor directory if expreset

rm -r hello/vendor 2> /dev/null

cd hello || exit

go mod vendor

#cleanup

cd "${starting_dir}" || exit




