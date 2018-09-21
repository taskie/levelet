# levelet

too simple LevelDB manipulator

## install

```sh
go get github.com/taskie/levelet/cmd/levelet
```

## usage

### get

```sh
levelet -f foo.ldb g key
# env LEVELET_DB_PATH=foo.ldb levelet g key
```

### put

```sh
echo value | levelet -f foo.ldb p key
```

### delete

```sh
levelet -f foo.ldb d key
```

## dependency

![dependency](images/dependency.png)

## license

Apache License 2.0
