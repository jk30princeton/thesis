#! /bin/bash

nix-store --tree -q $(nix-instantiate default.nix)