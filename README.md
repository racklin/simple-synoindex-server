Simple Synology NAS Media Index Server
-----------------------------
Simple Media Index Server is a web service wrapper for Synology NAS native `synoindex`.


Introduction
====
Since Synology DSM 6.0 comes Docker support (https://www.synology.com/en-global/dsm/6.0).
Users run many media services(ex. sickrage / couchpotato / transmission) in the docker.
But we can't notify Synology NAS to reindexing new files in the docker's container.

So, with `simple-synoindex-client` , now you can run `synoindex` inside the docker's container and request simple-synoindex-server to calling native `synoindex` to reindexing your new files.


Install
====
1. Download Pre-Build Binary and Unzip it to any directory (ex. /volume1/homes/admin) .
2. Runing `simple-synoindex-server` by Task Scheduler when Boot-Up or `/etc/rc.local`.


Docker Container Settings
======
1. Add Volume `[your unzip director]/bin` to `/usr/syno/bin`
2. Modify `simple-synoindex-server.ini` `SERVER_IP` if your docker's network bridge ip is not `172.17.0.1` [OPTIONAL].


Volume Mapping Settings
======
You may add some volumes to docker's container with difference names (ex. /video/tv -> /tv ).
But Synology native `synoindex` need real pathname to indexing your files or folders.

With `[mappings]` section settings, you can setting mapping rules, `simple-synoindex-server` will remapping docker's volume to real path before calling `synoindex`.


Sample INI Settings
=======
```ini
[main]
SERVER_IP=172.17.0.1
SERVER_PORT=32699

[mappings]
/tv=/volume1/video/tv
/movies=/volume1/video/movies
```

