#! /bin/bash

nix-store -qR $(nix-instantiate default.nix)

#nix-store -q --requisites $(nix-instantiate default.nix) 