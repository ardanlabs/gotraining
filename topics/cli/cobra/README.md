## Cobra

Cobra is both a library for creating powerful modern CLI applications as well as a program to generate applications and command files. Cobra is also an application that will generate your application scaffolding to rapidly develop a Cobra-based application.

## Notes

* Easy subcommand-based CLIs: app server, app fetch, etc.
* Fully POSIX-compliant flags (including short & long versions)
* Nested subcommands
* Global, local and cascading flags
* Easy generation of applications & commands with cobra create appname & cobra add cmdname
* Intelligent suggestions (app srver... did you mean app server?)
* Automatic help generation for commands and flags
* Automatic detailed help for app help [command]
* Automatic help flag recognition of -h, --help, etc.
* Automatically generated bash autocomplete for your application
* Automatically generated man pages for your application
* Command aliases so you can change things without breaking them
* The flexibilty to define your own help, usage, etc.
* Optional tight integration with viper for 12-factor apps

## Links

https://github.com/spf13/cobra  
http://spf13.com/post/announcing-cobra

## Code Review

[main](main.go)  
[commands](cmduser/commands.go)  
[create command](cmduser/create.go)

## Exercises

### Exercise 1

Add a new user command called get to retrieve a user record by email. For now reuse the User type and create a value and return some fake data. Then display the user value in the terminal. The purpose of the exercise it to wire up a new command. Use the create command as your template.

[Answer](exercises/exercise1/exercise1.go)
___
All material is licensed under the [Apache License Version 2.0, January 2004](http://www.apache.org/licenses/LICENSE-2.0).