# Paperchan.club

This is the source code of paperchan.club, an imageboard where all posts are drawn and handwritten by hand.

This software requires golang, postgresql, base64 and imagemagick (for image conversion, see themagicpipe)

## How to install

- Create a postgres database, populate it with the stuff in database/db.sql
- Create .env with DBSTRING to connect to the database
- `go build`
- launch the executable to start the server
