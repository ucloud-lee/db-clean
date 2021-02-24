package operator

import (
	"flag"
	"fmt"

	log "github.com/gogap/logrus"
)

func (c *Clean) StartCleanDB() {
	c.getflag()

	log.Info("start: ", c.start, "  end: ", c.end, "  cluster: ", c.cluster)
	if c.cluster == "" || c.start == "" || c.end == "" {
		return
	}

	c.db = c.clientMysql()
	defer c.db.Close()

	for _, v := range sqls {
		sql := fmt.Sprintf(v, c.start, c.end, c.cluster)
		log.Info(sql)
		c.deleteData(sql)
	}
}

func (c *Clean) deleteData(sql string) {
	ret, err := c.db.Exec(sql)
	if err != nil {
		log.Error("err: ", err)
		return
	}
	num, err := ret.RowsAffected()
	if err != nil {
		log.Error("fetch row affected failed:", err.Error())
		return
	}
	log.Info("delete record number", num)
}

func (c *Clean) getflag() {
	flag.StringVar(&c.start, "start", "", "delete db data start time")
	flag.StringVar(&c.end, "end", "", "delete db data end time")
	flag.StringVar(&c.cluster, "cluster", "", "delete db which cluster data")
	flag.Parse()
}
