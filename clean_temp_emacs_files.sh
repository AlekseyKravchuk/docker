#!/usr/bin/env bash

# delete all files which names end with '~' (delete all temporary emacs files)
find ./ -name "*~" -exec rm {} \;
