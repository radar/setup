# Setup

**This is only a prototype.**

This tool can be used to setup local development environments. It is designed to replace bespoke, hand-crafted `bin/setup` scripts.

Setup is clever enough to inspect the directory structure of your application and determine what things should be installed. Think of it like Heroku's `git push`, but on your local machine.

![Setup in action](/doc/install.png)

## Walkthrough

Here's a walkthrough of all the things that setup will check for.

### Homebrew

If there is a `Brewfile` file in the current directory, setup will:

1. run `brew bundle check` to see if the specified dependencies are installed.
    1. If they aren't, it will recommend to run `brew bundle install`
    2. (Setup will not run `brew bundle install` automatically, because Homebrew changes might break other dependencies)

### Languages (with ASDF)

A `.tool-versions` file is the standard used by the [asdf](https://github.com/asdf-vm/asdf) project.

Setup will check for a `.tool-versions` file and will attempt to install any language specified in that file. Below are a list of supported languages by setup:

* Elixir
* Elm
* Erlang
* Node
* Ruby

#### Elixir

Within an Elixir project with a `.tool-versions` file like this:

```
elixir 1.7.2
```

`setup` will check:

1. If Erlang is installed
    1. If not: it will recommend installing it via either Homebrew or `asdf`
1. If the `elixir` asdf plugin is installed
    1. If not: install it automatically for you
1. If that version of Elixir is installed via the `asdf` plugin
    1. If not: install it automatically for you
1. Attempts to install Hex for you (used for dependencies)
1. Attempt to install Rebar for you (used for compiling some dependencies)
1. Checks if a `mix.exs` file exists.
    1. Will check if your dependencies are installed with `mix deps`
    1. If some are missing: installs them with `mix deps.get`
    1. Runs `mix compile` to ensure the application is compiled
    1. Runs `mix ecto.setup` (if available)

#### Elm

Within an Elm project with a `.tool-versions` file like this:

```
elm 0.18.0
```

`setup` will check:

1. If the `elm` asdf plugin is installed
    1. If not: install it automatically for you
1. If that version of Elm is installed via the `asdf` plugin
    1. If not: install it automatically for you
1. Runs `elm-package install --yes` to install Elm dependencies


### Node

Within an Node project with a `.tool-versions` file like this:

```
nodejs 8.14.0
```

`setup` will check:

1. If the `nodejs` asdf plugin is installed
    1. If not: installs it automatically for you
1. If that version of nodejs is installed via the `asdf` plugin
    1. If not: installs it automatically for you
1. Checks if Yarn is installed
    1. If not: installs it automatically for you
1. Installs project dependencies with `yarn install`

#### Ruby

Within a Ruby project with a `.tool-versions` file like this:

```
ruby 2.6.2
```

`setup` will check:

1. If the `ruby` asdf plugin is installed
    1. If not: installs it automatically for you
1. If that version of ruby is installed via the `asdf` plugin
    1. If not: installs it automatically for you
1. If the `bundler` gem is installed
    1. If not: installs it automatically for you
1. If the application's dependencies are installed (with `bundle check`)
    1. If not: will run `bundle install` to install them
1. If the bundle includes `mongoid`
    1. If so: Checks if a MongoDB server is running at `localhost:27017`
        1. If not: Recommends starting it with `brew services`
1. If the bundle includes `pg`
    1. If so: Checks if a PostgreSQL server is running at `localhost:5432`
        1. If not: Recommends starting it with `brew services`

### bin/setup

If there is a `bin/setup` file located in the current directory, this file will be executed once the languages are setup. This script could do things like:

* Run migrations
* Re-build static assets
* etc.

Really, whatever you put in this file will be executed as a bash script! It's only limited by your imagination.
