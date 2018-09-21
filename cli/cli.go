package cli

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/taskie/levelet"
	"github.com/urfave/cli"
	"io"
	"io/ioutil"
	"os"
)

var (
	log              = logrus.New()
	stdin  io.Reader = os.Stdin
	stdout io.Writer = os.Stdout
)

func prepare(c *cli.Context) (db *leveldb.DB, key []byte, err error) {
	if c.NArg() == 0 {
		return nil, nil, fmt.Errorf("you must specify key")
	}
	key = []byte(c.Args().Get(0))
	dbPath := c.GlobalString("f")
	if dbPath == "" {
		return nil, nil, fmt.Errorf("you must specify DB path (--dbPath/-f)")
	}
	db, err = leveldb.OpenFile(dbPath, nil)
	return
}

func getAction(c *cli.Context) error {
	db, key, err := prepare(c)
	if err != nil {
		return err
	}
	defer db.Close()
	data, err := db.Get(key, nil)
	if err != nil {
		return err
	}
	stdout.Write(data)
	return nil
}

func putAction(c *cli.Context) error {
	db, key, err := prepare(c)
	if err != nil {
		return err
	}
	defer db.Close()
	value, err := ioutil.ReadAll(stdin)
	if err != nil {
		return err
	}
	err = db.Put(key, value, nil)
	if err != nil {
		return err
	}
	return nil
}

func deleteAction(c *cli.Context) error {
	db, key, err := prepare(c)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Delete(key, nil)
	if err != nil {
		return err
	}
	return nil
}

func mainImpl() error {
	app := cli.NewApp()
	app.Name = "levelet"
	app.Version = levelet.Version
	app.Usage = "too simple LevelDB manipulator"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "dbPath, f",
			Usage:  "LevelDB database file",
			EnvVar: "LEVELET_DB_PATH",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:    "get",
			Aliases: []string{"g"},
			Usage:   "get value from DB",
			Action:  getAction,
		},
		{
			Name:    "put",
			Aliases: []string{"p"},
			Usage:   "put value to DB",
			Action:  putAction,
		},
		{
			Name:    "delete",
			Aliases: []string{"d"},
			Usage:   "delete key and value from DB",
			Action:  deleteAction,
		},
	}

	return app.Run(os.Args)
}

func Main() {
	err := mainImpl()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
