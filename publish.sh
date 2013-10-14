#!/bin/bash

st=$(git status --short)
if [ -n "$st" ]; then
    echo "error: uncommited changes"
    exit 1
fi

/opt/google-appengine-go/appcfg.py update .
git tag -f deployed
git push
git push --tags -f
