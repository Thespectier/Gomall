#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/cpayment"
exec "$CURDIR/bin/cpayment"
