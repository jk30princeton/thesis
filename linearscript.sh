#! /bin/bash

nix-store -qR $(nix-instantiate default.nix)