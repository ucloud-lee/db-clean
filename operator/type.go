package operator

import "database/sql"

type Clean struct {
	db      *sql.DB
	start   string
	end     string
	cluster string
}

var sqls = []string{
	"delete from daily_pod_bill where created_at>'%s' and created_at<'%s' and cluster_name='%s'",
	"delete from daily_pod_info where created_at>'%s' and created_at<'%s' and cluster_name='%s'",
	"delete from daily_project_bill where created_at>'%s' and created_at<'%s' and cluster_name='%s'",
	"delete from daily_project_info where created_at>'%s' and created_at<'%s' and cluster_name='%s'",
	"delete from daily_pvc_info where created_at>'%s' and created_at<'%s' and cluster_name='%s'",
	"delete from daily_sc_info where created_at>'%s' and created_at<'%s' and cluster_name='%s'",
	"delete from daily_service_info where created_at>'%s' and created_at<'%s' and cluster_name='%s'",
	"delete from daily_crd_info where created_at>'%s' and created_at<'%s' and cluster_name='%s'",
	"delete from node_usage_info where created_at>'%s' and created_at<'%s' and cluster='%s'",
	"delete from daily_ip_info where created_at>'%s' and created_at<'%s' and cluster='%s'",
	"delete from daily_istio where created_at>'%s' and created_at<'%s' and cluster='%s'",
	"delete from daily_cluster_info where created_at>'%s' and created_at<'%s' and name='%s'",
}
