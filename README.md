# Paperchan.club

This is the source code of paperchan.club, an imageboard where all posts are drawn and handwritten by hand.

This software requires golang, sqlite3, base64 and imagemagick (for image conversion, see themagicpipe)

## How to install

- Create a sqlite3 database and save it as paperdb.db3 in the server's working directory, populate it with the stuff in database/db.sql
  Or use https://github.com/vulonkaaz/paperchan-db-convert to import your old postgres database
- Create .env if you need a moderation password, check .env.example
- `go build`
- launch the executable to start the server
