-- Database structure of penchan.club
-- Every post is stored in the post table
-- there can be multiple boards, they don't need to have their own table, by default posts go to /b/
-- picture is a base64 png data URL
-- thread point to the OP of a thread, if thread is null the post begin a new thread
-- reply_to is if someone wanna reply to someone else
-- ip_address is poster's IP, in case someone spam CP and my server get seized
-- special is mostly to mark messages as being sent by VIPs (like mods or site admin)

BEGIN;

CREATE TABLE "post" (
	"id" INTEGER PRIMARY KEY AUTOINCREMENT,
	"board" TEXT DEFAULT 'b',
	"picture" TEXT NOT NULL,
	"thread" INTEGER REFERENCES "post"("id"),
	"reply_to" INTEGER REFERENCES "post"("id"),
	"ip_address" TEXT,
	"special" TEXT,
	"created_at" INTEGER DEFAULT (unixepoch()),
	FOREIGN KEY (thread) REFERENCES post(id),
	FOREIGN KEY (reply_to) REFERENCES post(id)
);

COMMIT;
