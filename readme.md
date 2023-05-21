`# Checkout Kata

## Running

You can either run this application natively by building in the application (see Makefile) and then running the compiled
binary, or you can use the docker-compose file to spin up the application inside a docker-container.
The application runs on port 8080

## Features

The application provides two endpoints `/checkout/scan/{itemSKU}` to 'scan' an item and `/checkout/total` to get the
total for all the scanned items.

### Thought Process / Assumptions

While the main objective of the task was to provide a simple application which provided the checkout features required, 
I decided to spend a bit more time than required to showcase how I may approach this if it was a more 'real world' implementation 
so I have included the api feature as well as Dockerized  the application etc. 

The spec was quite vague but of course there are many things missing in this application that would be in prod namely:
  - The lack of any storage ( such as in memory or a SQL database)
  - The lack of any 'basket state' to allow multiple users to use the application 
  - Logging is missing 
  - A linter such as go lint 
  - The basket cannot be cleared, so to start fresh the application needs to be restarted
  - Ability to add / remove / delete items

Anyway I hope you like what I have done, and I hope the extra work i've put in to make this more real world makes up for the delay
in me getting this back to you. 

Cheers! 