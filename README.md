# Restic exporter

Small prometheus exporter that does a single thing: It exports the time at which
the last [restic](github.com/restic/restic) backup has been done. That's it.

## Usage

This is not a standalone exporter - rather it generates `.prom` files that can
be picked up by [`node-exporter`](github.com/prometheus/node_exporter).

In order to use it, first setup `node-exporter` to watch a directory. If you are
on debian based systems, this can ususally be done by adding the
`-collector.textfile.directory` option to `/etc/default/prometheus-node-exporter`:

```
ARGS="-collector.textfile.directory=\"/var/lib/prometheus/node-exporter/\""
```

Next, create a configuration file for `restic-exporter`:

```yaml
# restic-exporter.yaml
'arbitrary-name-here':
    repository: '/path/to/repo/here'
    password: 'passwordToRepoGOesHere'
```

Now, you can let `restic-exporter` generate a `.prom` file which in turn is picked
up and exposed to prometheus by `node-exporter`.

This will generate a single stat:
```
restic_snapshot_timestamp{name="arbitrary-name-here"} 1.599849001e+09
```

## Why not make this a 'real' exporter

I went the `node-exporter` route here because I personally store backups on a NAS
with spinning hard drives that are spun down most of the time. Thus, I needed a way
to check the backup status at a specific time where the drives are already spun up.

The easiest way to achieve this was the file-generation route in combination with
a cronjob.
