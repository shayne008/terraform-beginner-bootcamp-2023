# Terraform Beginner Bootcamp 2023 - Week 2

## Working with Ruby

### Bundler

Bundler is package manager for Ruby. It is the primary way to install ruby packages ( known as gems) for Ruby.

#### Install Gems

You need to create a Gemfile and define your gems in that file.

```rb
source "https://rubygems.org"

gem 'sinatra'
gem 'rake'
gem 'pry'
gem 'puma'
gem 'activerecord'
```

Then you run the `bundle install` command

This will install the gems on the system globally (unlike nodejs which install packages in place in a folder called node_modules)

A Gemfile.lock will be created to lock down the versions used in the project.

#### Executing Ruby Scripts in the context of bundler.

We have to use `bundle exec` to tell future ruby scripts to use the gems we installed. This is the way we set context.


### Sinatra

Sinatra is a mirco web-framework for Ruby to build web-apps.

It is great for mock or development servers or simple projects.

You can create a web-server in a single file.

[Sinatra](https://sinatrarb.com/)


## Terratowns Mock Server

### Running the web server

We can run the web server by executing the following commands:

```rb
bundle install
bundle exec ruby server.rb
```

All of the code for our server is stored in the `server.rb` file.

## CRUD

Terraform Provider resources utilize CRUD.

CRUD stands for create, read,update and delete.

[CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete)

