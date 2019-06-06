# Setup

This tool can be used to setup local development environments. **This is only a prototype.**

Setup is clever enough to inspect the directory structure of your application and determine what things should be installed.

![Setup in action](/doc/install.png)

## Elixir

Within an Elixir project with a `.tool-versions` file like this:

```
elixir 1.7.2
```

`./setup` will check:

1. If that version of Elixir is installed
    1. If not: Will suggest to you to install it via `asdf install`
2. Will attempt to install Hex for you with `mix hex.local --if-missing` -- a no-op if it is already installed
3. Will check if your dependencies are installed with `mix deps`
    1. If some are missing: installs them with `mix deps.get`

## Ruby

Within a Ruby project with a `.tool-versions` file like this:

```
ruby 2.6.2
```

`./setup` will check:

1. If that version of Ruby is installed
   1. If not: Will suggest to you to install it via `asdf install`
2. Will attempt to install Bundler for you if that gem is missing
3. Will check if your dependencies are installed with `bundle check`
   1. If some are missing: installs them with `bundle install`
