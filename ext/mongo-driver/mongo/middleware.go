package mongo

import (
	"context"

	"github.com/asppj/lolita/ext/log-driver/log"

	"github.com/opentracing/opentracing-go"
)

func (d *Dao) cmdSpan(ctx context.Context, cmd string, companyID string) opentracing.Span {
	sp := opentracing.SpanFromContext(ctx)
	if sp == nil {
		sp = opentracing.StartSpan("MgoSpan")
		log.Warn("ctx has not span")
	}
	nsp := opentracing.GlobalTracer().StartSpan(
		"MongoDB-"+cmd,
		opentracing.ChildOf(sp.Context()),
		opentracing.Tags{},
	)
	nsp.SetTag("db.Type", "mongo")
	nsp.SetTag("db.ClientName", d.ClientName)
	nsp.SetTag("db.DBName", d.DBName)
	nsp.SetTag("db.CollName", d.CollName)
	nsp.SetTag("db.Cmd", cmd)
	nsp.SetTag("db.CompanyID", companyID)
	return nsp
}
