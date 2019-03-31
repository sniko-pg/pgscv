//
package main

var (
	metricsHelp = map[string]string{
		// pg_stat_database
		"pg_stat_database_xact_commit":    "Number of transactions that have been committed",
		"pg_stat_database_xact_rollback":  "Number of transactions that have been rolled back",
		"pg_stat_database_blks_read":      "Number of disk blocks read",
		"pg_stat_database_blks_hit":       "Number of times disk blocks were found already in the buffer cache",
		"pg_stat_database_tup_returned":   "Number of rows returned by queries",
		"pg_stat_database_tup_fetched":    "Number of rows fetched by queries",
		"pg_stat_database_tup_inserted":   "Number of rows inserted by queries",
		"pg_stat_database_tup_updated":    "Number of rows updated by queries",
		"pg_stat_database_tup_deleted":    "Number of rows deleted by queries",
		"pg_stat_database_conflicts":      "Number of queries canceled due to conflicts with recovery",
		"pg_stat_database_temp_files":     "Number of temporary files created by queries",
		"pg_stat_database_temp_bytes":     "Total amount of data written to temporary files by queries",
		"pg_stat_database_deadlocks":      "Number of deadlocks detected",
		"pg_stat_database_blk_read_time":  "Time spent reading data file blocks by backends, in milliseconds",
		"pg_stat_database_blk_write_time": "Time spent writing data file blocks by backends, in milliseconds",
		"pg_stat_database_db_size":        "Size of the database, in bytes",
		"pg_stat_database_stats_age_seconds": "Age of the collected statistics, in seconds",
		// pg_stat_user_tables
		"pg_stat_user_tables_seq_scan":            "Number of sequential scans initiated",
		"pg_stat_user_tables_seq_tup_read":        "Number of live rows fetched by sequential scans",
		"pg_stat_user_tables_idx_scan":            "Number of index scans initiated",
		"pg_stat_user_tables_idx_tup_fetch":       "Number of live rows fetched by index scans",
		"pg_stat_user_tables_n_tup_ins":           "Number of rows inserted",
		"pg_stat_user_tables_n_tup_upd":           "Number of rows updated (includes HOT updated rows)",
		"pg_stat_user_tables_n_tup_del":           "Number of rows deleted",
		"pg_stat_user_tables_n_tup_hot_upd":       "Number of rows HOT updated",
		"pg_stat_user_tables_n_live_tup":          "Estimated number of live rows",
		"pg_stat_user_tables_n_dead_tup":          "Estimated number of dead rows",
		"pg_stat_user_tables_n_mod_since_analyze": "Estimated number of rows modified since this table was last analyzed",
		"pg_stat_user_tables_vacuum_count":        "Number of times this table has been manually vacuumed (not counting VACUUM FULL)",
		"pg_stat_user_tables_autovacuum_count":    "Number of times this table has been vacuumed by the autovacuum daemon",
		"pg_stat_user_tables_analyze_count":       "Number of times this table has been manually analyzed",
		"pg_stat_user_tables_autoanalyze_count":   "Number of times this table has been analyzed by the autovacuum daemon",
		// pg_statio_user_tables
		"pg_statio_user_tables_heap_blks_read":  "Number of disk blocks read from this table",
		"pg_statio_user_tables_heap_blks_hit":   "Number of buffer hits in this table",
		"pg_statio_user_tables_idx_blks_read":   "Number of disk blocks read from all indexes on this table",
		"pg_statio_user_tables_idx_blks_hit":    "Number of buffer hits in all indexes on this table",
		"pg_statio_user_tables_toast_blks_read": "Number of disk blocks read from this table's TOAST table (if any)",
		"pg_statio_user_tables_toast_blks_hit":  "Number of buffer hits in this table's TOAST table (if any)",
		"pg_statio_user_tables_tidx_blks_read":  "Number of disk blocks read from this table's TOAST table indexes (if any)",
		"pg_statio_user_tables_tidx_blks_hit":   "Number of buffer hits in this table's TOAST table indexes (if any)",
		//pg_stat_user_indexes
		"pg_stat_user_indexes_idx_scan":      "Number of index scans initiated on this index",
		"pg_stat_user_indexes_idx_tup_read":  "Number of index entries returned by scans on this index",
		"pg_stat_user_indexes_idx_tup_fetch": "Number of live table rows fetched by simple index scans using this index",
		// pg_statio_user_indexes
		"pg_statio_user_indexes_idx_blks_read": "Number of disk blocks read from this index",
		"pg_statio_user_indexes_idx_blks_hit":  "Number of buffer hits in this index",
		// pg_stat_bgwriter
		"pg_stat_bgwriter_checkpoints_timed":     "Number of scheduled checkpoints that have been performed",
		"pg_stat_bgwriter_checkpoints_req":       "Number of requested checkpoints that have been performed",
		"pg_stat_bgwriter_checkpoint_write_time": "Total amount of time that has been spent in the portion of checkpoint processing where files are written to disk, in milliseconds",
		"pg_stat_bgwriter_checkpoint_sync_time":  "Total amount of time that has been spent in the portion of checkpoint processing where files are synchronized to disk, in milliseconds",
		"pg_stat_bgwriter_buffers_checkpoint":    "Number of buffers written during checkpoints",
		"pg_stat_bgwriter_buffers_clean":         "Number of buffers written by the background writer",
		"pg_stat_bgwriter_maxwritten_clean":      "Number of times the background writer stopped a cleaning scan because it had written too many buffers",
		"pg_stat_bgwriter_buffers_backend":       "Number of buffers written directly by a backend",
		"pg_stat_bgwriter_buffers_backend_fsync": "Number of times a backend had to execute its own fsync call",
		"pg_stat_bgwriter_buffers_alloc":         "Number of buffers allocated",
		// pg_stat_replication
		"pg_stat_replication_pg_wal_bytes":      "Amount of WAL generated, in bytes",
		"pg_stat_replication_pending_lag_bytes": "Amount of WAL generated but not sent, in bytes",
		"pg_stat_replication_write_lag_bytes":   "Amount of WAL sent but not written, in bytes",
		"pg_stat_replication_flush_lag_bytes":   "Amount of WAL written but not flushed, in bytes",
		"pg_stat_replication_replay_lag_bytes":  "Amount of WAL flushed but not replayed, in bytes",
		"pg_stat_replication_total_lag_bytes":   "Amount of WAL generated but not replayed, in bytes",
		"pg_stat_replication_write_lag_sec":     "Time elapsed between flushing recent WAL locally and receiving notification that this standby server has written it",
		"pg_stat_replication_flush_lag_sec":     "Time elapsed between flushing recent WAL locally and receiving notification that this standby server has written and flushed it",
		"pg_stat_replication_replay_lag_sec":    "Time elapsed between flushing recent WAL locally and receiving notification that this standby server has written, flushed and applied it",
		"pg_replication_standby_count":          "Total number of connected standbys",
		// pg_replication_slots
		"pg_replciation_slots_conn":              "Number of opened replication slots",
		"pg_replciation_slots_restart_lag_bytes": "Amount of WAL since last restart position, in bytes",
		// recovery status
		"pg_recovery_status": "Current recovery status of Postgres service",
		//pg_stat_user_functions
		"pg_stat_user_functions_calls":      "Number of times this function has been called",
		"pg_stat_user_functions_total_time": "Total time spent in this function and all other functions called by it, in milliseconds",
		"pg_stat_user_functions_self_time":  "Total time spent in this function itself, not including other functions called by it, in milliseconds",
		// pg_stat_database_conflicts
		"pg_stat_database_conflicts_total":      "Total number of recovery conflicts occurred",
		"pg_stat_database_conflicts_tablespace": "Total number of recovery conflicts occurred due to dropped tablespaces",
		"pg_stat_database_conflicts_lock":       "Total number of recovery conflicts occurred due to lock timeouts",
		"pg_stat_database_conflicts_snapshot":   "Total number of recovery conflicts occurred due to old snapshot",
		"pg_stat_database_conflicts_bufferpin":  "Total number of recovery conflicts occurred due to pinned buffers",
		"pg_stat_database_conflicts_deadlock":   "Total number of recovery conflicts occurred due to deadlocks",
		// pg_stat_archiver
		// pg_stat_sizes
		// pg_stat_activity
		"pg_stat_activity_conn_total":                           "Total number of backends",
		"pg_stat_activity_conn_idle_total":                      "Number of backends that waiting for a new client command",
		"pg_stat_activity_conn_idle_xact_total":                 "Number of backends that in a transaction (active or failed), but is not currently executing a query.",
		"pg_stat_activity_conn_active_total":                    "Number of backends that executing a query",
		"pg_stat_activity_conn_waiting_total":                   "Number of backends that waiting another beckends",
		"pg_stat_activity_conn_others_total":                    "Number of backends that executing a fast-path function or have another state",
		"pg_stat_activity_conn_prepared_total":                  "Total number of prepared statements",
		"pg_stat_activity_xact_max_duration":                    "Duration of the longest transaction, in seconds",
		"pg_stat_activity_autovac_workers_total":                "Total number of autovacuum workers",
		"pg_stat_activity_autovac_antiwraparound_workers_total": "Total number of anti-wraparound autovacuum workers",
		"pg_stat_activity_autovac_user_vacuum_total":            "Total number of vacuum workers started by user",
		"pg_stat_activity_autovac_max_duration":                 "Duration of the longest (auto)vacuum worker, in seconds",
		// pg_stat_statements
		"pg_stat_statements_calls":               "Number of times query has been executed",
		"pg_stat_statements_total_time":          "Total time spent in the statement, in milliseconds",
		"pg_stat_statements_rows":                "Total number of rows retrieved or affected by the statement",
		"pg_stat_statements_shared_blks_hit":     "Total number of shared block cache hits by the statement",
		"pg_stat_statements_shared_blks_read":    "Total number of shared blocks read by the statement",
		"pg_stat_statements_shared_blks_dirtied": "Total number of shared blocks dirtied by the statement",
		"pg_stat_statements_shared_blks_written": "Total number of shared blocks written by the statement",
		"pg_stat_statements_local_blks_hit":      "Total number of local block cache hits by the statement",
		"pg_stat_statements_local_blks_read":     "Total number of local blocks read by the statement",
		"pg_stat_statements_local_blks_dirtied":  "Total number of local blocks dirtied by the statement",
		"pg_stat_statements_local_blks_written":  "Total number of local blocks written by the statement",
		"pg_stat_statements_temp_blks_read":      "Total number of temp blocks read by the statement",
		"pg_stat_statements_temp_blks_written":   "Total number of temp blocks written by the statement",
		"pg_stat_statements_blk_read_time":       "Total time the statement spent reading blocks, in milliseconds",
		"pg_stat_statements_blk_write_time":      "Total time the statement spent writing blocks, in milliseconds",
		// pg_stat_basebackup
		"pg_stat_basebackup_count":                "Total number of basebackups currently running",
		"pg_stat_basebackup_duration_seconds_max": "Duration of the longest basebackup",
		"pg_wal_directory_size_bytes":             "Total size of WAL directory",
		"pg_wal_directory":                        "A metric with a constant '1' value with details about WAL directory",
		"pg_data_directory":                       "A metric with a constant '1' value with details about DATA directory",
		"pg_catalog_size_bytes":                   "Total size of pg_catalog, in bytes",
		// active temp files
		"pg_stat_current_temp_files_total":                 "Total number of active temp files",
		"pg_stat_current_temp_bytes_total":                 "Total amount of space used by temp files, in bytes",
		"pg_stat_current_temp_oldest_file_age_seconds_max": "Total time when oldest temp file is created, in seconds",
		// pg_class
		// pg_locks
		// pg_settings
		"pg_settings_guc": "Postgres configuration settings",
		// pg schema
		"pg_schema_non_pk_table_exists": "A metric with a constant '1' value labeled by datname, schemaname, relname of relation with no primary key",
		"pg_schema_invalid_index_bytes": "Total size of invalid index, in bytes",
		"pg_schema_non_indexed_fkey_exists": "A metric with a constant '1' value labeled by datname, schemaname, relname, colnames, constraint, referenced of foreign key constrint which has no index",
		"pg_schema_redundant_index_bytes": "Size occupied by redundant index, in bytes",
		"pg_schema_sequence_fullness_ratio": "Fullness ratio of the sequence, in percent",
		"pg_schema_fkey_columns_mismatch_exists": "A metric with a constant '1' value labeled by datname, schemanames, relnames, colnames of fkey whose columns have different type",
		// node cpu metrics
		"node_cpu_usage_time": "Node CPU usage, in ticks",
		// node /proc/diskstats metrics
		"node_diskstats_rcompleted":   "Total number of read operations completed successfully",
		"node_diskstats_rmerged":      "Total number of merged read operations",
		"node_diskstats_rsectors":     "Total number of sectors read",
		"node_diskstats_rspent":       "Total time the device spent reading, in milliseconds",
		"node_diskstats_wcompleted":   "Total number of write operations completed successfully",
		"node_diskstats_wmerged":      "Total number of merged write operations",
		"node_diskstats_wsectors":     "Total number of sectors written",
		"node_diskstats_wspent":       "Total time the device spent writing, in milliseconds",
		"node_diskstats_ioinprogress": "Number of I/Os currently in progress",
		"node_diskstats_tspent":       "Total time the device spent doing I/Os, in milliseconds",
		"node_diskstats_tweighted":    "Total weighted time the device spent doing I/Os, in milliseconds",
		"node_diskstats_uptime":       "Total number of CPU ticks",
		// node /proc/net/dev metrics
		"node_netdev_rbytes":      "Total number of received bytes",
		"node_netdev_rpackets":    "Total number of received packets",
		"node_netdev_rerrs":       "Total number of receive errors",
		"node_netdev_rdrop":       "Total number of dropped packets",
		"node_netdev_rfifo":       "Total number of fifo buffers errors",
		"node_netdev_rframe":      "Total number of packet framing errors",
		"node_netdev_rcompressed": "Total number of received compressed packets",
		"node_netdev_rmulticast":  "Total number of received multicast packets",
		"node_netdev_tbytes":      "Total number of transmitted bytes",
		"node_netdev_tpackets":    "Total number of transmitted packets",
		"node_netdev_terrs":       "Total number of transmitted errors",
		"node_netdev_tdrop":       "Total number of dropped packets",
		"node_netdev_tfifo":       "Total number of fifo buffers errors",
		"node_netdev_tcolls":      "Total number of detected collisions",
		"node_netdev_tcarrier":    "Total number of carrier losses",
		"node_netdev_tcompressed": "Total number of received multicast packets",
		"node_netdev_saturation":  "Total number of errors seen for the interface",
		"node_netdev_uptime":      "Total number of CPU ticks",
		"node_netdev_speed":       "An interface network speed",
		"node_netdev_duplex":      "An interface duplex",
		// node /proc/meminfo metrics
		"node_memory_usage_bytes": "Node memory usage, in bytes",
		// node filesystems
		"node_filesystem_bytes":  "Node filesystem usage, in bytes",
		"node_filesystem_inodes": "Node filesystem usage, in inodes",
		// node various settings
		"node_settings_sysctl": "Node sysctl variables",
		// node hardware
		"node_hardware_cores_total":             "Total number of CPU cores",
		"node_hardware_scaling_governors_total": "Total number of scaling governors used by CPU cores",
		"node_hardware_numa_nodes":              "Total number of NUMA nodes",
		"node_hardware_storage_rotational":      "Type of the connected storage",
		// pgbouncer
		"pgbouncer_pool_cl_active":       "Client connections that are linked to server connection and can process queries",
		"pgbouncer_pool_cl_waiting":      "Client connections have sent queries but have not yet got a server connection",
		"pgbouncer_pool_sv_active":       "Server connections that linked to client",
		"pgbouncer_pool_sv_idle":         "Server connections that unused and immediately usable for client queries",
		"pgbouncer_pool_sv_used":         "Server connections that have been idle more than server_check_delay, so they needs server_check_query to run on it before it can be used",
		"pgbouncer_pool_sv_tested":       "Server connections that are currently running either server_reset_query or server_check_query",
		"pgbouncer_pool_sv_login":        "Server connections currently in logging in process",
		"pgbouncer_pool_maxwait":         "How long the first (oldest) client in queue has waited, in seconds",
		"pgbouncer_pool_maxwait_us":      "How long the first (oldest) client in queue has waited, in microseconds",
		"pgbouncer_stats_xact_count":     "Total number of SQL transactions pooled by pgbouncer",
		"pgbouncer_stats_query_count":    "Total number of SQL queries pooled by pgbouncer",
		"pgbouncer_stats_bytes_received": "Total volume in bytes of network traffic received by pgbouncer",
		"pgbouncer_stats_bytes_sent":     "Total volume in bytes of network traffic sent by pgbouncer",
		"pgbouncer_stats_xact_time":      "Total number of microseconds spent by pgbouncer when connected to PostgreSQL in a transaction, either idle in transaction or executing queries",
		"pgbouncer_stats_query_time":     "Total number of microseconds spent by pgbouncer when actively connected to PostgreSQL, executing queries",
		"pgbouncer_stats_wait_time":      "Time spent by clients waiting for a server in microseconds",
	}
)
