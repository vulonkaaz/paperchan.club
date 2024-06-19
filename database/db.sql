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
	"id" int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
	"board" text DEFAULT 'b',
	"picture" text NOT NULL,
	"thread" int REFERENCES "post"("id"),
	"reply_to" int REFERENCES "post"("id"),
	"ip_address" text,
	"special" text,
	"created_at" timestamptz NOT NULL DEFAULT now(),
	"updated_at" timestamptz
);

COMMIT;
